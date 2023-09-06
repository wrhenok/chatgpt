package controller

import (
	"chatgpt3/src/config"
	"chatgpt3/src/domain"
	"chatgpt3/src/mapper"
	"chatgpt3/src/service"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

func HttpServer() {

	//chatgpt3.5
	http.HandleFunc("/up", func(writer http.ResponseWriter, request *http.Request) {
		domain.Start = time.Now()
		var params map[string]string //接收前端参数的容器
		body := json.NewDecoder(request.Body)
		body.Decode(&params)
		fmt.Println("input: ", params)
		chat, err := service.GetResp(params["params"])
		if err != nil {
			fmt.Printf("get failed,err :%s", err)
			return
		}
		end := time.Since(domain.Start).String()
		fmt.Println(end)
		insert, _ := mapper.Insert(&chat, end)
		if insert <= 0 {
			fmt.Println("insert data failed")
			return
		}
		fmt.Println("insert data success, chatID: ", insert)
		io.WriteString(writer, chat.Answer)
		defer request.Body.Close()
	})

	//chatgpt3
	http.HandleFunc("/chat", func(writer http.ResponseWriter, request *http.Request) {
		domain.Start = time.Now()
		client := service.NewClient(config.ApiKey)
		ctx := context.Background()
		var params map[string]string //接收前端参数的容器
		body := json.NewDecoder(request.Body)
		body.Decode(&params)
		//fmt.Println("input", params)
		chat, err := client.GetAnswer(ctx, params["prompt"])
		if err != nil {
			fmt.Printf("getAnswer failed,err :%s", err)
			return
		}
		end := time.Since(domain.Start).String()
		fmt.Println(end)
		insert, _ := mapper.Insert(&chat, end)
		if insert <= 0 {
			fmt.Println("insert data failed")
			return
		}
		fmt.Println("insert data success, chatID: ", insert)
		io.WriteString(writer, chat.Answer)
		defer request.Body.Close()
	})

	http.HandleFunc("/common", func(writer http.ResponseWriter, request *http.Request) {
		files, err := ioutil.ReadFile("./templates/page.html")
		if err != nil {
			fmt.Fprintln(writer, "Read failed, err : ", err)
		}
		writer.Write(files)
	})

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		files, err := ioutil.ReadFile("./templates/stream.html")
		if err != nil {
			fmt.Fprintln(writer, "Read failed, err : ", err)
		}
		writer.Write(files)
	})

	http.HandleFunc("/ccc", func(writer http.ResponseWriter, request *http.Request) {
		chatAll := mapper.Select()
		var chat []string
		for _, chatlist := range chatAll {
			marshal, err := json.Marshal(chatlist)
			if err != nil {
				fmt.Println("json failed", err)
				return
			}
			chat = append(chat, string(marshal))
		}
		fmt.Println(chat)
	})

	http.ListenAndServe(config.Server.Port, nil)
}
