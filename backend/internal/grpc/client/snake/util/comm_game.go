package util

// CommGame 与游戏系统通信的接口，之后无论与哪个游戏通信，只需要实现这个接口即可
type CommGame interface {
	Send() error    // 发送消息
	Receive() error // 接收消息
}
