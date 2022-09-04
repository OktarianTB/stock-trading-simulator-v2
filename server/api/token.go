package api

import (
	"database/sql"
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type renewAccessTokenRequest struct {
	RefreshToken string `json:"refresh_token" binding:"required"`
}

type renewAccessTokenResponse struct {
	AccessToken          string    `json:"access_token"`
	AccessTokenExpiresAt time.Time `json:"access_token_expires_at"`
}

func (server *Server) renewAccessToken(ctx *gin.Context) {
	var req renewAccessTokenRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		//errResponse := errors.New("invalid input for renewing access token")
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	refreshPayload, err := server.tokenMaker.VerifyToken(req.RefreshToken)
	if err != nil {
		errResponse := errors.New("invalid token")
		ctx.JSON(http.StatusUnauthorized, errorResponse(errResponse))
		return
	}

	session, err := server.store.GetSession(ctx, refreshPayload.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			errResponse := errors.New("no existing session for user")
			ctx.JSON(http.StatusNotFound, errorResponse(errResponse))
			return
		}
		errResponse := errors.New("failed to login user, please try again")
		ctx.JSON(http.StatusInternalServerError, errorResponse(errResponse))
		return
	}

	if session.IsBlocked {
		errResponse := errors.New("blocked session")
		ctx.JSON(http.StatusUnauthorized, errorResponse(errResponse))
		return
	}

	if session.Username != refreshPayload.Username {
		errResponse := errors.New("incorrect session user")
		ctx.JSON(http.StatusUnauthorized, errorResponse(errResponse))
		return
	}

	if session.RefreshToken != req.RefreshToken {
		errResponse := errors.New("mismatched session token")
		ctx.JSON(http.StatusUnauthorized, errorResponse(errResponse))
		return
	}

	if time.Now().After(session.ExpiresAt) {
		errResponse := errors.New("expired session")
		ctx.JSON(http.StatusUnauthorized, errorResponse(errResponse))
		return
	}

	accessToken, accessPayload, err := server.tokenMaker.CreateToken(
		refreshPayload.Username,
		server.config.AccessTokenDuration,
	)
	if err != nil {
		errResponse := errors.New("failed to login user, please try again")
		ctx.JSON(http.StatusInternalServerError, errorResponse(errResponse))
		return
	}

	rsp := renewAccessTokenResponse{
		AccessToken:          accessToken,
		AccessTokenExpiresAt: accessPayload.ExpiredAt,
	}
	ctx.JSON(http.StatusOK, rsp)
}
