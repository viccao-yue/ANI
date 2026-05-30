# ANI · 当前冲刺上手指南

> 新开发者（人类或 AI 工具）的第一个入口文件。本文只描述当前真实执行状态；历史完成批次查 `repo/development-records/README.md`。

> 历史门禁保留：Sprint 4 的 `SPEC-SPLIT-A`、`SPEC-CORE-BETA`、`SPEC-COMPAT-A`、`SDK-BETA-A`、`SDK-BETA-B`、`SDK-BETA-C`、`SDK-BETA-D`、`SDK-MOCK-SMOKE-A`、`SDK-MOCK-SMOKE-B`、`SDK-MOCK-SMOKE-C`、`SDK-MOCK-SMOKE-D`、`MOCK-A`、`DOC-API-A`、`SPRINT4-CLOSURE-A` 和 `validate-sprint4-closure` 仍是提交前回归门禁，不作为当前任务清单。

> **仓库范围：仅 ANI Core。** ANI Services 已冻结并移交外部产品团队，本仓库不再开发任何 Services 代码（旧 Services 骨架只读保留）。外部团队预计 2026-06-10 前后给出清晰的产品功能/交互/API 定义，Core 据此以 AI Coding 快速循环实现支撑。
> **当前重心：真实 live gate 推进（物理服务器 2026-05-29 到位）+ guard 冻结。** 执行清单见下方"下一步"。

## 当前冲刺

| 字段 | 值 |
|---|---|
| **冲刺编号** | Sprint 5（收敛中） |
| **主题** | K8s 集群管理 + 加解密主链路 + **真实底座 live gate 推进**（契约/local 已完成，转入真实验证） |
| **当前状态** | 🔄 local profile 主链路、K8s proxy forwarding adapter/target resolver/metadata 持久化/Gateway 注入接线与 forwarding_static/forwarding_metadata runtime 选择、vCluster Helm/kubeconfig/upgrade provider 代码边界、K8s node pool local profile 与 Cluster API provider 代码边界、Kube-OVN/KubeVirt/controller HA/KMS-SM4/Secrets live contract gates 与 evidence JSON 输出、REAL-K8S-LAB-A component live runner/report/preflight/evidence guard 系列已收敛到 `M1-REAL-LAB-KX` KMS/SM4 live gate command type guard；三台物理开发服务器已完成 Kubernetes `v1.36.1` + Kube-OVN `v1.15.8` + KubeVirt `v1.8.2` 最小组件部署；真实 provider live 验证仍未完成 |
| **已由代码证明完成** | `M1-K8S-A/B/C/D/E/F/G`、`M1-K8S-PROXY-A/B/C/D/E/F`、`M1-K8S-LIVE-A/B/C/D/E/F`、`M1-NETWORK-LIVE-A/B`、`M1-KUBEVIRT-LIVE-A/B`、`M1-ENCRYPT-A/B/C/D`、`M1-ENCRYPT-LIVE-A/B`、`M1-SECRETS-A/B/C/D`、`M1-SECRETS-LIVE-A/B`、`M1-RECONCILE-A/B/C/D/E`、`M1-RECONCILE-LIVE-A/B`、`REAL-K8S-LAB-A`、`M1-REAL-LAB-B` 至 `M1-REAL-LAB-KX` 的验证器、文档记录和测试闭环、`REAL-K8S-LAB physical base environment` 最小物理基础环境记录、`REAL-K8S-LAB K8s/Kube-OVN/KubeVirt bootstrap` 真实底座组件部署记录；完整批次索引见 `repo/development-records/README.md` |
| **不可标记为完成** | vCluster live Helm 安装验证、真实 kubeconfig 可用性、live proxy 验证、live vCluster 升级验证真实执行结果、真实 Kube-OVN Vpc/Subnet 与 NetworkPolicy/Service LB 验证结果、KubeVirt VM lifecycle 与 console/VNC live 验证结果、真实节点池 live 扩缩容、GPU 节点池真实调度验证、controller 多副本 live HA failover 验证真实执行结果、KMS/SM4 live backend 验证、对象存储 + KMS/SM4 provider streaming 端到端验收、Kubernetes Secret live 写入验证、实例 Secret env/file/VM volume live 注入验证的真实执行结果 |
| **关联历史门禁** | Sprint 4 `SPEC-CORE-BETA` / `api/core-beta-readiness.yaml`、`DOC-API-A` 仍需保持 API docs 与 OpenAPI 同步；`SDK-BETA-A`、`SDK-BETA-B`、`SDK-BETA-C`、`SDK-BETA-D` 仍需保持 SDK helper 与新增 Core API 同步 |
| **最后校准日期** | 2026-05-29 |

