package request

type Content struct {
	ContentId int64  `json:"contentId"`
	Title     string `json:"title"`
	Story     string `json:"story"`
	UserId    int64  `json:"userId"`
}
