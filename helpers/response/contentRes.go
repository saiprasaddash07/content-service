package response

type Content struct {
	ContentId int64  `json:"contentId,omitempty"`
	Title     string `json:"title,omitempty"`
	Story     string `json:"story,omitempty"`
	UserId    int64  `json:"userId,omitempty"`
}