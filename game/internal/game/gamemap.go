package game

import (
	"fmt"
	"math/rand"
	"snake/conf/mysql"
	"snake/internal/game/util"
	code "snake/internal/grpc/client/coderuning"
	codePb "snake/internal/grpc/client/coderuning/pb"
	"snake/internal/grpc/client/result"
	resultPb "snake/internal/grpc/client/result/pb"
	"snake/internal/models"
	snakePb "snake/internal/pb"
	"snake/pkg/mw"
	"strconv"
	"strings"
	"sync"
	"time"

	"go.uber.org/zap"
	"golang.org/x/net/context"
)

var (
	dx   = []int{-1, 0, 1, 0}
	dy   = []int{0, 1, 0, -1}
	lock sync.Mutex
)

type GameMap struct {
	rows           int
	cols           int
	innerWallCount int
	idA            int
	botA           *models.Bot
	idB            int
	botB           *models.Bot
	g              [][]int32
	playerA        util.Player
	playerB        util.Player
	nextStepA      int32
	nextStepB      int32
	status         string
	loser          string
	loserId        string
	winnerId       string
	MoveMessage    chan *snakePb.SetNextStepResp
}

func NewGameMap(rows, cols, innerWallCount, idA int, botA *models.Bot, idB int, botB *models.Bot) *GameMap {
	botIdA, botIdB := -1, -1
	botCodeA, botCodeB := "", ""
	if botA.ID != 0 {
		botIdA = int(botA.ID)
		botCodeA = botA.Code
	}

	if botB.ID != 0 {
		botIdB = int(botB.ID)
		botCodeB = botB.Code
	}

	array := make([][]int32, rows)
	for i := range array {
		array[i] = make([]int32, cols)
	}

	gameMap := &GameMap{
		rows:           rows,
		cols:           cols,
		innerWallCount: innerWallCount,
		idA:            idA,
		botA:           botA,
		idB:            idB,
		botB:           botB,
		g:              array,
		playerA:        util.Player{Id: idA, BotId: botIdA, BotCode: botCodeA, Sx: rows - 2, Sy: 1, Steps: make([]int32, 0, 20)},
		playerB:        util.Player{Id: idB, BotId: botIdB, BotCode: botCodeB, Sx: 1, Sy: cols - 2, Steps: make([]int32, 0, 20)},
		nextStepA:      -1,
		nextStepB:      -1,
		status:         "playing",
		loser:          "",
		MoveMessage:    make(chan *snakePb.SetNextStepResp, 100),
	}

	return gameMap
}

func (g *GameMap) SetNestStepA(direction int32) {
	lock.Lock()
	defer lock.Unlock()
	g.nextStepA = direction
}

func (g *GameMap) SetNestStepB(direction int32) {
	lock.Lock()
	defer lock.Unlock()
	g.nextStepB = direction
}

func (g *GameMap) GetPlayerA() util.Player {
	return g.playerA
}
func (g *GameMap) GetPlayerB() util.Player {
	return g.playerB
}

func (g *GameMap) GetGameMap() [][]int32 {
	return g.g
}

func (g *GameMap) GetGameMapString() string {
	resp := strings.Builder{}

	for i := 0; i < g.rows; i++ {
		for j := 0; j < g.cols; j++ {
			resp.WriteString(fmt.Sprintf("%d", g.g[i][j]))
		}
	}
	return resp.String()
}

func (g *GameMap) CreateMap() {
	for i := 0; i < 1000; i++ {
		if draw(g) {
			break
		}
	}
}

