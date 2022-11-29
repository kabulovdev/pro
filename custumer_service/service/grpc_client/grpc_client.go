package grpcclient

//import (
//	"fmt"

//	"custumer/config"

	//rp "custumer/genproto/post"
	//rs "custumer/genproto/reating"

//	"google.golang.org/grpc"

//	"google.golang.org/grpc/credentials/insecure"
//)

//type Cleints interface {
//	Rewiew() rs.ReatingServiceClient
//}

//type ServiceManager struct {
//	config config.Config

//	reatingService rs.ReatingServiceClient
//	postService    rp.PostServiceClient
//}

//func New(c config.Config) (Cleints, error) {

//	reating, err := grpc.Dial(

//		fmt.Sprintf("%s:%d", c.ReatingServiceHost, c.ReatingServicePort),

//		grpc.WithTransportCredentials(insecure.NewCredentials()),
//	)
//	post, err := grpc.Dial(

//		fmt.Sprintf("%s:%d", c.PostServiceHost, c.PostServicePort),

//		grpc.WithTransportCredentials(insecure.NewCredentials()),
//	)

//	if err != nil {

//		return &ServiceManager{}, nil

//	}

//	return &ServiceManager{

//		config: c,

//		reatingService: rs.NewReatingServiceClient(reating),
//		postService:    rp.NewPostServiceClient(post),
//	}, nil

//}

//func (s *ServiceManager) Rewiew() rs.ReatingServiceClient {

//	return s.reatingService

//}
//func (s *ServiceManager) Post() rp.PostServiceClient {
//	return s.postService
//}
