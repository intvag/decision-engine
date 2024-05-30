GENERATED_DIR ?= service
GENERATED_FILES := $(GENERATED_DIR)/decisions.pb.go \
		   $(GENERATED_DIR)/decisions_grpc.pb.go
BINARY := decision-engine


$(BINARY): $(GENERATED_FILES) *.go go.mod go.sum
	CGO_ENABLED=0 go build -ldflags="-s -w" -trimpath -o $@
	-upx $@

$(GENERATED_DIR)/%_grpc.pb.go $(GENERATED_DIR)/%.pb.go: proto/%.proto | $(GENERATED_DIR)
	protoc -I proto/ $< --go_out=module=github.com/intvag/decision-engine/service:$(GENERATED_DIR) --go-grpc_out=module=github.com/intvag/decision-engine/service:$(GENERATED_DIR)

$(GENERATED_DIR):
	mkdir -p $@
