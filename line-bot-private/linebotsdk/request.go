package linebotsdk

type Request struct {
	Events []*RequestEvent
}

type RequestEvent struct {
	Message struct {
		Text string `json:"text"`
	} `json:"message"`
	Source struct {
		UserID string `json:"userId"`
	} `json:"source"`
	ReplyToken string `json:"replyToken"`
}
