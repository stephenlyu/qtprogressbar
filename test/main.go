package main

import (
	"github.com/therecipe/qt/widgets"
	"os"
	"github.com/z-ray/log"
	"runtime"
	"bytes"
	"github.com/stephenlyu/goqtreactor/reactor"
	"github.com/stephenlyu/qtwidgets"
)

const DATA_DIR = "data"

func PanicTrace(kb int) []byte {
	s := []byte("/src/runtime/panic.go")
	e := []byte("\ngoroutine ")
	line := []byte("\n")
	stack := make([]byte, kb<<10) //4KB
	length := runtime.Stack(stack, true)
	start := bytes.Index(stack, s)
	stack = stack[start:length]
	start = bytes.Index(stack, line) + 1
	stack = stack[start:]
	end := bytes.LastIndex(stack, line)
	if end != -1 {
		stack = stack[:end]
	}
	end = bytes.Index(stack, e)
	if end != -1 {
		stack = stack[:end]
	}
	stack = bytes.TrimRight(stack, "\n")
	return stack
}

func main() {
	writer, _ := os.Create("app.log")

	log.SetOutput(writer)
	defer func() {
		if err := recover(); err != nil {
			log.Println(string(PanicTrace(2)))
			log.Error(err)
			writer.Close()
			os.Exit(-2)
		}
	}()

	reactor.Initialize()

	app := widgets.NewQApplication(len(os.Args), os.Args)

	dialog := qtwidgets.CreateWaitingDialog(nil)
	dialog.Start()

	os.Exit(app.Exec())
	writer.Close()
}