// draw 随机生成地图
func draw(this *GameMap) bool {
	for i := 0; i < this.rows; i++ {
		for j := 0; j < this.cols; j++ {
			this.g[i][j] = 0
		}
	}

	for r := 0; r < this.rows; r++ {
		this.g[r][0], this.g[r][this.cols-1] = 1, 1
	}
	for c := 0; c < this.cols; c++ {
		this.g[0][c], this.g[this.rows-1][c] = 1, 1
	}

	for i := 0; i < this.innerWallCount/2; i++ {
		for j := 0; j < 1000; j++ {
			r := rand.Intn(this.rows)
			c := rand.Intn(this.cols)

			if this.g[r][c] == 1 || this.g[this.rows-1-r][this.cols-1-c] == 1 {
				continue
			}
			if r == this.rows-2 && c == 1 || r == 1 && c == this.cols-2 {
				continue
			}
			this.g[r][c], this.g[this.rows-1-r][this.cols-1-c] = 1, 1
			break
		}
	}

	return checkConnectivity(this, this.rows-2, 1, 1, this.cols-2)
}

// checkConnectivity 判断生成的地图左下角和右上角是否是联通的
func checkConnectivity(this *GameMap, sx, sy, tx, ty int) bool {
	if sx == tx && sy == ty {
		return true
	}
	this.g[sx][sy] = 1

	for i := 0; i < 4; i++ {
		x, y := sx+dx[i], sy+dy[i]
		if x >= 0 && x < this.rows && y >= 0 && y < this.cols && this.g[x][y] == 0 {

			if checkConnectivity(this, x, y, tx, ty) {
				this.g[sx][sy] = 0
				return true
			}
		}
	}

	this.g[sx][sy] = 0
	return false
}

func (g *GameMap) nextStep() bool {
	// 用代码去操作蛇的移动的时候会计算很多次  相当于加一个施法后摇，防止输入次数过多
	// 前端设定每秒走5个格子如果输入太多步数就会被覆盖，所以在每次计算之前睡一个最小值  1s / 5 = 200ms
	// 这样保证每次走一格最多只会有依次输入
	time.Sleep(time.Millisecond * 200)

	// 接下来可能会有一个Bot自动运行的系统
	// 所以还要有一个函数，去清求那个系统的API让bot自动运行
	g.sendBotCode(g.playerA)
	g.sendBotCode(g.playerB)

	for i := 0; i < 50; i++ {
		time.Sleep(time.Millisecond * 100)
		lock.Lock()

		if g.nextStepA != -1 && g.nextStepB != -1 {
			g.playerA.Steps = append(g.playerA.Steps, g.nextStepA)
			g.playerB.Steps = append(g.playerB.Steps, g.nextStepB)
			lock.Unlock()
			return true
		}
		lock.Unlock()
	}

	return false
}

func (g *GameMap) getInput(player util.Player) string {
	var me, you util.Player
	if g.playerA.Id == player.Id { // 说明自己是A
		me = g.playerA
		you = g.playerB
	} else {
		me = g.playerB
		you = g.playerA
	}

	return g.GetGameMapString() + "#" +
		strconv.Itoa(me.Sx) + "#" +
		strconv.Itoa(me.Sy) + "#(" +
		me.GetStepsString() + ")#" +
		strconv.Itoa(you.Sx) + "#" +
		strconv.Itoa(you.Sy) + "#(" +
		you.GetStepsString() + ")#"
}

func (g *GameMap) sendBotCode(player util.Player) {
	if player.BotId == -1 { // 说明是亲自出马
		return
	}

	// 与codeRunning进行通信
	ctx := context.Background()
	_, err := code.AddBot(ctx, &codePb.AddBotReq{
		UserId:  int32(player.Id),
		BotCode: player.BotCode,
		Input:   g.getInput(player),
	})
	if err != nil {
		zap.L().Error(err.Error())
		return
	}
}

// checkValid 判断最后一步传进来的这一步是否有效
func (g *GameMap) checkValid(cellsA, cellsB []util.Cell) bool {
	n := len(cellsA)
	cell := cellsA[n-1]

	if g.g[cell.X][cell.Y] == 1 {
		return false // 如果最后一位等于墙则判输
	}

	for i := 0; i < n-1; i++ { // 判断cellsA有没有和自己相撞
		if cell.X == cellsA[i].X && cell.Y == cellsA[i].Y {
			return false
		}
	}

	for i := 0; i < n-1; i++ { // 判断cellsB有没有和自己相撞
		if cell.X == cellsB[i].X && cell.Y == cellsB[i].Y {
			return false
		}
	}

	return true

}

