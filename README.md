# go-GRPC
Basic of Go GRPC 
step 1:
    #Install Go by downloading the tarball, extracting it, and adding the Go binary directory to your PATH.
    #Verify the installation by checking the Go version.
    go installtion:
    wget https://go.dev/dl/go1.22.2.linux-amd64.tar.gz
    rm -rf /usr/local/go && tar -C /usr/local -xzf go1.22.2.linux-amd64.tar.gz
    export PATH=$PATH:/usr/local/go/bin
    go version

step 2:
    #Install protoc-gen-go and protoc-gen-go-grpc using the go install command.
    #Add the Go binary directory to your PATH to make the installed binaries accessible.
    #Verify the installation by checking the versions of protoc-gen-go and protoc-gen-go-grpc.
    go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
    go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
    vagrant@ubuntu-focal:~/GRPC$ export PATH=$PATH:$(go env GOPATH)/bin
    vagrant@ubuntu-focal:~/GRPC$ which protoc-gen-go
    /home/vagrant/go/bin/protoc-gen-go
    vagrant@ubuntu-focal:~/GRPC$ which protoc-gen-go-grpc
    /home/vagrant/go/bin/protoc-gen-go-grpc
    vagrant@ubuntu-focal:~/GRPC$ protoc-gen-go-grpc --version
    protoc-gen-go-grpc 1.2.0
    vagrant@ubuntu-focal:~/GRPC$ protoc-gen-go --version
    protoc-gen-go v1.33.0
step 3:
    #Define your protocol buffer file hello.proto.
step 4:
    #Use protoc to generate Go code from your protocol buffer file.
    #The --go_out flag generates Go code for protocol buffer messages, and the --go-grpc_out flag generates Go code for gRPC service interfaces.
    #Verify that the code generation succeeded by checking for the generated Go files and directories.
    vagrant@ubuntu-focal:~/GRPC$protoc -I=. --go_out=. --go-grpc_out=. hello.proto
    vagrant@ubuntu-focal:~/GRPC$ ls
    grpc-go  hello.proto  veltris.com[contains auto generated files]
step 5:
    #Proceed with creating your gRPC server and client code based on the generated Go files.
    #Test your server and client to ensure they are functioning as expected
    docker build -t shubhamrathore777/my-grpc-client:v1 .
    docker run --network host 1c075bee517e
output client:
vagrant@ubuntu-focal:~/GRPC/go-GRPC/GRPC$ docker run --network host 1c075bee517e
2024/04/05 16:20:06 sent req with name shubham hold for 20sec waiting to come resp from server
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
output server:
vagrant@ubuntu-focal:~/GRPC/server$ docker run --network host 7b43c17bf87f
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
