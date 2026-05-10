# 本地开发环境

## 前提条件

- Docker Desktop 4.x+ 或 Docker Engine 24+（含 Compose V2）
- 可用内存 ≥ 8GB（Milvus 吃内存）
- 磁盘空间 ≥ 20GB

## 快速启动

```bash
# 从仓库根目录执行
make deps          # 启动所有依赖服务

# 验证服务就绪
make deps-status
```

## 服务访问

| 服务 | 地址 | 账号/密码 |
|---|---|---|
| PostgreSQL | localhost:5432 | ani / ani_dev_password |
| MinIO Console | http://localhost:9001 | ani-admin / ani_dev_password |
| NATS Monitor | http://localhost:8222 | — |
| Redis | localhost:6379 | 密码: ani_dev_password |
| Milvus | localhost:19530 | — |
| Milvus Attu | http://localhost:3000（需 `--profile tools`）| — |

## 启动可选工具

```bash
# 启动 Milvus Attu（Web UI 管理 Milvus）
docker compose -f deploy/docker/docker-compose.yml --profile tools up -d attu

# 启动 Dex（OIDC 认证，完整认证流程测试）
docker compose -f deploy/docker/docker-compose.yml --profile auth up -d dex
```

## 环境变量

复制 `.env.example` 为 `.env`，按注释修改后各服务自动加载：

```bash
cp .env.example .env
```

## 关闭和清理

```bash
make deps-down        # 停止服务，保留数据卷
make deps-clean       # 停止服务并删除所有数据（危险！）
```