## 已完成切片

本节保留当前 Sprint 的可执行事实，完整历史清单以 `repo/development-records/README.md` 为唯一归档索引。

1. `M1-K8S-A/B/C/D/E/F/G`：K8s 集群 CRUD/kubeconfig/proxy local profile、vCluster Helm provider、vCluster kubeconfig provider、cluster upgrade provider、node pool local profile 和 Cluster API node pool provider 代码边界已完成；不代表真实 vCluster 或节点池 live 验证完成。
2. `M1-K8S-PROXY-A/B/C/D/E/F`：proxy forwarding adapter、target resolver/store、metadata 持久化、Gateway 注入接线和 `forwarding_static` / `forwarding_metadata` runtime 选择已完成；不代表 live proxy 验证完成。
3. `M1-K8S-LIVE-A/B/C/D/E/F`：vCluster、`M1-K8S-LIVE-B` node pool（含 Cluster API 和 GPU 调度检查）、`M1-K8S-LIVE-C` vCluster upgrade（含 Helm `controlPlane.distro.k8s.version` 检查）live contract gates 与 evidence JSON 输出已完成；live 模式待真实 lab 执行。
4. `M1-NETWORK-LIVE-A/B`：`M1-NETWORK-LIVE-A` Kube-OVN `Vpc/Subnet`、NetworkPolicy、Service/LB live contract gate 与 evidence JSON 输出已完成；live 模式待真实 lab 执行。
5. `M1-KUBEVIRT-LIVE-A/B`：KubeVirt VM lifecycle、console/VNC live contract gate 与 evidence JSON 输出已完成；live 模式待真实 lab 执行。
6. `M1-ENCRYPT-A/B/C/D` 与 `M1-ENCRYPT-LIVE-A/B`：Encryption local profile、KMS/SM4 HTTP provider 代码边界、对象内容 SM4-GCM 流式加解密边界、KMS/SM4 live gate 与 evidence JSON 输出已完成；真实 KMS/SM4 backend 和 objectstore round trip 待执行。
7. `M1-SECRETS-A/B/C/D` 与 `M1-SECRETS-LIVE-A/B`：Secret local profile、Kubernetes Secret provider 写入代码边界、容器/Job env/file 注入、VM Secret volume 注入、Secrets live gate 与 evidence JSON 输出已完成；真实 Pod/VM 可见性待执行。
8. `M1-RECONCILE-A/B/C/D/E` 与 `M1-RECONCILE-LIVE-A/B`：controller adapter、opt-in bootstrap、目标级失败退避、Prometheus metrics、独立 worker、metadata-backed leader election、controller HA live gate 与 evidence JSON 输出已完成；多副本 HA failover live 待执行。
9. `REAL-K8S-LAB-A` 真实底座验证门禁 + `M1-REAL-LAB guard series`（B~KX，237 个 guard validators）：REAL-K8S-LAB-A 组件级 contract gate、`--live` evidence JSON、component live runner、preflight/env/report/evidence/provenance/diagnostic guards，以及 kubeconfig/vCluster/vCluster upgrade/node pool/Kube-OVN/KubeVirt/controller HA/KMS-SM4 live gate 前置校验系列已完成；完整 guard 列表见 [`repo/development-records/guard-series/REAL-K8S-LAB-guard-index.md`](development-records/guard-series/REAL-K8S-LAB-guard-index.md)；最新完成 ID：`M1-REAL-LAB-KX`（KMS/SM4 live gate command type guard）；不代表组件级 `--live` 已在真实 lab 执行成功。
10. `REAL-K8S-LAB physical base environment`：三台物理开发服务器已完成最小基础软件环境，包含 `containerd`、Kubernetes v1.36 bootstrap 工具、CRI 工具、Kubernetes 所需 OS 包、国内 APT 源和常见境外容器镜像仓库的国内 mirror hosts；未执行 `kubeadm init/join`，未安装 Helm、vCluster、Kube-OVN、KubeVirt、KMS/SM4 backend 或其它上层组件；记录见 `repo/development-records/real-k8s-lab-physical-base-environment.md`。
11. `REAL-K8S-LAB K8s/Kube-OVN/KubeVirt bootstrap`：三台物理开发服务器已完成 Kubernetes `v1.36.1` 集群、Kube-OVN `v1.15.8` 和 KubeVirt `v1.8.2` 最小部署；Kube-OVN CNI/CoreDNS Ready，KubeVirt phase `Deployed`；记录见 `repo/development-records/real-k8s-lab-k8s-kubeovn-kubevirt-bootstrap.md`。该记录不代表 Kube-OVN network live gate 或 KubeVirt VM lifecycle live gate 已通过。

