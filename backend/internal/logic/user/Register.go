package user

import (
	"backend/conf/mysql"
	userPublic "backend/internal/handlers"
	"backend/internal/models"
	"backend/pkg/encryption"
	"strings"

	"github.com/pkg/errors"
	"golang.org/x/net/context"
)

func RegisterService(username, password, confirmedPassword string) (*userPublic.RegisterResponse, error) {
	// TODO: 写服务端的注册代码

	var resp userPublic.RegisterResponse
	ctx := context.Background()

	// 根据username在查找数据库中查找，username是否重复
	var User = mysql.Q.User
	user, err := User.WithContext(ctx).Where(User.Username.Eq(username)).Limit(1).Find()
	if err != nil {
		return &resp, err
	}
	if len(user) != 0 {
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
