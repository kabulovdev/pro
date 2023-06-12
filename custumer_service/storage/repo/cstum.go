package repo

import (
	"context"
	pb "gitlab.com/pro/custumer_service/genproto/custumer_proto"
)

type CustumerInfoI interface {
	Create(context.Context, *pb.CustumerForCreate) (*pb.CustumerInfo, error)
	GetByCustumId(*pb.GetId) (*pb.CustumerInfo, error)
	Update(*pb.CustumerInfo) (*pb.CustumerInfo, error)
	DeletCustum(*pb.GetId) (*pb.Empty, error)
	ListAllCustum(*pb.Empty) (*pb.CustumerAll, error)
	CheckField(*pb.CheckFieldReq) (*pb.CheckFieldRes, error)
	GetAdmin(*pb.GetAdminReq) (*pb.GetAdminRes, error)
	GetModer(*pb.GetAdminReq) (*pb.GetAdminRes, error)
}