## 当前事实边界

- K8s 集群当前默认仍是 local profile 模拟服务；runtime 层已有 vCluster Helm/kubeconfig/upgrade provider、node pool provider 和 proxy forwarding 代码边界，但尚未完成 REAL-K8S-LAB-A live Helm 安装、live 升级、真实节点池扩缩容或 GPU 调度验证。
- Kubeconfig 默认仍指向模拟 vCluster endpoint；`vcluster_helm` provider mode 已具备通过 `vcluster connect --print` 获取 kubeconfig 的代码边界，但尚未在真实 vCluster 中验证 kubectl/Helm 可用性。
- K8s proxy 当前默认仍是 Core API 契约和 local profile 响应；`M1-K8S-LIVE-A` 已提供 `validate-vcluster-live-gate` 用于后续验证 live proxy，但尚未完成 live vCluster 验证。
- Encryption 当前已有 local profile、KMS/SM4 HTTP provider 代码边界、对象内容 SM4-GCM 流式加解密代码边界和 `validate-kms-sm4-live-gate`；尚未完成 live KMS/SM4 backend 验证，也未完成对象存储 + provider streaming 端到端验收。
- Secret 当前已有 Kubernetes Secret provider 写入代码边界和实例 binding manifest 注入边界；尚未完成 Kubernetes Secret live 写入、Pod env/file 可见性和 VM guest 内可见性的真实执行记录。
- Reconcile controller 当前完成 metadata-backed leader election 代码边界和 `validate-reconcile-ha-live-gate`；尚未做多副本 live HA failover 真实执行验证。
- REAL-K8S-LAB-A 当前只完成验证门禁定义与大量前置 guard；三台物理开发服务器已完成 Kubernetes `v1.36.1`、Kube-OVN `v1.15.8`、KubeVirt `v1.8.2` 最小组件部署，但不代表 Kube-OVN network live gate、KubeVirt VM lifecycle live gate、vCluster、KMS/SM4 或 Kubernetes Secret 已真实跑通。
- Sprint 6 不能作为当前执行入口，直到 Sprint 5 未完成项被代码、API 契约和测试共同证明完成或明确重新排期。

## 真实底座门禁

从 Sprint 5 起，涉及 K8s、Kube-OVN、KubeVirt、vCluster、KMS/SM4、K8s Secret 注入等真实组件的能力，不能再只靠 local profile 宣称完成。local profile 只能证明 API、SDK、状态机和调用边界正确；真实运行能力必须由真实组件环境、固定验证命令或批次记录证明。当前固定入口是 `REAL-K8S-LAB-A` 和 `make validate-real-k8s-profile`。

