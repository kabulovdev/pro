package models

import (
	pc "gitlab.com/pro/exam_api/genproto/custumer_proto"
	pp "gitlab.com/pro/exam_api/genproto/post_proto"
	pr "gitlab.com/pro/exam_api/genproto/reating_proto"
)

type CustumerCreateForAdmin struct {
	First_name  string //`json:"first_name"`
	Last_name   string //`json:"last_name"`
	Email       string //`json:"email"`
	PhoneNumber string //`json:"phoneNumber"`
	Password    string //`json:"password"`
	Bio         string //`json:"bio"`
	Adres       []Addresss
	Posts       PostForCreate
}

type  PostForCreate  struct {
    Poster_id int64 //`json:"poster_id"`
    Description_post string //`json:"description_post"`
    Medias  []Media 
}

type Media  struct{
	Name string //`json:"name"`
	Link string //`json:"link"`
	Types string //`json:"type"`
 }

type CustumerForCreate struct {
	First_name  string //`json:"first_name"`
	Last_name   string //`json:"last_name"`
	Email       string //`json:"email"`
	PhoneNumber string //`json:"phoneNumber"`
	Password    string //`json:"password"`
	Bio         string //`json:"bio"`
	Adres       []Addresss
}
type Addresss struct {
	Street      string //`json:"street"`
	Home_adress string //`json:"home_adress"`
}

type UpdateAccessToken struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type ResponseError struct {
	Error interface{} `json:"error"`
}

type ServerError struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

type CustumerAllInfo struct {
	Custumer pc.CustumerInfo
	Posts    []Postc
}
type Postc struct {
	Post     pp.PostInfo
	Reatings pr.Reatings
}
type Result struct {
	Custum pc.CustumerInfo
	Post   []Posts
}

type PostIfos struct {
	poster pc.CustumerInfo
	posth  pp.Posts
}

type Posts struct {
	Pos     pp.PostInfo
	Reating []*pr.ReatingInfo
}

type Error struct {
	Code        int
	Error       error
	Description string
}
