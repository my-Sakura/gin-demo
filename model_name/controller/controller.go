package controller

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Controller struct{}

func New() *Controller {
	return &Controller{}
}

func (c *Controller) Regist(r gin.IRouter) error {
	r.GET("/demo", c.demo)

	return nil
}

func (con *Controller) demo(c *gin.Context) {
	var req struct {
		ID string `json:"id" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest})
		log.Printf("Error bind request: %v\n", err)
		return
	}

	fmt.Println(req.ID, "receive succeed")

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK})
}
