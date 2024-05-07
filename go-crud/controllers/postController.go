package controllers

import (
	"github.com/MaaHiN15/go-practice/go-crud/initializers"
	"github.com/MaaHiN15/go-practice/go-crud/models"
	"github.com/gin-gonic/gin"
)

func PingPong(con *gin.Context) { con.JSON(200, gin.H{ "message": "pong" }) };

func PostCreate(con *gin.Context) {
	var PostBody struct { Title string; Body string }
	con.Bind(&PostBody)
	post := models.Post{Title: PostBody.Title, Body: PostBody.Body}
	result := initializers.DB.Create(&post)
	if result.Error != nil { con.Status(400); return }
	con.JSON(200, gin.H{ "post": post })
};

func GetPosts(con *gin.Context) {
	var posts []models.Post
	initializers.DB.Find(&posts)
	con.JSON(200, gin.H{ "posts": posts })
};

func PostShow(con *gin.Context) {
	id := con.Param("id")
	var post models.Post
	initializers.DB.Find(&post, id)
	con.JSON(200, gin.H{ "post": post })
};

func PostUpdate(con *gin.Context) {
	var PostBody struct { Title string; Body string }
	con.Bind(&PostBody)
	id := con.Param("id")
	var post models.Post
	initializers.DB.Find(&post, id)
	initializers.DB.Model(&post).Updates(models.Post{Title: PostBody.Title, Body: PostBody.Body})
	con.JSON(200, gin.H{ "post": post })
};

func PostDelete(con *gin.Context) {
	id := con.Param("id")
	initializers.DB.Delete(&models.Post{}, id)
	con.JSON(200, gin.H{ "post": "Deleted" })
}