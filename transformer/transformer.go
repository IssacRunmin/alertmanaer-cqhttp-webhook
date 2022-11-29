package transformer

import (
	"bytes"
	"fmt"
	"strings"
	"time"


	"github.com/IssacRunmin/alertmanaer-cqhttp-webhook/model"
)

// TransformToMarkdown transform alertmanager notification to cqhttp message
func TransformToCQmessage(notification model.Notification) (message *model.CQMessage, robotURL string, err error) {

	// groupKey := notification.GroupKey
	status := notification.Status

	annotations := notification.CommonAnnotations
	robotURL = annotations["cqRobot"]

	var buffer bytes.Buffer
	var info string
  var time0 time.Time
	chinaLocal, _ := time.LoadLocation("PRC")
	for _, alert := range notification.Alerts {
		annotations := alert.Annotations
		summary_slices := strings.Split(annotations["summary"], " ")
		time0 = alert.StartsAt
		if summary_slices[1] == "down" {
			if status == "firing" {
				info = "DOWN"
			} else{
				info = "UP"
				time0 = alert.EndsAt
			}
		} else {
			info = annotations["summary"]
		}
		buffer.WriteString(fmt.Sprintf("【%s】 %s since %s\n", summary_slices[0],
    info, time0.In(chinaLocal).Format("2006-01-02 15:04:05")))
	}

	// buffer.WriteString(fmt.Sprintf("### 通知组%s(当前状态:%s) \n", groupKey, status))
	//
	// buffer.WriteString(fmt.Sprintf("#### 告警项:\n"))
	//
	// for _, alert := range notification.Alerts {
	// 	annotations := alert.Annotations
	// 	buffer.WriteString(fmt.Sprintf("##### %s\n > %s\n", annotations["summary"], annotations["description"]))
	// 	buffer.WriteString(fmt.Sprintf("\n> 开始时间：%s\n", alert.StartsAt.Format("15:04:05")))
	// }

	message = &model.CQMessage{
		MsgType: "cqhttp",
		Message: buffer.String(),
	}

	//markdown = &model.DingTalkMarkdown{
	//	MsgType: "markdown",
	//	Markdown: &model.Markdown{
	//		Title: fmt.Sprintf("通知组：%s(当前状态:%s)", groupKey, status),
	//		Text:  buffer.String(),
	//	},
	//	At: &model.At{
	//		IsAtAll: false,
	//	},
	//}

	return
}
