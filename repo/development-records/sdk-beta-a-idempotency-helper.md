# SDK-BETA-A — SDK Idempotency Helper

完成日期：2026-05-20
对应 Sprint：Sprint 4（2026-05-20 提前启动；计划窗口 2026-07-01 ~ 2026-07-15）
验证结果：`make validate-sdk-beta`、`make validate-sdk-alpha` 通过

## 实现了什么

完成 SDK Beta 加固的首个切片：四语言 SDK 生成物统一提供 `idempotency_key` 辅助能力，并在 SDK metadata 中列出 Core API 中必须携带幂等键的操作。

这让 Services 团队在调用 Core 创建/生命周期 API 时，可以由 SDK 生成或注入幂等键，减少重试时重复创建资源的风险。

## 关键文件改动

| 文件 | 新增/修改 | 说明 |
|---|---|---|
| `scripts/gen_sdk_alpha.py` | 修改 | 从 API 契约识别 `idempotency_key` 操作，并生成四语言 helper |
| `scripts/validate_sdk_alpha.py` | 修改 | 将 helper 与 metadata 纳入 SDK smoke |
| `scripts/validate_sdk_beta.py` | 新增 | 校验 SDK Beta 幂等 helper、metadata 与文档闭环 |
| `sdks/core/**`、`sdks/services/**` | 修改 | 重新生成 Go/Python/TypeScript/Java SDK helper 和 smoke |
| `Makefile` | 修改 | 新增 `make validate-sdk-beta` |

## 四语言 helper 命名

| SDK | Helper |
|---|---|
| Go | `NewIdempotencyKey`、`WithIdempotencyKey`、`IdempotencyOperations` |
| Python | `new_idempotency_key`、`with_idempotency_key`、`IDEMPOTENCY_OPERATIONS` |
| TypeScript | `newIdempotencyKey`、`withIdempotencyKey`、`idempotencyOperations` |
| Java | `newIdempotencyKey`、`withIdempotencyKey`、`IDEMPOTENCY_OPERATIONS` |

## 完工标准达成

- [x] Core SDK metadata 标出所有需要 `idempotency_key` 的 Core 操作。
- [x] Services SDK 不声明 Core 幂等操作清单，但保留通用 helper 方便 Services 层自用。
- [x] 四语言 smoke 覆盖 helper 生成和 payload 注入。
- [x] `make validate-sdk-beta` 通过。

## 备注

本批次不引入完整 HTTP client、重试策略或分页迭代器；这些属于后续 SDK Beta 加固切片。
