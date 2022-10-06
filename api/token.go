package api

import (
	"github.com/gin-gonic/gin"
	"github.com/go-study/rest-api-go/db/gorm/models"
	"net/http"
)

type CreateTokenInput struct {
	Title  string `json:"title" binding:"required"`
	Author string `json:"author" binding:"required"`
}

type UpdatetokenInput struct {
	Title  string `json:"title"`
	Author string `json:"author"`
}

// GET /tokens
// Find all tokens
func (server *Server) FindTokens(c *gin.Context) {
	var tokens []models.Token
	server.gorm.DbTokenFind(&tokens)
	//gorm := models.Gorm{}
	//gorm.DbTokenFind(&tokens)
	//models.Gorm{}.DB.Table("token_info_table'").Find(&tokens)

	c.JSON(http.StatusOK, gin.H{"data": tokens})
}

// GET /tokens/:id
// Find a token
func FindToken(c *gin.Context) {
	// Get model if exist
	var token models.Token
	//if err := models.DB.Where("id = ?", c.Param("id")).First(&token).Error; err != nil {
	//	c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
	//	return
	//}

	c.JSON(http.StatusOK, gin.H{"data": token})
}

// POST /tokens
// Create new token
//func CreateToken(c *gin.Context) {
//	// Validate input
//	var input CreateTokenInput
//	if err := c.ShouldBindJSON(&input); err != nil {
//		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//		return
//	}
//
//	// Create token
//	token := models.Token{Title: input.Title, Author: input.Author}
//	models.DB.Create(&token)
//
//	c.JSON(http.StatusOK, gin.H{"data": token})
//}

// PATCH /tokens/:id
// Update a token
//func UpdateToken(c *gin.Context) {
//	// Get model if exist
//	var token models.Token
//	if err := models.DB.Where("id = ?", c.Param("id")).First(&token).Error; err != nil {
//		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
//		return
//	}
//
//	// Validate input
//	var input UpdateTokenInput
//	if err := c.ShouldBindJSON(&input); err != nil {
//		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//		return
//	}
//
//	models.DB.Model(&token).Updates(input)
//
//	c.JSON(http.StatusOK, gin.H{"data": token})
//}

// DELETE /tokens/:id
// Delete a token
//func DeleteToken(c *gin.Context) {
//	// Get model if exist
//	var token models.Token
//	if err := models.DB.Where("id = ?", c.Param("id")).First(&token).Error; err != nil {
//		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
//		return
//	}
//
//	models.DB.Delete(&token)
//
//	c.JSON(http.StatusOK, gin.H{"data": true})
//}
