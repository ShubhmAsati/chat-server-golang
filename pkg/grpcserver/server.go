package grpcserver

import(
	"net"
	"fmt"
	"google.golang.org/grpc"
	grpcMessage "chatbot.com/internal/grpc/service/message"
	grpcUser "chatbot.com/internal/grpc/service/user"
	"chatbot.com/pkg/message"
	"chatbot.com/pkg/user"
)
const(
	port = ":7000"
	hostpath = "localhost"
)

//Server struct for server
type Server struct{
	CleanUpChan chan int
	GrpcServer *grpc.Server
}

//GrpcServer grpc server for storing server details
var GrpcServer *Server

//StartServer starting grpc server
func StartServer(cleanUpChan chan int){
	GrpcServer = &Server{
      CleanUpChan:cleanUpChan ,
	  GrpcServer: grpc.NewServer(),
	}
	registerGrpcServices(GrpcServer.GrpcServer)
	if err := GrpcServer.GrpcServer.Serve(getListner(port)); err != nil {
		panic(err)
	}
}

func registerGrpcServices(grpcServer *grpc.Server){
	userService := user.NewUsers()
	messageService := message.GetMessageClient()
	grpcUser.RegisterUserServiceServer(grpcServer,userService)
	grpcMessage.RegisterMessageServer(grpcServer,messageService)
	fmt.Println("service registered")
}

func getListner(port string) net.Listener {
	lis, err := net.Listen("tcp", port)
	if err != nil {

		panic(err)
	}
	fmt.Println("listener created ")
	return lis
}

func CleanUpServer(cleanUpChan chan int){
	user.SaveUsersDataIntoFile()
    cleanUpChan<-1
}