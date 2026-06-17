package callbackrouters

import (
	"github.com/gin-gonic/gin"
	"github.com/juggleim/jugglechat-server-assistant/callbackapis"
)

var AssistantUrlPrefix string

func Route(eng *gin.Engine, prefix string) *gin.RouterGroup {
	AssistantUrlPrefix = prefix
	group := eng.Group("/" + prefix)

	group.POST("/msgcallback", callbackapis.AssistantMsgCallback)

	return group
}