当前必须并行准备的真实验证环境：

| 组件 | 进入时机 | 验证目标 |
|---|---|---|
| K8s 测试集群 | Sprint 5 当前起 | API Server、Namespace、RBAC、ServiceAccount、StorageClass 基础可用 |
| Kube-OVN | Sprint 5 当前起 | VPC、Subnet、NetworkPolicy、Service/LB 可创建、可观察 |
| KubeVirt | Sprint 5 当前起 | VM 创建、启动、停止、删除、console/VNC 可运行 |
| vCluster | Sprint 5 当前起 | K8s 集群创建、kubeconfig、proxy 能真实访问租户集群 |
| KMS/SM4 + K8s Secret | Sprint 5~6 | 加解密、密钥轮换、Secret 写入和实例注入真实跑通 |

## ⛔ Guard 冻结令（2026-05-30 起）

REAL-K8S-LAB guard 系列（`M1-REAL-LAB-*`，已 299 个）**冻结，不再新增**。唯一例外：真实 live gate 执行中**实际复现**的缺陷，且只能用一个新 guard 防回归——此时新增一个并在 guard-index 注明对应真实缺陷。禁止再为"假设可能出现"的字段/空值/空格/类型边角预防性批量生成 guard。**当前唯一重心是把下面的真实 live gate 跑通，不是扩充校验器。**

## 下一步：真实 live gate 推进顺序（唯一执行清单）

> 物理服务器 2026-05-29 到位，K8s `v1.36.1` + Kube-OVN `v1.15.8` + KubeVirt `v1.8.2` 已最小部署。
> 本仓库**只做 Core**；Services 已冻结移交外部团队，不在此推进。
> 执行原则：① 组件已就绪的先做；② 产品关键链路优先；③ 每个 gate 跑通后立刻 `--evidence-output` 归档 JSON 证据，并在 `repo/development-records/` 写一条 live 结果记录（Feature batch 规则）。
> 通用前置（每个 gate 进入前）：`--component-env-template-output` 生成 env 模板 → 填好 `contract_gates[].required_env` → `--component-preflight --component-env-file ... --evidence-output ...` 校验配置完整 → `--component-gate <id>` 重跑单个失败项。

> 命令说明：`make validate-*-live-gate` 只跑 **contract 模式**（无真实集群）；真实 **live 模式**必须直接运行脚本并带 `--live --evidence-output <path>`（脚本均支持这两个 flag）。

| 顺序 | Live gate | live 命令 | 前置/需补装 | 通过判据 |
|---|---|---|---|---|
| 1 | Kube-OVN 网络（组件已 Ready，无需补装） | `python scripts/validate_kubeovn_network_live_gate.py --live --evidence-output <path>` | 无（Kube-OVN 已部署） | Vpc/Subnet 创建可观察、NetworkPolicy 生效、Service/LB 可达 |
| 2 | KubeVirt VM 生命周期（组件已 Deployed） | `python scripts/validate_kubevirt_vm_live_gate.py --live --evidence-output <path>` | virtctl/CDI（仅在本步补装） | VM create/start/stop/delete、console/VNC 可用 |
| 3 | vCluster 集群 + kubeconfig + live proxy | `python scripts/validate_vcluster_live_gate.py --live --evidence-output <path>` | Helm + vCluster（本步补装） | Helm 装出 vCluster、kubeconfig 可用、kubectl `/version` OK、Core proxy 转发成功 |
| 4 | vCluster 升级 | `python scripts/validate_vcluster_upgrade_live_gate.py --live --evidence-output <path>` | 依赖第 3 步 | Helm `controlPlane.distro.k8s.version` 升级生效、升级后 kubeconfig/proxy 仍可用 |
| 5 | 节点池扩缩容 + GPU 调度 | `python scripts/validate_k8s_node_pool_live_gate.py --live --evidence-output <path>` | Cluster API（本步补装）、GPU 节点 | MachineDeployment 真实扩缩容、GPU workload 调度成功 |
| 6 | Kubernetes Secret 写入 + 注入 | `python scripts/validate_secrets_live_gate.py --live --evidence-output <path>` | 依赖 K8s 基座 | Secret 真实写入、Pod env/file 可见、KubeVirt VM volume 可见 |
| 7 | KMS/SM4 加解密 + 对象存储 round trip | `python scripts/validate_kms_sm4_live_gate.py --live --evidence-output <path>` | KMS 或 SM4 backend + 对象存储（本步补装） | 密钥生命周期、SM4-GCM 流式 seal/open、objectstore sealed 内容回读一致 |
| 8 | Controller HA failover | `python scripts/validate_reconcile_ha_live_gate.py --live --evidence-output <path>` | 双副本 worker 部署 | `control_plane_leases` holder 切换、删 leader pod 后 follower 接管 |

