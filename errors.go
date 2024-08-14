package kavenegar

import (
	"fmt"
)

func (e APIError) Error() string {
	return fmt.Sprintf("<APIError> status=%d, msg=%s", e.Return.Status, e.Return.Message)
}

type APIError struct {
	Return struct {
		Status  int    `json:"status,omitempty"`
		Message string `json:"message,omitempty"`
	} `json:"return"`
	Entries struct {
		Datetime string `json:"datetime,omitempty"`
		Year     int    `json:"year,omitempty"`
		Month    int    `json:"month,omitempty"`
		Day      int    `json:"day,omitempty"`
		Hour     int    `json:"hour,omitempty"`
		Minute   int    `json:"minute,omitempty"`
		Second   int    `json:"second,omitempty"`
		Unixtime int    `json:"unixtime,omitempty"`
	} `json:"entries"`
}
