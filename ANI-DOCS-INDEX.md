# KuberCloud ANI · 文档导航与一致性矩阵

> 最后更新：2026-05-29
> 目的：让人类开发者和 AI 工具在 5 分钟内判断当前开发阶段、文档职责、下一步入口和闭环规则。

---

## 当前结论

```text
当前阶段：Phase 1 / Sprint 5 收敛中
当前不是 Phase 2：Phase 2 指 2026-10 以后延期能力
当前入口：repo/CURRENT-SPRINT.md
当前真实底座门禁：REAL-K8S-LAB-A / make validate-real-k8s-profile
```

Sprint 5 当前仍在收敛 K8s、Kube-OVN、KubeVirt、vCluster、KMS/SM4、Kubernetes Secret 与 controller HA 的真实 provider 主链路。local profile 只能证明 API、SDK、状态机和调用边界正确；真实运行能力必须由真实组件环境、固定验证命令或批次记录证明。

刚完成的当前批次是 `M1-REAL-LAB-KX` KMS/SM4 live gate command type guard：`validate_kms_sm4_live_gate.py --gate <kms-command-non-string-gate>` 会在 contract 校验阶段拒绝非字符串 `live_checks[].command`，避免真实 lab 手工复跑时把 YAML boolean KMS/SM4 command 带入后续 live 执行阶段。该能力不代表任何真实 lab live gate 已执行成功。

三台物理开发服务器已完成最小真实底座组件部署：Kubernetes `v1.36.1` 集群、Kube-OVN `v1.15.8` 和 KubeVirt `v1.8.2` 已在物理环境中达到组件 Ready/Deployed 状态；记录见 `repo/development-records/real-k8s-lab-k8s-kubeovn-kubevirt-bootstrap.md`。该状态不代表 Kube-OVN network live gate、KubeVirt VM lifecycle live gate、vCluster、KMS/SM4 或 Kubernetes Secret live gates 已通过。

M1-K8S-LIVE-A / validate-vcluster-live-gate 对应的 vCluster live Helm 安装验证、真实 kubeconfig 可用性、live proxy 验证，M1-K8S-LIVE-C / validate-vcluster-upgrade-live-gate 对应的 live vCluster upgrade 和 Helm `controlPlane.distro.k8s.version` 真实执行结果，M1-K8S-LIVE-B / validate-k8s-node-pool-live-gate 对应的真实 Cluster API 节点池 live 扩缩容和 GPU 节点池真实调度验证，M1-NETWORK-LIVE-A / validate-kubeovn-network-live-gate 对应的真实 Kube-OVN Vpc/Subnet 与 NetworkPolicy/Service LB 验证结果、M1-KUBEVIRT-LIVE-A / validate-kubevirt-vm-live-gate 对应的 KubeVirt VM lifecycle 与 console/VNC live 验证结果、M1-RECONCILE-LIVE-A / validate-reconcile-ha-live-gate 对应的 `control_plane_leases` holder 切换与 controller 多副本 live HA failover 验证真实执行结果、M1-ENCRYPT-LIVE-A / validate-kms-sm4-live-gate 对应的 KMS/SM4 live backend 验证、对象存储 + KMS/SM4 provider streaming 端到端验收、Kubernetes Secret live 写入验证和实例 Secret env/file/VM volume live 注入验证真实执行结果尚未完成。

---

## 唯一真实来源矩阵

