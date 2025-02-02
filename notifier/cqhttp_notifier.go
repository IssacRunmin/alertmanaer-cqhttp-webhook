package notifier

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/IssacRunmin/alertmanaer-cqhttp-webhook/model"
	"github.com/IssacRunmin/alertmanaer-cqhttp-webhook/transformer"
)

// Send send markdown message to cqhttp
func Send(notification model.Notification, defaultRobot string) (err error) {

	markdown, robotURL, err := transformer.TransformToCQmessage(notification)

	if err != nil {
		return
	}

	data, err := json.Marshal(markdown)
	if err != nil {
		return
	}

	var cqRobotURL string

	if robotURL != "" {
		cqRobotURL = robotURL
	} else {
		cqRobotURL = defaultRobot
	}

	if len(cqRobotURL) == 0 {
		return nil
	}

	req, err := http.NewRequest(
		"POST",
		cqRobotURL,
		bytes.NewBuffer(data))

	if err != nil {
		fmt.Println("cqhttp robot url not found ignore:")
		return
	}

	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return
	}

	defer resp.Body.Close()
	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)

	return
}
