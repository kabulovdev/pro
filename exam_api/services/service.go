package services

import (
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/resolver"

	"gitlab.com/pro/exam_api/config"
	pbs "gitlab.com/pro/exam_api/genproto/custumer_proto"
	pb "gitlab.com/pro/exam_api/genproto/post_proto"
	pr "gitlab.com/pro/exam_api/genproto/reating_proto"
)

type IServiceManager interface {
	PostService() pb.PostServiceClient
	CustumerService() pbs.CustomServiceClient
	ReatingService() pr.ReatingServiceClient
}

type serviceManager struct {
	postService  pb.PostServiceClient
	custumerService pbs.CustomServiceClient
	reatingService pr.ReatingServiceClient
}

func (s *serviceManager) CustumerService() pbs.CustomServiceClient {
	return s.custumerService
}

func (s *serviceManager) PostService() pb.PostServiceClient {
	return s.postService
}

func (s *serviceManager) ReatingService() pr.ReatingServiceClient {
	return s.reatingService
}

func NewServiceManager(conf *config.Config) (IServiceManager, error) {
	resolver.SetDefaultScheme("dns")

	connCustum, err := grpc.Dial(
		fmt.Sprintf("%s:%d", conf.CustumerServiceHost, conf.CustumerServicePort),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	connPost, err := grpc.Dial(
		fmt.Sprintf("%s:%d", conf.PostServiceHost, conf.PostServicePort),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	connReating, err := grpc.Dial(
		fmt.Sprintf("%s:%d", conf.ReatingServiceHost, conf.ReatingServicePort),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}	
	serviceManager := &serviceManager{
		postService: pb.NewPostServiceClient(connPost),
		custumerService: pbs.NewCustomServiceClient(connCustum),
		reatingService: pr.NewReatingServiceClient(connReating),
	}

	return serviceManager, nil
}
