### grpc


# Prerequisites

[Protocol Buffers](https://github.com/protocolbuffers/protobuf/releases)

# Usage

```bash
f7k grpc-build --verbose
```

This command above read the conf/grpc.yaml file and build a full grpc server from .proto files.

Once built, you can  
run the grpc server with :
```bash
go run grpc/server/server.go --verbose
```

then run the rest api with grpc-gateway :
```bash
go run rest/main.go --verbose
```
