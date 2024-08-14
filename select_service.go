package kavenegar

import (
	"context"
	"net/http"
)

type GetSelectService struct {
	c         *Client
	messageID int64
}

func (s *GetSelectService) MessageID(messageID int64) *GetSelectService {
	s.messageID = messageID
	return s
}

func (s *GetSelectService) Do(ctx context.Context, opts ...RequestOption) (res *Select, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/v1/%s/sms/select.json",
	}
	r.setParam("messageid", s.messageID)

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return &Select{}, err
	}
	res = new(Select)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return &Select{}, err
	}
	return res, nil
}

type Select struct {
	Return struct {
		Status  int    `json:"status"`
		Message string `json:"message"`
	} `json:"return"`
	Entries []struct {
		MessageID  int    `json:"messageid"`
		Message    string `json:"message"`
		Status     int    `json:"status"`
		StatusText string `json:"statustext"`
		Sender     string `json:"sender"`
		Receptor   string `json:"receptor"`
		Date       int    `json:"date"`
		Cost       int    `json:"cost"`
	} `json:"entries"`
}
