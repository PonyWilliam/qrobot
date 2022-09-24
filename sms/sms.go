package sms

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type Result struct {
	Status int64  `json:"status"` //返回1说明查询成果，其它均为查询失败，进入查询失败处理逻辑
	Msg    string `json:"msg"`    //直接作为机器人消息返回
}

var res Result

func StartSmsboom(phone string) Result {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", "https://service-6poquwqo-1257689370.gz.apigw.tencentcs.com/phone/"+phone+"?key=qrobot_dadi", nil)
	req.Header.Add("Authorization", "APPCODE 73fb5aa43a784e6e8b4aab05947ad6df")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(body, &res)
	return res
}
