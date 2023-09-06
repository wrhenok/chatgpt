package controller

import (
	"chatgpt3/src/Util"
	"chatgpt3/src/config"
	"chatgpt3/src/domain"
	"chatgpt3/src/mapper"
	"chatgpt3/src/service"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/sashabaranov/go-openai"
	"io"
	"log"
	"net/http"
	"time"
)

func StreamServer() {
	//chatgpt3.5 stream
	http.HandleFunc("/stream", func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "text/event-stream")
		writer.Header().Set("Cache-Control", "no-cache")
		writer.Header().Set("Connection", "keep-alive")
		writer.Header().Set("Access-Control-Allow-Origin", "*")

		domain.Start = time.Now()
		body := request.URL.Query()["prompt"]
		params, _ := json.Marshal(body)
		fmt.Println("input: ", string(params))

		flusher, ok := writer.(http.Flusher)
		if !ok {
			log.Panic("server not support")
		}

		c := openai.NewClient(config.ApiKey)
		ctx := context.Background()
		req := openai.ChatCompletionRequest{
			Model:     openai.GPT3Dot5Turbo,
			MaxTokens: 2048,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: string(params),
				},
			},
			Stream:      true,
			Temperature: 0.9,
		}
		stream, err := c.CreateChatCompletionStream(ctx, req)
		if err != nil {
			fmt.Printf("ChatCompletionStream error: %v\n", err)
			return
		}
		defer stream.Close()
		fmt.Printf("Stream response: ")
		var (
			resp     []byte
			response openai.ChatCompletionStreamResponse
			answer   string
		)
		for {
			count := 0
			response, err = stream.Recv()
			if errors.Is(err, io.EOF) {
				fmt.Println("\nStream finished")
				fmt.Fprintf(writer, "event: close\ndata: close\n\n")

				end := time.Since(domain.Start).String()
				fmt.Println(end)
				key := config.ApiKey
				ip, _ := Util.GetIp()

				chat := domain.Chatgpt3{
					Username:         "admin",
					Userkey:          key,
					Model:            openai.GPT3Dot5Turbo,
					Prompt:           string(params),
					Answer:           answer,
					Isend:            0,
					RequestIpaddress: ip,
					ResponseJson:     string(resp),
					PromptId:         response.ID,
					FinishReason:     "stop",
					RequestUrl:       service.ResqUrl(),
					MaxTokens:        req.MaxTokens,
					Temperature:      float64(req.Temperature),
					ThreadID:         Util.GetGoID(),
				}
				count++
				fmt.Println("chat: ", chat)
				insert, _ := mapper.Insert(&chat, end)
				if insert <= 0 {
					fmt.Println("insert data failed")
					return
				}
				fmt.Println("insert data success, chatID: ", insert)

				flusher.Flush()
				return
			}
			if err != nil {
				fmt.Printf("\nStream error: %v\n", err)
				return
			}
			fmt.Println(response.Choices[0].Delta.Content)
			fmt.Fprintf(writer, "data: %s\n\n", response.Choices[0].Delta.Content)
			answer += response.Choices[0].Delta.Content
			count++ //下标
			flusher.Flush()
			marshal, _ := json.Marshal(response)
			resp = append(resp, marshal...)
		}
		defer request.Body.Close()
	})
}
