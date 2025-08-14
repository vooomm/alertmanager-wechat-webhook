package pkg

import (
	"alertmanager-wechat-webhook/model"
	"fmt"
	"strings"
	"time"
)

// 转换函数：Alertmanager → 企业微信
func ConvertAlerts(alerts []model.Alert) model.WechatMessage {
	var sb strings.Builder

	// 转换成本地时间
	loc, _ := time.LoadLocation("Asia/Shanghai")

	// 把告警循环遍历出来，放到结构体一起发送
	for _, alert := range alerts {

		status := alert.Status
		severity := alert.Labels["severity"]
		alertname := alert.Labels["alertname"]
		instance := alert.Labels["instance"]
		summary := alert.Annotations["summary"]
		desc := alert.Annotations["description"]

		// 告警模板
		if alert.Status != "resolved" {
			sb.WriteString("<font color=\"#FFA500\">**@告警通知**</font>:\n")
			fmt.Fprintf(&sb, "**- 告警级别**: %s\n**- 告警类型**: %s\n**- 故障主机**: %s\n**- 告警主题**: %s\n**- 告警详情**: %s\n", severity, alertname, instance, summary, desc)
			fmt.Fprintf(&sb, "**- 告警状态**: %s\n**- 触发时间**: <font color=\"red\"> %s</font>\n", status, alert.StartsAt.In(loc).Format("2006-01-02 15:04:05"))
		} else {
			// 恢复模板
			sb.WriteString("<font color=\"green\">**@告警恢复**</font>:\n")
			fmt.Fprintf(&sb, "**- 告警级别**: %s\n**- 告警类型**: %s\n**- 故障主机**: %s\n**- 告警主题**: %s\n**- 告警详情**: %s\n", severity, alertname, instance, summary, desc)
			fmt.Fprintf(&sb, "**- 恢复状态**: %s\n**- 触发时间**: <font color=\"red\"> %s</font>\n**- 恢复时间**: <font color=\"red\"> %s</font>\n", status, alert.StartsAt.In(loc).Format("2006-01-02 15:04:05"), alert.EndsAt.In(loc).Format("2006-01-02 15:04:05"))
		}
	}

	return model.WechatMessage{
		MsgType: "markdown",
		// 这里没使用@用户
		Markdown: model.WechatText{
			Content: sb.String(),
		},
	}
}
