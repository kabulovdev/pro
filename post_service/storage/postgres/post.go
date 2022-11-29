package postgres

import (
	"fmt"
	pb "gitlab.com/pro/post_service/genproto/post_proto"

	"github.com/jmoiron/sqlx"
)

type postRepo struct {
	db *sqlx.DB
}

func NewPostRepo(db *sqlx.DB) *postRepo {
	return &postRepo{db: db}
}

func (r *postRepo) Create(req *pb.PostForCreate) (*pb.PostInfo, error) {
	result := pb.PostInfo{}
	err := r.db.QueryRow(`insert into posts (
	poster_id,
	description_post
	) values ($1, $2) returning id,description_post, poster_id`, req.PosterId, req.DescriptionPost).Scan(
		&result.Id,
		&result.DescriptionPost,
		&result.PosterId)
	if err != nil {
		return &pb.PostInfo{}, err

	}
	medias := []*pb.Media{}
	for _, md := range req.Medias {
		medResp := pb.Media{}
		err = r.db.QueryRow(`
		insert into medias ( 
		post_id,
		name,
		link,
		type) values ($1,$2,$3,$4)
		returning name, link, type`, result.Id, md.Name, md.Link, md.Type).Scan(
			&medResp.Name, &medResp.Link, &medResp.Type)
		if err != nil {
			return &pb.PostInfo{}, err
		}
		medias = append(medias, &medResp)
	}
	fmt.Println(medias)
	result.Medias = medias
	return &result, nil
}

func (p *postRepo) GetByOwnerId(req *pb.Id) (*pb.Posts, error) {
	result := pb.Posts{}
	temp1 := []*pb.PostInfo{}
	rows, err := p.db.Query(`select 
	id,
	poster_id,
	description_post
	from posts where poster_id = $1 and deleted_at is null`, req.Id)
	if err != nil {
		return &pb.Posts{}, err
	}

	for rows.Next() {
		temp := &pb.PostInfo{}
		err = rows.Scan(
			&temp.Id,
			&temp.PosterId,
			&temp.DescriptionPost,
		)
		if err != nil {
			return &pb.Posts{}, err
		}
		rowlar, err := p.db.Query(`
	select  
    name,
    link,
    type
	from medias where post_id=$1
	`, temp.Id)
		if err != nil {
			return &pb.Posts{}, err
		}
		medias := []*pb.Media{}
		for rowlar.Next() {
			med := pb.Media{}
			err = rowlar.Scan(
				&med.Name,
				&med.Link,
				&med.Type)
			if err != nil {
				return &pb.Posts{}, err
			}
			medias = append(medias, &med)
		}
		fmt.Println(temp)
		temp.Medias = medias
		temp1 = append(temp1, temp)
		fmt.Println(temp1)
	}
	result.Posts = temp1
	return &result, nil
}

func (r *postRepo) GetPosterInfo(req *pb.Id) (*pb.PostInfo, error) {
	result := pb.PostInfo{}

	err := r.db.QueryRow(`select 
	id,
	poster_id,
	description_post
	from posts where id = $1 and deleted_at is null`, req.Id).Scan(
		&result.Id,
		&result.PosterId,
		&result.DescriptionPost,
	)
	if err != nil {
		return &pb.PostInfo{}, err
	}
	rows, err := r.db.Query(`
	select  
    name,
    link,
    type
	from medias where post_id=$1
	`, result.Id)
	if err != nil {
		return &pb.PostInfo{}, err
	}
	medias := []*pb.Media{}
	for rows.Next() {
		med := pb.Media{}
		err = rows.Scan(
			&med.Name,
			&med.Link,
			&med.Type)
		if err != nil {
			return &pb.PostInfo{}, err
		}
		medias = append(medias, &med)
	}
	result.Medias = medias
	return &result, nil
}

func (r *postRepo) GetPost(req *pb.Id) (*pb.PostInfo, error) {
	result := pb.PostInfo{}

	err := r.db.QueryRow(`select 
	id,
	poster_id,
	description_post
	from posts where id = $1 and deleted_at is null`, req.Id).Scan(
		&result.Id,
		&result.PosterId,
		&result.DescriptionPost,
	)
	fmt.Println(result)
	if err != nil {
		return &pb.PostInfo{}, err
	}
	rows, err := r.db.Query(`
	select  
    name,
    link,
    type
	from medias where post_id=$1
	`, result.Id)
	if err != nil {
		fmt.Println(err)
		return &pb.PostInfo{}, err
	}
	medias := []*pb.Media{}
	for rows.Next() {
		med := pb.Media{}
		err = rows.Scan(
			&med.Name,
			&med.Link,
			&med.Type)
		if err != nil {
			fmt.Println(err)
			return &pb.PostInfo{}, err
		}
		fmt.Println(med)
		medias = append(medias, &med)
	}
	result.Medias = medias
	fmt.Println(result)
	return &result, nil
}

func (r *postRepo) Update(req *pb.PostInfo) (*pb.PostInfo, error) {
	postResp := pb.PostInfo{}

	err := r.db.QueryRow(`
	UPDATE posts
	SET
	poster_id = $1,
	description_post = $2
	WHERE id = $3 
	returning id, poster_id, description_post`,
		req.PosterId, req.DescriptionPost, req.Id).
		Scan(&postResp.Id, &postResp.PosterId, &postResp.DescriptionPost)
	if err != nil {
		return &pb.PostInfo{}, err
	}
	adder := []*pb.Media{}
	for _, adress := range req.Medias {
		address := pb.Media{}
		err = r.db.QueryRow(`UPDATE medias
	SET 
	name = $1,
    link = $2,
    type = $3 WHERE post_id = $4 returning name, link, type`, adress.Name, adress.Link, adress.Type, req.Id).Scan(
			&address.Name,
			&address.Link,
			&address.Type)

		adder = append(adder, adress)
		fmt.Println(adder)
	}
	if err != nil {
		return &pb.PostInfo{}, err
	}
	postResp.Medias = adder
	return &postResp, nil
}
func (r *postRepo) Delet(req *pb.Id) (*pb.EmptyPost, error) {
	_, err := r.db.Query(`UPDATE posts SET deleted_at = NOW() WHERE id =$1`, req.Id)
	if err != nil {
		return &pb.EmptyPost{}, err
	}
	return &pb.EmptyPost{}, nil
}
