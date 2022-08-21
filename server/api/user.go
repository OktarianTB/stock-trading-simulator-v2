package api

import (
	"database/sql"
	"errors"
	"net/http"
	"time"

	db "github.com/OktarianTB/stock-trading-simulator-golang/db/sqlc"
	util "github.com/OktarianTB/stock-trading-simulator-golang/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/lib/pq"
)

type createUserRequest struct {
	Username string `json:"username" binding:"required,alphanum"`
	Password string `json:"password" binding:"required,min=6"`
}

type userResponse struct {
	Username  string    `json:"username"`
	CreatedAt time.Time `json:"created_at"`
	Balance   float64   `json:"balance"`
}

func newUserResponse(user db.User) userResponse {
	return userResponse{
		Username:  user.Username,
		CreatedAt: user.CreatedAt,
		Balance:   user.Balance,
	}
}

func (server *Server) createUser(ctx *gin.Context) {
	var req createUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		errResponse := errors.New("invalid input for creating user")
		ctx.JSON(http.StatusBadRequest, errorResponse(errResponse))
		return
	}

	hashedPassword, err := util.HashPassword(req.Password)
	if err != nil {
		errResponse := errors.New("failed to create user, please try again")
		ctx.JSON(http.StatusInternalServerError, errorResponse(errResponse))
		return
	}

	arg := db.CreateUserParams{
		Username:       req.Username,
		HashedPassword: hashedPassword,
	}

	user, err := server.store.CreateUser(ctx, arg)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "unique_violation":
				errResponse := errors.New("username already exists, please try a different username")
				ctx.JSON(http.StatusBadRequest, errorResponse(errResponse))
				return
			}
		}
		errResponse := errors.New("failed to create user, please try again")
		ctx.JSON(http.StatusInternalServerError, errorResponse(errResponse))
		return
	}

	rsp := newUserResponse(user)
	ctx.JSON(http.StatusOK, rsp)
}

type loginUserRequest struct {
	Username string `json:"username" binding:"required,alphanum"`
	Password string `json:"password" binding:"required,min=6"`
}

type loginUserResponse struct {
	SessionID             uuid.UUID    `json:"session_id"`
	AccessToken           string       `json:"access_token"`
	AccessTokenExpiresAt  time.Time    `json:"access_token_expires_at"`
	RefreshToken          string       `json:"refresh_token"`
	RefreshTokenExpiresAt time.Time    `json:"refresh_token_expires_at"`
	User                  userResponse `json:"user"`
}

func (server *Server) loginUser(ctx *gin.Context) {
	var req loginUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		errResponse := errors.New("invalid input for creating user")
		ctx.JSON(http.StatusBadRequest, errorResponse(errResponse))
		return
	}

	user, err := server.store.GetUser(ctx, req.Username)
	if err != nil {
		if err == sql.ErrNoRows {
			errResponse := errors.New("user does not exist")
			ctx.JSON(http.StatusNotFound, errorResponse(errResponse))
			return
		}
		errResponse := errors.New("failed to login user, please try again")
		ctx.JSON(http.StatusInternalServerError, errorResponse(errResponse))
		return
	}

	err = util.CheckPassword(req.Password, user.HashedPassword)
	if err != nil {
		errResponse := errors.New("invalid password, please try again")
		ctx.JSON(http.StatusUnauthorized, errorResponse(errResponse))
		return
	}

	accessToken, accessPayload, err := server.tokenMaker.CreateToken(
		user.Username,
		server.config.AccessTokenDuration,
	)
	if err != nil {
		errResponse := errors.New("failed to login user, please try again")
		ctx.JSON(http.StatusInternalServerError, errorResponse(errResponse))
		return
	}

	refreshToken, refreshPayload, err := server.tokenMaker.CreateToken(
		user.Username,
		server.config.RefreshTokenDuration,
	)
	if err != nil {
		errResponse := errors.New("failed to login user, please try again")
		ctx.JSON(http.StatusInternalServerError, errorResponse(errResponse))
		return
	}

	session, err := server.store.CreateSession(ctx, db.CreateSessionParams{
		ID:           refreshPayload.ID,
		Username:     user.Username,
		RefreshToken: refreshToken,
		UserAgent:    ctx.Request.UserAgent(),
		ClientIp:     ctx.ClientIP(),
		IsBlocked:    false,
		ExpiresAt:    refreshPayload.ExpiredAt,
	})
	if err != nil {
		errResponse := errors.New("failed to login user, please try again")
		ctx.JSON(http.StatusInternalServerError, errorResponse(errResponse))
		return
	}

	rsp := loginUserResponse{
		SessionID:             session.ID,
		AccessToken:           accessToken,
		AccessTokenExpiresAt:  accessPayload.ExpiredAt,
		RefreshToken:          refreshToken,
		RefreshTokenExpiresAt: refreshPayload.ExpiredAt,
		User:                  newUserResponse(user),
	}
	ctx.JSON(http.StatusOK, rsp)
}