// judge 检查两名玩家操作是否合法
func (g *GameMap) judge() {

	zap.L().Debug("judge")
	cellsA := g.playerA.GetCells()
	cellsB := g.playerB.GetCells()

	validA := g.checkValid(cellsA, cellsB)
	validB := g.checkValid(cellsB, cellsA)

	if !validB || !validA {
		g.status = "finished"
		if !validA && !validB {
			g.loser = "all"
		} else if !validA {
			g.loser = "A"
			g.loserId = strconv.Itoa(g.playerA.Id)
			g.winnerId = strconv.Itoa(g.playerB.Id)
		} else {
			g.loser = "B"
			g.loserId = strconv.Itoa(g.playerB.Id)
			g.winnerId = strconv.Itoa(g.playerA.Id)
		}
	}

}

func setLoser(g *GameMap) {
	lock.Lock()
	defer lock.Unlock()
	if g.nextStepA == -1 && g.nextStepB == -1 {
		g.loser = "all"
	} else if g.nextStepA == -1 {
		g.loser = "A"
		g.loserId = strconv.Itoa(g.playerA.Id)
		g.winnerId = strconv.Itoa(g.playerB.Id)
	} else {
		g.loser = "B"
		g.loserId = strconv.Itoa(g.playerB.Id)
		g.winnerId = strconv.Itoa(g.playerA.Id)
	}
	g.sendResult()
}

// 把结果发送给api层
func (g *GameMap) sendMove() {
	lock.Lock()
	defer lock.Unlock()

	zap.L().Debug("sendMove")

	resp := &snakePb.SetNextStepResp{
		Event:      "move",
		ADirection: strconv.Itoa(int(g.nextStepA)),
		BDirection: strconv.Itoa(int(g.nextStepB)),
	}

	g.nextStepA, g.nextStepB = -1, -1

	g.MoveMessage <- resp
	g.MoveMessage <- resp
}

// 向两位玩家公布结果
func (g *GameMap) sendResult() {
	resp := &resultPb.ResultReq{
		EventType: 1,
		GameResult: &resultPb.GameResult{
			Loser:    g.loser,
			LoserId:  g.loserId,
			WinnerId: g.winnerId,
		},
	}

	g.saveGameRecord()
	mw.SugarLogger.Debug(resp)
	_, err := result.Result(context.Background(), resp)
	if err != nil {
		return
	}
}

func (g *GameMap) saveGameRecord() {
	record := models.Record{
		AID:        int32(g.GetPlayerA().Id),
		ASx:        int32(g.GetPlayerA().Sx),
		ASy:        int32(g.GetPlayerA().Sy),
		BID:        int32(g.GetPlayerB().Id),
		BSx:        int32(g.GetPlayerB().Sx),
		BSy:        int32(g.GetPlayerB().Sy),
		ASteps:     g.playerA.GetStepsString(),
		BSteps:     g.playerB.GetStepsString(),
		Map:        g.GetGameMapString(),
		Loser:      g.loser,
		Createtime: time.Now(),
	}

	fmt.Println("saveGameRecord")
	fmt.Println(g.playerA.GetStepsString())
	fmt.Println(g.playerB.GetStepsString())
	fmt.Println(g.GetGameMapString())

	recordDao := mysql.Q.Record

	recordDao.WithContext(context.Background()).Create(&record)
}

func (g *GameMap) Start() {
	for i := 0; i < 1000; i++ {
		if g.nextStep() {
			g.judge()
			if g.status == "playing" {
				g.sendMove()
			} else {
				g.sendResult()
				break
			}
		} else {
			zap.L().Debug("no setNestStep")
			g.status = "finished"
			setLoser(g)
			break
		}
	}
}
