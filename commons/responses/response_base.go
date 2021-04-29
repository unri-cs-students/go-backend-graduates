package responses

type ResponseBase struct {
	Code    int    `json:"status"`
	Message string `json:"status_message"`
}

type ResponseData struct {
	ResponseBase
	Data interface{} `json:"data"`
}

