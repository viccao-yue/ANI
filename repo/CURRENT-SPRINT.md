# ANI · 当前冲刺上手指南

> **新开发者（人类或 AI 工具）的第一个入口文件。**
> 读完本文件，5 分钟内明确：当前做什么、怎么开始、怎么验证完成。
> 已完成批次只查 `repo/development-records/README.md`，不要把历史细节塞回本文。

---

## 当前冲刺

| 字段 | 值 |
|---|---|
| **冲刺编号** | Sprint 4（收尾） |
| **时间范围** | 2026-05-20 提前启动；计划窗口 2026-07-01 → 2026-07-15 |
| **主题** | API Beta 准备 + 四语言 SDK 加固 + Mock Server |
| **当前状态** | ✅ 开发与验收完成，待提交 GitHub；提交完成后再切换下一 Sprint |
| **核心批次** | SPEC-SPLIT-A + SPEC-CORE-BETA + SPEC-COMPAT-A + SDK-GO/PY/TS/JAVA + MOCK-A + DOC-API-A |
| **前置验证** | Sprint 3 已完成：网络/存储/向量/Workload Identity/SDK Alpha/Core Dev Profile Ready |

---

## 本冲刺目标

Sprint 4 的目标是在 Core Dev Profile Ready 之后，进入 API Beta 准备和 SDK/Mock Server 加固。

1. 完成 Core/Services API 分层收口，Core API 不承载 Services 业务路径。
2. 审查 Sprint 1-3 所有 Core P0 路径的 schema、分页、幂等、错误码、状态机和 RBAC scope。
3. 加固 Go/Python/TypeScript/Java SDK，确保 Services 团队可稳定使用。
4. 准备 Mock Server 和 API 文档生成，支持 Services 团队并行开发。

> 当前 Sprint 4 的功能开发、文档闭环和自动化验收已完成；仓库仍处于“待提交 GitHub”的流程状态，因此本文暂不切换到 Sprint 5。

---

## P0：SPEC-SPLIT-A

**状态：✅ 已完成，下一步进入 SPEC-CORE-BETA**

**主题：Core/Services API 分层收口**

### 最小实现切片

1. `/models`、`/inference-services`、`/knowledge-bases` 只维护在 `api/openapi/services/v1.yaml`。
2. Core API `api/openapi/v1.yaml` 不再包含 Services 业务路径或 Services tags。
3. Gateway 将 Services 过渡 stub 挂到 `/api/v1/svc/*`，不再挂到 Core `/api/v1/*`。
4. SDK Alpha 生成不再依赖 Core 侧排除列表；Core/Services SDK metadata 按各自 API 契约自然生成。
5. 新增合同守卫，避免 Services 业务路径重新混回 Core API、Core SDK 或 Core Gateway group。

### 验收命令

```bash
make gen-core-sdk
make validate-spec-split
make validate-sdk-alpha
python scripts/validate_yaml.py api/openapi/v1.yaml api/openapi/services/v1.yaml
make test
make validate-architecture
git diff --check
```

### 本批次完成内容

1. Core API 移除 `/models`、`/inference-services`、`/knowledge-bases` 业务路径和过渡 tags。
2. Services API 使用 `https://{host}/api/v1/svc` 作为 base path，并承载 models、inference-services、knowledge-bases path/schema。
3. Gateway Services 过渡 stub 改挂 `/api/v1/svc/*`，RBAC resource 推导跳过 `svc` 前缀。
4. SDK Alpha 生成和校验不再依赖 Core SDK 排除列表；Core/Services SDK metadata 按各自契约自然分层。
5. 新增 `make validate-spec-split` 合同守卫。

---

## P0：SPEC-CORE-BETA

**状态：✅ 开发与验收完成，待提交 GitHub（Beta 准备矩阵、Core API v1 兼容性基线、SDK/Mock/API 文档加固、Python/TypeScript/Go/Java SDK-Mock 联动烟测和提交前审查均已通过）**

**主题：Core API Beta 准备**

### 最小实现切片

