package v1

import (
	"github.com/freelifer/coolgo/app/service"
	"github.com/freelifer/coolgo/models"
	"github.com/freelifer/coolgo/pkg/redis"
	"github.com/freelifer/coolgo/pkg/utils"
	"github.com/gin-gonic/gin"
)

func WeiXinLogin(c *gin.Context) {
	code := c.Query("code")

}
