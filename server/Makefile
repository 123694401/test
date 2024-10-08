ROOT_DIR    = $(shell pwd)
NAMESPACE   = "default"
DEPLOY_NAME = "template-single"
DOCKER_NAME = "template-single"
ADMIN_RESOURCE_PATH = "/resource/public/admin/"

# 一键编译。编译web前端，并将编译后的包移动到服务端对应静态资源路径下，最后编译服务端
.PHONY: build
build:
	@rm -rf ./$(ADMIN_RESOURCE_PATH)
	@mkdir ./$(ADMIN_RESOURCE_PATH)
	@cd ../web && pnpm run build && \cp -rf ./dist/*  ../server$(ADMIN_RESOURCE_PATH)
	@cd ../server
	@echo "y" | gf build

# 通过热编译启动所有服务
.PHONY: all
all:
	gf run main.go --args "all"

.PHONY: http
http:
	gf run main.go --args "http"

.PHONY: queue
queue:
	gf run main.go --args "queue"

.PHONY: cron
cron:
	gf run main.go --args "cron"

.PHONY: auth
auth:
	gf run main.go --args "auth"

# 启动web服务
.PHONY: web
web:
	@cd ../web && pnpm run dev

# 刷新casbin权限
.PHONY: refresh
refresh:
	@go run main.go tools -m=casbin -a1=refresh

# 清理casbin权限
.PHONY: clear
clear:
	@go run main.go tools -m=casbin -a1=clear

# 运行代码质量分析工具
# https://github.com/ywanbing/golangci
.PHONY: lint
lint:
	golangci-lint run

.PHONY: killmain
killmain:
	@kill -9 $(ps -ef|grep main|grep -v grep|awk '{print $2}')

# Install/Update to the latest CLI tool.
.PHONY: cli
cli:
	@set -e; \
	wget -O gf https://github.com/gogf/gf/releases/latest/download/gf_$(shell go env GOOS)_$(shell go env GOARCH) && \
	chmod +x gf && \
	./gf install -y && \
	rm ./gf

# Check and install CLI tool.
.PHONY: cli.install
cli.install:
	@set -e; \
	gf -v > /dev/null 2>&1 || if [[ "$?" -ne "0" ]]; then \
  		echo "GoFame CLI is not installed, start proceeding auto installation..."; \
		make cli; \
	fi;

# Generate Go files for DAO/DO/Entity.
.PHONY: dao
dao: cli.install
	@gf gen dao

# Generate Go files for Service.
.PHONY: service
service: cli.install
	@gf gen service

# Build image, deploy image and yaml to current kubectl environment and make port forward to local machine.
.PHONY: start
start:
	@set -e; \
	make image; \
	make deploy; \
	make port;

# Build docker image and commit to the repository.
# example: make image tag=v0.0.1
.PHONY: image
image:
	@echo "y" | gf docker main.go -p -tn hotgo:$(tag)

# Deploy image and yaml to current kubectl environment.
.PHONY: deploy
deploy:
	$(eval _TAG = $(if ${TAG},  ${TAG}, develop))
	@set -e; \
	mkdir -p $(ROOT_DIR)/temp/kustomize;\
	cd $(ROOT_DIR)/manifest/deploy/kustomize/overlays/${_TAG};\
	kustomize build > $(ROOT_DIR)/temp/kustomize.yaml;\
	kubectl   apply -f $(ROOT_DIR)/temp/kustomize.yaml; \
	kubectl   patch -n $(NAMESPACE) deployment/$(DEPLOY_NAME) -p "{\"spec\":{\"template\":{\"metadata\":{\"labels\":{\"date\":\"$(shell date +%s)\"}}}}}";
