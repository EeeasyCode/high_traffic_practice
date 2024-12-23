package network

import (
	"github.com/gin-gonic/gin"
	"high-traffic-practice/types"
	"net/http"
)

func (n *Network) Login(c *gin.Context) {
	var req types.LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	} else if res, err := n.service.CreateAuth(req.Username); err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
	} else {
		c.JSON(http.StatusOK, res)
	}
}

func (n *Network) verify(c *gin.Context) {
	c.JSON(http.StatusOK, "success")
}
