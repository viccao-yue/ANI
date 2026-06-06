# CORE-HISTORICAL-DOC-MARKER-COMPAT-A — Core historical doc marker compatibility

完成日期：2026-06-06
对应 Sprint：Sprint 11（Core Real Deployment Validation / historical gate maintenance）
验证结果：`make validate-sprint11-real-deployment`、`make validate-doc-entrypoints`、`make validate-architecture`、`make test`、`git diff --check`

## 背景

Sprint 11 聚合门禁依赖 Sprint 10 release-prep；Sprint 10 又串联 Sprint 9 和 Sprint 8 历史门禁。Sprint 11 入口文档已收敛为当前真实部署完成状态，只保留 Sprint 8/9/10 的历史门禁和已完成归档表达，不再保留这些 Sprint 当时作为“当前态”的精确短语。

因此历史文档一致性 validator 不能要求入口文档停留在旧 Sprint 当前态；它们应接受当前入口文档中的历史归档表达，同时继续拒绝 stale current marker、缺失 Makefile target 和缺失 development records 的情况。

## 关键变更

| 文件 | 说明 |
|---|---|
| `scripts/validate_core_doc_consistency.py` | Sprint 8 文档 marker 兼容当前 Sprint 11 入口文档中的 `Sprint 6-10` 汇总与历史门禁表达 |
| `scripts/validate_sprint9_core_doc_consistency.py` | Sprint 9 文档 marker 兼容当前历史 RC readiness / records 归档表达 |
| `scripts/validate_sprint10_core_doc_consistency.py` | Sprint 10 文档 marker 兼容当前 release-prep 历史门禁表达，并继续要求不是实际 v1.0.0 发布 |
| `scripts/validate_core_doc_consistency_test.py` | 新增 Sprint 8 后续 Sprint 历史 marker 兼容测试 |
| `scripts/validate_sprint9_core_doc_consistency_test.py` | 新增 Sprint 9 后续 Sprint 历史 marker 兼容测试 |
| `scripts/validate_sprint10_core_doc_consistency_test.py` | 新增 Sprint 10 后续 Sprint 历史 marker 兼容测试 |

## 边界

- 本批次只维护 ANI Core 历史文档一致性门禁。
- 不新增 Core OpenAPI path，不修改 Services、RAG、Console、BOSS、model-service、kb-service、ai、operators 或 frontends。
- 不改变 Sprint 11 真实部署完成结论；也不代表实际 v1.0.0 发布或完整 production ready。
