package v1

import (
	"context"
	"gitlab.com/pro/exam_api/api/handlers/models"
	"fmt"
	"net/http"
	"strings"
	email "gitlab.com/pro/exam_api/email"
	"encoding/json"
	"strconv"
	"github.com/spf13/cast"
	"time"
 	etc "gitlab.com/pro/exam_api/pkg/etc"
	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/encoding/protojson"

	pbs "gitlab.com/pro/exam_api/genproto/custumer_proto"
	pb "gitlab.com/pro/exam_api/genproto/post_proto"
	pr "gitlab.com/pro/exam_api/genproto/reating_proto"

	//pr "exam_api/genproto/reating_proto"
	l "gitlab.com/pro/exam_api/pkg/logger"
)

type AllthingPost struct {
	Postinfo   pb.Posts
	Posterinfo pbs.CustumerInfo
}
// GetReatingAavarage get Reating
// @Summary      Get  posts reating api
// @Description this api get posts reating by id
// @Tags Post
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param   id   path  int  true  "Post ID"
// @Success 200 {object} reating.Reatings
// @Router /v1/post/get/reatings/avarage/{id}  [get]
func (h *handlerV1) GetPostReatingNew(c *gin.Context) {
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true
	guid := c.Param("id")
	fmt.Println(guid)
	id, err := strconv.ParseInt(guid, 10, 64)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed parse string to int", l.Error(err))
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()
	result, err := h.serviceManager.ReatingService().GetPostReating(ctx, &pr.Id{Id: id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to create user", l.Error(err))
		return
	}
	var reatinglar int64
	var count int64
	for _, reatings := range result.Reatins{
		reatinglar=reatinglar+reatings.Reating
		count++
	}

	c.JSON(http.StatusOK, reatinglar/count)
}

// GetReating get Reating
// @Summary      Get  posts reating api
// @Description this api get posts reating by id
// @Tags Post
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param   id   path  int  true  "Post ID"
// @Success 200 {object} reating.Reatings
// @Router /v1/post/get/reatings/{id}  [get]
func (h *handlerV1) GetPostReating(c *gin.Context) {
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true
	guid := c.Param("id")
	id, err := strconv.ParseInt(guid, 10, 64)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed parse string to int", l.Error(err))
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()
	result, err := h.serviceManager.ReatingService().GetPostReating(ctx, &pr.Id{Id: id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to create user", l.Error(err))
		return
	}
	c.JSON(http.StatusOK, result)
}

// GetCustumers get Custumers
// @Summary      Get only custumers api
// @Description this api get custumers
// @Tags Custumer
// @Security BearerAuth
// @Accept json
// @Produce json
// @Success 200 {object} custumer.CustumerAll
// @Router /v1/custumer/getList [get]
func (h *handlerV1) GetListCustumers(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()
	result, err := h.serviceManager.CustumerService().ListAllCustum(ctx, &pbs.Empty{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to create user", l.Error(err))
		return
	}
	c.JSON(http.StatusOK, result)
}

// CreatePost create post
// @Summary      create post api
// @Description this api create post
// @Tags Post
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param request body post.PostForCreate true "Custumer"
// @Success 200 {object} post.PostInfo
// @Router /v1/post/create [post]
func (h *handlerV1) CreatePost(c *gin.Context) {
	var (
		body        pb.PostForCreate
		jspbMarshal protojson.MarshalOptions
	)
	jspbMarshal.UseProtoNames = true

	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to bind json", l.Error(err))
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()
	resStore, err := h.serviceManager.PostService().Create(ctx, &body)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to create post", l.Error(err))
		return
	}

	c.JSON(http.StatusCreated, resStore)

}

// DeleteCustumer delete Post
// @Summary      delete Post api
// @Description this api posts by id
// @Tags Post
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param   id   path  int  true  "Post id"
// @Success 200 {object} post.EmptyPost
// @Router /v1/post/delet/{id}  [delete]
func (h *handlerV1) DeletePost(c *gin.Context) {
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true
	guid := c.Param("id")
	id, err := strconv.ParseInt(guid, 10, 64)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed parse string to int", l.Error(err))
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()
	resStore, err := h.serviceManager.PostService().Delet(ctx, &pb.Id{Id: id})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to create user", l.Error(err))
		return
	}
	c.JSON(http.StatusOK, resStore)
}

// CreateStore create store
// @Summary      create store api
// @Description this api create store
// @Tags Custumer
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param request body custumer.CustumerForCreate true "Custumer"
// @Success 200 {object} custumer.CustumerInfo
// @Router /v1/custumer/create [post]
func (h *handlerV1) CreateCustumer(c *gin.Context) {
	var (
		body        pbs.CustumerForCreate
		jspbMarshal protojson.MarshalOptions
	)
	jspbMarshal.UseProtoNames = true

	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to bind json", l.Error(err))
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()
	resStore := &pbs.CustumerInfo{}
	resStore, err = h.serviceManager.CustumerService().Create(ctx, &body)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to create user", l.Error(err))
		return
	}

	c.JSON(http.StatusCreated, resStore)

}

