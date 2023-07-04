package snake

import (
	pb "backend/internal/grpc/client/snake/pb"
	shape "backend/pkg/share_space"
	"strings"

	"golang.org/x/net/context"
	"golang.org/x/sync/errgroup"
)

var (
	eg errgroup.Group
)

// WithSnake 只有一种消息，设置下一步，以便之后进行拓展
type WithSnake struct {
	Msg chan shape.Pair
}

func (s *WithSnake) Send() error {
	go ReadWebMsg()

	eg.Go(func() error {
		for {
			select {
			case msg := <-s.Msg:
				req := &pb.SetNextStepReq{Direction: msg.Direction, PlayerId: msg.PlayerId}

				eg.Go(func() error {
					resp, err := SetNextStep(context.Background(), req)
					if err != nil {
						return err
					}

					if strings.EqualFold("move", resp.GetEvent()) {
						getMove(resp)
					}
					return nil
				})
			}
		}
	})

	return eg.Wait()
}

// ReadWebMsg 死循环读取前端传来的消息
func ReadWebMsg() {
	for {
		select {
		case msg := <-Space.ClientDirection:
			SnakeMd.Msg <- msg
		}
	}
}

func getMove(resp *pb.SetNextStepResp) {

	pair := shape.Pair{
		Event:      "move",
		ADirection: resp.GetADirection(),
		BDirection: resp.GetBDirection(),
	}

	Space.ServiceDirection <- pair
}
