package types

type JSONDataResponse struct {
	Success bool `json:"success"`
	Data    any  `json:"data"`
}

type JSONMessageResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}
