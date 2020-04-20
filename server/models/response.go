package models

// ResponsePattern return object  general response pattern
type ResponsePattern struct {
	Status        string              `json:"status"`
	Data          interface{}         `json:"data"`
	Message       string              `json:"message"`
	SystemMessage string              `json:"system_message"`
	Code          int                 `json:"code"`
	Meta          ResponsePatternMeta `json:"meta"`
}

// ResponsePatternMeta return object for  reposne meta
type ResponsePatternMeta struct {
	Limit *int `json:"limit"`
	Page  *int `json:"page"`
	Total *int `json:"total"`
	Count *int `json:"count"`
}
