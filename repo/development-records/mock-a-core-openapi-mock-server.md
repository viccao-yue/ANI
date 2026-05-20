# MOCK-A — Core OpenAPI Mock Server

完成日期：2026-05-20
对应 Sprint：Sprint 4（2026-05-20 提前启动；计划窗口 2026-07-01 ~ 2026-07-15）
验证结果：`make validate-mock-a`、`make test`、`make validate-architecture` 通过

## 实现了什么

完成 MOCK-A 的首个切片：新增由 Core API 契约驱动的本地 Mock Server，默认监听 `http://127.0.0.1:4010/api/v1`。

它用于 Services 团队和 SDK 使用方在无真实 Gateway 时验证 Core API 路径、成功响应结构和统一错误结构；不提供 models、inference-services、knowledge-bases 等 Services 业务 mock。

## 关键文件改动

| 文件 | 新增/修改 | 说明 |
|---|---|---|
| `scripts/serve_core_mock.py` | 新增 | 从 `api/openapi/v1.yaml` 读取路径和 schema，启动本地 Core Mock Server |
| `scripts/validate_mock_server_contract.py` | 新增 | 校验 Mock Server 覆盖所有 Core API operation，且 mock body 可 JSON 序列化 |
| `Makefile` | 修改 | 新增 `make validate-mock-a` 和 `make dev-core-mock` |
| `CURRENT-SPRINT.md` | 修改 | 将 MOCK-A 首个切片和验收命令纳入 Sprint 4 当前状态 |

## 完工标准达成

- [x] Mock Server 基于 `api/openapi/v1.yaml`，不手写 Core 路径清单。
- [x] 全量 Core API operation 均能生成 mock route。
- [x] 统一错误结构 `ErrorResponse` 能生成标准 mock body。
- [x] `make validate-mock-a` 通过。

## 备注

本批次使用 Python 标准库实现本地 mock 入口，避免引入额外安装依赖。后续若团队决定使用 Prism，可以保留 `make validate-mock-a` 作为契约覆盖门禁，再将 `make dev-core-mock` 切换到 Prism 启动方式。
