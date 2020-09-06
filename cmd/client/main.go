package main

import (
	"fmt"
	"io"
	"time"

	"context"

	grpcMessage "chatbot.com/internal/grpc/service/message"
	grpcUser "chatbot.com/internal/grpc/service/user"
	"google.golang.org/grpc"
)

var conn *grpc.ClientConn

func init(){
	serverAddress := "localhost:7000"
	var err error
	conn,err = grpc.Dial(serverAddress,grpc.WithInsecure())
    if err !=nil{
		fmt.Println(err)
	}
}

func main(){
	defer conn.Close()

    userId,name := getUserIdCtx()
	 
	msgClient := grpcMessage.NewMessageClient(conn)
    userData := &grpcUser.User{
		Name: name,
		UserId: userId,
	}
	stream,err := msgClient.Connect(context.Background(),userData)
	if err != nil{
		fmt.Println(err)
	}
    
	waitc := make(chan struct{})
	fmt.Println(userId,name)
//sendMessageInInterval(userId,msgClient)
	go func(){
		for {
			in,err := stream.Recv()
			if err== io.EOF{
				close(waitc)
				return
			}
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println("recieved")
			fmt.Println(in)
		}
	}()
	stream.CloseSend()
	<-waitc
}

func getUserIdCtx()(id uint32,name string){

	userClient := grpcUser.NewUserServiceClient(conn)
	userData := &grpcUser.User{
		Name: "shubham",
		MobileNo: "1212121",
	}
	 resp,err := userClient.Add(context.Background(),userData)
	 if err!=nil{
		fmt.Println(err)
	}
	fmt.Println(resp)
   return resp.Id,userData.Name

}


func sendMessageInInterval(userId uint32,msgClient grpcMessage.MessageClient){
	for i:=0;i<10;i++{
		go func(userId uint32){
		newMessage := grpcMessage.Msg{
			From: userId,
			To: userId,
			Message: "hello bro",
		}
		resp ,err := msgClient.SendMessage(context.Background(),&newMessage)
		if err != nil{
			fmt.Println(err)
		}
		fmt.Println(resp)
	}(userId)
	time.Sleep(2*time.Second)
	}
}