// CreateStore create store
// @Summary      create reating api
// @Description this api create reating
// @Tags Reating
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param request body reating.ReatingForCreate true "Custumer"
// @Success 200 {object} reating.ReatingInfo
// @Router /v1/reating/create [post]
func (h *handlerV1) CreateReating(c *gin.Context) {
	var (
		body        pr.ReatingForCreate
		jspbMarshal protojson.MarshalOptions
	)
	jspbMarshal.UseProtoNames = true

	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to bind json", l.Error(err))
		return
	}
	if body.Reating > 5 {
		c.JSON(http.StatusBadRequest, "reating false")
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()
	resStore := &pr.ReatingInfo{}
	resStore, err = h.serviceManager.ReatingService().Create(ctx, &body)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to create user", l.Error(err))
		return
	}

	c.JSON(http.StatusCreated, resStore)

}

// GetReating get Reating
// @Summary      Get reating api
// @Description this api get reating by id
// @Tags Reating
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param   id   path  int  true  "reating id"
// @Success 200 {object} reating.ReatingInfo
// @Router /v1/reating/get/{id}  [get]
func (h *handlerV1) GetReating(c *gin.Context) {
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true
	guid := c.Param("id")
	id, err := strconv.ParseInt(guid, 10, 64)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed parse string to int", l.Error(err))
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()
	result, err := h.serviceManager.ReatingService().GetReating(ctx, &pr.Id{Id: id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to create user", l.Error(err))
		return
	}
	c.JSON(http.StatusOK, result)
}

// DeleteCustumer delete Custumer
// @Summary      delete Reating api
// @Description this api delet reating by id
// @Tags Reating
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param   id   path  int  true  "reating id"
// @Success 200 {object} reating.EmptyReating
// @Router /v1/reating/delete/{id}  [delete]
func (h *handlerV1) DeleteReating(c *gin.Context) {
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true
	guid := c.Param("id")
	id, err := strconv.ParseInt(guid, 10, 64)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed parse string to int", l.Error(err))
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()
	resStore, err := h.serviceManager.ReatingService().Delet(ctx, &pr.Id{Id: id})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to create user", l.Error(err))
		return
	}
	c.JSON(http.StatusOK, resStore)
}

// DeleteCustumer delete Custumer
// @Summary      delete Custumer api
// @Description this api delet custumer with posts by id
// @Tags Custumer
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param   id   path  int  true  "Custumer id"
// @Success 200 {object} custumer.Empty
// @Router /v1/custumer/delete/{id}  [delete]
func (h *handlerV1) DeleteCustumer(c *gin.Context) {
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true
	guid := c.Param("id")
	id, err := strconv.ParseInt(guid, 10, 64)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed parse string to int", l.Error(err))
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	resStore, err := h.serviceManager.CustumerService().DeletCustum(ctx, &pbs.GetId{Id: id})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to create user", l.Error(err))
		return
	}
	c.JSON(http.StatusOK, resStore)
}

