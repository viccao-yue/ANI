# SDK-MOCK-SMOKE-C — Go SDK Mock Server Smoke

完成日期：2026-05-21
对应 Sprint：Sprint 4（2026-05-20 提前启动；计划窗口 2026-07-01 ~ 2026-07-15）
验证结果：`make validate-sdk-mock-smoke`、`make validate-sprint4-closure` 通过

## 实现了什么

完成 Sprint 4 SDK 与 Mock Server 联动烟测的第三个切片：Core Go SDK 生成物提供基于 `net/http` 的 `Request()` 能力，自动烟测会启动由 `api/openapi/v1.yaml` 驱动的 Core Mock Server，并用 Go SDK 调用 Core API。

这让 SDK-Mock 联动从 Python、TypeScript 继续扩展到 Go，覆盖 Core 团队和后端服务更常用的 SDK 使用方式。

## 关键文件改动

| 文件 | 新增/修改 | 说明 |
|---|---|---|
| `scripts/gen_sdk_alpha.py` | 修改 | 生成 Core/Services Go SDK 的 `net/http` Request 能力和 `APIError` error 实现 |
| `scripts/validate_sdk_mock_smoke.py` | 修改 | 在 Python/TypeScript 烟测基础上增加 Go SDK 调 Mock Server |
| `Makefile` | 修改 | `make validate-sdk-mock-smoke` 覆盖 Python/TypeScript/Go 三条 SDK-Mock 联动 |
| `sdks/core/go/anisdk/client.go` | 生成更新 | Core Go SDK 暴露 `Client.Request()` |
| `CURRENT-SPRINT.md`、`CLAUDE.md`、`ANI-06-开发计划.md`、`ANI-DOCS-INDEX.md` | 修改 | 同步 `SDK-MOCK-SMOKE-C` 状态和验收命令 |

## 完工标准达成

- [x] Core Go SDK 可用 `net/http` 发起 HTTP 请求，不引入新依赖。
- [x] SDK 调用 Mock Server 的 `/healthz` 成功响应并解析 JSON。
- [x] SDK 调用 Mock Server 的分页列表响应并识别 `items`。
- [x] SDK 将 Mock Server 的标准错误响应转换为 `APIError`。
- [x] `make validate-sdk-mock-smoke` 通过。

## 备注

本批次只补 Go SDK-Mock 联动能力，不生成完整资源方法；完整多语言资源级 SDK 仍按后续 SDK 生成方案继续推进。
