package main

import (
    "github.com/del-xiong/miniblink"
    "log"
)

func main() {
    err := miniblink.InitBlink()
    if err != nil {
        log.Fatal(err)
    }
    view := miniblink.NewWebView(false, 1366, 920)
    view.LoadURL("C:/Users/Administrator/Desktop/goexe/main.html")
    // 设置窗体标题(会被web页面标题覆盖)
    view.SetWindowTitle("miniblink window")
    // 移动到屏幕中心位置
    view.MoveToCenter()
    // 显示窗口
    view.ShowWindow()
	
	value, err := view.Invoke("main")
	if err != nil {
		log.Println(err)
	} else {
		log.Println(value)
	}
	

	
    <- make(chan bool)
}