// GetCustumer get Custumer
// @Summary      get Custumer api
// @Description this api get custumer with posts by id
// @Tags Custumer
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param   id   path  int  true  "Custumer id"
// @Success 200 {object} custumer.CustumerInfo
// @Router /v1/custumer/get/{id}  [get]
func (h *handlerV1) GetCustumer(c *gin.Context) {
	reul := models.CustumerAllInfo{}
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true
	guid := c.Param("id")
	id, err := strconv.ParseInt(guid, 10, 64)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed parse string to int", l.Error(err))
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	response, err := h.serviceManager.CustumerService().GetByCustumId(ctx, &pbs.GetId{
		Id: id,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to create user", l.Error(err))
		return
	}
	reul.Custumer = *response
	response2, err := h.serviceManager.PostService().GetByOwnerID(ctx, &pb.Id{Id: id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to create user", l.Error(err))
		return
	}
	azaz := reul.Posts
	aza := models.Postc{}
	//reatings := []*pr.ReatingInfo{}
	for _, post := range response2.Posts {
		aza.Post = *post
		response3, err := h.serviceManager.ReatingService().GetPostReating(ctx, &pr.Id{Id: post.Id})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
		}
		aza.Reatings = *response3
	}
	azaz = append(azaz, aza)

	reul.Posts = azaz

	c.JSON(http.StatusCreated, reul)
}

// CreateStore create store
// @Summary      update reating api
// @Description this api update reating
// @Tags Reating
//@Security BearerAuth
// @Accept json
// @Produce json
// @Param request body reating.ReatingInfo true "reating"
// @Success 200 {object} reating.ReatingInfo
// @Router /v1/reating/update [put]
func (h *handlerV1) UpdateReating(c *gin.Context) {
	var (
		body        pr.ReatingInfo
		jspbMarshal protojson.MarshalOptions
	)
	jspbMarshal.UseProtoNames = true

	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to bind json", l.Error(err))
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()
	resStore := &pr.ReatingInfo{}
	resStore, err = h.serviceManager.ReatingService().Update(ctx, &body)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to create user", l.Error(err))
		return
	}

	c.JSON(http.StatusCreated, resStore)

}

// GetOnlyPost get Post
// @Summary      Get  post api
// @Description this api get post by id
// @Tags Post
//@Security BearerAuth
// @Accept json
// @Produce json
// @Param   id   path  int  true  "Post ID"
// @Success 200 {object} post.PostInfo
// @Router /v1/post/get/{id}  [get]
func (h *handlerV1) GetPost(c *gin.Context) {
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true
	guid := c.Param("id")
	id, err := strconv.ParseInt(guid, 10, 64)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed parse string to int", l.Error(err))
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()
	result, err := h.serviceManager.PostService().GetPost(ctx,&pb.Id{Id: id} )
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to create user", l.Error(err))
		return
	}
	c.JSON(http.StatusOK, result)
}


// CreateStore create store
// @Summary      update custumer api
// @Description this api update custumer
// @Tags Custumer
//@Security BearerAuth
// @Accept json
// @Produce json
// @Param request body custumer.CustumerInfo true "Custumer"
// @Success 200 {object} custumer.CustumerInfo
// @Router /v1/custumer/update [put]
func (h *handlerV1) UpdateCustumer(c *gin.Context) {
	var (
		body        pbs.CustumerInfo
		jspbMarshal protojson.MarshalOptions
	)
	jspbMarshal.UseProtoNames = true

	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to bind json", l.Error(err))
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()
	resStore := &pbs.CustumerInfo{}
	resStore, err = h.serviceManager.CustumerService().Update(ctx, &body)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to create user", l.Error(err))
		return
	}

	c.JSON(http.StatusCreated, resStore)

}

// CreateStore create store
// @Summary      update post api
// @Description this api update Post
// @Tags Post
//@Security BearerAuth
// @Accept json
// @Produce json
// @Param request body post.PostInfo true "Post"
// @Success 200 {object} post.PostInfo
// @Router /v1/post/update [put]
func (h *handlerV1) UpdatePost(c *gin.Context) {
	var (
		body        pb.PostInfo
		jspbMarshal protojson.MarshalOptions
	)
	jspbMarshal.UseProtoNames = true

	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to bind json", l.Error(err))
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()
	resStore := &pb.PostInfo{}
	resStore, err = h.serviceManager.PostService().Update(ctx, &body)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to create user", l.Error(err))
		return
	}

	c.JSON(http.StatusCreated, resStore)

}

