package devp

import (
	"bytes"
	"encoding/json"
	"fmt"
	"runtime"
	"time"
)

func GetObjectJsonWithStack(txt string, obj interface{}) string {
	if txt != "" {
		return fmt.Sprintf("%s %s\n%s\n", txt, getObjectJson(obj), getStackInfo())
	} else {
		return fmt.Sprintf("%s\n%s\n", getObjectJson(obj), getStackInfo())
	}
}

func GetObjectJson(txt string, obj interface{}) string {
	if txt != "" {
		return txt + " " + getObjectJson(obj)
	} else {
		return getObjectJson(obj)
	}
}

func getObjectJson(obj interface{}) string {
	b, err := json.Marshal(obj)
	json := ""
	if err != nil {
		json = fmt.Sprintf("json Marshal error: %s", err)
	}
	json = string(b)
	return json
}

func PrintObjectJsonWithStack(txt string, obj interface{}) {
	fmt.Printf("[%s] %s\n%s\n%s\n", getDateTimeNow(), txt, getObjectJson(obj), getStackInfo())
}

func PrintObjectJson(txt string, obj interface{}) {
	fmt.Printf("[%s] %s\n%s\n", getDateTimeNow(), txt, getObjectJson(obj))
}

func PrintStack(txt string) {
	if txt != "" {
		fmt.Printf("[%s] %s\n%s\n", getDateTimeNow(), txt, getStackInfo())
	} else {
		fmt.Printf("[%s]\n%s\n", getDateTimeNow(), getStackInfo())
	}
}

func Stack() string {
	return getStackInfo()
}

func getDateTimeNow() string {
	return time.Now().Format("2006-01-02 15:04:05.99999")
}

func getStackInfo() string {
	buf := make([]byte, 1<<16)
	runtime.Stack(buf, false)

	buf = bytes.Trim(buf, "\x00")

	return string(buf)
}

