package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type createTokenRequest struct {
	Name string `json:"name" binding:"required""`
	//Id string `json:"id" binding:"required,oneof=USD EUR""`
	Id     string `json:"id" binding:"required"`
	Amount int    `json:"default_Amount" binding:"required"`
}

func (server *Server) createToken(ctx *gin.Context) {
	var req createTokenRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	//arg := db.CreateAccountParams{
	//	Name:     req.Name,
	//	Currency: req.Currency,
	//	Balance:  0,
	//}

	//account, err := server.store.CreateAccount(ctx, arg)
	//if err != nil {
	//	ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	//	return
	//}
	//ctx.JSON(http.StatusOK, account)
}
