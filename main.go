package main

import (
	"context"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

type Post_ig struct {
	Username    string `json:"username"`
	Caption     string `json:"caption"`
	Jumlahlike  string `json:"jumlah like"`
	Jumlahkomen string `json:"jumlah komen"`
}

func main() {
	r := gin.Default()

	p := r.Group("/post_ig")
	{
		p.GET("/", Get)
		p.GET("/:id", GetByID)
		p.POST("/create", Create)
		p.PUT("/:id", Update)
		p.DELETE("/:id", Delete)
	}
	r.Run(":8080")
}

func Get(c *gin.Context) {
	post, err := integration.Mongodb()
	if err != nil {
		c.String(500, err.Error())
		return
	}

	result, err := post.Database("integration").Collection("post").Find(context.Background(), bson.M{})
	if err != nil {
		c.String(500, err.Error())
		return
	}

	var data []map[string]interface{}
	result.All(context.Background(), &data)

	c.JSON(200, data)
}

func GetByID(c *gin.Context) {
	id := c.Param("id")

	post, err := integration.Mongodb()
	if err != nil {
		c.String(500, err.Error())
		return
	}

	result, err := post.Database("integration").Collection("post").Find(context.Background(), bson.M{"username": id})
	if err != nil {
		c.String(500, err.Error())
		return
	}

	var data []map[string]interface{}
	result.All(context.Background(), &data)

	c.JSON(200, data)
}

func Create(c *gin.Context) {
	var post_ig Post_ig
	c.BindJSON(&post)

	post, err := integration.Mongodb()
	if err != nil {
		c.String(500, err.Error())
		return
	}

	_, err = post.Database("integration").Collection("post").InsertOne(context.Background(), post_ig)
	if err != nil {
		c.String(500, err.Error())
		return
	}

	c.String(200, "success")
}

func Update(c *gin.Context) {
	username := c.Param("id")

	var post_ig Post_ig
	c.BindJSON(&post_ig)

	post, err := integration.Mongodb()
	if err != nil {
		c.String(500, err.Error())
		return
	}

	_, err = post.Database("integration").Collection("post").UpdateOne(context.Background(), bson.M{"username": username}, bson.M{"$set": post_ig})
	if err != nil {
		c.String(500, err.Error())
		return
	}

	c.String(200, "success")
}

func Delete(c *gin.Context) {
	id := c.Param("id")

	post, err := integration.Mongodb()
	if err != nil {
		c.String(500, err.Error())
		return
	}

	_, err = post.Database("integration").Collection("post").DeleteOne(context.Background(), bson.M{"username": id})
	if err != nil {
		c.String(500, err.Error())
		return
	}

	c.String(200, "success")
}
