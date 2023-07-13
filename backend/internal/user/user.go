package user

import "github.com/gin-gonic/gin"

type Service interface {
	GetInfoService(ctx *gin.Context, req *GetUserInfoRequest) (*GetUserInfoResponse, error)
	LoginService(ctx *gin.Context, req *LoginRequest) (*LoginResponse, error)
	RegisterService(ctx *gin.Context, req *RegisterRequest) (*RegisterResponse, error)
}

type GetUserInfoRequest struct {
	UserId string `form:"user_id" query:"user_id" json:"user_id"`
}

type GetUserInfoResponse struct {
	Message  string `json:"message"`
	UserId   int32  `json:"user_id"`
	Username string `json:"username"`
	Photo    string `json:"photo"`
}

type LoginRequest struct {
	UserName string `json:"username" binding:"required" form:"username"`
	PassWord string `json:"password" binding:"required" form:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

type RegisterRequest struct {
	UserName          string `form:"username" json:"username" binding:"required"`
	PassWord          string `form:"password" json:"password" binding:"required"`
	ConfirmedPassword string `form:"confirmed_password" json:"confirmed_password" binding:"required"`
}

type RegisterResponse struct {
	Message string `json:"message"`
}
