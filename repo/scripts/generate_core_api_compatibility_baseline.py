#!/usr/bin/env python3
"""Generate the Core API v1 compatibility baseline used by Sprint 4 guards."""

from __future__ import annotations

from datetime import date
from pathlib import Path
from typing import Any

import yaml


ROOT = Path(__file__).resolve().parents[1]
SPEC_PATH = ROOT / "api/openapi/v1.yaml"
BASELINE_PATH = ROOT / "api/core-v1-compatibility-baseline.yaml"


def schema_ref(schema: dict[str, Any]) -> str:
    ref = schema.get("$ref")
    return ref.rsplit("/", 1)[-1] if isinstance(ref, str) else ""


def normalize_schema(schema: Any) -> Any:
    if not isinstance(schema, dict):
        return schema
    normalized: dict[str, Any] = {}
    for key in ("$ref", "type", "format", "nullable", "minimum", "maximum", "minLength", "maxLength", "pattern"):
        if key in schema:
            normalized[key] = schema[key]
    if "items" in schema:
        normalized["items"] = normalize_schema(schema["items"])
    if "additionalProperties" in schema:
        normalized["additionalProperties"] = normalize_schema(schema["additionalProperties"])
    for key in ("allOf", "anyOf", "oneOf"):
        if key in schema:
            normalized[key] = [normalize_schema(item) for item in schema[key]]
    return normalized


def operation_schema_name(operation: dict[str, Any], section: str, code: str | None = None) -> str:
    if section == "request":
        content = operation.get("requestBody", {}).get("content", {}).get("application/json", {})
    else:
        content = operation.get("responses", {}).get(code or "", {}).get("content", {}).get("application/json", {})
    schema = content.get("schema", {})
    if not isinstance(schema, dict):
        return ""
    return schema_ref(schema) or yaml.safe_dump(normalize_schema(schema), sort_keys=True).strip()


def parameter_signature(parameter: dict[str, Any]) -> dict[str, Any]:
    return {
        "in": parameter.get("in"),
        "name": parameter.get("name"),
        "required": bool(parameter.get("required", False)),
        "schema": normalize_schema(parameter.get("schema", {})),
    }


def operation_baseline(operation: dict[str, Any]) -> dict[str, Any]:
    responses: dict[str, Any] = {}
    for code, response in sorted(operation.get("responses", {}).items()):
        if not isinstance(response, dict):
            continue
        responses[str(code)] = {
            "schema": operation_schema_name(operation, "response", str(code)),
        }
    return {
        "operation_id": operation.get("operationId", ""),
        "rbac_scope": operation.get("x-ani-rbac-scope", ""),
        "parameters": [parameter_signature(param) for param in operation.get("parameters", [])],
        "request_body_required": bool(operation.get("requestBody", {}).get("required", False)),
        "request_schema": operation_schema_name(operation, "request"),
        "responses": responses,
    }


def schema_baseline(schema: dict[str, Any]) -> dict[str, Any]:
    properties = schema.get("properties", {})
    return {
        "type": schema.get("type", ""),
        "required": sorted(schema.get("required", [])),
        "properties": {
            name: {
                "signature": normalize_schema(prop),
                "enum": list(prop.get("enum", [])) if isinstance(prop, dict) and "enum" in prop else [],
            }
            for name, prop in sorted(properties.items())
            if isinstance(prop, dict)
        },
    }


def main() -> None:
    spec = yaml.safe_load(SPEC_PATH.read_text(encoding="utf-8"))
    paths = spec.get("paths", {})
    schemas = spec.get("components", {}).get("schemas", {})
    baseline = {
        "batch": "SPEC-COMPAT-A",
        "phase": "Phase 1 / Sprint 4",
        "source": "api/openapi/v1.yaml",
        "generated_at": date.today().isoformat(),
        "compatibility_rules": {
            "allowed": [
                "Add optional request fields.",
                "Add response fields.",
                "Add endpoints.",
                "Add enum values.",
            ],
            "forbidden": [
                "Remove paths or HTTP methods.",
                "Change operationId.",
                "Remove existing parameters.",
                "Remove required request fields.",
                "Remove response fields.",
                "Change existing field schema signatures.",
            ],
        },
        "paths": {
            path: {
                method: operation_baseline(operation)
                for method, operation in sorted(path_item.items())
                if method in {"get", "post", "put", "patch", "delete"} and isinstance(operation, dict)
            }
            for path, path_item in sorted(paths.items())
            if isinstance(path_item, dict)
        },
        "schemas": {
            name: schema_baseline(schema)
            for name, schema in sorted(schemas.items())
            if isinstance(schema, dict)
        },
    }
    BASELINE_PATH.write_text(yaml.safe_dump(baseline, sort_keys=False, allow_unicode=True), encoding="utf-8")
    print(f"generated {BASELINE_PATH.relative_to(ROOT)}")


if __name__ == "__main__":
    main()
