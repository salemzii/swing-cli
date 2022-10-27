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
