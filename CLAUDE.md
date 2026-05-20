# CLAUDE.md

This file provides mandatory guidance for Claude Code / Codex / Cursor / GPT coding agents working in this repository.

> 本文件是 AI/人类开发的**轻量入口和强制规则索引**。动态进度、批次细节和历史记录不在本文重复维护；首次加载必须按本文的读取顺序进入真实来源。

---

## 0. 当前状态

```text
当前阶段：Phase 1 / Sprint 4 收尾
当前状态：开发与验收完成，待提交 GitHub；提交后再切换下一 Sprint
当前不是 Phase 2：Phase 2 是 2026-10 以后延期能力
下一步入口：repo/CURRENT-SPRINT.md
文档导航：ANI-DOCS-INDEX.md
首个正式版本目标：v1.0.0，2026-09-30
```

当前 Sprint 4 已完成并需保持可追溯的批次标记：`SPEC-SPLIT-A`、`SPEC-CORE-BETA`、`SPEC-COMPAT-A`、`SDK-BETA-A`、`SDK-BETA-B`、`SDK-BETA-C`、`SDK-BETA-D`、`SDK-MOCK-SMOKE-A`、`SDK-MOCK-SMOKE-B`、`SDK-MOCK-SMOKE-C`、`SDK-MOCK-SMOKE-D`、`MOCK-A`、`DOC-API-A`、`SPRINT4-CLOSURE-A`。批次细节以 `repo/CURRENT-SPRINT.md` 和 `repo/development-records/README.md` 为准。

---

## 1. 首次加载顺序

每次新会话、切换模型、上下文压缩或准备继续开发时，按顺序读取：

```text
1. CLAUDE.md
2. ANI-DOCS-INDEX.md
3. repo/CURRENT-SPRINT.md
4. ANI-06-开发计划.md 的 Section 零和当前 Sprint
5. 本次任务直接相关的 API 契约、代码、测试和批次记录
```

只根据 `repo/CURRENT-SPRINT.md` 判断当前要做什么。不要从历史 handoff、旧批次记录、路线图远期规划或 Phase 2 内容倒推当前任务。

---

## 2. 真实来源

| 问题 | 真实来源 |
|---|---|
| 文档导航和读取顺序 | `ANI-DOCS-INDEX.md` |
| 当前 Sprint、状态、验收命令 | `repo/CURRENT-SPRINT.md` |
| 全局阶段、Services 解锁门禁、延期项 | `ANI-06-开发计划.md` |
| 已完成批次索引 | `repo/development-records/README.md` |
| 单批次实现与验证细节 | `repo/development-records/*.md` |
| Core API 契约 | `repo/api/openapi/v1.yaml` |
| Services API 契约 | `repo/api/openapi/services/v1.yaml` |
| Core API Beta 准备矩阵 | `repo/api/core-beta-readiness.yaml` |
| Core API v1 兼容性基线 | `repo/api/core-v1-compatibility-baseline.yaml` |
| 产品边界、路线图、版本、代码规范、ports/adapters 细则 | `ANI-02`、`ANI-03`、`ANI-11`、`ANI-12`、`ANI-13` |

进度状态冲突时，以 `ANI-06-开发计划.md` Section 零和 `repo/CURRENT-SPRINT.md` 为准。工程约定冲突时，以本文的强制规则为准。

---

## 3. 强制架构边界

1. ANI 分为 ANI Core 和 ANI Services 两层。ANI Core 只负责基础设施平台能力，只输出 REST API、SDK、CLI；不得包含模型推理、RAG、PaaS 业务逻辑。
2. ANI Services 只能通过 Core REST API / SDK 调用 Core；禁止 import Core 代码包或直接操作底层组件。
3. `repo/services/model-service/` 和 `repo/services/kb-service/` 逻辑属于 ANI Services，暂存于 monorepo。Core 服务禁止调用它们。
4. Services 业务资源如 `models`、`inference-services`、`knowledge-bases` 必须维护在 `repo/api/openapi/services/v1.yaml`，不得回流到 Core API。
5. `CORE-DEV-PROFILE-A` 是 Core dev/local profile，不是 Services 业务 mock。Services 团队如需业务 mock，应在 Services 层自行建设。

---

## 4. API 与 SDK 强制规则

