package service

import (
	"context"
	"log"

	jrpc "github.com/ybbus/jsonrpc/v3"
)

var (
	serverEndpoint = "http://localhost:8081"
)

func allRecords(token string) {
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
	for k, v := range records {
		log.Printf("Record %d, \n Level: %s, \n FunctionName: %s, \n Stacktrace: %s", k+1, v.Level, v.Function, v.StackTrace)
	}
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

	for k, v := range records {
		log.Printf("Record %d, \n Level: %s, \n FunctionName: %s, \n Stacktrace: %s", k+1, v.Level, v.Function, v.StackTrace)
	}
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

	for k, v := range records {
		log.Printf("Record %d, \n Level: %s, \n FunctionName: %s, \n Stacktrace: %s", k+1, v.Level, v.Function, v.StackTrace)
	}
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

	for k, v := range records {
		log.Printf("Record %d, \n Level: %s, \n FunctionName: %s, \n Stacktrace: %s", k+1, v.Level, v.Function, v.StackTrace)
	}
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

	for k, v := range records {
		log.Printf("Record %d, \n Level: %s, \n FunctionName: %s, \n Stacktrace: %s", k+1, v.Level, v.Function, v.StackTrace)
	}
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

	for k, v := range records {
		log.Printf("Record %d, \n Level: %s, \n FunctionName: %s, \n Stacktrace: %s", k+1, v.Level, v.Function, v.StackTrace)
	}
}
