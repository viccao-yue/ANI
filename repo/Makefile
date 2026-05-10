# ANI Monorepo — 统一构建入口
# 使用方式：make <target>

.PHONY: help deps deps-down deps-status deps-clean \
        tools-proto gen-api gen-proto \
        build build-gateway build-model-service build-task-service \
        test test-go test-python test-cover \
        lint lint-go lint-python lint-ts \
        dev-gateway dev-console \
        image image-gateway \
        clean

# ── 帮助 ──────────────────────────────────────────────────────────────────────
help:
	@echo "ANI Monorepo 构建命令"
	@echo ""
	@echo "  开发环境:"
	@echo "    make deps           启动本地依赖服务（PG/MinIO/NATS/Redis/Milvus）"
	@echo "    make deps-down      停止本地依赖服务"
	@echo "    make deps-status    查看依赖服务状态"
	@echo "    make deps-clean     停止并清除所有数据（危险）"
	@echo ""
	@echo "  代码生成:"
	@echo "    make tools-proto    安装本地 Protobuf 生成插件"
	@echo "    make gen-api        从 OpenAPI Spec 生成 Go/TS 代码"
	@echo "    make gen-proto      从 Protobuf 生成 Go gRPC 代码"
	@echo ""
	@echo "  构建:"
	@echo "    make build          构建所有 Go 服务和工具"
	@echo "    make build-gateway  仅构建 ANI Gateway"
	@echo "    make build-model-service  仅构建 Model Service"
	@echo "    make build-task-service   仅构建 Task Service"
	@echo ""
	@echo "  测试:"
	@echo "    make test           运行所有测试（Go + Python）"
	@echo "    make test-cover     生成测试覆盖率报告"
	@echo ""
	@echo "  代码检查:"
	@echo "    make lint           运行所有 linter"
	@echo ""
	@echo "  本地开发:"
	@echo "    make dev-gateway    启动 ANI Gateway（热重载）"
	@echo "    make dev-console    启动 Console 前端（Vite HMR）"

# ── 依赖服务 ──────────────────────────────────────────────────────────────────
COMPOSE = docker compose -f deploy/docker/docker-compose.yml
GO_CACHE_ENV = GOCACHE=/private/tmp/ani-go-build GOMODCACHE=$(CURDIR)/.cache/gomod
BUF_CACHE_ENV = BUF_CACHE_DIR=$(CURDIR)/.cache/buf
GO_PACKAGES = ./pkg/... ./services/ani-gateway/... ./services/model-service/... ./services/task-service/...
TOOLS_BIN = $(CURDIR)/.bin

deps:
	@echo "→ 启动本地依赖服务..."
	$(COMPOSE) up -d
	@echo "→ 等待服务就绪..."
	@$(COMPOSE) exec postgres pg_isready -U ani -t 30 && echo "✅ PostgreSQL 就绪"
	@echo "✅ 所有依赖服务已启动"
	@echo ""
	@echo "  MinIO Console: http://localhost:9001  (ani-admin / ani_dev_password)"
	@echo "  NATS Monitor:  http://localhost:8222"
	@echo "  Milvus:        localhost:19530"

deps-down:
	$(COMPOSE) down

deps-status:
	$(COMPOSE) ps

deps-clean:
	@echo "⚠️  这将删除所有本地开发数据，是否确认？[y/N]" && read ans && [ $${ans:-N} = y ]
	$(COMPOSE) down -v

# ── 代码生成 ──────────────────────────────────────────────────────────────────
tools-proto:
	@echo "→ 安装 Protobuf 生成插件到 .bin..."
	@mkdir -p .bin
	GOBIN=$(TOOLS_BIN) $(GO_CACHE_ENV) go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.33.0
	GOBIN=$(TOOLS_BIN) $(GO_CACHE_ENV) go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.3.0
	GOBIN=$(TOOLS_BIN) $(GO_CACHE_ENV) go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@v2.19.1
	@echo "✅ Protobuf 工具已安装"

