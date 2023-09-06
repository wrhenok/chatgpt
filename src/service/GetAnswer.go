package service

import (
	"chatgpt3/src/Util"
	"chatgpt3/src/config"
	"chatgpt3/src/domain"
	"context"
	"encoding/json"
	"fmt"
	gpt "github.com/sashabaranov/go-openai"
	"strings"
)

type Client struct {
	g *gpt.Client
}

func NewClient(key string) *Client {
	return &Client{gpt.NewClient(key)}
}

func (c *Client) fullURL(suffix string) string {
	return fmt.Sprintf("%s%s", config.Url.BaseUrl, suffix)
}

func (c *Client) GetAnswer(ctx context.Context, prompt string) (chat domain.Chatgpt3, err error) {
	req := gpt.CompletionRequest{
		User:        "admin",
		Model:       gpt.GPT3TextDavinci003,
		MaxTokens:   3000, // usually works maxTokens - len(prompt)/3
		Prompt:      prompt,
		Temperature: 0.9,
		Suffix:      c.fullURL(config.Url.UrlSuffix),
	}
	resp, err := c.g.CreateCompletion(ctx, req)
	if err != nil {
		fmt.Println("createCompletion failed,err: ", err)
		return
	} else {
		answer := strings.TrimSpace(resp.Choices[0].Text)
		marshal, _ := json.Marshal(resp)
		key := config.ApiKey
		ip, _ := Util.GetIp()
		println()

		chat = domain.Chatgpt3{
			Username:         "admin",
			Userkey:          key,
			Model:            gpt.GPT3TextDavinci003,
			Prompt:           prompt,
			Answer:           answer,
			Isend:            resp.Choices[0].Index,
			RequestIpaddress: ip,
			ResponseJson:     string(marshal),
			PromptId:         resp.ID,
			FinishReason:     resp.Choices[0].FinishReason,
			RequestUrl:       c.fullURL(config.Url.UrlSuffix),
			MaxTokens:        req.MaxTokens,
			Temperature:      float64(req.Temperature),
			ThreadID:         Util.GetGoID(),
		}
		return
	}
}
