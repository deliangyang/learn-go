package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	//for i := 0; i < 100; i++ {
	//	newPostMultiGift()
	//}
	http.HandleFunc("/", postMultiGift)
	err := http.ListenAndServe(":12345", nil)
	if err != nil {
		fmt.Println("ListenAndServe: ", err)
	}
}

type GiftPack struct {
	GiftId int `json:"giftId"`
	Number int `json:"number"`
}

func newPostMultiGift() {

	client := &http.Client{}
	giftPack := [1]GiftPack{
		GiftPack{
			GiftId: 1,
			Number: 10,
		},
	}
	gifts, err := json.Marshal(giftPack)
	if err != nil {
		panic(err)
	}
	content := strings.NewReader("userId=11577607&roomId=8&taskId=xxx&gifts=" + string(gifts))
	req, err := http.NewRequest("POST", "http://www.ydl.com:9090/api/user/multi/gifts", content)
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("app-version", "2.1.0")
	req.Header.Set("X-HTTP-Method-Override", "PATCH")
	req.Header.Set("x-api-test", "1")
	req.Header.Set("authorize-token", "MTA4ODIyOTksMTU2MzUxMDQ5MCxmOWE4MWYzYmZhNzFkNGMxOGY1NGU5ZWRlM2FkODA5Yw==")
	resp, err := client.Do(req)
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	println(string(body))
}

func postMultiGift(w http.ResponseWriter, reqa *http.Request) {

	client := &http.Client{}
	giftPack := [1]GiftPack{
		GiftPack{
			GiftId: 1,
			Number: 10,
		},
	}
	gifts, err := json.Marshal(giftPack)
	if err != nil {
		panic(err)
	}
	content := strings.NewReader("userId=11577607&roomId=8&taskId=xxx&gifts=" + string(gifts))
	req, err := http.NewRequest("POST", "http://www.ydl.com:9090/api/user/multi/gifts", content)
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("app-version", "2.1.0")
	req.Header.Set("X-HTTP-Method-Override", "PATCH")
	req.Header.Set("x-api-test", "1")
	req.Header.Set("authorize-token", "MTA4ODIyOTksMTU2MzUxMDQ5MCxmOWE4MWYzYmZhNzFkNGMxOGY1NGU5ZWRlM2FkODA5Yw==")
	resp, err := client.Do(req)
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	println(string(body))
	w.Write(body)
}
