# KuberCloud ANI · 文档导航与一致性矩阵

> 最后更新：2026-05-21
> 目的：让人类开发者和 AI 工具在 5 分钟内判断当前开发阶段、文档职责、下一步入口和闭环规则。

---

## 当前结论

```text
当前阶段：Phase 1 / Sprint 4 收尾
当前不是 Phase 2：Phase 2 指 2026-10 以后延期能力
当前优先级：SPEC-CORE-BETA → 开发与验收完成，待提交 GitHub；提交后再切换下一 Sprint（SPEC-SPLIT-A 已完成）
刚完成：Sprint 3（2026-05-20，网络/存储/向量/Workload Identity/SDK Alpha/CORE-DEV-PROFILE-A）
下一步入口：repo/CURRENT-SPRINT.md（先提交 Sprint 4，再切换下一 Sprint）
```

Sprint 4 的核心任务是在 Core Dev Profile Ready 之后进入 API Beta 准备：Core/Services API 分层收口、Core API Beta 准备矩阵、Core API v1 兼容性基线、SDK 加固、四语言 SDK-Mock 联动烟测、Mock Server、API 文档生成和 Sprint 4 关联性闭环已完成；提交前审查命令已通过。当前状态是开发与验收完成、待提交 GitHub；提交完成后再切换下一 Sprint。

`SPEC-SPLIT-A` 已完成：`/models`、`/inference-services`、`/knowledge-bases` 只保留在 Services API，Core API 和 Core SDK 不再承载这些业务路径。

`SPEC-CORE-BETA` 已完成首个切片：`repo/api/core-beta-readiness.yaml` 和 `make validate-core-beta` 用于持续校验 Core P0 path/schema、分页、幂等、状态机、RBAC scope 和 Core/Services 关联边界。

`SPEC-COMPAT-A` 已完成首个切片：`repo/api/core-v1-compatibility-baseline.yaml` 和 `make validate-core-api-compatibility` 用于持续保护 Core API v1 的 path/method/operationId/参数/响应/schema 字段，允许新增可选能力但阻止破坏性变更。

`SDK-BETA-A` 已完成首个切片：四语言 SDK 已生成 `idempotency_key` helper，并通过 `make validate-sdk-beta` 持续校验。

`SDK-BETA-B` 已完成首个切片：四语言 SDK 已生成 cursor 分页 helper，并在 SDK metadata 中标出支持 `limit/cursor` 的 Core 列表操作。

`SDK-BETA-C` 已完成首个切片：四语言 SDK 已生成统一 API error helper，并在 SDK metadata 中标出 API 契约声明的标准错误码。

`SDK-BETA-D` 已完成首个切片：四语言 SDK 已生成 basic example，覆盖 client 初始化、幂等、cursor 分页和 API error helper 的组合用法。

`SDK-MOCK-SMOKE-A` 已完成首个切片：Core Python SDK 已提供标准库 HTTP `request()` 能力，并通过 `make validate-sdk-mock-smoke` 调用由 API 契约驱动的 Core Mock Server。

`SDK-MOCK-SMOKE-B` 已完成首个切片：Core TypeScript SDK 已提供基于 `fetch` 的 `request()` 能力，并通过 `make validate-sdk-mock-smoke` 调用同一个 Core Mock Server。

`SDK-MOCK-SMOKE-C` 已完成首个切片：Core Go SDK 已提供基于 `net/http` 的 `Request()` 能力，并通过 `make validate-sdk-mock-smoke` 调用同一个 Core Mock Server。

`SDK-MOCK-SMOKE-D` 已完成首个切片：Core Java SDK 已提供基于 `java.net.http.HttpClient` 的 `request()` 能力；有 JDK 时调用同一个 Core Mock Server，无 JDK 时执行 source smoke。

`MOCK-A` 已完成首个切片：Core Mock Server 由 `repo/api/openapi/v1.yaml` 驱动，`make validate-mock-a` 校验全量 Core path 可 mock。

`DOC-API-A` 已完成首个切片：Core/Services 静态 API 文档由 API 契约生成到 `repo/docs/api/`，`make validate-doc-api` 校验 operation 和 schema 覆盖。

`SPRINT4-CLOSURE-A` 已完成首个切片：`make validate-sprint4-closure` 统一校验 Sprint 4 API/SDK/Mock/Docs/Records 关联性闭环。

