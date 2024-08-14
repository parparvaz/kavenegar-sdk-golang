package kavenegar

import (
	"context"
	"net/http"
	"strings"
)

type GetCancelService struct {
	c         *Client
	messageID string
}

func (s *GetCancelService) MessageID(messageID []string) *GetCancelService {
	s.messageID = strings.Join(messageID, ",")
	return s
}

func (s *GetCancelService) Do(ctx context.Context, opts ...RequestOption) (res *Cancel, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/v1/%s/sms/cancel.json",
	}
	r.setParam("messageid", s.messageID)

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return &Cancel{}, err
	}
	res = new(Cancel)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return &Cancel{}, err
	}
	return res, nil
}

type Cancel struct {
	Return struct {
		Status  int    `json:"status"`
		Message string `json:"message"`
	} `json:"return"`
	Entries []struct {
		Messageid  int    `json:"messageid"`
		Status     int    `json:"status"`
		Statustext string `json:"statustext"`
	} `json:"entries"`
}
