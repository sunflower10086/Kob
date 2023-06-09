package match

import (
	"context"
	"matching/internal/match/logic"

	pb "matching/internal/pb/matchingServer"
)

type MatchingSystemServerImpl struct {
	pb.UnimplementedMatchingSystemServer
}

func (m *MatchingSystemServerImpl) AddUser(ctx context.Context, user *pb.User) (*pb.Response, error) {
	addUser, err := logic.AddUser(ctx, user.GetUserId(), user.GetBotId())
	if err != nil {
		return nil, err
	}
	return addUser, nil
}

func (m *MatchingSystemServerImpl) Remove(ctx context.Context, user *pb.User) (*pb.Response, error) {
	remove, err := logic.Remove(ctx, user.GetUserId())
	if err != nil {
		return nil, err
	}
	return remove, nil
}
