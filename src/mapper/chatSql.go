package mapper

import (
	"chatgpt3/src/config"
	"chatgpt3/src/domain"
	_ "database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"log"
)

var DB *sqlx.DB

func init() {
	config.InitConf("../../application.yml")
	dsn := config.Database.UserName + ":" + config.Database.Password + "@tcp(" + config.Database.Host + ":" + config.Database.Port + ")/" + config.Database.DbName
	db, err := sqlx.Open(config.Database.Type, dsn)
	if err != nil {
		fmt.Println("connect mysql failed :", err)
	}
	//defer db.Close()
	DB = db
	log.Printf("database connection success: %s", dsn)
}

// 数据库插入操作
func Insert(c *domain.Chatgpt3, runtime string) (i int64, err error) {
	sql := "insert into chatgpt3(username,userkey,model,prompt,answer,isend,request_ipaddress,response_json,prompt_id,finish_reason,request_url,run_time,max_tokens,temperature,threadID) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)"
	exec, err := DB.Exec(sql, c.Username, c.Userkey, c.Model, c.Prompt, c.Answer, c.Isend, c.RequestIpaddress, c.ResponseJson, c.PromptId, c.FinishReason, c.RequestUrl, runtime, c.MaxTokens, c.Temperature, c.ThreadID)
	if err != nil {
		fmt.Println("insert data failed, ", err)
		return
	}
	i, err = exec.LastInsertId()
	if err != nil {
		fmt.Println("insert failed", err)
		return
	}
	return
}

// 数据库查询
func Select() (chat []domain.Chatgpt3) {
	sql := "select * from chatgpt3 order by chatID desc limit 10"
	err := DB.Select(&chat, sql)
	if err != nil {
		fmt.Println("select data failed, err :", err)
		return
	}

	fmt.Println("chat: ", chat)
	return
}
