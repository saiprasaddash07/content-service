package request

type Content struct {
	ContentId int64  `json:"contentId,omitempty"`
	Title     string `json:"title"`
	Story     string `json:"story"`
	UserId    int64  `json:"userId"`
}

type ContentFetchReq struct {
	Size      int `json:"size"`
	NextToken int `json:"nextToken"`
}
