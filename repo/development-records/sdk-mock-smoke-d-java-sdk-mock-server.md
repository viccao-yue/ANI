# SDK-MOCK-SMOKE-D — Java SDK Mock Server Smoke

完成日期：2026-05-21
对应 Sprint：Sprint 4（2026-05-20 提前启动；计划窗口 2026-07-01 ~ 2026-07-15）
验证结果：`make validate-sdk-mock-smoke`、`make validate-sprint4-closure` 通过

## 实现了什么

完成 Sprint 4 SDK 与 Mock Server 联动烟测的第四个切片：Core Java SDK 生成物提供基于 `java.net.http.HttpClient` 的 `request()` 能力。有 JDK 时自动烟测会启动由 `api/openapi/v1.yaml` 驱动的 Core Mock Server，并用 Java SDK 调用 Core API；无 JDK 时执行 Java source smoke，明确不伪装成真实 Java 联动已运行。

这让 Sprint 4 四语言 SDK 都具备最小 HTTP 调用面，并让 SDK/Mock/API 契约的联动关系被统一门禁覆盖。

## 关键文件改动

| 文件 | 新增/修改 | 说明 |
|---|---|---|
| `scripts/gen_sdk_alpha.py` | 修改 | 生成 Core/Services Java SDK 的 `HttpClient` request 能力、`APIException` 和标准错误解析 |
| `scripts/validate_sdk_mock_smoke.py` | 修改 | 在 Python/TypeScript/Go 烟测基础上增加 Java SDK 调 Mock Server；无 JDK 时执行 source smoke |
| `Makefile` | 修改 | `make validate-sdk-mock-smoke` 覆盖四语言 SDK-Mock 联动 |
| `sdks/core/java/src/main/java/com/kubercloud/ani/core/ApiClient.java` | 生成更新 | Core Java SDK 暴露 `request()` |
| `CURRENT-SPRINT.md`、`CLAUDE.md`、`ANI-06-开发计划.md`、`ANI-DOCS-INDEX.md` | 修改 | 同步 `SDK-MOCK-SMOKE-D` 状态和验收命令 |

## 完工标准达成

- [x] Core Java SDK 可用 `java.net.http.HttpClient` 发起 HTTP 请求，不引入第三方依赖。
- [x] 有 JDK 时，SDK 可调用 Mock Server 的 `/healthz` 成功响应并读取正文。
- [x] 有 JDK 时，SDK 可调用 Mock Server 的分页列表响应并识别 `items`。
- [x] 有 JDK 时，SDK 可将 Mock Server 的标准错误响应转换为 `APIException(APIError)`。
- [x] 无 JDK 时，门禁执行 Java source smoke，并明确保留环境限制说明。

## 备注

本批次只补 Java SDK-Mock 联动能力，不生成完整资源方法；完整多语言资源级 SDK 仍按后续 SDK 生成方案继续推进。