// GetPost get Post
// @Summary      get Post api
// @Description this api get Post by id
// @Tags Post
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param   id   path  int  true  "Poster id"
// @Success 200 "object"
// @Router /v1/post/allInfo/{id}  [get]
func (h *handlerV1) GetPostAllInfo(c *gin.Context) {
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true
	guid := c.Param("id")
	id, err := strconv.ParseInt(guid, 10, 64)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed parse string to int", l.Error(err))
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	response, err := h.serviceManager.PostService().GetByOwnerID(ctx, &pb.Id{Id: id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to get user", l.Error(err))
		return

	}
	var azaz int64
	for _, poster := range response.Posts {
		azaz = poster.PosterId
	}

	fmt.Println(azaz)
	response2, err := h.serviceManager.CustumerService().GetByCustumId(ctx, &pbs.GetId{Id: azaz})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to get user", l.Error(err))
		return
	}
	natija := AllthingPost{}
	natija.Posterinfo = *response2
	natija.Postinfo = *response
	c.JSON(http.StatusOK, natija)
}
// Verify user
// @Summary      Verify custumer
// @Description  Verifys custumer
// @Tags         Register
// @Security BearerAuth
// @Accept       json
// @Produce      json
// @Param        code   path string true "code"
// @Param        email  path string true "email"
// @Success      200  {object}  custumer.CustumerInfo
// @Router       /v1/verify/{email}/{code} [patch]
func (h *handlerV1) Verify(c *gin.Context) {
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	var (
		code  = c.Param("code")
		email = c.Param("email")
	)
	fmt.Println(email, code)
	userBody, err := h.redis.Get(email)
	if err != nil {
		c.JSON(http.StatusGatewayTimeout, gin.H{
			"info":  "Your time has expired",
			"error": err.Error(),
		})
		h.log.Error("Error while getting user from redis", l.Any("redis", err))
	}
	fmt.Printf(">>",userBody)
	userBodys := cast.ToString(userBody)

	body := pbs.CustumerForCreate{}
	fmt.Println(string(userBodys))
	err = json.Unmarshal([]byte(userBodys), &body)
	fmt.Println(body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("Error while unmarshaling from json to user body", l.Any("json", err))
		return
	}
	if body.Code != code {
		c.JSON(http.StatusConflict, gin.H{
			"info": "Wrong code",
		})
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	id := strconv.FormatInt(int64(body.Id), 10)

	// Genrating refresh and jwt tokens
	h.jwthandler.Iss = "user"
	h.jwthandler.Sub = id
	h.jwthandler.Role = "authorized"
	h.jwthandler.Aud = []string{"exam-app"}
	h.jwthandler.SigninKey = h.cfg.SignKey
	h.jwthandler.Log = h.log
	tokens, err := h.jwthandler.GenerateAuthJWT()
	accessToken := tokens[0]
	refreshToken := tokens[1]

	if err != nil {
		h.log.Error("error occured while generating tokens")
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "something went wrong,please try again",
		})
		return
	}
	body.AccessToken = accessToken
	res, err := h.serviceManager.CustumerService().Create(ctx, &body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("Error while creating user", l.Any("post", err))
		return
	}

	fmt.Println(refreshToken,"\n","\n")
	fmt.Println(accessToken)

	c.JSON(http.StatusOK, res)
}


