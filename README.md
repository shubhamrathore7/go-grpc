Steps to Setup Go GRPC
Step 1: Install Go
Download the Go tarball, extract it, and add the Go binary directory to your PATH. Verify the installation by checking the Go version.
# Install Go
wget https://go.dev/dl/go1.22.2.linux-amd64.tar.gz
rm -rf /usr/local/go && tar -C /usr/local -xzf go1.22.2.linux-amd64.tar.gz
export PATH=$PATH:/usr/local/go/bin
go version

Step 2: Install protoc-gen-go and protoc-gen-go-grpc
Install the protoc-gen-go and protoc-gen-go-grpc using the go install command. Add the Go binary directory to your PATH. Verify the installation by checking the versions of protoc-gen-go and protoc-gen-go-grpc.

# Install protoc-gen-go and protoc-gen-go-grpc
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2

# Add Go binary directory to PATH
export PATH=$PATH:$(go env GOPATH)/bin

# Verify installation
which protoc-gen-go
which protoc-gen-go-grpc
protoc-gen-go-grpc --version
protoc-gen-go --version

Step 3: Define your protocol buffer file (hello.proto).

Step 4: Generate Go code from your protocol buffer file
Use protoc to generate Go code from your protocol buffer file. The --go_out flag generates Go code for protocol buffer messages, and the --go-grpc_out flag generates Go code for gRPC service interfaces.
protoc -I=. --go_out=. --go-grpc_out=. hello.proto
ls

Step 5: Create your gRPC server and client code
Proceed with creating your gRPC server and client code based on the generated Go files.

Step 6: Test your server and client
Test your server and client to ensure they are functioning as expected.
# Build Docker images for server and client
docker build -t shubhamrathore777/my-grpc-client:v1 .
docker build -t shubhamrathore777/my-grpc-server:v1 .

# Run the client container
docker run --network host 1c075bee517e

# Run the server container
docker run --network host 7b43c17bf87f

Output
Client
2024/04/05 16:20:06 Response from server: Hello other2
2024/04/05 16:20:06 Response from server: Hello other4
2024/04/05 16:20:06 Response from server: Hello other8
2024/04/05 16:20:06 Response from server: Hello other3
2024/04/05 16:20:06 Response from server: Hello other0
2024/04/05 16:20:06 Response from server: Hello other7
2024/04/05 16:20:06 Response from server: Hello other5
2024/04/05 16:20:06 Response from server: Hello other6
2024/04/05 16:20:06 Result of addition: 30
2024/04/05 16:20:06 Response from server: Hello other1
2024/04/05 16:20:26 Response from server: Hello shubham

Server
2024/04/05 16:26:01 Server is running on port 50053...
2024/04/05 16:26:07 Received: other2
2024/04/05 16:26:07 Received: other3
2024/04/05 16:26:07 Received: other4
2024/04/05 16:26:07 Received: other5
2024/04/05 16:26:07 Received: other6
2024/04/05 16:26:07 Received: other7
2024/04/05 16:26:07 Received: other8
2024/04/05 16:26:07 Received: shubham
2024/04/05 16:26:07 Received: other1
2024/04/05 16:26:07 Received: other0