---

## 唯一真实来源矩阵

| 问题 | 先看哪里 | 说明 |
|---|---|---|
| 当前做什么 | `repo/CURRENT-SPRINT.md` | 当前 Sprint 的执行入口，状态、任务、验收命令以它为准 |
| 全局开发节奏 | `ANI-06-开发计划.md` | Sprint 计划、Services 解锁门禁、延期项以它为准 |
| 产品功能边界 | `ANI-02-产品功能设计.md` | Core/Services 分层、v1.0.0 P0 能力边界以它为准 |
| 路线图阶段 | `ANI-03-产品路线图.md` | Phase 1/2/3 与版本号关系以它为准 |
| 工程约定和 AI 工作规则 | `CLAUDE.md` | AI/人类开发前必须先读；架构、提交、闭环规则以它为准 |
| API 契约 | `repo/api/openapi/v1.yaml` | Core REST API 唯一真实来源 |
| API Beta 准备矩阵 | `repo/api/core-beta-readiness.yaml` | Core P0 Beta 审查、兼容性和自动校验矩阵 |
| API 兼容性基线 | `repo/api/core-v1-compatibility-baseline.yaml` | Core API v1 已交付 path/schema 的防破坏基线 |
| Services API 契约 | `repo/api/openapi/services/v1.yaml` | Services 层业务 API 契约 |
| 已完成批次 | `repo/development-records/README.md` | 历史完成记录索引，不作为当前任务清单 |
| 单批次细节 | `repo/development-records/*.md` | 追溯实现、验证和关键文件时再读 |
| 审查提示词模板 | `ANI-10-GPT审查提示词集.md` | 只作为审查问题模板；内置示例不得作为当前事实来源 |

---

## 推荐阅读路径

### 人类开发者

1. `ANI-DOCS-INDEX.md`
2. `CLAUDE.md` 的 5 分钟快速上手
3. `repo/CURRENT-SPRINT.md`
4. `ANI-06-开发计划.md` Section 零和 Sprint 4
5. `repo/api/openapi/v1.yaml` + 相关代码入口

### AI 编码工具

1. 必须先读 `CLAUDE.md`
2. 再读 `repo/CURRENT-SPRINT.md`
3. 开发前检查 `ANI-06-开发计划.md` Section 零
4. 涉及接口时先改 `repo/api/openapi/v1.yaml`
5. 完成后按 `CLAUDE.md` 的进度更新规约闭环

---

## 当前开发门禁

| 日期 | 门禁 | 当前影响 |
|---|---|---|
| 2026-05-31 | P0 依赖矩阵冻结 | 已完成 instance P0 冻结矩阵；Sprint 3 继续网络/存储/向量依赖矩阵 |
| 2026-06-10 | Core API Alpha Freeze | Sprint 2 已冻结 instances；Sprint 3 扩充网络/存储/向量 API |
| 2026-06-20 | SDK Alpha | Go/Python/TypeScript/Java SDK 必须可生成、可 import |
| 2026-06-30 | Core Dev Profile Ready | Services 团队可基于 SDK + dev/local profile 做端到端开发 |
| 2026-09-30 | v1.0.0 Final Delivery | ANI Core v1.0.0 + ANI Services P0 |

---

## 文档维护规则

1. 当前阶段变更时，必须同步 `ANI-DOCS-INDEX.md`、`ANI-06-开发计划.md` 和 `repo/CURRENT-SPRINT.md`。
2. 批次完成时，必须新增或更新 `repo/development-records/{批次名}.md`，并更新 `repo/development-records/README.md`。
3. 历史归档文档允许保留当时日期和上下文，不反向改写为当前态。
4. 若 `CLAUDE.md` 与其它文档冲突，以 `CLAUDE.md` 的工程规则为准；若是进度状态冲突，以 `ANI-06-开发计划.md` Section 零和 `repo/CURRENT-SPRINT.md` 为准。
5. 不把大段完成细节堆到入口文档；入口文档只保留当前状态、下一步和链接。
6. 更换 AI 模型或工具时，必须先重新读取本文件、`CLAUDE.md` 和 `repo/CURRENT-SPRINT.md`，不得依赖上一个会话的记忆。
