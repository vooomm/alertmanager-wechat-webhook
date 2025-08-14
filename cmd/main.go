package main

import (
	"alertmanager-wechat-webhook/api"
	"alertmanager-wechat-webhook/model"
	"alertmanager-wechat-webhook/pkg"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
)

var (
	H        bool
	RobotKey string
	Addr     string
)

// 构建命令行参数
func init() {
	flag.BoolVar(&H, "h", false, "help")
	flag.StringVar(&RobotKey, "robotkey", "", "wechatrobot token")
	flag.StringVar(&Addr, "addr", ":5001", "listen addr")
}

// 接收 Alertmanager Webhook
func alertHandler(w http.ResponseWriter, r *http.Request) {
	// 获取url中的key
	key := r.URL.Query().Get("key")

	// 判断token的来源
	if key == "" && RobotKey == "" {
		http.Error(w, "请输入key", http.StatusInternalServerError)
	} else if key != "" {
		RobotKey = key
	} else if RobotKey != "" {

	}

	var payload model.AlertWrapper
	// r.Body是一个流，使用Docode存到结构体中
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		log.Println("JSON解析失败:", err)
		http.Error(w, "Please use the POST request invalid json", http.StatusBadRequest)
		return
	}

	if len(payload.Alerts) == 0 {
		log.Println("收到空告警")
		return
	}

	msg := pkg.ConvertAlerts(payload.Alerts)
	err := api.SendToWechat(msg, RobotKey) // 你的key
	if err != nil {
		log.Println("发送到企业微信失败:", err)
		http.Error(w, "send failed", http.StatusInternalServerError)
		return
	}

	log.Println(r.RemoteAddr, r.Method, r.URL)
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Message sent successfully")
}

func main() {
	flag.Parse()

	if H {
		flag.Usage()
		return
	}
	http.HandleFunc("/webhook", alertHandler)
	log.Printf("启动服务，监听端口:%s/webhook", Addr)
	log.Fatal(http.ListenAndServe(":"+Addr, nil))
}
