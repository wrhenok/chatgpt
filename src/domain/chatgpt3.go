package domain

import "time"

type Chatgpt3 struct {
	ChatID           int64       `json:"chatID,omitempty" db:"chatID"` //id
	Username         string      `json:"username,omitempty"`           //用户名
	Userkey          string      `json:"userkey,omitempty"`
	Model            string      `json:"model,omitempty"`
	Prompt           string      `json:"prompt,omitempty"`           //问题
	Answer           string      `json:"answer,omitempty"`           //答案
	Isend            int         `json:"isend,omitempty"`            //是否结束
	RequestIpaddress string      `json:"requestIpAddress,omitempty"` //ip地址
	ResponseJson     interface{} `json:"responseJson,omitempty"`     //返回的response的json
	PromptId         string      `json:"promptId,omitempty"`
	FinishReason     string      `json:"finishReason,omitempty"`
	RequestUrl       string      `json:"requestUrl,omitempty"`
	MaxTokens        int         `json:"maxTokens,omitempty"`
	Temperature      float64     `json:"temperature,omitempty"`
	ThreadID         uint64      `json:"threadID,omitempty"`
	Run_time         string      `json:"run_Time,omitempty"`
	Createtime       string      `json:"createtime,omitempty"` //创建时间
}

var Start time.Time
