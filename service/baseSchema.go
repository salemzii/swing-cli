package service

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/TwiN/go-color"
)

var Tojson bool
var Tail bool

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
		time.Sleep(30 * time.Millisecond)

		switch strings.ToLower(v.Level) {
		case "debug":
			println(color.Colorize(color.Cyan, v.Created.String()), color.InGray(v.Function), color.Colorize(color.Blue, v.Message),
				color.Colorize(color.Cyan, v.Level), color.Colorize(color.Purple, v.StackTrace), color.Colorize(color.Black, v.Logger))
		case "info":
			println(color.Colorize(color.Cyan, v.Created.String()), color.InGray(v.Function), color.Colorize(color.Blue, v.Message),
				color.Colorize(color.Green, v.Level), color.Colorize(color.Purple, v.StackTrace), color.Colorize(color.Black, v.Logger))
		case "critical":
			println(color.Colorize(color.Cyan, v.Created.String()), color.InGray(v.Function), color.Colorize(color.Blue, v.Message),
				color.Colorize(color.Yellow, v.Level), color.Colorize(color.Purple, v.StackTrace), color.Colorize(color.Black, v.Logger))
		case "warning":
			println(color.Colorize(color.Cyan, v.Created.String()), color.InGray(v.Function), color.Colorize(color.Blue, v.Message),
				fmt.Sprintf("\033[38;5;214m %s", v.Level), color.Colorize(color.Purple, v.StackTrace), color.Colorize(color.Black, v.Logger))
		case "error":
			println(color.Colorize(color.Cyan, v.Created.String()), color.InGray(v.Function), color.Colorize(color.Blue, v.Message),
				color.Colorize(color.Red, v.Level), color.Colorize(color.Purple, v.StackTrace), color.Colorize(color.Black, v.Logger))
		}
	}
	return
}

func consoleLogsJson(records []LogRecord) {

	for _, v := range records {
		time.Sleep(30 * time.Millisecond)

		switch strings.ToLower(v.Level) {
		case "debug":
			b, err := json.MarshalIndent(v, "", "  ")
			if err != nil {
				log.Fatalf("Error Pretty Printing %v", err)
			}
			println(color.Colorize(color.Cyan, string(b)))

		case "info":
			b, err := json.MarshalIndent(v, "", "  ")
			if err != nil {
				log.Fatalf("Error Pretty Printing %v", err)
			}
			println(color.Colorize(color.Green, string(b)))

		case "critical":
			b, err := json.MarshalIndent(v, "", "  ")
			if err != nil {
				log.Fatalf("Error Pretty Printing %v", err)
			}
			println(color.Colorize(color.Yellow, string(b)))

		case "warning":
			b, err := json.MarshalIndent(v, "", "  ")
			if err != nil {
				log.Fatalf("Error Pretty Printing %v", err)
			}
			fmt.Sprintf("\033[38;5;214m %s", string(b))

		case "error":
			b, err := json.MarshalIndent(v, "", "  ")
			if err != nil {
				log.Fatalf("Error Pretty Printing %v", err)
			}
			println(color.Colorize(color.Red, string(b)))

		}
	}
	return
}

func consoleAccounts(resp LoginResponse) {
	println(color.Colorize(color.Cyan, " Username: "+resp.User.Username+"\n"),
		color.Colorize(color.Blue, "Email: "+resp.User.Email+"\n"),
		color.InPurple("Token: "+resp.Token))
}

func consoleInvalidMailError(e error) {
	println(color.Red, "encountered error validating your mail: ", e.Error())
}
