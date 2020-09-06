package user

import (
	"context"
	"encoding/json"
	"fmt"

	"io/ioutil"
	"sync"
	"time"

	"log"

	"chatbot.com/internal/grpc/service/user"
	"github.com/google/uuid"
)

const (
	USER_DATE_FILE_PATH = "server_data/users_data.json"
)

type User struct {
	Name       string                 `json:"name"`
	UserId     uint32                 `json:"userId"`
	ProfilePic string                 `json:"profilePic"`
	MobileNo   string                 `json:"mobileNo"`
	Status     string                 `json:"status"`
	Metadata   map[string]interface{} `json:"metadata"`
}

//Users implementation
type Users struct {
	//map to MobileNoToId
	MobileNoToId map[string]uint32 `json:"mobileNoToId"`
	//map to userId to userdata
	UsersData map[uint32]*User `json:"usersData"`
	mu        sync.Mutex
}

var users *Users

var once sync.Once

func init() {
	log.Print("reading users.json file")
	NewUsers()
	file, err := ioutil.ReadFile(USER_DATE_FILE_PATH)
	if err != nil {
		fmt.Println(err)
		log.Printf("%+v", err)
		log.Printf("error while reading users.json")
		return
	}
	err = json.Unmarshal([]byte(file), users)
	if err != nil {
		fmt.Println(err)
		log.Printf("error reading initial users")
	}
	fmt.Println(users)
	//go printUsers()
}


//NewUsers get users struct
func NewUsers() *Users {
	once.Do(func() {
		fmt.Println("this should be executed only once")
		users = &Users{
			MobileNoToId: make(map[string]uint32, 0),
			UsersData:    make(map[uint32]*User, 0),
		}
	})
	return users
}

//Add a new user
func (u *Users) Add(ctx context.Context, in *user_grpc.User) (*user_grpc.AddUserResponse, error) {
	fmt.Println(in)
	fmt.Print(in.GetName())
	userID := uuid.New().ID()
	fmt.Println("user added ", userID)
	//return userID if same mobile no tries to re-register
	if userID,ok := u.MobileNoToId[in.MobileNo];ok{
		return &user_grpc.AddUserResponse{
			Id:userID,
		},nil
	}
	u.mu.Lock()
	u.MobileNoToId[in.MobileNo] = userID
	u.UsersData[userID] = &User{
		Name:       in.Name,
		MobileNo:   in.MobileNo,
		ProfilePic: in.ProfilePic,
		UserId:     userID,
	}
	u.mu.Unlock()
	return &user_grpc.AddUserResponse{
		Id: userID,
	}, nil
}

//GetUserByMobile gets user details from mobile no
func (u *Users) GetUserByMobileNos(ctx context.Context, in *user_grpc.GetUserByMobileNosRequest) (*user_grpc.GetUserByMobileNosResponse, error) {
	response := &user_grpc.GetUserByMobileNosResponse{
		User: make([]*user_grpc.User, 0),
	}
	fmt.Println(in)
	for _, mobileNo := range in.MobileNos {
		if userID, ok := u.MobileNoToId[mobileNo]; ok {
			if userData, ok := u.UsersData[userID]; ok {
				user := &user_grpc.User{
					Name:       userData.Name,
					ProfilePic: userData.ProfilePic,
					Status:     userData.Status,
					UserId:     userID,
					MobileNo:   mobileNo,
				}
				response.User = append(response.User, user)
			} else {
				log.Printf("cannot find user details for mobile %s", mobileNo)
			}
		} else {
			log.Printf("Cannot find mobile no %s", mobileNo)
		}
	}
	return response, nil
}

//UploadProfilePic upload user profile pic
func (u *Users) UploadProfilePic(ctx context.Context, in *user_grpc.UploadProfilePicRequest) (*user_grpc.StatusOkResponse, error) {
	fmt.Println("uploading profile pic")
	return &user_grpc.StatusOkResponse{}, nil
}

//SetMyStatus set my status to online/offline
func (u *Users) SetMyStatus(ctx context.Context, in *user_grpc.UserStatus) (*user_grpc.StatusOkResponse, error) {
	fmt.Println("setting my status")
	return &user_grpc.StatusOkResponse{}, nil
}

//IamTyping captures the typing event
func (u *Users) IamTyping(ctx context.Context, in *user_grpc.Typing) (*user_grpc.StatusOkResponse, error) {
	fmt.Println("iam typing ")
	return &user_grpc.StatusOkResponse{}, nil
}

//LoadMyMessages load a users messages which are not delivered
func (u *Users) LoadMyMessages(ctx context.Context, in *user_grpc.User) (*user_grpc.MyMessages, error) {
	fmt.Println("loading my messages")
	return &user_grpc.MyMessages{}, nil
}

func (u *Users) MustEmbedUnimplementedUserServiceServer() {}


//SaveUsersDataIntoFile saves the inmemory user data into a file
func SaveUsersDataIntoFile() {
	fmt.Println("save data into file")
	fmt.Println(users)
	usersData, err := json.MarshalIndent(users, "", " ")
	if err != nil {
		log.Printf("unmarshal error data")
		return
	}
	fmt.Println(string(usersData))
	err = ioutil.WriteFile(USER_DATE_FILE_PATH, usersData, 0644)
	if err != nil {
		log.Printf("error while saving users data")
		return
	}

}

func printUsers() {
	for {
		time.Sleep(1 * time.Second)
		fmt.Println(users)
	}
}

func (u *Users)GetUserByUserId(context context.Context, in *user_grpc.GetUserByUserIdRequest)(*user_grpc.User,error){
   fmt.Println(in)
   if userData,ok := u.UsersData[in.GetUserId()];ok{
    return &user_grpc.User{
	  UserId: userData.UserId,
	  MobileNo: userData.MobileNo,
	  Name: userData.Name,
	  ProfilePic: userData.ProfilePic,
	  Status: userData.Status,
	},nil
   }
	   return nil, fmt.Errorf("The user %d does not exsits",in.GetUserId)
}
