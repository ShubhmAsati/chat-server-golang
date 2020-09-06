package message

import (
	"context"
	"fmt"
	_ "io"
	"log"
	"sync"

	"chatbot.com/internal/grpc/service/message"
)

//MessageService wrapper for message
type MessageService struct {
	connectInfo map[uint32]ConnectInfo
	mu          sync.Mutex
}

type ConnectInfo struct {
	stream   *message_grpc.Message_ConnectServer
	isActive bool
	userName string
	userId   uint32
	err      chan error
}

var once sync.Once

//NewMessage creates anew message obkect
func GetMessageClient() *MessageService {
	var msg *MessageService
	once.Do(func() {
		msg = &MessageService{
			connectInfo: make(map[uint32]ConnectInfo, 0),
		}
	})
	return msg
}

//Connect creates a stream from server to client for messaging
func (m *MessageService) Connect(in *message_grpc.UserById, stream message_grpc.Message_ConnectServer) error {
	fmt.Println("inside connect")
	m.mu.Lock()
	fmt.Println(in.UserId);
	m.connectInfo[in.UserId] = ConnectInfo{
		stream:   &stream,
		isActive: true,
		userId:   in.UserId,
		err:      make(chan error),
	}
	fmt.Println("connection established")
	m.mu.Unlock()
	return <-m.connectInfo[in.UserId].err
}

//BroadCastMessage sends a message to every user
func (m *MessageService) BroadCastMessage(ctx context.Context, in *message_grpc.UserChatMessage) (*message_grpc.MsgSent, error) {

	fmt.Println("sending message")
	for _, userConn := range m.connectInfo {
		if err := (*userConn.stream).Send(in); err != nil {
			fmt.Println(err)
		}
	}
	return &message_grpc.MsgSent{
		IsSent:           true,
		SentUtcTimeStamp: "121212",
	}, nil
}

//SendMessage sends message to stream
func (m *MessageService) SendMessage(ctx context.Context, in *message_grpc.UserChatMessage) (*message_grpc.MsgSent, error) {
	fmt.Println("sending message")
	toUserId := in.ToUserId
	fmt.Print(toUserId)
	//TODO: validate toId
	fmt.Println("after printing to userid")
	fmt.Println(m.connectInfo[toUserId])
	if userInfo, ok := m.connectInfo[toUserId]; ok {
		if userInfo.isActive {
			if err := (*userInfo.stream).Send(in); err != nil {
				fmt.Println(err)
				fmt.Println("fail to send error")
			} else {
				fmt.Println("message sent successfully")
			}
		} else {
			//user is not active buffer this message
			log.Printf("%+v",in)
			log.Printf("used is not active add this is buffer")
		}
	} else {
		//TODO: we should be able to know if user exist or not
		//at this point the user does not exist of we dont have an active connection
		//buffer this user message , will send this later when connection is established
		log.Printf("%+v",in)
		log.Printf("used is not active add this is buffer")
	}
	return &message_grpc.MsgSent{}, nil
}

//RemoveConnection deleting user connection permanently
func (m *MessageService) RemoveConnection(ctx context.Context, in *message_grpc.UserById) (*message_grpc.StatusOkResponse, error) {
	if _, ok := m.connectInfo[in.UserId]; ok {
		m.mu.Lock()
		delete(m.connectInfo, in.UserId)
		fmt.Println("deleted existing connection")

		fmt.Println("connection established")
		m.mu.Unlock()

	}
	return &message_grpc.StatusOkResponse{}, nil
}

//SendMessageToAll testing function
func (m *MessageService) SendMessageToAll(msg string) {
	fmt.Println("sending message bravo")
	for _, userConn := range m.connectInfo {
		fmt.Println("sending message to ", userConn)
		fmt.Println(userConn.userName)
		newMessage := message_grpc.UserChatMessage{
			Nt: message_grpc.UserChatMessage_MESSAGE,
			IsGroup: false,
			FromUserId: 12121,
			ToUserId: 12121,
		    Msg: &message_grpc.Msg{
				Message: "hello there",
				MsgId: "new mefd",
			},	
		}
		if err := (*userConn.stream).Send(&newMessage); err != nil {
			//user is offline/ not reachable , stream connection broke or something
			fmt.Println(err)
		}
	}
}

func (m *MessageService) MustEmbedUnimplementedMessageServer() {}
