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

func Route(group *gin.RouterGroup) {
	//registe user register event
	events.RegisteUserRegisteEvent(func(appkey string, user jchatModels.User) {
		appinfo, exist := appinfos.GetAppInfo(appkey)
		if exist && appinfo != nil {
			if exist, obj := appinfo.GetExt("open_ai_assistant"); exist && obj != nil {
				objStr := obj.(string)
				openAssistant := String2Bool(objStr)
				if openAssistant {
					//register assistant
					assistantId := "assistant_" + user.UserId
					sdk := imsdk.GetImSdk(appkey)
					sdk.RegisterBot(juggleimsdk.BotInfo{
						BotId:    assistantId,
						Nickname: user.Nickname + "'s Assistant",
						Portrait: user.UserPortrait,
						BotConf: &juggleimsdk.BotConf{
							Url: configures.Config.CallbackBaseUrl + "/" + callbackrouters.AssistantUrlPrefix + "/msgcallback",
						},
					})
					sdk.SendPrivateMsg(juggleimsdk.Message{
						SenderId:   assistantId,
						TargetId:   user.UserId,
						MsgType:    "jg:text",
						MsgContent: `{"content":"hello"}`,
					})
				}
			}
		}
	})
}

func String2Bool(str string) bool {
	b, err := strconv.ParseBool(str)
	if err == nil {
		return b
	}
	return false
}
