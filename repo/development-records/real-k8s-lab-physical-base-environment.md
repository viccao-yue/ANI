# REAL-K8S-LAB physical base environment

日期：2026-05-29

## 背景

本记录归档三台物理开发服务器的最小基础软件环境。目标是让当前 Sprint 5 的真实底座验证后续可以从 Kubernetes bootstrap/preflight 开始推进，同时遵循最小安装原则。

本记录不包含服务器连接地址、登录信息或其它凭据。敏感访问信息只保留在本地排除目录。

## 最小安装边界

已安装或配置：

- Ubuntu APT 源切换为国内镜像源。
- Docker CE 国内 APT 源，仅用于安装 `containerd.io`。
- Kubernetes v1.36 国内 APT 源，仅用于安装 kubeadm/kubelet/kubectl/CRI 工具。
- `containerd` 启用 systemd cgroup。
- `containerd` 配置常见境外镜像仓库的国内 mirror hosts，覆盖 `registry.k8s.io`、`docker.io`、`quay.io`、`ghcr.io`、`gcr.io` 的后续拉取入口。
- `crictl` runtime/image endpoint 固定到 `unix:///run/containerd/containerd.sock`。
- 关闭 swap，并写入 Kubernetes 所需内核模块和 sysctl。
- 安装 Kubernetes bootstrap 所需的基础 OS 包：`conntrack`、`ebtables`、`ethtool`、`iptables`、`jq`、`nfs-common`、`open-iscsi`、`socat` 等。

未安装或未部署：

- 未安装 Helm、vCluster CLI、virtctl、Docker Engine。
- 未初始化 Kubernetes 集群。
- 未安装 Kube-OVN、KubeVirt、vCluster、Harbor、MinIO、Milvus、NATS、Redis、PostgreSQL、Volcano、GPU device plugin 或 KMS/SM4 provider。
- 未格式化、分区或挂载数据盘。
- 未执行任何 REAL-K8S-LAB-A live gate。

## 三台节点共同状态

| 项 | 值 |
|---|---|
| OS | Ubuntu 24.04.3 LTS |
| Kernel | `6.8.0-71-generic` |
| `containerd.io` | `2.2.4-1~ubuntu.24.04~noble` |
| `kubeadm` | `1.36.1-1.1` |
| `kubelet` | `1.36.1-1.1` |
| `kubectl` | `1.36.1-1.1` |
| `cri-tools` | `1.36.0-1.1` |
| `kubernetes-cni` | `1.9.1-1.1` |
| `containerd` service | `active/enabled` |
| `kubelet` service | `inactive/enabled`，尚未执行 `kubeadm init/join` |
| swap | `0` active entries |
| `crictl` endpoint | `unix:///run/containerd/containerd.sock` |

## 软件源

三台节点最终保留的软件源：

- `https://mirrors.aliyun.com/ubuntu/`
- `https://mirrors.aliyun.com/docker-ce/linux/ubuntu`
- `https://mirrors.aliyun.com/kubernetes-new/core/stable/v1.36/deb/`

## 节点磁盘状态

| 节点 | 系统盘 | 未挂载数据盘 |
|---|---|---|
| `dev-phys-01` | 约 893.8G SSD | 约 3.5T SSD |
| `dev-phys-02` | 约 893.8G SSD | 2 x 约 3.5T SSD |
| `dev-phys-03` | 约 893.8G SSD | 2 x 约 3.5T SSD；1 x 约 3.6T HDD |

数据盘仍保持未挂载状态。后续如需 Longhorn、Ceph、MinIO、Harbor、Milvus 或 containerd 数据目录迁移，应先单独确认磁盘规划，再执行格式化或挂载。

## 后续安装原则

- 只有当对应开发阶段或 live gate 需要时，才继续安装 Helm、vCluster、Kube-OVN、KubeVirt、Harbor、MinIO、Milvus、NATS、Redis、PostgreSQL、Volcano、GPU 组件或 KMS/SM4 provider。
- 所有需要 GitHub、Quay、Docker Hub、`registry.k8s.io` 等境外资源的组件，优先使用国内镜像源、镜像仓库重写或后续私有 Harbor 缓存。
- 在 Kubernetes 集群初始化前，不把当前环境标记为 REAL-K8S-LAB-A live ready。

## 验证

- 三台节点只读核验：`containerd` active/enabled，`kubeadm/kubelet/kubectl/cri-tools/kubernetes-cni` 版本一致，`crictl info` 可访问 containerd。
- 三台节点只读核验：`helm`、`vcluster`、`virtctl`、`docker` 均为 absent。
- 三台节点只读核验：swap active entries 为 `0`。
