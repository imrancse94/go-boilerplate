package logger

import (
	"encoding/json"
	"go-boilerplate/Helper"
	"go-boilerplate/constant"
	"log"
	"os"
	"path/filepath"
)

type LogData struct {
	Action string      `json:"action"`
	Data   interface{} `json:"data"`
}

func openLogFile(folder string, filename string) (*os.File, error) {
	folder = filepath.Join(".", folder)
	err := os.MkdirAll(folder, os.ModePerm)
	if err != nil {
		panic(err)
	}
	logFile, err := os.OpenFile(folder+"/"+filename, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0777)
	if err != nil {
		return nil, err
	}
	return logFile, nil
}

func CreateLog(data LogData) {

	file, err := openLogFile("logs", "log-"+Helper.Date_Format("now", constant.Format("Y-m-d"))+".log")
	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(file)
	//log.SetFlags(log.LstdFlags | log.Lshortfile | log.Lmicroseconds)
	b, err := json.Marshal(&data)
	log.Println(string(b))
}
