package constant

import "time"

type ApiResponse struct {
	Txid      string      `json:"txid"`
	Code      string      `json:"code"`
	Message   string      `json:"message"`
	Data      interface{} `json:"data"`
	TimeStamp time.Time   `json:"timestamp"`
}

func NewApiResponse() *ApiResponse {
	return &ApiResponse{Txid: GenTxID(), TimeStamp: time.Now()}
}
