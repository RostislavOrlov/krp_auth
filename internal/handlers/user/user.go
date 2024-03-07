package user

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"krp_project/internal/dto"
	"krp_project/internal/entities"
	"krp_project/internal/handlers/request"
	"krp_project/internal/middleware"
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
	router.POST("/update_access_token",
		middleware.AuthMiddleware(), h.UpdateAccessToken)

	//router.Use(cors.New(cors.Config{
	//	AllowOrigins:     []string{"http://localhost:8081"},
	//	AllowMethods:     []string{"GET", "POST"},
	//	AllowHeaders:     []string{"Origin"},
	//	AllowCredentials: true,
	//}))

	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowCredentials = true
	config.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"}
	config.AllowHeaders = []string{"Access-Control-Allow-Headers", "Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization", "Accept", "Origin", "Cache-Control", "X-Requested-With"}
	router.Use(cors.New(config))
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
	usr, err := h.userService.Register(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "register service error", "text": err.Error()})
		return
	}
	logrus.Debug(usr)

	resp := dto.RegisterResponse{
		Id:         usr.Id,
		LastName:   usr.LastName,
		FirstName:  usr.FirstName,
		MiddleName: usr.MiddleName,
		Email:      usr.Email,
		Password:   usr.Password,
		Role:       usr.Role,
	}
	c.JSON(http.StatusOK, gin.H{"status": "ok", "data": resp})
}

func (h *UserHandler) UpdateAccessToken(c *gin.Context) {
	//req, ok := request.GetRequest[dto.UpdateAccessTokenRequest](c)
	accessToken := c.MustGet("access_token").(string)
	refreshToken, err := c.Cookie("refresh_token")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "error refresh token parsing from cookie", "text": err.Error()})
		return
	}

	//logrus.Debug(req)
	//if !ok {
	//	c.JSON(http.StatusBadRequest, gin.H{"error": "update access token request error", "text": ok})
	//	return
	//}

	tokens, err := h.userService.UpdateAccessToken(&entities.Tokens{
		AccessTokenString:  accessToken,
		RefreshTokenString: refreshToken,
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "update access token service error", "text": err.Error()})
		return
	}
	logrus.Debug(tokens)

	resp := dto.UpdateAccessTokenResponse{
		AccessTokenString:  tokens[0].TokenString,
		RefreshTokenString: tokens[1].TokenString,
	}

	c.JSON(http.StatusOK, gin.H{"status": "ok", "data": resp})
}
