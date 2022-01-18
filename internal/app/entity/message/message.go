package ChatMessage

type Message struct {
	ID         int64  `json:"id"`
	Chat       int64  `json:"chat"`
	Author     int64  `json:"author"`
	Text       string `json:"text"`
	Created_at string `json:"create_at"`
}
