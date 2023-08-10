package model

type ResponseError struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Errors  interface{} `json:"errors"`
}

type ResponseValue struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Value   interface{} `json:"value"`
}
