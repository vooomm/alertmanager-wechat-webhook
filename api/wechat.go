package api

import (
	"alertmanager-wechat-webhook/model"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// 发送到企业微信机器人
func SendToWechat(msg model.WechatMessage, webhookKey string) error {
	// 拼接url
	url := fmt.Sprintf("https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=%s", webhookKey)
	body, _ := json.Marshal(msg)

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("微信返回状态码: %d", resp.StatusCode)
	}
	return nil
}
