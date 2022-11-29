package v1

import (
	"context"
	"strconv"
	"net/http"
	"time"

	models "gitlab.com/pro/exam_api/api/models"
	pbs "gitlab.com/pro/exam_api/genproto/custumer_proto"
	"gitlab.com/pro/exam_api/pkg/logger"

	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/encoding/protojson"
)

// Get Accsess token
// @Summary      Get Accsess token
// @Description  This API generates new access token
// @Tags         Token
// @Accept       json
// @Produce      json
// @Param        token  query string true "Refresh Token"
// @Success      200  "object"
// @Router      /token [get]
func (h *handlerV1) GetAccesToken(c *gin.Context) {
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	var (
		refreshToken1 = c.Query("token")
	)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()
	res, err := h.serviceManager.CustumerService().GetByRefreshToken(ctx, &pbs.RefreshTokenReq{
		Token: refreshToken1,
	})
	id := strconv.FormatInt(int64(res.Id), 10)
  
	if err != nil {
		c.JSON(http.StatusGatewayTimeout, gin.H{
			"info":  "Wrong Token",
			"error": err.Error(),
		})
		h.log.Error("Error while refreshing token", logger.Any("refresh token", err))
	}
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
	response := models.UpdateAccessToken{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	c.JSON(http.StatusOK, response)
}

