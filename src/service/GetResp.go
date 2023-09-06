package service

import (
	"chatgpt3/src/Util"
	"chatgpt3/src/config"
	"chatgpt3/src/domain"
	"context"
	"encoding/json"
	"fmt"
	gpt "github.com/sashabaranov/go-openai"
	"log"
	"strings"
)

var client *gpt.Client

func ResqUrl() string {
	return fmt.Sprintf("%s%s", config.Url.BaseUrl, config.Url.SuffixUrl)
}

func GetResp(prompt string) (chat domain.Chatgpt3, err error) {
	//conf := config.DefaultConfig(config.ApiKey)
	//proxyUrl, err := url.Parse(ResqUrl())
	//if err != nil {
	//	panic(err)
	//}
	//transport := &http.Transport{
	//	Proxy: http.ProxyURL(proxyUrl),
	//}
	//conf.HTTPClient = &http.Client{
	//	Transport: transport,
	//}
	//fmt.Println(proxyUrl)
	//client = gpt.NewClientWithConfig(conf)
	client = gpt.NewClient(config.ApiKey)
	fmt.Printf("client: %v\n", *client)

	req := gpt.ChatCompletionRequest{
		Model: gpt.GPT3Dot5Turbo,
		Messages: []gpt.ChatCompletionMessage{
			{
				Role:    gpt.ChatMessageRoleUser,
				Content: prompt,
			},
		},
		MaxTokens:   2048,
		Temperature: 0.1,
	}

	resp, err := client.CreateChatCompletion(context.Background(), req)

	if err != nil {
		fmt.Printf("CreateChatCompletion error: %v\n", err)
		log.Println(err)
		return
	}
	fmt.Println("responseJSON:", resp)
	answer := strings.TrimSpace(resp.Choices[0].Message.Content)
	marshal, _ := json.Marshal(resp)
	ip, _ := Util.GetIp()
	chat = domain.Chatgpt3{
		Username:         "admin",
		Userkey:          config.ApiKey,
		Model:            gpt.GPT3Dot5Turbo,
		Prompt:           prompt,
		Answer:           answer,
		Isend:            resp.Choices[0].Index,
		RequestIpaddress: ip,
		ResponseJson:     string(marshal),
		PromptId:         resp.ID,
		FinishReason:     resp.Choices[0].FinishReason,
		RequestUrl:       ResqUrl(),
		MaxTokens:        req.MaxTokens,
		Temperature:      float64(req.Temperature),
		ThreadID:         Util.GetGoID(),
	}
	fmt.Println(resp.Choices[0].Message.Content)
	return
}

//func (c *Client) CreateChatCompletion(
//	ctx context.Context,
//	request ChatCompletionRequest,
//) (response ChatCompletionResponse, err error) {
//	model := request.Model
//	if model != GPT3Dot5Turbo0301 && model != GPT3Dot5Turbo {
//		err = ErrChatCompletionInvalidModel
//		return
//	}
//
//	var reqBytes []byte
//	reqBytes, err = json.Marshal(request)
//	if err != nil {
//		return
//	}
//
//	urlSuffix := "/chat/completions"
//	req, err := http.NewRequestWithContext(ctx, http.MethodPost, c.fullURL(urlSuffix), bytes.NewBuffer(reqBytes))
//	if err != nil {
//		return
//	}
//
//	err = c.sendRequest(req, &response)
//	return
//}
