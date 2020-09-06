package utils

import(
	"sync"
)
type MyMessage struct{
	 Msg string
	 Sender int32
}
type Notification struct {
	userNotification map[string] chan MyMessage
	mu         sync.Mutex 
}

var n *Notification

func init(){
	n = &Notification{
		userNotification: make(map[string]chan MyMessage,0),
	}
}

func AddUser(userId string){
	n.mu.Lock()
	//setting default buffer size to 10 for a user
	n.userNotification[userId] = make(chan MyMessage,10)
    n.mu.Unlock()
}

func AddMessage(m MyMessage,userId string){
	n.mu.Lock()
	//setting default buffer size to 10 for a user
	n.userNotification[userId]<-m
    n.mu.Unlock()
}
