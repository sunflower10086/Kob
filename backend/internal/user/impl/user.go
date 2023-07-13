package impl

import (
	"backend/conf/mysql"
	"backend/conf/redis"
	"backend/internal"
	"backend/internal/models"
	"backend/internal/user"
	"backend/pkg/encryption"
	"backend/pkg/ioc"
	"backend/pkg/util"
	"encoding/json"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	impl = &UserServiceImpl{}
)

func init() {
	ioc.RegistryImpl(impl)
}

type UserServiceImpl struct {
	l *log.Logger
}

func (h *UserServiceImpl) Name() string {
	return internal.UserServiceName
}

func (h *UserServiceImpl) Config() {
	h.l = log.New(os.Stderr, "  [user] ", log.Ldate|log.Ltime|log.Lshortfile)
}

func NewUserServiceImpl() *UserServiceImpl {
	return &UserServiceImpl{
		l: log.New(os.Stderr, "  [user] ", log.Ldate|log.Ltime|log.Lshortfile),
	}
}

func (h *UserServiceImpl) GetInfoService(ctx *gin.Context, req *user.GetUserInfoRequest) (*user.GetUserInfoResponse, error) {
	var resp user.GetUserInfoResponse
	strUserId := req.UserId
	intUserId, _ := strconv.Atoi(strUserId)

	// 1.先去redis中查，查到直接返回
	userRedis, _ := redis.RDB.Get(ctx, "cache:kob:user:"+strUserId).Result()
	if len(userRedis) != 0 {
		var redisUser models.User

		err := json.Unmarshal([]byte(userRedis), &redisUser)
		if err != nil {
			return &resp, err
		}

		resp.UserId = redisUser.ID
		resp.Username = redisUser.Username
		resp.Photo = redisUser.Photo
		return &resp, nil
	}

	// 2.不存在先去数据库中查，查到之后写入redis
	var User = mysql.Q.User
	sqlUser, err := User.WithContext(ctx).Where(User.ID.Eq(int32(intUserId))).Limit(1).Find()
	if err != nil {
		return &resp, err
	}

	// 3.存入redis
	go func() {
		userJson, err := json.Marshal(sqlUser[0])
		if err != nil {
			return
		}
		redis.RDB.Set(ctx, "cache:kob:user:"+strUserId, userJson, time.Second*60*60*24)
	}()

	resp.UserId = sqlUser[0].ID
	resp.Username = sqlUser[0].Username
	resp.Photo = sqlUser[0].Photo

	return &resp, err
}

func (h *UserServiceImpl) LoginService(ctx *gin.Context, req *user.LoginRequest) (*user.LoginResponse, error) {
	var resp user.LoginResponse
	username, password := req.UserName, req.PassWord

	//根据username去数据库查寻信息，有用户名相同的用户直接返回
	var User = mysql.Q.User
	sqlUser, err := User.WithContext(ctx).Where(User.Username.Eq(username)).Limit(1).First()
	if errors.As(err, &gorm.ErrRecordNotFound) { // 没有查找到用户名为username的用户
		zap.L().Error(err.Error())
		return &resp, err
	}

	//查到此人判断密码加密后和数据库的密码是否相同，不相同直接返回错误，
	if encryption.EqualPassword(password, sqlUser.Password) {
		return &resp, errors.New("密码错误")
	}

	//如果相同则颁发token
	token, err := util.CreateToken(int(sqlUser.ID))
	if err != nil {
		zap.L().Error(err.Error())
		return &resp, err
	}

	resp.Token = token
	//返回token
	return &resp, err
}

func (h *UserServiceImpl) RegisterService(ctx *gin.Context, req *user.RegisterRequest) (*user.RegisterResponse, error) {
	var resp user.RegisterResponse

	username, password, confirmedPassword := req.UserName, req.PassWord, req.ConfirmedPassword

	// 根据username在查找数据库中查找，username是否重复
	var User = mysql.Q.User
	sqlUser, err := User.WithContext(ctx).Where(User.Username.Eq(username)).Limit(1).Find()
	if err != nil {
		return &resp, err
	}
	if len(sqlUser) != 0 {
		err = errors.New("该用户已存在")
		return &resp, err
	}

	// 用户名不能为空
	if len(username) == 0 || len(strings.Trim(username, " ")) == 0 {
		err = errors.New("用户名不能为空")
		return &resp, err
	}

	// 判断两个密码是否相等
	if password != confirmedPassword {
		err = errors.New("两次密码不一致")
		return &resp, err
	}

	// 密码不能为空
	if len(password) == 0 || len(strings.Trim(password, " ")) == 0 {
		err = errors.New("密码不能为空")
		return &resp, err
	}

	// 密码长度超过1000
	if len(password) >= 1000 {
		err = errors.New("密码长度超过1000")
		return &resp, err
	}

	// 用户名长度超过1000
	if len(username) >= 1000 {
		err = errors.New("用户名长度超过1000")
		return &resp, err
	}

	encodePassword := encryption.PasswordBcrypt(password)
	photo := "https://cdn.acwing.com/media/user/profile/photo/160535_lg_e4534d8e65.jpg"

	registerUser := models.User{
		Username: username,
		Password: encodePassword,
		Photo:    photo,
	}

	err = User.WithContext(ctx).Create(&registerUser)
	if err != nil {
		return &resp, err
	}

	resp.Message = "success"
	return &resp, err
}
