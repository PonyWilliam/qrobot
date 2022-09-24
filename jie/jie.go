package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func init() {

}
func GetHost() {
	req, _ := http.NewRequest("GET", "https://api.julym.com/xjj/", nil)
	c := http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
		Timeout: 5 * time.Second,
	}
	resp, err := c.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resp)
}
func main() {
	GetHost()
}
