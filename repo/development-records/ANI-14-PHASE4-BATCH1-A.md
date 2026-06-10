# ANI-14-PHASE4-BATCH1-A

**批次类型**：Feature batch  
**完成日期**：2026-06-10  
**工作范围**：ani-gateway Phase 4 第一批 handler 骨架（TASK-SVC-001/007/011/017/019/021、TASK-CORE-001/002）

---

## 完成内容

### 新建文件（8 个 handler 文件）

| 文件 | 路由组 | 路由数 |
|------|--------|--------|
| `router/model_resources.go` | `/api/v1/svc/models` | 7 |
| `router/inference_resources.go` | `/api/v1/svc/inference-services` | 6 |
| `router/kb_resources.go` | `/api/v1/svc/knowledge-bases` | 9 |
| `router/gpu_container_resources.go` | `/api/v1/svc/gpu-containers` | 10 |
| `router/sandbox_resources.go` | `/api/v1/svc/sandboxes` | 9 |
| `router/tenant_resources.go` | `/api/v1/svc/tenant` | 9 |
| `router/branding_resources.go` | `/api/v1/branding` | 3 |
| `router/task_resources.go` | `/api/v1/tasks` | 2 |

### 修改文件（2 个）

- `stubs.go`：删除 registerBranding、registerModels、registerInferenceServices、registerKnowledgeBases、registerTasks 五个函数；删除不再使用的 `route` import；保留 `notImplemented` 和 `inferenceProxy`。
- `router.go`：追加 `registerGpuContainers(svc)`、`registerSandboxes(svc)`、`registerTenant(svc)` 三行调用。

### 实现说明

第一批均为骨架实现：列表端点返回 `{"items":[],"next_cursor":null}`（200），单资源 GET 返回 `{"id":"<path_param>"}`，变更操作返回占位 JSON，DELETE 返回 204，SSE 端点设置正确 Content-Type 头。所有 handler 使用 `middleware.GetTenantID(c)` 提取租户 ID。禁止 import `pkg/adapters`、`pkg/ports`（架构边界合规）。

---

## 验收结果

```
go build ./services/ani-gateway/...     → BUILD OK
make test                               → 所有测试通过
make validate-architecture              → architecture guardrails valid
git diff --check                        → EXIT 0
```

curl 验证（ANI_AUTH_MODE=dev）：7 个端点全部从 501 → 200。

---

## 下一步（第二批）

参考 `TASK-DEPENDENCY-MAP.md`：TASK-SVC-002~006（Model CRUD）、TASK-SVC-008~010（InferenceService CRUD）、TASK-SVC-012~013（KnowledgeBase CRUD）、TASK-SVC-018/020（GPU/Sandbox 高级接口）。第二批需要下游微服务地址就绪（通过 env var 配置）。
