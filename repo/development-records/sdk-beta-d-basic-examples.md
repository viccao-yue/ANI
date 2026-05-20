# SDK-BETA-D — SDK Basic Examples

完成日期：2026-05-20
对应 Sprint：Sprint 4（2026-05-20 提前启动；计划窗口 2026-07-01 ~ 2026-07-15）
验证结果：`make gen-core-sdk`、`make validate-sdk-beta`、`make test`、`make validate-architecture` 通过

## 实现了什么

完成 SDK Beta 加固的第四个切片：四语言 SDK 生成物统一提供 basic example，展示 client 初始化、幂等键注入、cursor 分页参数和 API error helper 的组合用法。

这让 Services 团队不只看到 helper 名称，还能直接参考最小可运行示例接入 Core SDK。

## 关键文件改动

| 文件 | 新增/修改 | 说明 |
|---|---|---|
| `scripts/gen_sdk_alpha.py` | 修改 | 生成 Go/Python/TypeScript/Java basic example |
| `scripts/validate_sdk_alpha.py` | 修改 | 将 basic example 纳入 SDK smoke；有 JDK 时编译运行 Java example |
| `scripts/validate_sdk_beta.py` | 修改 | 校验四语言 examples 都覆盖幂等、分页和错误 helper |
| `sdks/core/**/examples/`、`sdks/services/**/examples/` | 新增 | 四语言 Core/Services SDK 示例 |

## 完工标准达成

- [x] 四语言 SDK 均有 basic example。
- [x] 示例覆盖 `idempotency_key`、cursor 分页和 API error helper。
- [x] `make validate-sdk-beta` 通过。

## 备注

本批次仍不实现完整 HTTP 请求发送；示例聚焦 SDK Beta helper 的组合用法。后续若引入完整 HTTP client，再在 examples 中补充真实请求和 Mock Server 联调示例。
