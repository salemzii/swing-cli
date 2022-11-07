package service

import (
	"context"
	"log"

	"github.com/badoux/checkmail"
	jrpc "github.com/ybbus/jsonrpc/v3"
)

var (
	serverEndpoint       = "http://localhost:8081"
	publicserverEndpoint = "https://swinglogs.herokuapp.com"
)

func allRecords(token string) {
	log.Println(token)
	client := jrpc.NewClient(publicserverEndpoint)
	var records []LogRecord
	resp, err := client.Call(context.Background(), "records.all", &AllRecordStruct{Tokenid: token, Limit: 500})
	if err != nil {
		handleFailedCalls(err)
	}

	err = resp.GetObject(&records)
	if err != nil || len(records) < 1 {
		log.Fatal(err)
	}

	if Tojson == true {
		consoleLogsJson(records)
		return
	}

	consoleLogs(records)
}

func GetRecordsWithLineNum(token string, line int) {
	client := jrpc.NewClient(publicserverEndpoint)
	var records []LogRecord
	resp, err := client.Call(context.Background(), "records.lineno", &RecordLineNum{Tokenid: token, Line: line, Limit: 500})

	if err != nil {
		handleFailedCalls(err)
	}

	err = resp.GetObject(&records) // expects a rpc-object result value like: {"id": 123, "name": "alex", "age": 33}
	if err != nil || len(records) < 1 {
		log.Fatal(err)
	}

	if Tojson == true {
		consoleLogsJson(records)
		return
	}

	consoleLogs(records)
}

func GetRecordsWithFunction(token string, function string) {
	client := jrpc.NewClient(publicserverEndpoint)
	var records []LogRecord
	resp, err := client.Call(context.Background(), "records.function", &RecordFunction{Tokenid: token, Function: function})

	if err != nil {
		handleFailedCalls(err)
	}

	err = resp.GetObject(&records)
	if err != nil || len(records) < 1 {
		log.Fatal(err)
	}

	if Tojson == true {
		consoleLogsJson(records)
		return
	}

	consoleLogs(records)
}

func GetRecordsWithLogLevel(token string, level string) {
	client := jrpc.NewClient(publicserverEndpoint)
	var records []LogRecord
	resp, err := client.Call(context.Background(), "records.level", &RecordLevel{Tokenid: token, Level: level})

	if err != nil {
		handleFailedCalls(err)
	}

	err = resp.GetObject(&records)
	if err != nil || len(records) < 1 {
		log.Fatal(err)
	}
	if Tojson == true {
		consoleLogsJson(records)
		return
	}
	consoleLogs(records)
}

func GetRecordsLast15(token string) {
	client := jrpc.NewClient(publicserverEndpoint)
	var records []LogRecord
	resp, err := client.Call(context.Background(), "records.duration.15", &XRecords{Tokenid: token, Minutes: 15})

	if err != nil {
		handleFailedCalls(err)
	}

	err = resp.GetObject(&records)
	if err != nil || len(records) < 1 {
		log.Fatal(err)
	}

	if Tojson == true {
		consoleLogsJson(records)
		return
	}
	consoleLogs(records)
}

func GetRecordsLastX(token string, minutes int) {
	client := jrpc.NewClient(publicserverEndpoint)
	var records []LogRecord
	resp, err := client.Call(context.Background(), "records.duration.x", &XRecords{Tokenid: token, Minutes: minutes})

	if err != nil {
		handleFailedCalls(err)
	}

	err = resp.GetObject(&records)
	if err != nil || len(records) < 1 {
		log.Fatal(err)
	}

	if Tojson == true {
		consoleLogsJson(records)
		return
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
	client := jrpc.NewClient(publicserverEndpoint)
	var accountResp LoginResponse
	account := User{Username: username, Email: email, Password: password}
	resp, err := client.Call(context.Background(), "users.create", &account)
	if err != nil {
		handleFailedCalls(err)
	}

	err = resp.GetObject(&accountResp)
	if err != nil {
		log.Fatal(err)
	}
	if !WriteToken(accountResp.Token) {
		log.Fatal("Token storage unsuccessful!")
	}

	consoleAccounts(accountResp)
}

func Login(email, password string) {
	client := jrpc.NewClient(publicserverEndpoint)
	var loginResp LoginResponse
	resp, err := client.Call(context.Background(), "users.login", &LoginUser{Email: email, Password: password})
	if err != nil {
		handleFailedCalls(err)
	}

	err = resp.GetObject(&loginResp)
	if err != nil {
		log.Fatal(err)
	}

	consoleAccounts(loginResp)
}
