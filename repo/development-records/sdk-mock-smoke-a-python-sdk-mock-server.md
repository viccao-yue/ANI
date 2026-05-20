# SDK-MOCK-SMOKE-A — Python SDK Mock Server Smoke

完成日期：2026-05-20
对应 Sprint：Sprint 4（2026-05-20 提前启动；计划窗口 2026-07-01 ~ 2026-07-15）
验证结果：`make validate-sdk-mock-smoke`、`make validate-sprint4-closure` 通过

## 实现了什么

完成 Sprint 4 SDK 与 Mock Server 联动烟测的首个切片：Core Python SDK 生成物提供标准库 HTTP `request()` 能力，自动烟测会启动由 `api/openapi/v1.yaml` 驱动的 Core Mock Server，并用 SDK 调用 Core API。

这让“Services 团队能用 SDK 调 Mock Server”从文档描述变成可执行门禁，避免 Mock Server、SDK helper 和 API 契约三者各自通过但联动断开。

## 关键文件改动

| 文件 | 新增/修改 | 说明 |
|---|---|---|
| `scripts/gen_sdk_alpha.py` | 修改 | 生成 Core/Services Python SDK 的标准库 HTTP `request()` 能力 |
| `scripts/validate_sdk_mock_smoke.py` | 新增 | 启动 Core Mock Server，并用 Core Python SDK 调用 `/healthz`、分页列表和 404 错误响应 |
| `Makefile` | 修改 | 新增 `make validate-sdk-mock-smoke`，并纳入 `make validate-sprint4-closure` |
| `sdks/core/python/kubercloud_ani_core/client.py` | 生成更新 | Core Python SDK 暴露 `Client.request()` |
| `CURRENT-SPRINT.md`、`CLAUDE.md`、`ANI-06-开发计划.md`、`ANI-DOCS-INDEX.md` | 修改 | 同步 `SDK-MOCK-SMOKE-A` 状态和验收命令 |

## 完工标准达成

- [x] Core Python SDK 可用标准库发起 HTTP 请求，不引入新依赖。
- [x] SDK 调用 Mock Server 的 `/healthz` 成功响应并解析 JSON。
- [x] SDK 调用 Mock Server 的分页列表响应并识别 `items`。
- [x] SDK 将 Mock Server 的标准错误响应转换为 `APIError`。
- [x] `make validate-sdk-mock-smoke` 通过。

## 备注

本批次只补最小 SDK-Mock 联动能力，不生成完整资源方法；完整多语言资源级 SDK 仍按后续 SDK 生成方案继续推进。
