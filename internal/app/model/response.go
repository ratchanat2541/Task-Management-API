package model

type Response struct {
	Code         int         `json:"code"`
	Message      string      `json:"message"`
	CountAll     *int64      `json:"count_all,omitempty"`
	Count        *int64      `json:"count,omitempty"`
	Page         *int64      `json:"page,omitempty"`
	TotalPage    *int64      `json:"total_page,omitempty"`
	NextPage     *int64      `json:"next_page,omitempty"`
	PreviousPage *int64      `json:"previous_page,omitempty"`
	Limit        *int64      `json:"limit,omitempty"`
	Data         interface{} `json:"data,omitempty"`
}

func (err *Response) Error() string {
	message := err.Message
	return message
}
