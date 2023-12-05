package route

type BaseResponse struct {
	Success bool        `json:"success"`
	Msg     interface{} `json:"data"`
}

type ResponseOptions func(response BaseResponse)

func WithData(data interface{}) ResponseOptions {
	return func(response BaseResponse) {
		response.Msg = &data
	}
}

func WithOptions(options ...ResponseOptions) BaseResponse {
	var response BaseResponse
	response.Success = true
	for _, op := range options {
		op(response)
	}
	return response
}
