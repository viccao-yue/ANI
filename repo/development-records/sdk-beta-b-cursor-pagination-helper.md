# SDK-BETA-B — SDK Cursor Pagination Helper

完成日期：2026-05-20
对应 Sprint：Sprint 4（2026-05-20 提前启动；计划窗口 2026-07-01 ~ 2026-07-15）
验证结果：`make validate-sdk-beta`、`make validate-sdk-alpha` 通过

## 实现了什么

完成 SDK Beta 加固的第二个切片：四语言 SDK 生成物统一提供 cursor 分页辅助能力，并在 SDK metadata 中列出 Core API 中支持 `limit/cursor` 的列表操作。

这让调用方可以稳定构造列表请求参数，并能从 metadata 中知道哪些操作遵循 `{ items, total, next_cursor }` 这一分页约定。

## 关键文件改动

| 文件 | 新增/修改 | 说明 |
|---|---|---|
| `scripts/gen_sdk_alpha.py` | 修改 | 从 API 契约识别 `limit/cursor` 列表操作，并生成四语言分页 helper |
| `scripts/validate_sdk_alpha.py` | 修改 | 将分页 helper 与 `cursorPaginationOperations` 纳入 SDK smoke |
| `scripts/validate_sdk_beta.py` | 修改 | 校验 Core SDK 分页操作清单、四语言 helper 与文档闭环 |
| `sdks/core/**`、`sdks/services/**` | 修改 | 重新生成 Go/Python/TypeScript/Java SDK helper 和 smoke |

## 四语言 helper 命名

| SDK | Helper |
|---|---|
| Go | `CursorParams`、`CursorPaginationOperations` |
| Python | `cursor_params`、`CURSOR_PAGINATION_OPERATIONS` |
| TypeScript | `cursorParams`、`cursorPaginationOperations` |
| Java | `cursorParams`、`CURSOR_PAGINATION_OPERATIONS` |

## 完工标准达成

- [x] Core SDK metadata 标出所有支持 `limit/cursor` 的 Core 列表操作。
- [x] 四语言 smoke 覆盖分页参数 helper。
- [x] `make validate-sdk-beta` 通过。

## 备注

本批次只提供分页参数和操作清单，不实现完整 HTTP 请求发送、自动翻页迭代器或重试策略；这些仍属于后续 SDK Beta 切片。
