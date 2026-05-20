# SPEC-CORE-BETA-A — Core API Beta Readiness Matrix

完成日期：2026-05-20
对应 Sprint：Sprint 4（2026-05-20 提前启动；计划窗口 2026-07-01 ~ 2026-07-15）
验证结果：`make validate-core-beta` 通过

## 实现了什么

完成 Core API Beta 准备的首个可验证切片：新增机器可读 Beta 准备矩阵，覆盖 Sprint 1-3 的 Core P0 path/schema、分页、幂等、状态机、dev/local profile 标记、RBAC scope 和 Core/Services 分层边界。

同时补齐网络、存储、向量创建接口的 `idempotency_key` 契约与本地 profile 重试语义，避免 Services 团队在重试创建资源时得到重复资源。

## 关键文件改动

| 文件 | 新增/修改 | 说明 |
|---|---|---|
| `api/core-beta-readiness.yaml` | 新增 | Core API Beta 准备矩阵，列出允许新增、禁止破坏性变更和 P0 资源守卫 |
| `scripts/validate_core_beta_contract.py` | 新增 | 校验矩阵、Core API 契约、SDK metadata、Gateway/Core P0 stub 边界和文档引用 |
| `api/openapi/v1.yaml` | 修改 | 网络/存储/向量创建请求补齐 `idempotency_key`；列表响应补齐 `next_cursor`；列表路径补齐 `limit/cursor` |
| `pkg/ports/*resources.go`、`pkg/ports/vector_store.go` | 修改 | Create request port 增加 `IdempotencyKey` |
| `pkg/adapters/runtime/*_service.go` | 修改 | local profile 按 `(tenant_id, idempotency_key)` 返回同一创建结果 |
| `services/ani-gateway/internal/router/*resources.go` | 修改 | Gateway request/response 与 API 契约对齐 |
| `Makefile` | 修改 | 新增 `make validate-core-beta` |

## 完工标准达成

- [x] Core API 继续排除 Services 业务路径。
- [x] Core P0 创建接口请求体统一包含 `idempotency_key`。
- [x] Core P0 列表接口统一暴露 `limit/cursor` 并返回 `next_cursor`。
- [x] 资源响应 schema 保留 `tenant_id/state/created_at/updated_at/dev_profile`。
- [x] Core P0 Gateway/local profile 不包含无 owner/date 的 `NOT_IMPLEMENTED`。
- [x] `make validate-core-beta` 通过。

## 备注

本批次是 `SPEC-CORE-BETA` 的首个切片，不代表 Sprint 4 全部完成。后续仍需继续 SDK Beta 加固、Mock Server 和 API 文档生成。
