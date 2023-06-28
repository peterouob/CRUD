package service

import (
	"golang-mysql/mysql"
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Address string `json:"address"`
}

func (u *User) TableName() string {
	return "user"
}

func ShowAllUser(c *gin.Context) {
	db := mysql.DB
	var result []User
	db.Find(&result)
	c.JSON(http.StatusOK, gin.H{
		"result": result,
	})
}

func CreateUser(c *gin.Context) {
	db := mysql.DB
	user := &User{}
	c.ShouldBind(&user)
	db.Create(user)
	c.JSON(http.StatusOK, gin.H{
		"msg": "success",
	})
}

func UpdateUser(c *gin.Context) {
	userID := c.Param("id")
	db := mysql.DB
	var userData User
	if err := c.ShouldBindJSON(&userData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.Model(&User{}).Where("id = ?", userID).Updates(userData).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User updated successfully"})
}

func DeleteUser(c *gin.Context) {
	userID := c.Param("id")
	db := mysql.DB
	var user User
	if err := db.First(&user, "id=?", userID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	if err := db.Where("id=?", userID).Delete(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}
