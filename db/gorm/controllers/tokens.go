package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-study/rest-api-go/db/gorm/models"
)

// GET /tokens
// Get all tokens
func FindTokens(c *gin.Context) {
	var tokens []models.Token
	fmt.Println(tokens)
	g := models.Gorm{}
	fmt.Println(g)

	//g._db 원래는 이렇게 하면 잡혀야 되는 거 아닌가욥?

}
