# SPEC-COMPAT-A — Core API v1 兼容性基线

完成日期：2026-05-21
对应 Sprint：Sprint 4（2026-05-20 提前启动；计划窗口 2026-07-01 ~ 07-15）
验证结果：make validate-core-api-compatibility passed；make validate-sprint4-closure passed；make test passed；make validate-architecture passed

## 实现了什么

建立 Core API v1 的兼容性基线，保护已交付的 path、HTTP method、operationId、参数、响应和 schema 字段。后续允许新增可选字段、响应字段、端点和枚举值，但误删或修改既有契约会被校验拦截。

## 关键文件改动

| 文件 | 新增/修改 | 说明 |
|---|---|---|
| `api/core-v1-compatibility-baseline.yaml` | 新增 | Core API v1 兼容性基线 |
| `scripts/generate_core_api_compatibility_baseline.py` | 新增 | 有意更新 Beta 基线时使用的生成脚本 |
| `scripts/validate_core_api_compatibility.py` | 新增 | 日常门禁校验脚本 |
| `Makefile` | 修改 | 新增 `gen-core-api-compat-baseline` 和 `validate-core-api-compatibility` |
| `scripts/validate_sprint4_closure.py` | 修改 | 将 `SPEC-COMPAT-A` 纳入 Sprint 4 关联性闭环 |
| `CLAUDE.md` / `ANI-DOCS-INDEX.md` / `ANI-06-开发计划.md` / `CURRENT-SPRINT.md` | 修改 | 同步当前 Sprint 状态和文档入口 |

## 完工标准达成

- [x] `make validate-core-api-compatibility` 通过。
- [x] `make validate-sprint4-closure` 已纳入并通过。
- [x] `make test` 全通。
- [x] `make validate-architecture` 通过。
- [x] 文档入口、当前 Sprint 和开发记录索引已同步。

## 备注

兼容性基线不是自动随每次校验重建的文件；只有确认 API 契约允许演进时，才运行 `make gen-core-api-compat-baseline` 更新基线。
