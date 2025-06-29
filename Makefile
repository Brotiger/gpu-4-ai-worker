PROTO_DIR=proto
PROTO_FILE=$(PROTO_DIR)/worker.proto

proto:
	protoc --go_out=paths=source_relative:$(PROTO_DIR) --go-grpc_out=paths=source_relative:$(PROTO_DIR) $(PROTO_FILE) 