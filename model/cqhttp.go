package model

type CQMessage struct {
	MsgType string `json:"msgType"`
	Message string `json:"message"`
}

//type At struct {
//	AtMobiles []string `json:"atMobiles"`
//	IsAtAll   bool     `json:"isAtAll"`
//}
//
//type DingTalkMarkdown struct {
//	MsgType  string    `json:"msgtype"`
//	At       *At       `json:at`
//	Markdown *Markdown `json:"markdown"`
//}
//
//type Markdown struct {
//	Title string `json:"title"`
//	Text  string `json:"text"`
//}
