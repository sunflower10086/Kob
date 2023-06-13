package share_space

// Space 游戏系统与webSocket的共享空间，游戏系统向其中传入数据，WebSocket从其中取出数据
type Space struct {
	Game             chan *SnakeGame // 游戏信息
	ClientDirection  chan Pair       // 前端传来的下一步
	ServiceDirection chan Pair       // 后端传来的下一步
	Result           chan Result     // 游戏结果
}

// NewSpace New一个新的公共空间
func NewSpace() *Space {
	return &Space{
		Game:             make(chan *SnakeGame, 100), // 同时可以有一百盘游戏
		ClientDirection:  make(chan Pair, 10000),     // 前端最多传10000步
		ServiceDirection: make(chan Pair, 10000),     // 后端最多传10000步
		Result:           make(chan Result, 100),     // 和游戏的盘数做对应
	}
}

// Result 游戏的结果的结构体
type Result struct {
	Event string `json:"event"`
	Loser string `json:"loser"`
}

type Pair struct {
	Event      string `json:"event"`
	PlayerId   string
	Direction  string
	ADirection string `json:"a_direction"`
	BDirection string `json:"b_direction"`
}

type baseGameStruct struct {
	Map [][]int32 `json:"map"`
	AId string    `json:"a_id"`
	BId string    `json:"b_id"`
}

// SnakeGame snake游戏的信息
type SnakeGame struct {
	baseGameStruct
	ASx     string `json:"a_sx"`
	ASy     string `json:"a_sy"`
	BSx     string `json:"b_sx"`
	BSy     string `json:"b_sy"`
	PlayerA Player
	PlayerB Player
}

type Player struct {
	Photo    string `json:"photo"`
	Username string `json:"username"`
	UserID   int32  `json:"user_id"`
}

func NewSnakeGame(AId, ASx, ASy, BId, BSx, BSy string, Map [][]int32, playerA, playerB Player) *SnakeGame {
	return &SnakeGame{
		baseGameStruct: baseGameStruct{
			Map: Map,
			AId: AId,
			BId: BId,
		},
		ASx:     ASx,
		ASy:     ASy,
		BSx:     BSx,
		BSy:     BSy,
		PlayerA: playerA,
		PlayerB: playerB,
	}
}
