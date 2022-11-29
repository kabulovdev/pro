package postgres

import (
	pb "gitlab.com/pro/reating_service/genproto/reating"

	"github.com/jmoiron/sqlx"
)

type reatingRepo struct {
	db *sqlx.DB
}

func NewPostRepo(db *sqlx.DB) *reatingRepo {
	return &reatingRepo{db: db}
}

func (r *reatingRepo) Create(req *pb.ReatingForCreate) (*pb.ReatingInfo, error) {
	result := pb.ReatingInfo{}
	err := r.db.QueryRow(`insert into reating (
		post_id,
		custumer_id,
    	rating,
    	description
		) values ($1, $2, $3, $4) returning id, post_id,custumer_id,rating,description`, req.PostId, req.CustumerId, req.Reating, req.Description).Scan(
		&result.Id,
		&result.PostId,
		&result.CustumerId,
		&result.Reating,
		&result.Description)
	if err != nil {
		return &pb.ReatingInfo{}, err

	}
	return &result, nil
}

func (r *reatingRepo) GetPostReating(req *pb.Id) (*pb.Reatings, error) {
	result := pb.Reatings{}
	rows, err := r.db.Query(`select 
	id,
	custumer_id,
    rating,
    description
	from reating where post_id = $1 and deleted_at is null`, req.Id)
	if err != nil {
		return &pb.Reatings{}, err
	}
	reating := []*pb.ReatingInfo{}
	for rows.Next() {
		red := pb.ReatingInfo{}
		err = rows.Scan(
			&red.Id,
			&red.CustumerId,
			&red.Reating,
			&red.Description)
		if err != nil {
			return &pb.Reatings{}, err
		}
		reating = append(reating, &red)
	}
	result.Reatins = reating
	return &result, nil

}

func (r *reatingRepo) GetReating(req *pb.Id) (*pb.ReatingInfo, error) {
	result := pb.ReatingInfo{}

	err := r.db.QueryRow(`select 
	post_id,
	custumer_id,
    rating,
    description
	from reating where id = $1 and deleted_at is null`, req.Id).Scan(
		&result.PostId,
		&result.CustumerId,
		&result.Reating,
		&result.Description)
	if err != nil {
		return &pb.ReatingInfo{}, err
	}
	return &result, nil
}
func (r *reatingRepo) Update(req *pb.ReatingInfo) (*pb.ReatingInfo, error) {
	reatingResp := pb.ReatingInfo{}

	err := r.db.QueryRow(`
	UPDATE reating
	SET
	post_id = $1,
	custumer_id=$2,
	rating = $3,
	description = $4
	WHERE id = $5 
	returning id, post_id,custumer_id, rating, description`,
		req.PostId, req.CustumerId, req.Reating, req.Description, req.Id).
		Scan(&reatingResp.Id, &reatingResp.PostId, &reatingResp.CustumerId, &reatingResp.Reating, &reatingResp.Description)
	if err != nil {
		return &pb.ReatingInfo{}, err
	}
	return &reatingResp, nil
}
func (r *reatingRepo) Delet(req *pb.Id) (*pb.EmptyReating, error) {
	_, err := r.db.Query(`UPDATE reating SET deleted_at = NOW() WHERE id =$1`, req.Id)
	if err != nil {
		return &pb.EmptyReating{}, err
	}
	return &pb.EmptyReating{}, nil
}
