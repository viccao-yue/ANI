# SDK-MOCK-SMOKE-B — TypeScript SDK Mock Server Smoke

完成日期：2026-05-21
对应 Sprint：Sprint 4（2026-05-20 提前启动；计划窗口 2026-07-01 ~ 2026-07-15）
验证结果：`make validate-sdk-mock-smoke`、`make validate-sprint4-closure` 通过

## 实现了什么

完成 Sprint 4 SDK 与 Mock Server 联动烟测的第二个切片：Core TypeScript SDK 生成物提供基于 `fetch` 的 `request()` 能力，自动烟测会启动由 `api/openapi/v1.yaml` 驱动的 Core Mock Server，并用 TypeScript SDK 调用 Core API。

这让 SDK-Mock 联动不只停留在 Python，而是覆盖 Services/Console 更容易使用的 TypeScript 侧，继续保持“多语言 SDK 统一可用”的方向。

## 关键文件改动

| 文件 | 新增/修改 | 说明 |
|---|---|---|
| `scripts/gen_sdk_alpha.py` | 修改 | 生成 Core/Services TypeScript SDK 的 `fetch` request 能力 |
| `scripts/validate_sdk_mock_smoke.py` | 修改 | 在原 Python 烟测基础上增加 TypeScript SDK 调 Mock Server |
| `Makefile` | 修改 | `make validate-sdk-mock-smoke` 覆盖 Python/TypeScript 两条 SDK-Mock 联动 |
| `sdks/core/typescript/src/index.mjs`、`sdks/core/typescript/src/index.ts` | 生成更新 | Core TypeScript SDK 暴露 `Client.request()` |
| `CURRENT-SPRINT.md`、`CLAUDE.md`、`ANI-06-开发计划.md`、`ANI-DOCS-INDEX.md` | 修改 | 同步 `SDK-MOCK-SMOKE-B` 状态和验收命令 |

## 完工标准达成

- [x] Core TypeScript SDK 可用 `fetch` 发起 HTTP 请求，不引入新依赖。
- [x] SDK 调用 Mock Server 的 `/healthz` 成功响应并解析 JSON。
- [x] SDK 调用 Mock Server 的分页列表响应并识别 `items`。
- [x] SDK 将 Mock Server 的标准错误响应转换为 API error 对象。
- [x] `make validate-sdk-mock-smoke` 通过。

## 备注

本批次只补 TypeScript SDK-Mock 联动能力，不生成完整资源方法；完整多语言资源级 SDK 仍按后续 SDK 生成方案继续推进。
