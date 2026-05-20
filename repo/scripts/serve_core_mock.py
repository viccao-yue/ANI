#!/usr/bin/env python3
"""Serve a lightweight Core API mock from the OpenAPI contract.

This is a dependency-free local fallback for Sprint 4 MOCK-A. It keeps the
mock behavior driven by api/openapi/v1.yaml so a future Prism entrypoint can
replace the server without changing the contract checks.
"""

from __future__ import annotations

from dataclasses import dataclass
from http.server import BaseHTTPRequestHandler, HTTPServer
from pathlib import Path
from urllib.parse import urlparse
import argparse
import json
import re
from typing import Any

import yaml


HTTP_METHODS = {"get", "post", "put", "patch", "delete"}
ROOT = Path(__file__).resolve().parents[1]
DEFAULT_SPEC = ROOT / "api/openapi/v1.yaml"


@dataclass(frozen=True)
class MockRoute:
    method: str
    path_template: str
    operation_id: str
    status_code: int
    content_type: str
    body: Any
    pattern: re.Pattern[str]


def load_spec(path: Path) -> dict[str, Any]:
    with path.open(encoding="utf-8") as handle:
        data = yaml.safe_load(handle)
    if not isinstance(data, dict):
        raise ValueError(f"{path} must parse to an object")
    return data


def server_base_path(spec: dict[str, Any]) -> str:
    servers = spec.get("servers") or []
    if not servers:
        return ""
    path = urlparse(servers[0].get("url", "")).path.rstrip("/")
    return path if path != "/" else ""


def build_routes(spec: dict[str, Any]) -> list[MockRoute]:
    routes: list[MockRoute] = []
    for path, path_item in sorted(spec.get("paths", {}).items()):
        if not isinstance(path_item, dict):
            continue
        for method, operation in sorted(path_item.items()):
            if method not in HTTP_METHODS or not isinstance(operation, dict):
                continue
            response = choose_success_response(spec, operation)
            status_code = response["status_code"]
            media = response.get("media") or {}
            schema = media.get("schema", {})
            content_type = response.get("content_type") or "application/json"
            body = None if status_code == 204 else mock_value(spec, schema, operation.get("operationId", "mock"))
            routes.append(
                MockRoute(
                    method=method.upper(),
                    path_template=path,
                    operation_id=operation.get("operationId") or fallback_operation_id(method, path),
                    status_code=status_code,
                    content_type=content_type,
                    body=body,
                    pattern=compile_path(path),
                )
            )
    return routes


def choose_success_response(spec: dict[str, Any], operation: dict[str, Any]) -> dict[str, Any]:
    responses = operation.get("responses", {})
    if not isinstance(responses, dict):
        raise ValueError(f"{operation.get('operationId', '<unknown>')} missing responses")
    for status in sorted(responses):
        if not str(status).startswith("2"):
            continue
        response = resolve_ref(spec, responses[status])
        content = response.get("content", {}) if isinstance(response, dict) else {}
        if "application/json" in content:
            return {
                "status_code": int(status),
                "content_type": "application/json",
                "media": content["application/json"],
            }
        if str(status) == "204":
            return {"status_code": 204, "content_type": "", "media": {}}
        first_content_type = next(iter(content), "application/json")
        return {
            "status_code": int(status),
            "content_type": first_content_type,
            "media": content.get(first_content_type, {}),
        }
    raise ValueError(f"{operation.get('operationId', '<unknown>')} missing 2xx response")


def mock_value(spec: dict[str, Any], schema: dict[str, Any], operation_id: str) -> Any:
    schema = resolve_ref(spec, schema or {})
    if not isinstance(schema, dict) or not schema:
        return {"operation_id": operation_id, "mock": True}
    if "example" in schema:
        return schema["example"]
    if "enum" in schema and schema["enum"]:
        return schema["enum"][0]
    if "oneOf" in schema:
        return mock_value(spec, schema["oneOf"][0], operation_id)
    if "anyOf" in schema:
        return mock_value(spec, schema["anyOf"][0], operation_id)
    if "allOf" in schema:
        merged: dict[str, Any] = {}
        for item in schema["allOf"]:
            value = mock_value(spec, item, operation_id)
            if isinstance(value, dict):
                merged.update(value)
        return merged
    schema_type = schema.get("type")
    if isinstance(schema_type, list):
        schema_type = next((item for item in schema_type if item != "null"), schema_type[0])
    if schema_type == "object" or "properties" in schema:
        properties = schema.get("properties", {})
        if not properties:
            return {}
        return {name: mock_value(spec, child, operation_id) for name, child in properties.items()}
    if schema_type == "array":
        return [mock_value(spec, schema.get("items", {}), operation_id)]
    if schema_type == "integer":
        return 1
    if schema_type == "number":
        return 1.0
    if schema_type == "boolean":
        return True
    if schema.get("format") == "uuid":
        return "00000000-0000-4000-8000-000000000001"
    if schema.get("format") == "date-time":
        return "2026-05-20T00:00:00Z"
    return "mock"


