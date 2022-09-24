package express

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type HttpResponse struct {
	Status string         `json:"status"`
	Msg    string         `json:"msg"`
	Result HttpRespResult `json:"result"`
}
type HttpRespResult struct {
	Number     string     `json:"number"`
	Type       string     `json:"type"`
	List       []RespList `json:"list"`
	ExpName    string     `json:"expName"`
	UpdateTime string     `json:"updateTime"`
	TakeTime   string     `json:"takeTime"`
}
type RespList struct {
	Time   string `json:"time"`
	Status string `json:"status"`
}

type Result struct {
	Status  int64  //返回1说明查询成果，其它均为查询失败，进入查询失败处理逻辑
	Res_str string //直接作为机器人消息返回
}

var res *Result
var httpres HttpResponse

func init() {
	res = new(Result)
}
func GetInfo(no string) *Result {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", "http://wuliu.market.alicloudapi.com/kdi?no="+no, nil)
	req.Header.Add("Authorization", "APPCODE 73fb5aa43a784e6e8b4aab05947ad6df")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(body, &httpres)
	if err != nil {
		log.Fatal(err)
	}
	if httpres.Status == "0" {
		//查询成果，返回序列化信息
		res.Status = 1
		res.Res_str = "您的快递单号:" + no + "\n查询结果如下:\n"
		res.Res_str += "快递类型:" + httpres.Result.ExpName + "\n"
		res.Res_str += "总用时:" + httpres.Result.TakeTime + "\n"
		for i := range httpres.Result.List {
			res.Res_str += httpres.Result.List[i].Time + ":" + httpres.Result.List[i].Status + "\n"
		}
	} else {
		res.Status = 0
	}
	return res
}