完成一个就把该能力从"local profile / 代码边界"升级标记为"real-provider 已验证"（见 ANI-06 真实底座门禁规则），并更新本文"已完成切片"与"不可标记为完成"两节。八个全部通过 = Sprint 5 真实收敛完成，方可进入 Sprint 6。

> 注：也可用 `make validate-real-k8s-profile` 的组件级聚合模式（`--component-live --component-env-file ... --component-evidence-dir ... --evidence-output ...`）一次跑多个 gate；上表的单 gate 脚本命令用于聚焦推进与排错。

## 文档入口边界

- `CLAUDE.md` 只维护稳定强制规则、读取顺序、架构边界、提交门禁和 Karpathy 五条开发原则。
- 当前 Sprint 的详细完成项、未完成项、验收命令、下一步和真实底座边界以本文为准。
- 批次实现细节只写入 `repo/development-records/*.md`，不得把每日开发流水账或 API path 长列表写回 `CLAUDE.md`。
- 修改入口文档后必须运行 `make validate-doc-entrypoints`。

## 验收命令

```bash
make validate-doc-entrypoints
python scripts/validate_yaml.py api/openapi/v1.yaml api/openapi/services/v1.yaml deploy/real-k8s-lab/profile.yaml deploy/real-k8s-lab/vcluster-live-gate.yaml deploy/real-k8s-lab/vcluster-upgrade-live-gate.yaml deploy/real-k8s-lab/k8s-node-pool-live-gate.yaml deploy/real-k8s-lab/kubeovn-network-live-gate.yaml deploy/real-k8s-lab/kubevirt-vm-live-gate.yaml deploy/real-k8s-lab/reconcile-ha-live-gate.yaml deploy/real-k8s-lab/kms-sm4-live-gate.yaml deploy/real-k8s-lab/secrets-live-gate.yaml
make validate-mock-a
make validate-doc-api
make validate-sdk-beta
make validate-sdk-mock-smoke
make validate-real-k8s-profile
make validate-vcluster-live-gate
make validate-vcluster-upgrade-live-gate
make validate-k8s-node-pool-live-gate
make validate-kubeovn-network-live-gate
make validate-kubevirt-vm-live-gate
make validate-reconcile-ha-live-gate
make validate-kms-sm4-live-gate
make validate-secrets-live-gate
go test ./services/ani-gateway/internal/router ./pkg/adapters/runtime
go test ./pkg/adapters/runtime ./pkg/bootstrap -run 'TestLocalWorkloadReconcileController|TestNewCapabilitiesDefaults|TestConfigEnvironmentOverridesWorkloadReconcileController|TestStartWorkloadReconcileControllerRequiresOptIn' -v
git diff --check
```

> 在没有联网依赖缓存时，`go test` 可能需要下载 Go module；本地可复用 `Makefile` 中的 `GOCACHE`/`GOMODCACHE` 设置。
