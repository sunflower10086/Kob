package consumer

import "encoding/json"

type Opt func(resp *RespForWeb)

// RespForWeb 用户向前端返回消息，一个函数式编程，隐藏了操作，提高代码的复用性
type RespForWeb struct {
	Event            string      `json:"event"`
	OpponentUsername string      `json:"opponent_username"`
	OpponentPhoto    string      `json:"opponent_photo"`
	Game             interface{} `json:"game"`
}

func (resp *RespForWeb) SendMsg(c *Client) error {
	marshal, err := json.Marshal(resp)
	if err != nil {
		return err
	}

	c.Send <- marshal

	return nil
}

func NewRespForWeb(opts ...Opt) *RespForWeb {
	resp := &RespForWeb{}
	for _, opt := range opts {
		opt(resp)
	}

	return resp
}

func WithEvent(event string) Opt {
	return func(resp *RespForWeb) {
		resp.Event = event
	}
}

func WithOpponentUsername(username string) Opt {
	return func(resp *RespForWeb) {
		resp.OpponentUsername = username
	}
}

func WithOpponentPhoto(photo string) Opt {
	return func(resp *RespForWeb) {
		resp.OpponentPhoto = photo
	}
}

func WithGame(game interface{}) Opt {
	return func(resp *RespForWeb) {
		resp.Game = game
	}
}
