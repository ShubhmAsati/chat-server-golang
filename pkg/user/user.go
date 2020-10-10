package user

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"io/ioutil"
	"sync"
	"time"

	"chatbot.com/pkg/logger"

	"errors"

	"chatbot.com/internal/grpc/service/user"
	"github.com/google/uuid"
)

const (
	USER_DATE_FILE_PATH = "server_data/users_data.json"
	INIT_CONTEXT        = "users-init"
	USER_CONTEXT        = "users"
)

type User struct {
	FirstName  string `json:"firstName"`
	UserName   string `json:"userName"`
	UserId     string `json:"userId"`
	ProfilePic string `json:"profilePic"`

	MobileNo string                 `json:"mobileNo"`
	Status   string                 `json:"status"`
	Metadata map[string]interface{} `json:"metadata"`
}

//Users implementation
type Users struct {
	//map to MobileNoToId
	MobileNoToId map[string]string `json:"mobileNoToId"`
	//map to userId to userdata
	UsersData map[string]*User `json:"usersData"`
	mu        sync.Mutex
}

var users *Users

var once sync.Once

func init() {
	logger.Log(INIT_CONTEXT).Debugf("Loading existing users")
	NewUsers()
	file, err := ioutil.ReadFile(USER_DATE_FILE_PATH)
	if err != nil {
		logger.Log(INIT_CONTEXT).Errorf("Error while reading '%s' file ", USER_DATE_FILE_PATH)
		return
	}
	err = json.Unmarshal([]byte(file), users)
	if err != nil {
		logger.Log(INIT_CONTEXT).Warnf("Invalid user data present in '%s'", USER_DATE_FILE_PATH)
	}
	logger.Log(INIT_CONTEXT).Debugf("Loaded these user details %+v", users)
}

//NewUsers get users struct
func NewUsers() *Users {
	once.Do(func() {
		fmt.Println("this should be executed only once")
		users = &Users{
			MobileNoToId: make(map[string]string, 0),
			UsersData:    make(map[string]*User, 0),
		}
	})
	return users
}

//Add a new user
func (u *Users) Add(ctx context.Context, in *user_grpc.User) (*user_grpc.AddUserResponse, error) {
	logger.Log(USER_CONTEXT).Debugf("Starting request to add user %+v", in)

	validationErr := validateAddOrUpdateUser(in)
	if validationErr != nil {
		return nil, validationErr
	}
	//return userID if same mobile no tries to re-register
	if userID, ok := u.MobileNoToId[in.MobileNo]; ok {
		logger.Log(USER_CONTEXT).Debugf("User '%s' already exists hence not creating a new user", in.UserName)
		return &user_grpc.AddUserResponse{
			UserId: userID,
		}, nil
	}
	
	for _, users := range u.UsersData {
		if users.UserName == in.UserName {
			return nil, fmt.Errorf("Username already taken")
		}
	}
	
	userID := fmt.Sprintf("%d", uuid.New().ID())
	u.mu.Lock()
	u.MobileNoToId[in.MobileNo] = userID
	u.UsersData[userID] = &User{
		FirstName:  in.FirstName,
		MobileNo:   in.MobileNo,
		ProfilePic: in.ProfilePic,
		UserId:     userID,
		UserName:   in.UserName,
		Status:     in.Status,
	}
	u.mu.Unlock()
	logger.Log(USER_CONTEXT).Debugf("Successfully added user %+v with userId %s", in, userID)
	return &user_grpc.AddUserResponse{
		UserId: userID,
	}, nil
}

func validateAddOrUpdateUser(in *user_grpc.User) error {
	if len(in.FirstName) < 2 {
		return errors.New("Invalid name")
	}
	if len(in.UserName) < 2 {
		return errors.New("Invalid username")
	}
	if len(in.MobileNo) < 10 {
		return errors.New("Invalid mobile number")
	}

	return nil
}

//Update update user details like name and status
func (u *Users) Update(ctx context.Context, in *user_grpc.User) (*user_grpc.StatusOkResponse, error) {
	logger.Log("UPDATE").Debugf("Updating user details %+v", in)
	u.mu.Lock()
	defer u.mu.Unlock()
	if userData, ok := u.UsersData[in.UserId]; ok {
		if in.FirstName != "" {
			userData.FirstName = in.FirstName
		}
		if in.Status != "" {
			userData.Status = in.Status
		}
		if in.MobileNo != "" {
			userData.MobileNo = in.MobileNo
		}
	} else {
		return nil, fmt.Errorf("User not found %s", in.UserId)
	}
	return &user_grpc.StatusOkResponse{
		Message:     "Success fully updated user",
		SuccessCode: 1,
	}, nil
}

//UploadProfilePic upload user profile pic
func (u *Users) UploadProfilePic(ctx context.Context, in *user_grpc.UploadProfilePicRequest) (*user_grpc.StatusOkResponse, error) {
	fmt.Println("uploading profile pic")
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
	logger.Log("SHUTDOWNTASK").Debugf("Saving users data into file")
	usersData, err := json.MarshalIndent(users, "", " ")
	if err != nil {
		logger.Log("SHUTDOWNTASK").Errorf("error while marshalling user data %+v",err)
		return
	}
	err = ioutil.WriteFile(USER_DATE_FILE_PATH, usersData, 0644)
	if err != nil {
		logger.Log("SHUTDOWNTASK").Errorf("error while saving user data %+v",err)
		return
	}

}

func printUsers() {
	for {
		time.Sleep(1 * time.Second)
		fmt.Println(users)
	}
}

//GetUserByUserId fetch user by userid
func (u *Users) GetUserByUserId(context context.Context, in *user_grpc.GetUserByUserIdRequest) (*user_grpc.User, error) {
	fmt.Println(in)
	if userData, ok := u.UsersData[in.GetUserId()]; ok {
		return &user_grpc.User{
			UserId:     userData.UserId,
			MobileNo:   userData.MobileNo,
			FirstName:  userData.FirstName,
			ProfilePic: userData.ProfilePic,
			Status:     userData.Status,
			UserName:   userData.UserName,
		}, nil
	}
	return nil, fmt.Errorf("The user %d does not exsits", in.GetUserId)
}

//SearchUsersByUserName regex search by user name
func (u *Users) SearchUsersByUserName(context context.Context, in *user_grpc.GetUserByUserNameRequest) (*user_grpc.SearchUsersByUserNameResponse, error) {
	response := &user_grpc.SearchUsersByUserNameResponse{
		User: make([]*user_grpc.User, 0),
	}
	for _, userData := range u.UsersData {
		if strings.Contains(userData.UserName, in.UserName) {
			respUser := &user_grpc.User{
				UserId:     userData.UserId,
				MobileNo:   userData.MobileNo,
				FirstName:  userData.FirstName,
				ProfilePic: userData.ProfilePic,
				Status:     userData.Status,
				UserName:   userData.UserName,
			}
			response.User = append(response.User, respUser)
		}
	}
	return response, nil
}

//GetExistingUserNames returns a list of user names
func (u *Users) GetExistingUserNames(context context.Context, in *user_grpc.GetExistingUserNamesRequest) (*user_grpc.GetExistingUserNamesResponse, error) {
	logger.Log("GET").Debugf("Getting existing user names")
	response := &user_grpc.GetExistingUserNamesResponse{
		UserName: make([]*user_grpc.UserName, 0),
	}
	for _, userData := range u.UsersData {
		userName := &user_grpc.UserName{
			UserName: userData.UserName,
		}
		response.UserName = append(response.UserName, userName)
	}
	return response, nil
}
