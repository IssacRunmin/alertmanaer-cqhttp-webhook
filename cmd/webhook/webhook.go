package main

import (
	"flag"
	"net/http"
	"os"
	"github.com/gin-gonic/gin"
	model "github.com/IssacRunmin/alertmanaer-cqhttp-webhook/model"
	"github.com/IssacRunmin/alertmanaer-cqhttp-webhook/notifier"
)

var (
	h            bool
	defaultRobot string
	listenPort string
)

func init() {
	flag.BoolVar(&h, "h", false, "help")
	flag.StringVar(&defaultRobot, "defaultRobot", "", "global cqhttp robot webhook, you can overwrite by alert rule with annotations cqRobot")
	flag.StringVar(&listenPort, "port", "5601", "global webhook listening port, default 5601")
}

func main() {

	flag.Parse()

	if h {
		flag.Usage()
		return
	}
	err := os.Setenv("PORT", listenPort)
	if err != nil {
		return 
	}
	router := gin.Default()
	router.POST("/webhook", func(c *gin.Context) {
		var notification model.Notification

		err := c.BindJSON(&notification)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		err = notifier.Send(notification, defaultRobot)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		}

		c.JSON(http.StatusOK, gin.H{"message": "send to cqhttp successful!"})

	})
	router.Run()
}
