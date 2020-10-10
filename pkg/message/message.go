package message

import (
	"context"
	"fmt"
	_ "io"
	"log"
	"sync"

	"chatbot.com/internal/grpc/service/message"
	"errors"
)

type BufferedMessagesForUser struct {
	Messages map[string][]*message_grpc.UserChatMessage
	mu       sync.Mutex
}

var bufferedMessagesForUser BufferedMessagesForUser

//MessageService wrapper for message
type MessageService struct {
	connectInfo map[string]ConnectInfo
	mu          sync.Mutex
}

type ConnectInfo struct {
	stream   *message_grpc.Message_ConnectServer
	isActive bool
	userName string
	userId   string
	err      chan error
}

var once sync.Once

//NewMessage creates anew message obkect
func GetMessageClient() *MessageService {
	var msg *MessageService
	once.Do(func() {
		msg = &MessageService{
			connectInfo: make(map[string]ConnectInfo, 0),
		}
	})
	return msg
}

func init() {
	bufferedMessagesForUser.Messages = make(map[string][]*message_grpc.UserChatMessage, 0)
}

//Connect creates a stream from server to client for messaging
func (m *MessageService) Connect(in *message_grpc.UserById, stream message_grpc.Message_ConnectServer) error {
	fmt.Println("inside connect")
	m.mu.Lock()
	fmt.Println(in.UserId)
	m.connectInfo[in.UserId] = ConnectInfo{
		stream:   &stream,
		isActive: true,
		userId:   in.UserId,
		err:      make(chan error),
	}
	fmt.Println("connection established")
	m.mu.Unlock()
	m.sendAllBufferedMessagesToUser(in.UserId)
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
	fmt.Println(in.GetChatId())
	if userInfo, ok := m.connectInfo[toUserId]; ok {
		if userInfo.isActive {
			if err := (*userInfo.stream).Send(in); err != nil {
				fmt.Println(err)
				fmt.Println("fail to send error")
				addMessageToBuffer(in)
			} else {
				fmt.Println("message sent successfully")
			}
		} else {
			//user is not active buffer this message
			log.Printf("%+v", in)
			log.Printf("used is not active add this is buffer")
			addMessageToBuffer(in)
		}
	} else {
		//TODO: we should be able to know if user exist or not
		//at this point the user does not exist of we dont have an active connection
		//buffer this user message , will send this later when connection is established
		log.Printf("%+v", in)
		log.Printf("used is not active add this is buffer")
         addMessageToBuffer(in)
	}
	return &message_grpc.MsgSent{}, nil
}

func addMessageToBuffer(in *message_grpc.UserChatMessage) {
	toUserId := in.ToUserId
	if val, ok := bufferedMessagesForUser.Messages[toUserId]; ok {
		if len(val) > 0 {
			bufferedMessagesForUser.Messages[toUserId] = append(bufferedMessagesForUser.Messages[toUserId], in)
		} else {
			bufferedMessagesForUser.Messages[toUserId] = make([]*message_grpc.UserChatMessage, 0)
			bufferedMessagesForUser.Messages[toUserId] = append(bufferedMessagesForUser.Messages[toUserId], in)
		}
	} else {
		bufferedMessagesForUser.Messages[toUserId] = make([]*message_grpc.UserChatMessage, 0)
		bufferedMessagesForUser.Messages[toUserId] = append(bufferedMessagesForUser.Messages[toUserId], in)
	}
}

func (m *MessageService)sendAllBufferedMessagesToUser(userId string){
     if val,ok := bufferedMessagesForUser.Messages[userId];ok{
		 log.Printf("%+v",val)
		 retryBuffer := make([]*message_grpc.UserChatMessage,0)
             for i:=0;i<len(val);i++{
				if err := (*m.connectInfo[userId].stream).Send(val[i]); err != nil {
					fmt.Println(err)
					fmt.Println("fail to send buffered messages ")
					retryBuffer = append(retryBuffer,val[i])
				} else {
					fmt.Println("buffered message sent successfully")
				}
			 }
			 log.Printf("retry buffer %+v for user %s",retryBuffer,userId)
			 bufferedMessagesForUser.Messages[userId] = retryBuffer
	 }else{
		 log.Printf("No buffered message for user %s",userId)
	 }
}

//RemoveConnection deleting user connection permanently
func (m *MessageService) RemoveConnection(ctx context.Context, in *message_grpc.UserById) (*message_grpc.StatusOkResponse, error) {
	if val, ok := m.connectInfo[in.UserId]; ok {
		val.err <- errors.New("removing users connection")
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
			Nt:         message_grpc.UserChatMessage_MESSAGE,
			IsGroup:    false,
			FromUserId: "12121",
			ToUserId:   "12121",
			Msg: &message_grpc.Msg{
				Message: "hello there",
				MsgId:   "new mefd",
			},
		}
		if err := (*userConn.stream).Send(&newMessage); err != nil {
			//user is offline/ not reachable , stream connection broke or something
			fmt.Println(err)
		}
	}
}

func (m *MessageService) MustEmbedUnimplementedMessageServer() {}
