package routers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	juggleimsdk "github.com/juggleim/imserver-sdk-go"
	"github.com/juggleim/jugglechat-server-assistant/callbackrouters"
	"github.com/juggleim/jugglechat-server/commons/appinfos"
	"github.com/juggleim/jugglechat-server/commons/configures"
	"github.com/juggleim/jugglechat-server/commons/imsdk"
	"github.com/juggleim/jugglechat-server/events"
	jchatModels "github.com/juggleim/jugglechat-server/storages/models"
)

func Route(eng *gin.Engine, prefix string) *gin.RouterGroup {
	//registe user register event
	events.RegisteUserRegisteEvent(func(appkey string, user jchatModels.User) {
		appinfo, exist := appinfos.GetAppInfo(appkey)
		if exist && appinfo != nil {
			if exist, obj := appinfo.GetExt("open_ai_assistant"); exist && obj != nil {
				objStr := obj.(string)
				openAssistant := String2Bool(objStr)
				if openAssistant {
					//register assistant
					sdk := imsdk.GetImSdk(appkey)
					sdk.RegisterBot(juggleimsdk.BotInfo{
						BotId:    "assistant_" + user.UserId,
						Nickname: user.Nickname + "'s Assistant",
						Portrait: user.UserPortrait,
						BotConf: &juggleimsdk.BotConf{
							Url: configures.Config.CallbackUrl + callbackrouters.AssistantUrlPrefix + "/msgcallback",
						},
					})
				}
			}
		}
	})

	group := eng.Group("/" + prefix)

	return group
}

func String2Bool(str string) bool {
	b, err := strconv.ParseBool(str)
	if err == nil {
		return b
	}
	return false
}