1. ✅ 审查 Sprint 1-3 所有 Core P0 path/schema，首轮补齐分页、幂等、状态机、dev_profile 和 RBAC scope 守卫。
2. ✅ 输出机器可读 Beta 检查矩阵 `api/core-beta-readiness.yaml`，明确 allowed additive changes 和 forbidden breaking changes。
3. ✅ 新增 `make validate-core-beta`，确保 Services P0 依赖路径没有无 owner/date 的 stub、mock success 或 `NOT_IMPLEMENTED`。
4. ✅ `SPEC-COMPAT-A` 首个切片完成：新增 `api/core-v1-compatibility-baseline.yaml` 和 `make validate-core-api-compatibility`，保护 Core API path/method/operationId/参数/响应/schema 字段不被误破坏。
5. ✅ `SDK-BETA-A` 首个切片完成：四语言 SDK 提供 `idempotency_key` helper，并通过 `make validate-sdk-beta` 固化。
6. ✅ `SDK-BETA-B` 首个切片完成：四语言 SDK 提供 cursor 分页 helper，并通过 `make validate-sdk-beta` 固化。
7. ✅ `SDK-BETA-C` 首个切片完成：四语言 SDK 提供统一 API error helper，并通过 `make validate-sdk-beta` 固化。
8. ✅ `SDK-BETA-D` 首个切片完成：四语言 SDK 提供可执行 basic example，覆盖 client 初始化、幂等、分页和错误 helper。
9. ✅ `SDK-MOCK-SMOKE-A` 首个切片完成：Core Python SDK 提供标准库 HTTP `request()` 能力，并通过 `make validate-sdk-mock-smoke` 调用本地 Core Mock Server。
10. ✅ `SDK-MOCK-SMOKE-B` 首个切片完成：Core TypeScript SDK 提供基于 `fetch` 的 `request()` 能力，并通过 `make validate-sdk-mock-smoke` 调用同一本地 Core Mock Server。
11. ✅ `SDK-MOCK-SMOKE-C` 首个切片完成：Core Go SDK 提供基于 `net/http` 的 `Request()` 能力，并通过 `make validate-sdk-mock-smoke` 调用同一本地 Core Mock Server。
12. ✅ `SDK-MOCK-SMOKE-D` 首个切片完成：Core Java SDK 提供基于 `java.net.http.HttpClient` 的 `request()` 能力；有 JDK 时真实调用 Mock Server，无 JDK 时执行 source smoke。
13. ✅ `MOCK-A` 首个切片完成：Core Mock Server 由 `api/openapi/v1.yaml` 驱动，默认监听 `http://127.0.0.1:4010/api/v1`，并通过 `make validate-mock-a` 固化覆盖。
14. ✅ `DOC-API-A` 首个切片完成：Core/Services 静态 API 文档由 API 契约生成到 `docs/api/`，并通过 `make validate-doc-api` 固化覆盖。
15. ✅ `SPRINT4-CLOSURE-A` 首个切片完成：新增 `make validate-sprint4-closure`，统一校验 Sprint 4 API/SDK/Mock/Docs/Records 关联性。
16. ✅ 提交前人工审查和 GitHub 提交准备已通过：`make gen-core-sdk`、`make validate-core-api-compatibility`、`make validate-sdk-mock-smoke`、`make validate-sprint4-closure`、`make test`、`make validate-architecture`、YAML 校验和 `git diff --check` 均通过。
17. 下一步流程动作：提交并推送 GitHub；提交完成后再更新本文到下一 Sprint。

### 本批次完成内容

1. 网络、存储、向量创建请求补齐 `idempotency_key`，local profile 按 `(tenant_id, idempotency_key)` 返回同一创建结果。
2. 网络、存储、向量列表 API 补齐 `limit/cursor` 参数和 `next_cursor` 响应字段。
3. `api/core-beta-readiness.yaml` 记录 Core P0 Beta 准备矩阵、允许新增项和禁止破坏性变更项。
4. `api/core-v1-compatibility-baseline.yaml` 记录 Core API v1 兼容性基线，保护已交付 path/method/operationId/参数/响应/schema 字段。
5. `scripts/validate_core_beta_contract.py`、`scripts/validate_core_api_compatibility.py`、`make validate-core-beta` 和 `make validate-core-api-compatibility` 将矩阵、兼容性基线、OpenAPI、SDK metadata、Gateway/Core P0 stub 边界和文档引用纳入自动校验。

### 验收命令

```bash
make gen-core-sdk
make validate-core-beta
make validate-core-api-compatibility
make validate-sdk-beta
make validate-sdk-mock-smoke
make validate-mock-a
make validate-doc-api
make validate-sprint4-closure
make validate-spec-split
make validate-sdk-alpha
python scripts/validate_yaml.py api/openapi/v1.yaml api/openapi/services/v1.yaml api/core-beta-readiness.yaml api/core-v1-compatibility-baseline.yaml
make test
make validate-architecture
git diff --check
```

---

## 本冲刺不做

- 不进入 Phase 2 延期项。
- 不把 Services 业务实现写进 Core。
- 不在 Core API 里重新引入 models、inference-services、knowledge-bases 等 Services 业务路径。
- 不在没有 API 契约和测试的情况下直接生成实现代码。
- 不在 Sprint 4 未提交 GitHub 前直接把当前执行入口切换到 Sprint 5。

---

## 代码结构 10 分钟导航

```text
必读（按顺序）：
  1. CLAUDE.md
  2. ANI-DOCS-INDEX.md
  3. ANI-06-开发计划.md 的 Section 零和 Sprint 4
  4. api/openapi/v1.yaml
  5. api/openapi/services/v1.yaml
  6. scripts/validate_spec_split_contract.py
  7. api/core-beta-readiness.yaml
  8. scripts/validate_core_beta_contract.py

查历史：
  - repo/development-records/README.md
  - repo/development-records/sprint3-closure-a-contract.md
```

---

## 环境启动

```bash
cd /path/to/ANI/repo

make test
make validate-architecture
make validate-spec-split
make validate-core-beta
```

## 明日继续入口

```text
1. 重新读取 CLAUDE.md、ANI-DOCS-INDEX.md、ANI-06-开发计划.md 和本文件。
2. 先执行提交前复核：make validate-sprint4-closure、make test、make validate-architecture、git diff --check。
3. 若复核通过，提交并推送 Sprint 4 到 GitHub。
4. GitHub 状态确认后，再将 CURRENT-SPRINT.md 切换到下一 Sprint，并同步 ANI-06 和 ANI-DOCS-INDEX。
```

---

## 完工后必做

每完成一个批次或可独立验收切片，按顺序执行：

```text
1. 当前批次验收命令
2. make test
3. make validate-architecture
4. git diff --check
5. 新建或更新 repo/development-records/{批次名}.md
6. 更新 repo/development-records/README.md
7. 更新本文件对应批次状态
8. 更新 ANI-06-开发计划.md Section 零对应批次状态
```

完整规约说明：`CLAUDE.md`。

---

*Sprint 4 负责人：[填入]　最后更新：2026-05-21*