1. `repo/api/openapi/v1.yaml` 是 ANI Core 对外 REST API 的唯一真实来源；所有新 Core API 必须先改 API 契约，再写实现、测试和 SDK。
2. Core API `servers[0].url` 必须为 `https://{host}/api/v1`；Services API `servers[0].url` 必须为 `https://{host}/api/v1/svc`。
3. Proto 是内部 gRPC 实现细节；当 Proto 与 REST schema 描述同一资源冲突时，以 API 契约为准。
4. Core API v1 允许新增可选 request 字段、response 字段、端点和枚举值；删除字段、改字段类型、删除端点、修改 HTTP 方法、修改错误语义均属于破坏性变更。
5. 所有 POST 创建和有副作用的 PUT/PATCH 必须支持 `idempotency_key`；客户端重试必须复用同一个 key。
6. Core SDK 来源是 `repo/api/openapi/v1.yaml`；Services SDK 来源是 `repo/api/openapi/services/v1.yaml`。SDK 不得包含对方层资源类型。
7. Core Mock Server 本地默认地址是 `http://127.0.0.1:4010/api/v1`，只覆盖 Core API 契约，不提供 Services 业务 mock。
8. Sprint 4 的 SDK helper、Mock Server、静态 API 文档和兼容性基线细节不在本文展开；以 `repo/CURRENT-SPRINT.md`、`repo/development-records/README.md` 和对应校验脚本为准。

---

## 5. 组件边界强制规则

1. `pkg/ports/` 中的 port 指产品能力抽象/接口边界，不是 TCP/IP 端口。
2. 只要某组件承载 ANI 产品能力、会被 Core service / Services / API handler 依赖，或存在合理替换/多实现可能，就必须经过 `pkg/ports/` 和 `pkg/adapters/`。
3. 业务服务不得直接依赖 MinIO、Milvus、NATS JetStream、Redis、Harbor、CloudNativePG 等组件 SDK；直接导入必须有 allowlist、`coupling_level` 和保留理由。
4. Kubernetes API 属于 `bounded_direct`：允许在 adapter/controller/preflight 等边界内原生使用；禁止在 Gateway handler、Core domain service、ANI Services 业务服务中直接使用 K8s SDK 或拼装 provider 对象。
5. `ports` 不封装完整 Kubernetes SDK；`ports` 只表达 ANI 产品意图，例如 `WorkloadRuntime`、`WorkloadProviderApply`、`NetworkProviderRenderer`。
6. VM、Container、GPU Container、Sandbox、Batch Job、Notebook、K8s Cluster、BM、DPU 都必须先经过 `WorkloadRuntime` 能力抽象；异构 GPU 发现、分类和调度必须经过 `GPUInventory`。

---

## 6. 开发、验证、文档闭环

1. 代码生成批次使用可回溯命名，例如 `M2.1-TASK-A`、`M1-INSTANCE-U`。历史 `Stage 3A/3B/3C` 仅可作为旧名说明。
2. 每个批次完成时，必须通过当前批次验收命令、`make test`、`make validate-architecture`、`git diff --check`。
3. 每个批次完成时，必须更新 `repo/development-records/{批次名}.md`、`repo/development-records/README.md`、`repo/CURRENT-SPRINT.md`、`ANI-06-开发计划.md`。
4. Sprint 切换时，必须同步 `ANI-06-开发计划.md`、`repo/CURRENT-SPRINT.md`、`ANI-DOCS-INDEX.md`。
5. 当前 Sprint 4 尚未提交 GitHub 前，不得把当前执行入口切换到 Sprint 5。

---

## 7. 提交与版本

提交前至少执行：

```bash
cd repo
make test
make validate-architecture
git diff --check
```

发布或预发布必须遵守 `ANI-12-版本管理策略.md`。ANI 使用 SemVer，首个正式版本是 `v1.0.0`，目标日期 `2026-09-30`；当前不得标记为 `v1.0.0` 或 RC。

---

## 8. Karpathy 四条开发原则

### 原则一：先思考，再编码
**不要假设。不要掩饰困惑。要揭示取舍。**

- 如果需求有歧义，明确说出来并询问，而不是悄悄选一种猜测实现
- 存在多种合理方案时，列出并说明各方案的取舍，由人决策
- 面对复杂问题，先陈述理解再动手
- 遇到不确定的地方，停下来问，而不是带着疑惑继续

### 原则二：用能解决问题的最小代码
**拒绝一切带有猜想的实现。**

- 不实现没被要求的功能，哪怕"感觉以后用得到"
- 不为一次性代码创建抽象层
- 不加"灵活性""可配置性"等未被要求的扩展点
- 200 行能写成 50 行的，重写

### 原则三：只触碰你必须改动的部分
**只清理你自己制造的脏。**

- 不顺手"优化"任务范围之外的代码、注释或格式
- 不重构没坏的东西
- 保持现有代码风格，即使你有不同偏好
- 发现死代码，提出来，不要自作主张删除

### 原则四：定义成功标准，循环迭代直到验证通过
**把任务转化为可验证的目标。**

- 每个任务开始前明确"什么状态算完成"
- 多步骤任务先列出简短计划和验证步骤
- 优先给 Claude 成功标准而非操作指令：不是"修复这个 bug"，而是"写一个能复现 bug 的测试，再修复它"
