# DOC-API-A — Static API Docs

完成日期：2026-05-20
对应 Sprint：Sprint 4（2026-05-20 提前启动；计划窗口 2026-07-01 ~ 2026-07-15）
验证结果：`make gen-api-docs`、`make validate-doc-api`、`make test`、`make validate-architecture` 通过

## 实现了什么

完成 DOC-API-A 的首个切片：新增由 Core/Services API 契约生成的静态 API 文档，输出到 `docs/api/`。

文档入口区分 Core API 和 Services API，Core 文档只展示基础设施路径，Services 文档展示 models、inference-services、knowledge-bases 等业务路径。

## 关键文件改动

| 文件 | 新增/修改 | 说明 |
|---|---|---|
| `scripts/generate_api_docs.py` | 新增 | 从 Core/Services API 契约生成 `docs/api/index.html`、`core.html`、`services.html` |
| `scripts/validate_api_docs_contract.py` | 新增 | 校验文档覆盖 operation/schema，并确保 Core 文档不混入 Services 业务路径 |
| `docs/api/*.html` | 新增 | 生成后的静态 API 文档 |
| `Makefile` | 修改 | 新增 `make gen-api-docs` 和 `make validate-doc-api` |
| `CURRENT-SPRINT.md` | 修改 | 将 DOC-API-A 首个切片和验收命令纳入 Sprint 4 当前状态 |

## 完工标准达成

- [x] API 文档由 `api/openapi/v1.yaml` 和 `api/openapi/services/v1.yaml` 生成，不手写路径清单。
- [x] Core/Services 文档分层展示。
- [x] 每个 API operation 和 schema 都能在对应文档中找到。
- [x] `make validate-doc-api` 通过。

## 备注

本批次先提供离线静态 HTML 文档，避免依赖外部 CDN 或额外 Node 安装。后续若需要 Swagger UI / Redoc 外观，可在保留 `make validate-doc-api` 门禁的基础上替换渲染层。
