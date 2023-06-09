package user

import (
	"backend/conf/mysql"
	userPublic "backend/internal/handlers"
	"backend/pkg/encryption"
	"backend/pkg/util"

	"github.com/pkg/errors"
	"go.uber.org/zap"
	"golang.org/x/net/context"
	"gorm.io/gorm"
)

func LoginService(username, password string) (*userPublic.LoginResponse, error) {
	// TODO: 写服务端的登录的代码

	var resp userPublic.LoginResponse
	ctx := context.Background()
	//根据username去数据库查寻信息，有用户名相同的用户直接返回
	var User = mysql.Q.User
	user, err := User.WithContext(ctx).Where(User.Username.Eq(username)).Limit(1).First()
	if errors.As(err, &gorm.ErrRecordNotFound) { // 没有查找到用户名为username的用户
		zap.L().Error(err.Error())
		return &resp, err
	}

	//查到此人判断密码加密后和数据库的密码是否相同，不相同直接返回错误，
	if encryption.EqualPassword(password, user.Password) {
		return &resp, errors.New("密码错误")
	}

	//如果相同则颁发token
	token, err := util.CreateToken(int(user.ID))
	if err != nil {
		zap.L().Error(err.Error())
		return &resp, err
	}

	resp.Token = token
	//返回token
	return &resp, err
}
