package main

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"net/http"
	"qrobot/config"
	"qrobot/express"
	"qrobot/sms"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/tencent-connect/botgo"
	"github.com/tencent-connect/botgo/dto"
	"github.com/tencent-connect/botgo/event"
	"github.com/tencent-connect/botgo/openapi"
	"github.com/tencent-connect/botgo/token"
	"github.com/tencent-connect/botgo/websocket"
)

var ctx context.Context
var api openapi.OpenAPI

func init() {
}
func main() {
	conf := config.Conf
	botToken := token.BotToken(conf.AppID, conf.Token)
	api = botgo.NewOpenAPI(botToken)
	ctx = context.Background()
	wsInfo, err := api.WS(ctx, nil, "")
	if err != nil {
		log.Fatalln("wsinfo请求错误" + err.Error())
	}
	log.Printf("%+v, err:%v", wsInfo, err)
	go get_back(wsInfo, botToken)
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"code": 200,
			"msg":  "ok",
		})
	})
	r.Run(":9000")
}
func get_back(wsInfo *dto.WebsocketAP, botToken *token.Token) {
	intent := websocket.RegisterHandlers(
		ReadyHandler(),
		AtMessageHandler(),
	)
	if err := botgo.NewSessionManager().Start(wsInfo, botToken, &intent); err != nil {
		log.Fatalln(err)
	}
}
func ReadyHandler() event.ReadyHandler {
	return func(event *dto.WSPayload, data *dto.WSReadyData) {
		log.Println("ready event receive:", data)
	}
}

func AtMessageHandler() event.ATMessageEventHandler {
	return func(event *dto.WSPayload, data *dto.WSATMessageData) error {
		fmt.Println(data.Content)
		if strings.Contains(data.Content, "/快递查询") {
			temp := strings.Fields(data.Content)
			no := temp[len(temp)-1]
			if len(temp) != 3 {
				_, err := api.PostMessage(ctx, data.ChannelID, &dto.MessageToCreate{
					Content: "输入格式错误，格式: /快递查询 单号",
					MsgID:   data.ID,
				})
				if err != nil {
					return err
				}
				return nil
			}
			rsp := express.GetInfo(no)
			//回复消息
			if rsp.Status == 1 {
				_, err := api.PostMessage(ctx, data.ChannelID, &dto.MessageToCreate{
					Content: rsp.Res_str,
					MsgID:   data.ID,
				})
				if err != nil {
					return err
				}
			} else {
				_, err := api.PostMessage(ctx, data.ChannelID, &dto.MessageToCreate{
					Content: "查询失败，请检查单号，顺丰快递请在快递单号后加手机号后4位（单号:手机尾号），如123877813:1234",
					MsgID:   data.ID,
				})
				if err != nil {
					return err
				}
			}
			//快递查询逻辑结束
			return nil
		} else if strings.Contains(data.Content, "/看美女") {
			//发送模板消息
			reader := bytes.NewReader([]byte("123"))
			request, _ := http.NewRequest("GET", "https://qrobot.dadiqq.cn", reader)
			client := &http.Client{}
			client.Do(request)
			_, err := api.PostMessage(ctx, data.ChannelID, &dto.MessageToCreate{
				Content: "美女来咯",
				Image:   "https://qrobot.dadiqq.cn/pic/1.jpg",
				MsgID:   data.ID,
			})
			if err != nil {
				client.Do(request)
				time.Sleep(time.Second * 5)
				api.PostMessage(ctx, data.ChannelID, &dto.MessageToCreate{
					Content: "美女来咯",
					Image:   "https://qrobot.dadiqq.cn/pic/1.jpg",
					MsgID:   data.ID,
				})
			}
		} else if strings.Contains(data.Content, "/短信轰炸") {
			temp := strings.Fields(data.Content)
			phone := temp[len(temp)-1]
			res := sms.StartSmsboom(phone)
			api.PostMessage(ctx, data.ChannelID, &dto.MessageToCreate{
				Content: "status:" + fmt.Sprintf("%d", res.Status) + "\nMsg:" + res.Msg + "\n",
				MsgID:   data.ID,
			})
			//获取phone后直接请求boom

		}

		return nil
	}
}
