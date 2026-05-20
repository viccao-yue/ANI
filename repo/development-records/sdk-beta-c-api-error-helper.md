# SDK-BETA-C — SDK API Error Helper

完成日期：2026-05-20
对应 Sprint：Sprint 4（2026-05-20 提前启动；计划窗口 2026-07-01 ~ 2026-07-15）
验证结果：`make validate-sdk-beta`、`make validate-sdk-alpha` 通过

## 实现了什么

完成 SDK Beta 加固的第三个切片：四语言 SDK 生成物统一提供 API error 对象、错误码清单和错误码判断 helper。

这让 Services 团队和后续 CLI/控制台调用方可以按 API 契约统一识别 `{ code, message, request_id, details }` 错误响应，不需要各语言各自猜测错误结构。

## 关键文件改动

| 文件 | 新增/修改 | 说明 |
|---|---|---|
| `scripts/gen_sdk_alpha.py` | 修改 | 从 API 契约提取标准错误码，并生成四语言 error helper |
| `scripts/validate_sdk_alpha.py` | 修改 | 将 error helper、`errorCodes` 和四语言 smoke 纳入 SDK Alpha 校验 |
| `scripts/validate_sdk_beta.py` | 修改 | 校验 Core SDK errorCodes、四语言 helper 与文档闭环 |
| `sdks/core/**`、`sdks/services/**` | 修改 | 重新生成 Go/Python/TypeScript/Java SDK error helper 和 smoke |

## 四语言 helper 命名

| SDK | Helper |
|---|---|
| Go | `APIError`、`ErrorCodes`、`IsAPIErrorCode` |
| Python | `APIError`、`ERROR_CODES`、`is_api_error_code` |
| TypeScript | `APIError`、`errorCodes`、`isAPIErrorCode` |
| Java | `APIError`、`ERROR_CODES`、`isAPIErrorCode` |

## 完工标准达成

- [x] SDK metadata 标出 API 契约声明的标准错误码。
- [x] 四语言 smoke 覆盖统一错误对象和错误码判断 helper。
- [x] `make validate-sdk-beta` 通过。

## 备注

本批次只固化错误响应结构和错误码辅助能力，不实现完整 HTTP 请求发送、异常抛出策略或自动重试；这些仍属于后续 SDK Beta/CLI 切片。
