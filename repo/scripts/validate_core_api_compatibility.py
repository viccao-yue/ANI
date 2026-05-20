#!/usr/bin/env python3
"""Validate Core API v1 against the Sprint 4 compatibility baseline."""

from __future__ import annotations

from pathlib import Path
from typing import Any

import yaml


ROOT = Path(__file__).resolve().parents[1]
SPEC_PATH = ROOT / "api/openapi/v1.yaml"
BASELINE_PATH = ROOT / "api/core-v1-compatibility-baseline.yaml"


def fail(message: str) -> None:
    raise SystemExit(f"core api compatibility invalid: {message}")


def load_yaml(path: Path) -> dict[str, Any]:
    data = yaml.safe_load(path.read_text(encoding="utf-8"))
    if not isinstance(data, dict):
        fail(f"{path.relative_to(ROOT)} must parse to an object")
    return data


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


def schema_ref(schema: dict[str, Any]) -> str:
    ref = schema.get("$ref")
    return ref.rsplit("/", 1)[-1] if isinstance(ref, str) else ""


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


def validate_operations(spec: dict[str, Any], baseline: dict[str, Any]) -> None:
    paths = spec.get("paths", {})
    for path, baseline_path_item in baseline.get("paths", {}).items():
        current_path_item = paths.get(path)
        if not isinstance(current_path_item, dict):
            fail(f"missing protected path {path}")
        for method, baseline_operation in baseline_path_item.items():
            current_operation = current_path_item.get(method)
            if not isinstance(current_operation, dict):
                fail(f"missing protected method {method.upper()} {path}")
            if current_operation.get("operationId", "") != baseline_operation.get("operation_id", ""):
                fail(f"{method.upper()} {path} changed operationId")
            if current_operation.get("x-ani-rbac-scope", "") != baseline_operation.get("rbac_scope", ""):
                fail(f"{method.upper()} {path} changed x-ani-rbac-scope")

            current_parameters = {
                (parameter.get("in"), parameter.get("name")): parameter_signature(parameter)
                for parameter in current_operation.get("parameters", [])
                if isinstance(parameter, dict)
            }
            for baseline_parameter in baseline_operation.get("parameters", []):
                key = (baseline_parameter.get("in"), baseline_parameter.get("name"))
                if key not in current_parameters:
                    fail(f"{method.upper()} {path} removed parameter {key[0]}:{key[1]}")
                if current_parameters[key] != baseline_parameter:
                    fail(f"{method.upper()} {path} changed parameter {key[0]}:{key[1]}")

            if baseline_operation.get("request_body_required") and not current_operation.get("requestBody", {}).get("required", False):
                fail(f"{method.upper()} {path} made required request body optional")
            if operation_schema_name(current_operation, "request") != baseline_operation.get("request_schema", ""):
                fail(f"{method.upper()} {path} changed request schema")
            for code, baseline_response in baseline_operation.get("responses", {}).items():
                if code not in current_operation.get("responses", {}):
                    fail(f"{method.upper()} {path} removed response {code}")
                if operation_schema_name(current_operation, "response", code) != baseline_response.get("schema", ""):
                    fail(f"{method.upper()} {path} changed response {code} schema")


def request_schema_names(baseline: dict[str, Any]) -> set[str]:
    names: set[str] = set()
    for path_item in baseline.get("paths", {}).values():
        if not isinstance(path_item, dict):
            continue
        for operation in path_item.values():
            if not isinstance(operation, dict):
                continue
            request_schema = operation.get("request_schema", "")
            if isinstance(request_schema, str) and request_schema and "\n" not in request_schema and request_schema != "{}":
                names.add(request_schema)
    return names


def validate_schemas(spec: dict[str, Any], baseline: dict[str, Any]) -> None:
    schemas = spec.get("components", {}).get("schemas", {})
    protected_request_schemas = request_schema_names(baseline)
    for schema_name, baseline_schema in baseline.get("schemas", {}).items():
        current_schema = schemas.get(schema_name)
        if not isinstance(current_schema, dict):
            fail(f"missing protected schema {schema_name}")
        if current_schema.get("type", "") != baseline_schema.get("type", ""):
            fail(f"{schema_name} changed schema type")
        current_required = set(current_schema.get("required", []))
        baseline_required = set(baseline_schema.get("required", []))
        for field in baseline_schema.get("required", []):
            if field not in current_required:
                fail(f"{schema_name} removed required field {field}")
        if schema_name in protected_request_schemas:
            added_required = sorted(current_required - baseline_required)
            if added_required:
                fail(f"{schema_name} added required request fields: {added_required}")
        current_properties = current_schema.get("properties", {})
        for property_name, baseline_property in baseline_schema.get("properties", {}).items():
            current_property = current_properties.get(property_name)
            if not isinstance(current_property, dict):
                fail(f"{schema_name} removed property {property_name}")
            if normalize_schema(current_property) != baseline_property.get("signature", {}):
                fail(f"{schema_name}.{property_name} changed schema signature")
            baseline_enum = set(baseline_property.get("enum", []))
            current_enum = set(current_property.get("enum", []))
            if baseline_enum and not baseline_enum.issubset(current_enum):
                fail(f"{schema_name}.{property_name} removed enum values")


def main() -> None:
    spec = load_yaml(SPEC_PATH)
    baseline = load_yaml(BASELINE_PATH)
    if baseline.get("batch") != "SPEC-COMPAT-A":
        fail("api/core-v1-compatibility-baseline.yaml must be owned by SPEC-COMPAT-A")
    if baseline.get("source") != "api/openapi/v1.yaml":
        fail("compatibility baseline source must be api/openapi/v1.yaml")
    validate_operations(spec, baseline)
    validate_schemas(spec, baseline)
    print("core api compatibility valid")


if __name__ == "__main__":
    main()
