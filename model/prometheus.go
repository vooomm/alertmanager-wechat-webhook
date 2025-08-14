package model

import "time"

// Alertmanager 告警结构
type Alert struct {
	Status      string            `json:"status"`
	Labels      map[string]string `json:"labels"`
	Annotations map[string]string `json:"annotations"`
	StartsAt    time.Time         `json:"startsAt"`
	EndsAt      time.Time         `json:"endsAt"`
}

// Alertmanager Webhook 外层
type AlertWrapper struct {
	Alerts []Alert `json:"alerts"`
}

// 企业微信消息结构，下面两参数为@用户操作
type WechatText struct {
	Content         string   `json:"content"`
	MentionedList   []string `json:"mentioned_list,omitempty"`
	MentionedMobile []string `json:"mentioned_mobile_list,omitempty"`
}
type WechatMessage struct {
	MsgType  string     `json:"msgtype"`
	Markdown WechatText `json:"markdown"`
}
