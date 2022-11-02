package service

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/badoux/checkmail"
	"github.com/mitchellh/go-homedir"
	jrpc "github.com/ybbus/jsonrpc/v3"
)

var (
	serverEndpoint = "http://localhost:8081"
)

func allRecords(token string) {
	log.Println(token)
	client := jrpc.NewClient(serverEndpoint)
	var records []LogRecord
	resp, err := client.Call(context.Background(), "records.all", &AllRecordStruct{Tokenid: token, Limit: 500})
	if err != nil {
		log.Fatal(err)
	}

	err = resp.GetObject(&records) // expects a rpc-object result value like: {"id": 123, "name": "alex", "age": 33}
	if err != nil || len(records) < 1 {
		log.Fatal(err)
	}
	consoleLogs(records)
}

func GetRecordsWithLineNum(token string, line int) {
	client := jrpc.NewClient(serverEndpoint)
	var records []LogRecord
	resp, err := client.Call(context.Background(), "records.lineno", &RecordLineNum{Tokenid: token, Line: line, Limit: 500})

	if err != nil {
		log.Fatal(err)
	}

	err = resp.GetObject(&records) // expects a rpc-object result value like: {"id": 123, "name": "alex", "age": 33}
	if err != nil || len(records) < 1 {
		log.Fatal(err)
	}

	consoleLogs(records)
}

func GetRecordsWithFunction(token string, function string) {
	client := jrpc.NewClient(serverEndpoint)
	var records []LogRecord
	resp, err := client.Call(context.Background(), "records.function", &RecordFunction{Tokenid: token, Function: function})

	if err != nil {
		log.Fatal(err)
	}

	err = resp.GetObject(&records) // expects a rpc-object result value like: {"id": 123, "name": "alex", "age": 33}
	if err != nil || len(records) < 1 {
		log.Fatal(err)
	}
	consoleLogs(records)
}

func GetRecordsWithLogLevel(token string, level string) {
	client := jrpc.NewClient(serverEndpoint)
	var records []LogRecord
	resp, err := client.Call(context.Background(), "records.level", &RecordLevel{Tokenid: token, Level: level})

	if err != nil {
		log.Fatal(err)
	}

	err = resp.GetObject(&records) // expects a rpc-object result value like: {"id": 123, "name": "alex", "age": 33}
	if err != nil || len(records) < 1 {
		log.Fatal(err)
	}

	consoleLogs(records)
}

func GetRecordsLast15(token string) {
	client := jrpc.NewClient(serverEndpoint)
	var records []LogRecord
	resp, err := client.Call(context.Background(), "records.duration.15", &XRecords{Tokenid: token, Minutes: 15})

	if err != nil {
		log.Fatal(err)
	}

	err = resp.GetObject(&records) // expects a rpc-object result value like: {"id": 123, "name": "alex", "age": 33}
	if err != nil || len(records) < 1 {
		log.Fatal(err)
	}

	consoleLogs(records)
}

func GetRecordsLastX(token string, minutes int) {
	client := jrpc.NewClient(serverEndpoint)
	var records []LogRecord
	resp, err := client.Call(context.Background(), "records.duration.x", &XRecords{Tokenid: token, Minutes: minutes})

	if err != nil {
		log.Fatal(err)
	}

	err = resp.GetObject(&records) // expects a rpc-object result value like: {"id": 123, "name": "alex", "age": 33}
	if err != nil || len(records) < 1 {
		log.Fatal(err)
	}

	consoleLogs(records)
}

func ValidateMail(email string) error {
	err := checkmail.ValidateFormat(email)
	if err != nil {
		return err
	}

	return nil
}

func CreateAccount(username, email, password string) {

	if err := ValidateMail(email); err != nil {
		consoleInvalidMailError(err)
		return
	}
	client := jrpc.NewClient(serverEndpoint)
	var accountResp LoginResponse
	account := User{Username: username, Email: email, Password: password}
	resp, err := client.Call(context.Background(), "users.create", &account)
	if err != nil {
		log.Fatal(err)
	}

	err = resp.GetObject(&accountResp) // expects a rpc-object result value like: {"id": 123, "name": "alex", "age": 33}
	if err != nil {
		log.Fatal(err)
	}
	if !WriteToken(accountResp.Token) {
		log.Fatal("Token storage unsuccessful!")
	}
	consoleAccounts(accountResp)
}

func WriteToken(token string) bool {
	home, err := homedir.Dir()
	if err != nil {
		log.Fatal("Unable to find home directory")
		return false
	}
	appendToken := fmt.Sprintf("TOKEN=%s", token)
	swingEnvFile := home + "swing.env"
	file, err := os.OpenFile(swingEnvFile, os.O_APPEND|os.O_WRONLY|os.O_CREATE, os.ModeAppend)
	if err != nil {
		log.Fatal("Unable to open swing.env")
		return false
	}
	lineWritten, err := file.Write([]byte(appendToken))
	if err != nil {
		log.Fatal("Unable to write token to swing.env")
		return false
	}
	if lineWritten == 0 {
		log.Fatal("token write was unsuccesful")
		return false
	}
	return true
}

func Login(email, password string) {
	client := jrpc.NewClient(serverEndpoint)
	var loginResp LoginResponse
	resp, err := client.Call(context.Background(), "users.login", &LoginUser{Email: email, Password: password})
	if err != nil {
		log.Fatal(err)
	}

	err = resp.GetObject(&loginResp) // expects a rpc-object result value like: {"id": 123, "name": "alex", "age": 33}
	if err != nil {
		log.Fatal(err)
	}

	consoleAccounts(loginResp)
}