gen-api:
	@echo "→ 从 OpenAPI Spec 生成 Go 代码..."
	cd services/ani-gateway && \
		go generate ./...
	@echo "→ 从 OpenAPI Spec 生成 TypeScript SDK..."
	npx openapi-typescript api/openapi/v1.yaml -o frontends/console/src/api/types.ts
	@echo "✅ API 代码生成完成"

gen-proto: tools-proto
	@echo "→ 从 Protobuf 生成 Go gRPC 代码..."
	cd api/proto && PATH="$(TOOLS_BIN):$$PATH" $(BUF_CACHE_ENV) buf generate --template buf.gen.yaml .
	@echo "✅ Protobuf 代码生成完成"

# ── 构建 ──────────────────────────────────────────────────────────────────────
VERSION ?= $(shell git describe --tags --always --dirty 2>/dev/null || echo "dev")
LDFLAGS = -X main.Version=$(VERSION) -X main.BuildTime=$(shell date -u +%Y%m%dT%H%M%SZ)

build: build-gateway build-model-service build-task-service

build-gateway:
	@echo "→ 构建 ANI Gateway..."
	@mkdir -p bin
	cd services/ani-gateway && \
		$(GO_CACHE_ENV) CGO_ENABLED=0 go build -ldflags "$(LDFLAGS)" -o ../../bin/ani-gateway .
	@echo "✅ bin/ani-gateway"

build-model-service:
	@echo "→ 构建 Model Service..."
	@mkdir -p bin
	cd services/model-service && \
		$(GO_CACHE_ENV) CGO_ENABLED=0 go build -ldflags "$(LDFLAGS)" -o ../../bin/model-service .
	@echo "✅ bin/model-service"

build-task-service:
	@echo "→ 构建 Task Service..."
	@mkdir -p bin
	cd services/task-service && \
		$(GO_CACHE_ENV) CGO_ENABLED=0 go build -ldflags "$(LDFLAGS)" -o ../../bin/task-service .
	@echo "✅ bin/task-service"

# ── 测试 ──────────────────────────────────────────────────────────────────────
test: test-go test-python

test-go:
	@echo "→ 运行 Go 单元测试..."
	$(GO_CACHE_ENV) go test $(GO_PACKAGES) -v -timeout 120s

test-python:
	@echo "→ 检查 Python 服务语法（RAG）..."
	python -m compileall -q ai/rag-engine

test-cover:
	@echo "→ 生成 Go 测试覆盖率报告..."
	$(GO_CACHE_ENV) go test $(GO_PACKAGES) -coverprofile=coverage.out -covermode=atomic
	$(GO_CACHE_ENV) go tool cover -html=coverage.out -o coverage.html
	@echo "✅ coverage.html"

# ── 代码检查 ──────────────────────────────────────────────────────────────────
lint: lint-go lint-python lint-ts

lint-go:
	@echo "→ Go 静态分析..."
	golangci-lint run ./...
	gosec -severity high ./services/... ./operators/...

lint-python:
	@echo "→ Python 代码检查..."
	cd ai && ruff check . && mypy .

lint-ts:
	@echo "→ TypeScript 代码检查..."
	cd frontends && pnpm lint

# ── 本地开发（热重载）────────────────────────────────────────────────────────
dev-gateway:
	@echo "→ 启动 ANI Gateway（air 热重载）..."
	cd services/ani-gateway && air

dev-console:
	@echo "→ 启动 Console 前端（Vite HMR）..."
	cd frontends/console && pnpm dev

# ── 镜像构建 ──────────────────────────────────────────────────────────────────
REGISTRY ?= harbor.ani.internal/ani

image: image-gateway

image-gateway:
	docker build -t $(REGISTRY)/ani-gateway:$(VERSION) \
		-f services/ani-gateway/Dockerfile \
		services/ani-gateway

# ── 清理 ──────────────────────────────────────────────────────────────────────
clean:
	rm -rf bin/
	find . -name '*.test' -delete
	find . -name 'coverage.out' -delete
