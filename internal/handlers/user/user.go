package user

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"krp_project/internal/dto"
	"krp_project/internal/handlers/request"
	"krp_project/internal/services"
	"net/http"
)

type UserHandler struct {
	userService *services.UserService
}

func (h *UserHandler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.POST("/auth", h.Auth)
	router.POST("/register", h.Register)

	return router
}

func NewUserHandler(srv *services.UserService) (*UserHandler, error) {
	h := &UserHandler{
		userService: srv,
	}
	return h, nil
}

func (h *UserHandler) Auth(c *gin.Context) {
	req, ok := request.GetRequest[dto.AuthRequest](c)
	logrus.Debug(req)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "auth request error", "text": ok})
		return
	}
	user, userTokens, err := h.userService.Auth(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "auth service error", "text": err.Error()})
		return
	}

	resp := dto.AuthResponse{
		Email:                 user.Email,
		Password:              user.Password,
		AccessTokenString:     userTokens[0].TokenString,
		AccessTokenExpiresAt:  userTokens[0].ExpiresAt,
		AccessTokenIssuedAt:   userTokens[0].IssuedAt,
		RefreshTokenString:    userTokens[1].TokenString,
		RefreshTokenExpiresAt: userTokens[1].ExpiresAt,
		RefreshTokenIssuedAt:  userTokens[1].IssuedAt,
	}
	logrus.Debug(resp)
	c.JSON(http.StatusOK, gin.H{"status": "ok", "data": resp})
}

func (h *UserHandler) Register(c *gin.Context) {
	req, ok := request.GetRequest[dto.RegisterRequest](c)
	logrus.Debug(req)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "register request error", "text": ok})
		return
	}
	userResp, err := h.userService.Register(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "register service error", "text": err.Error()})
		return
	}
	logrus.Debug(userResp)
	c.JSON(http.StatusOK, gin.H{"status": "ok", "data": userResp})
}

// TODO: потом
func (h *UserHandler) UpdateAccessToken(c *gin.Context) {

}
