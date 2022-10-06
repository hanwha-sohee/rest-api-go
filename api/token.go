package api

import (
	"github.com/gin-gonic/gin"
	"github.com/go-study/rest-api-go/db/gorm/models"
	"net/http"
)

// GET /tokens
// Find all tokens
func (server *Server) getTokens(c *gin.Context) {
	var tokens []models.Token
	server.gorm.DbTokensFind(&tokens)
	//gorm := models.Gorm{}
	//gorm.DbTokenFind(&tokens)
	//models.Gorm{}.DB.Table("token_info_table'").Find(&tokens)

	c.JSON(http.StatusOK, gin.H{"data": tokens})
}

// GET /tokens/:id
// Find a token
func (server *Server) getToken(c *gin.Context) {
	// Get model if exist
	var token models.Token
	var req models.GetTokenRequest
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	server.gorm.DbTokenFind(&token, req.ID)

	c.JSON(http.StatusOK, gin.H{"data": token})
}

// POST /tokens
// Create new token
func (server *Server) CreateToken(c *gin.Context) {
	// Validate input
	var input models.CreateTokenInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create token
	token := models.Token{Id: input.Id, Name: input.Name, Default_amount: input.Default_amount}
	server.gorm.DbTokenCreate(&token)

	c.JSON(http.StatusOK, gin.H{"data": token})
}

// PATCH /tokens/:id
// Update a token
func (server *Server) UpdateToken(c *gin.Context) {
	// Get model if exist
	var token models.Token
	server.gorm.DbTokenFind(&token, c.Param("id"))

	// Validate input
	var input models.UpdateTokenInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	server.gorm.DbTokenUpdate(&token, &input)

	c.JSON(http.StatusOK, gin.H{"data": token})
}

// DELETE /tokens/:id
// Delete a token
func (server *Server) DeleteToken(c *gin.Context) {
	// Get model if exist
	var token models.Token
	server.gorm.DbTokenFind(&token, c.Param("id"))
	server.gorm.DbTokenDelete(&token)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