// Register Custumer
// @Summary      Register Custumer
// @Description  Registers Custumer
// @Tags         Register
// @Security BearerAuth
// @Accept       json
// @Produce      json
// @Param        user   body   custumer.CustumerForCreate   true  "Custumer"
// @Success      200  {object}  custumer.CustumerInfo
// @Router       /v1/register [post]
func (h *handlerV1) RegisterUser(c *gin.Context) {
	var body pbs.CustumerForCreate

	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"Error": err.Error(),
			"Hint":  "Check your data",
		})
		h.log.Error("Error while binding json", l.Any("json", err))
		return
	}
	body.Email = strings.TrimSpace(body.Email)
	body.FirstName = strings.TrimSpace(body.FirstName)

	body.Email = strings.ToLower(body.Email)

	body.Password, err = etc.HashPassword(body.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Something went wrong")
		h.log.Error("couldn't hash the password")
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	emailExists, err := h.serviceManager.CustumerService().CheckField(ctx, &pbs.CheckFieldReq{
		Field: "email",
		Value: body.Email,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Error": err.Error(),
		})
		h.log.Error("Error while cheking email uniqeness", l.Any("check", err))
		return
	}

	if emailExists.Exist {
		c.JSON(http.StatusConflict, gin.H{
			"info": "Email is already used",
		})
		return
	}

	usernameExists, err := h.serviceManager.CustumerService().CheckField(ctx, &pbs.CheckFieldReq{
		Field: "username",
		Value: body.FirstName,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Error": err.Error(),
		})
		h.log.Error("Error while cheking username uniqeness", l.Any("check", err))
		return
	}

	if usernameExists.Exist {
		c.JSON(http.StatusConflict, gin.H{
			"info": "Username is already used",
		})
		return
	}
	body.Code = etc.GenerateCode(6)
	fmt.Println(body)
	bodyByte, err := json.Marshal(body)
	if err != nil {
		h.log.Error("Error while marshaling to json", l.Any("json", err))
		return
	}
	
	fmt.Println(body.Code)
	msg := "Subject: Customer email verification\n Your verification code: " + body.Code
	err = email.SendEmail([]string{body.Email}, []byte(msg))

	c.JSON(http.StatusAccepted, gin.H{
		"info": "Your request is accepted we have sent you an email message, please check and verify",
	})
	fmt.Println(body.Code)
	fmt.Println(body.Email)
	fmt.Println(string(bodyByte))
	err = h.redis.SetWithTTL(body.Email, string(bodyByte), 300)
	fmt.Println(body.Email)
	if err != nil {
		h.log.Error("Error while marshaling to json", l.Any("json", err))
		return
	}
	fmt.Println(body)
}
// Login Admin
// @Summary      Login admin
// @Description  Login admin
// @Tags         Login
// @Security BearerAuth
// @Accept       json
// @Produce      json
// @Param        name  path string true "admin name"
// @Param        password  path string true "admin password"
// @Success      200  {object}  custumer.GetAdminRes
// @Router       /v1/admin/login/{name}/{password} [GET]
func (h *handlerV1) LoginAdmin(c *gin.Context) {
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	var (
		password  = c.Param("password")
		adminName = c.Param("name")
	)

	fmt.Println("--> ",password, adminName)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()
	fmt.Println(adminName)
	res, err := h.serviceManager.CustumerService().GetAdmin(ctx, &pbs.GetAdminReq{Name: adminName})
	fmt.Println(password,res.Password)
	if err != nil {
		c.JSON(http.StatusNotFound, models.Error{
			Code: http.StatusNotFound,
			Error:       err,
			Description: "Couln't find matching information, Have you registered before?",
		})
		h.log.Error("Error while getting admin by admin Name", l.Any("Get", err))
		return
	}
	fmt.Println(len(res.Password), len(password))
	fmt.Println("->",res)
	//if !etc.CheckPasswordHash(password, res.Password) {
	//	c.JSON(http.StatusConflict, models.Error{
	//		Description: "Password or adminName error",
	//		Code:        http.StatusConflict,
	//	})
	//	return
	//}

	h.jwthandler.Iss = "admin"
	h.jwthandler.Sub = res.Id
	h.jwthandler.Role = "admin"
	h.jwthandler.Aud = []string{"exam-app"}
	h.jwthandler.SigninKey = h.cfg.SignKey
	h.jwthandler.Log = h.log
	tokens, err := h.jwthandler.GenerateAuthJWT()
	accessToken := tokens[0]

	if err != nil {
		h.log.Error("error occured while generating tokens")
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "something went wrong,please try again",
		})
		return
	}
	res.AccesToken = accessToken
	res.Password = ""
	c.JSON(http.StatusOK, res)
}

// Login Moderator
// @Summary      Login moder
// @Description  Login moder
// @Tags         Login
// @Security BearerAuth
// @Accept       json
// @Produce      json
// @Param        name  path string true "Moderator name"
// @Param        password  path string true "Moderator password"
// @Success      200  {object}  custumer.GetAdminRes
// @Router       /v1/moder/login/{name}/{password} [GET]
func (h *handlerV1) LoginModerator(c *gin.Context) {
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	var (
		password  = c.Param("password")
		adminName = c.Param("name")
	)

	fmt.Println("--> ",password, adminName)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()
	fmt.Println(adminName)
	res, err := h.serviceManager.CustumerService().GetModer(ctx, &pbs.GetAdminReq{Name: adminName})
	fmt.Println(password,res.Password)
	if err != nil {
		c.JSON(http.StatusNotFound, models.Error{
			Code: http.StatusNotFound,
			Error:       err,
			Description: "Couln't find matching information, Have you registered before?",
		})
		h.log.Error("Error while getting admin by admin Name", l.Any("Get", err))
		return
	}
	fmt.Println(len(res.Password), len(password))
	fmt.Println("->",res)
	//if !etc.CheckPasswordHash(password, res.Password) {
	//	c.JSON(http.StatusConflict, models.Error{
	//		Description: "Password or adminName error",
	//		Code:        http.StatusConflict,
	//	})
	//	return
	//}

	h.jwthandler.Iss = "moder"
	h.jwthandler.Sub = res.Id
	h.jwthandler.Role = "moder"
	h.jwthandler.Aud = []string{"exam-app"}
	h.jwthandler.SigninKey = h.cfg.SignKey
	h.jwthandler.Log = h.log
	tokens, err := h.jwthandler.GenerateAuthJWT()
	accessToken := tokens[0]

	if err != nil {
		h.log.Error("error occured while generating tokens")
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "something went wrong,please try again",
		})
		return
	}
	res.AccesToken = accessToken
	res.Password = ""
	c.JSON(http.StatusOK, res)
}