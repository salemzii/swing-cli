package service

import (
	"log"
	"time"

	"github.com/TwiN/go-color"
)

type LogRecord struct {
	Id         int       `json:"id"`
	Level      string    `json:"level"`
	Message    string    `json:"message"`
	StackTrace string    `json:"stacktrace"`
	Function   string    `json:"function"`
	LineNumber int       `json:"linenum"`
	Process    int       `json:"process"`
	TimeStamp  time.Time `json:"timestamp"`
	Created    time.Time `json:"created"`
	Logger     string    `json:"logger"`
	TokenId    string    `json:"token"`
	UserId     int       `json:"userid"`
}

type AllRecordStruct struct {
	Tokenid string `json:"token"`
	Limit   int    `json:"limit"`
}
type RecordLineNum struct {
	Tokenid string `json:"token"`
	Limit   int    `json:"limit"`
	Line    int    `json:"line"`
}
type RecordFunction struct {
	Tokenid  string `json:"token"`
	Function string `json:"function"`
}
type RecordLevel struct {
	Tokenid string `json:"token"`
	Level   string `json:"level"`
}
type XRecords struct {
	Tokenid string `json:"token"`
	Minutes int    `json:"minutes"`
}

type User struct {
	Id       int       `json:"id"`
	Username string    `json:"username"`
	Email    string    `json:"email"`
	Password string    `json:"password"`
	Created  time.Time `json:"created"`
}

type LoginUser struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
type LoginResponse struct {
	User  User
	Token string
}

func consoleLogs(records []LogRecord) {
	for _, v := range records {
		log.Println(color.Colorize(color.Cyan, v.Message))
	}
	return
}

func consoleAccounts(user User) {
	println(color.Purple, user.Id, user.Username, user.Email)
}

func consoleAccountsLogin(resp LoginResponse) {
	println(color.Purple, resp.User.Id, resp.User.Email, resp.User.Username, resp.Token)
}