def resolve_ref(spec: dict[str, Any], value: Any) -> Any:
    if not isinstance(value, dict) or "$ref" not in value:
        return value
    ref = value["$ref"]
    if not ref.startswith("#/"):
        raise ValueError(f"unsupported external ref {ref}")
    current: Any = spec
    for part in ref[2:].split("/"):
        current = current[part]
    return current


def compile_path(path: str) -> re.Pattern[str]:
    escaped = re.escape(path)
    pattern = re.sub(r"\\\{[^/]+\\\}", r"[^/]+", escaped)
    return re.compile("^" + pattern + "$")


def fallback_operation_id(method: str, path: str) -> str:
    tokens = re.sub(r"[^a-zA-Z0-9]+", " ", path).title().replace(" ", "")
    return method.lower() + tokens


def find_route(routes: list[MockRoute], method: str, path: str, base_path: str) -> MockRoute | None:
    candidate = path
    if base_path and candidate.startswith(base_path + "/"):
        candidate = candidate[len(base_path) :]
    for route in routes:
        if route.method == method and route.pattern.match(candidate):
            return route
    return None


def make_handler(routes: list[MockRoute], base_path: str) -> type[BaseHTTPRequestHandler]:
    class CoreMockHandler(BaseHTTPRequestHandler):
        def do_OPTIONS(self) -> None:
            self.send_response(204)
            self.send_header("Access-Control-Allow-Origin", "*")
            self.send_header("Access-Control-Allow-Headers", "Authorization,Content-Type,X-API-Key")
            self.send_header("Access-Control-Allow-Methods", "GET,POST,PUT,PATCH,DELETE,OPTIONS")
            self.end_headers()

        def do_GET(self) -> None:
            self.respond()

        def do_POST(self) -> None:
            self.respond()

        def do_PUT(self) -> None:
            self.respond()

        def do_PATCH(self) -> None:
            self.respond()

        def do_DELETE(self) -> None:
            self.respond()

        def log_message(self, fmt: str, *args: Any) -> None:
            return

        def respond(self) -> None:
            path = urlparse(self.path).path
            route = find_route(routes, self.command, path, base_path)
            if route is None:
                self.send_json(404, {"code": "NOT_FOUND", "message": "mock route not found", "request_id": "mock"})
                return
            if route.status_code == 204:
                self.send_response(204)
                self.send_header("X-ANI-Mock-Operation", route.operation_id)
                self.end_headers()
                return
            self.send_json(route.status_code, route.body, route.content_type, route.operation_id)

        def send_json(
            self,
            status_code: int,
            body: Any,
            content_type: str = "application/json",
            operation_id: str = "",
        ) -> None:
            payload = json.dumps(body, ensure_ascii=False).encode("utf-8")
            self.send_response(status_code)
            self.send_header("Content-Type", content_type)
            self.send_header("Content-Length", str(len(payload)))
            self.send_header("Access-Control-Allow-Origin", "*")
            if operation_id:
                self.send_header("X-ANI-Mock-Operation", operation_id)
            self.end_headers()
            self.wfile.write(payload)

    return CoreMockHandler


def main() -> None:
    parser = argparse.ArgumentParser(description="Serve ANI Core API mock from api/openapi/v1.yaml")
    parser.add_argument("--spec", default=str(DEFAULT_SPEC), help="OpenAPI contract path")
    parser.add_argument("--host", default="127.0.0.1")
    parser.add_argument("--port", type=int, default=4010)
    args = parser.parse_args()

    spec = load_spec(Path(args.spec))
    base_path = server_base_path(spec)
    routes = build_routes(spec)
    server = HTTPServer((args.host, args.port), make_handler(routes, base_path))
    print(f"ANI Core mock server listening on http://{args.host}:{args.port}{base_path} ({len(routes)} routes)")
    server.serve_forever()


if __name__ == "__main__":
    main()
