# REAL-K8S-LAB K8s/Kube-OVN/KubeVirt bootstrap

日期：2026-05-29

## 背景

本记录归档三台物理开发服务器上的最小真实底座组件安装结果。目标是暂停业务开发后，先把 Sprint 5 最核心且短期不会频繁变化的 K8s、Kube-OVN、KubeVirt 底座部署起来，提前暴露依赖、镜像源、节点身份和 CNI 问题。

本记录不包含服务器管理地址、SSH 用户、密码、bootstrap token、certificate key 或 kubeconfig 认证材料。敏感访问和临时排障脚本只保留在本地排除目录。

## 最小安装边界

本次已部署：

- Kubernetes 控制平面和两个 worker 节点，版本 `v1.36.1`。
- `containerd.io` `2.2.4-1~ubuntu.24.04~noble`，三台节点均使用 systemd cgroup。
- Kube-OVN `v1.15.8`，使用默认 VPC/subnet 模型。
- KubeVirt `v1.8.2`，仅安装 operator 与 `KubeVirt` CR。

本次仍未部署：

- Helm、vCluster CLI、virtctl、CDI、Harbor、MinIO、Milvus、NATS、Redis、PostgreSQL、Volcano、GPU device plugin、KMS/SM4 backend。
- Kube-OVN Vpc/Subnet + NetworkPolicy + Service/LB live gate。
- KubeVirt VM lifecycle、console/VNC live gate。
- vCluster live Helm/kubeconfig/proxy、node pool live 扩缩容、controller HA、Encryption/Secrets live gates。

## 组件版本

| 组件 | 版本/状态 |
|---|---|
| OS | Ubuntu 24.04.3 LTS |
| Kernel | `6.8.0-71-generic` |
| `containerd.io` | `2.2.4-1~ubuntu.24.04~noble` |
| `kubeadm` / `kubelet` / `kubectl` | `1.36.1-1.1` |
| `cri-tools` | `1.36.0-1.1` |
| `kubernetes-cni` | `1.9.1-1.1` |
| Kubernetes node state | 1 control-plane + 2 workers Ready |
| Kube-OVN | `v1.15.8` |
| KubeVirt | `v1.8.2`, phase `Deployed` |

## 镜像和软件源处理

- Kubernetes kubeadm 镜像使用国内 registry override，避免 `registry.k8s.io` 后端访问超时。
- `containerd` v2 的 sandbox pinned image 修正为国内可拉取的 pause 镜像。
- `containerd` registry hosts 已覆盖 `registry.k8s.io`、`docker.io`、`quay.io`、`ghcr.io`、`gcr.io` 的国内 mirror 入口。
- Kube-OVN 镜像从 `docker.io/kubeovn/kube-ovn:v1.15.8` 拉取成功。
- KubeVirt 镜像从 `quay.io/kubevirt/*:v1.8.2` 拉取成功；事件显示 operator、virt-api、virt-controller、virt-launcher、virt-handler 镜像均完成拉取并运行。

## 关键问题和处理

1. 初次 `kubeadm init` 使用默认 `registry.k8s.io` 拉取镜像超时。
   - 处理：使用国内 Kubernetes 镜像仓库作为 kubeadm image repository。

2. `containerd` v2 仍尝试拉取默认 sandbox image。
   - 处理：修正 `containerd` v2 CRI 配置中的 `pinned_images.sandbox`，并重启 `containerd` / `kubelet`。

3. 三台服务器 OS hostname 初始相同，导致 worker join 后 Kubernetes node identity 冲突。
   - 处理：重置 worker join 状态，删除错误 Node 对象，两个 worker 使用显式 node name 重新 join。

4. Kube-OVN 控制平面 CNI 一度卡在等待 `ovn0` join 网关。
   - 诊断：控制平面 Node annotation 显示已分配，但 OVN NB/SB 中缺少对应 node logical switch port，且缺少对应 `ip.kubeovn.io` 对象。
   - 处理：清理控制平面 Node 的 `ovn.kubernetes.io/*` 残留 annotation，删除残留 IP 对象并重启 Kube-OVN controller/CNI pod，让 Kube-OVN 重新分配 node join port。

5. KubeVirt `virt-launcher` 镜像较大，worker 首次拉取耗时明显长于其它组件。
   - 处理：等待镜像拉取完成；最终 `virt-handler` 在两个 worker 上均 Ready。

## 当前验证结果

- `kubectl get nodes -o wide`：三台节点均为 Ready，控制平面角色正常。
- `kubectl -n kube-system get pod -o wide`：Kube-OVN CNI、OVS/OVN、controller、monitor、CoreDNS 均 Ready。
- `kubectl get --raw='/readyz?verbose'`：API Server readyz passed。
- `kubectl -n kubevirt wait kubevirt kubevirt --for=condition=Available`：通过。
- `kubectl -n kubevirt get kubevirt kubevirt`：phase `Deployed`。
- `kubectl -n kubevirt get pods`：`virt-operator`、`virt-api`、`virt-controller`、`virt-handler` 均 Ready。

## 当前边界

该记录证明真实物理底座上的 K8s、Kube-OVN、KubeVirt 最小组件已经部署成功，不等价于 ANI 的真实 provider live gates 已通过。后续仍需按既有 gate 执行并归档：

- `validate-kubeovn-network-live-gate --live --evidence-output ...`
- `validate-kubevirt-vm-live-gate --live --evidence-output ...`
- `validate-vcluster-live-gate --live --evidence-output ...`
- `validate-vcluster-upgrade-live-gate --live --evidence-output ...`
- `validate-k8s-node-pool-live-gate --live --evidence-output ...`
- `validate-reconcile-ha-live-gate --live --evidence-output ...`
- `validate-kms-sm4-live-gate --live --evidence-output ...`
- `validate-secrets-live-gate --live --evidence-output ...`

## 参考来源

- Kubernetes packages: `https://mirrors.aliyun.com/kubernetes-new/core/stable/v1.36/deb/`
- Kube-OVN install manifest source: `https://raw.githubusercontent.com/kubeovn/kube-ovn/refs/tags/v1.15.8/dist/images/install.sh`
- KubeVirt stable version source: `https://storage.googleapis.com/kubevirt-prow/release/kubevirt/kubevirt/stable.txt`
- KubeVirt release manifests: `https://github.com/kubevirt/kubevirt/releases/download/v1.8.2/kubevirt-operator.yaml` and `kubevirt-cr.yaml`
