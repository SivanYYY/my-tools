package mhelp

import (
	"bytes"
	"fmt"
	"log"
	"runtime/debug"
	"time"
)

func CatchPanic() {
	if err := recover(); err != nil {
		errs := debug.Stack()
		log.Printf("错误：%v", err)
		log.Printf("追踪：%s", string(errs))
	}
}

func ShowVersion(text string) {
	text += " " + time.Now().Format(time.DateTime)
	alen := len([]uint8(text))
	line := bytes.Buffer{}
	for i := 0; i < alen; i++ {
		line.Write([]byte("="))
	}
	output := line.String() + "\n" + text + "\n" + line.String() + "\n"
	fmt.Println(output)
}
