FROM golang:1.21.6
WORKDIR /app
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
COPY . ./
RUN go mod download
RUN go build -o server grpc-server/server.go
RUN go build -o client grpc-client/client.go
EXPOSE 50051
CMD ["./server"]
