package serializer



type ResponseFormatter struct {
	Data interface{} `json:"data"`
	Status int `json:"status"`
}
