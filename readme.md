go version 1.13

protobuf version 3.13.0


go lanf command to generate grpc files
protoc -I=d:\chatbot\server\internal\proto_files --go_out=d:\chatbot\server\internal\grpc\service d:\chatbot\server\internal\proto_files\user\user_service.proto
protoc -I=d:\chatbot\server\internal\proto_files --go_grpc_out=d:\chatbot\server\internal\grpc\service d:\chatbot\server\internal\proto_files\user\user_service.proto


go installation
wget https://dl.google.com/go/go1.13.4.linux-amd64.tar.gz
sudo tar -C /usr/local -xzf go1.13.4.linux-amd64.tar.gz
export PATH=$PATH:/usr/local/go/bin
git clone .... git path
go mod download
