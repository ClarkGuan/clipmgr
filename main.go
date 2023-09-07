package main

import (
	"bytes"
	_ "embed"
	"log"
	"regexp"

	"github.com/getlantern/systray"
	"golang.design/x/clipboard"
)

//go:embed icon.png
var icon []byte
var spaces = regexp.MustCompile(`\s+`)
var oneSpace = []byte(" ")

func main() {
	if err := clipboard.Init(); err != nil {
		log.Fatalln(err)
	}
	systray.Run(onReady, onExit)
}

func onExit() {
}

func onReady() {
	systray.SetIcon(icon)

	modifyItem := systray.AddMenuItem("修改剪切板内容", "修改剪切板内容")
	go func() {
		for range modifyItem.ClickedCh {
			content := bytes.TrimSpace(clipboard.Read(clipboard.FmtText))
			if len(content) == 0 {
				return
			}
			clipboard.Write(clipboard.FmtText, spaces.ReplaceAll(content, oneSpace))
		}
	}()

	quitItem := systray.AddMenuItem("退出", "退出")
	go func() {
		<-quitItem.ClickedCh
		systray.Quit()
	}()
}
