package handlers

// addBot
type AddBotReq struct {
	Title       string `json:"title" binding:"required" form:"title"`
	Code        string `json:"code" binding:"required" form:"code"`
	Description string `json:"description" binding:"required" form:"description"`
}

type AddBotResp struct {
	Message string `json:"message"`
}

// getBot
type GetListBotReq struct {
	UserId string `json:"user_id" binding:"required" form:"user_id"`
}

type GetListBotResp struct {
	BotList []*ResultBot `json:"botList"`
}

type ResultBot struct {
	ID          int32  `json:"id"`
	UserID      int32  `json:"user_id"`
	Title       string `json:"title"`
	Description string `son:"description"`
	Code        string `json:"code"`
	CreateTime  string `json:"create_time"`
	ModifyTime  string `json:"modify_time"`
}

// removeBot
type RemoveBotReq struct {
	BotId string `json:"bot_id" binding:"required" form:"bot_id"`
}

type RemoveBotResp struct {
	Message string `json:"message"`
}

// updateBot
type UpdateBotReq struct {
	BotId       string `form:"bot_id" binding:"required" json:"bot_id"`
	Title       string `json:"title" binding:"required" form:"title"`
	Code        string `json:"code" binding:"required" form:"code"`
	Description string `json:"description" binding:"required" form:"description"`
}

type UpdateBotResp struct {
	Message string `json:"message"`
}
