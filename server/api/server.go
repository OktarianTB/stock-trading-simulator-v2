package api

import (
	"fmt"

	db "github.com/OktarianTB/stock-trading-simulator-golang/db/sqlc"
	"github.com/OktarianTB/stock-trading-simulator-golang/token"
	util "github.com/OktarianTB/stock-trading-simulator-golang/utils"
	"github.com/gin-gonic/gin"
)

type Server struct {
	config     util.Config
	store      db.Store
	tokenMaker token.Maker
	router     *gin.Engine
}

func NewServer(config util.Config, store db.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %v", err)
	}
	server := &Server{
		config:      config,
		store:      store,
		tokenMaker: tokenMaker,
	}

	server.setupRouter()
	return server, nil
}

func (server *Server) setupRouter() {
	router := gin.Default()
	baseRouter := router.Group("/api/v1/")

	baseRouter.POST("/users", server.createUser)
	baseRouter.POST("/users/login", server.loginUser)
	baseRouter.POST("/tokens/renew_access", server.renewAccessToken)

	//authRouter := router.Group("/").Use(authMiddleware(server.tokenMaker))

	server.router = router
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