| 问题 | 先看哪里 | 说明 |
|---|---|---|
| 当前做什么 | `repo/CURRENT-SPRINT.md` | 当前 Sprint 的执行入口，状态、任务、验收命令以它为准 |
| 全局开发节奏 | `ANI-06-开发计划.md` | Sprint 计划、Services 解锁门禁、延期项以它为准 |
| 产品功能边界 | `ANI-02-产品功能设计.md` | Core/Services 分层、v1.0.0 P0 能力边界以它为准 |
| 系统架构图和模块边界 | `ANI-05-系统架构设计.md` | Core/Services、API/SDK、ports/adapters、local profile/real provider 的结构图以它为准 |
| 路线图阶段 | `ANI-03-产品路线图.md` | Phase 1/2/3 与版本号关系以它为准 |
| 工程约定和 AI 工作规则 | `CLAUDE.md` | AI/人类开发前必须先读；只维护稳定规则和入口，不维护批次流水账 |
| API 契约 | `repo/api/openapi/v1.yaml` | Core OpenAPI REST API 与 Core/Services 跨层控制面契约的唯一真实来源 |
| Services API 契约 | `repo/api/openapi/services/v1.yaml` | Services 层业务 API 契约 |
| 已完成批次 | `repo/development-records/README.md` | 历史完成记录索引，不作为当前任务清单 |
| 单批次细节 | `repo/development-records/*.md` | 追溯实现、验证和关键文件时再读 |
| guard 微批次系列详情 | `repo/development-records/guard-series/{series}-guard-index.md` | guard 系列完整 ID 列表和批次链接；新 guard 微批次只更新此文件，不更新主文档 |

---

## 推荐阅读路径

### 人类开发者

1. `ANI-DOCS-INDEX.md`
2. `CLAUDE.md` 的 5 分钟快速上手
3. `repo/CURRENT-SPRINT.md`
4. `ANI-06-开发计划.md` Section 零和当前 Sprint
5. `ANI-05-系统架构设计.md`
6. `repo/api/openapi/v1.yaml` + `repo/api/openapi/services/v1.yaml` + 相关代码入口

### AI 编码工具

1. 必须先读 `CLAUDE.md`
2. 再读 `ANI-DOCS-INDEX.md`
3. 再读 `repo/CURRENT-SPRINT.md`
4. 开发前检查 `ANI-06-开发计划.md` Section 零
5. 涉及架构边界时检查 `ANI-05-系统架构设计.md`
6. 涉及接口时先改 `repo/api/openapi/v1.yaml` 或 `repo/api/openapi/services/v1.yaml`
7. 完成后按 `CLAUDE.md` 的进度更新规约闭环

---

## 当前开发门禁

| 日期 | 门禁 | 当前影响 |
|---|---|---|
| 2026-05-31 | P0 依赖矩阵冻结 | 已完成历史批次归档，后续只按当前 Sprint 补缺口 |
| 2026-06-10 | Core API Alpha Freeze | 已完成 instances 等核心路径冻结；新增能力必须保持兼容性 |
| 2026-06-20 | SDK Alpha | 四语言 Core/Services SDK 已可生成，并由 SDK Beta/Mock smoke 持续校验 |
| 2026-06-30 | Core Dev Profile Ready | Core dev/local profile 边界已建立；Sprint 5 继续补真实 provider |
| 2026-07-31 | Core Real Path Beta | 当前关键门禁：K8s/Kube-OVN/KubeVirt/vCluster 真实底座验证和真实 provider 主链路 |
| 2026-09-30 | v1.0.0 Final Delivery | ANI Core v1.0.0 + ANI Services P0 |

---

## 文档维护规则

1. 当前阶段变更时，必须同步 `ANI-DOCS-INDEX.md`、`ANI-06-开发计划.md` 和 `repo/CURRENT-SPRINT.md`。
2. 批次完成时，必须新增或更新 `repo/development-records/{批次名}.md`，并更新 `repo/development-records/README.md`。
3. 历史归档文档允许保留当时日期和上下文，不反向改写为当前态。
4. 若 `CLAUDE.md` 与其它文档冲突，以 `CLAUDE.md` 的工程规则为准；若是进度状态冲突，以 `ANI-06-开发计划.md` Section 零和 `repo/CURRENT-SPRINT.md` 为准。
5. `CLAUDE.md` 只保留稳定强制规则、读取顺序、架构边界、提交门禁和 Karpathy 五条开发原则；禁止写入单批次完成清单、API path 长列表、文件级变更清单和每日开发流水账。
6. 动态进度只维护在 `repo/CURRENT-SPRINT.md`、`ANI-06-开发计划.md` Section 零和 `repo/development-records/*.md`；入口文档只保留当前状态、下一步和链接。
7. 更换 AI 模型或工具时，必须先重新读取本文件、`CLAUDE.md` 和 `repo/CURRENT-SPRINT.md`，不得依赖上一个会话的记忆。
8. 修改文档入口后必须运行 `make validate-doc-entrypoints`。
