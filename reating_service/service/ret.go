package service

import (
	"context"
	"fmt"

	pb "gitlab.com/pro/reating_service/genproto/reating"
	l "gitlab.com/pro/reating_service/pkg/logger"
	"gitlab.com/pro/reating_service/storage"

	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ReatingService struct {
	storage storage.IStorage
	logger  l.Logger
}

func NewReatingService(db *sqlx.DB, log l.Logger) *ReatingService {
	return &ReatingService{
		storage: storage.NewStoragePg(db),
		logger:  log,
	}
}

func (s *ReatingService) GetPostReating(cxt context.Context, req *pb.Id) (*pb.Reatings, error) {
	fmt.Println(req)
	Reatin, err := s.storage.Reating().GetPostReating(&pb.Id{Id: req.Id})
	if err != nil {
		s.logger.Error("error while geting post", l.Any("error geting post", err))
		return &pb.Reatings{}, status.Error(codes.Internal, "something went wrong")
	}
	return Reatin, nil
}

func (s *ReatingService) Create(cxt context.Context, req *pb.ReatingForCreate) (*pb.ReatingInfo, error) {
	fmt.Println(req)
	Reatin, err := s.storage.Reating().Create(req)
	if err != nil {
		s.logger.Error("error while creating post", l.Any("error creating post", err))
		return &pb.ReatingInfo{}, status.Error(codes.Internal, "something went wrong")
	}
	return Reatin, nil
}
func (s *ReatingService) GetReating(cxt context.Context, req *pb.Id) (*pb.ReatingInfo, error) {
	fmt.Println(req)
	Reatin, err := s.storage.Reating().GetReating(req)
	if err != nil {
		s.logger.Error("error while geting post", l.Any("error geting post", err))
		return &pb.ReatingInfo{}, status.Error(codes.Internal, "something went wrong")
	}
	return Reatin, nil
}
func (s *ReatingService) Update(cxt context.Context, req *pb.ReatingInfo) (*pb.ReatingInfo, error) {
	fmt.Println(req)
	Reatin, err := s.storage.Reating().Update(req)
	if err != nil {
		s.logger.Error("error while updating post", l.Any("error updating post", err))
		return &pb.ReatingInfo{}, status.Error(codes.Internal, "something went wrong")
	}
	return Reatin, nil
}
func (s *ReatingService) Delet(cxt context.Context, req *pb.Id) (*pb.EmptyReating, error) {
	fmt.Println(req)
	Post, err := s.storage.Reating().Delet(req)
	if err != nil {
		s.logger.Error("error while deleting post", l.Any("error deleting post", err))
		return &pb.EmptyReating{}, status.Error(codes.Internal, "something went wrong")
	}
	return Post, nil
}
