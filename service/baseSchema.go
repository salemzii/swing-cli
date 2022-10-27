package service

import "time"

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
