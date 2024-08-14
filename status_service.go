package kavenegar

import (
	"context"
	"net/http"
)

type StatusService struct {
	c         *Client
	messageID int64
}

func (s *StatusService) MessageID(messageID int64) *StatusService {
	s.messageID = messageID
	return s
}

func (s *StatusService) Do(ctx context.Context, opts ...RequestOption) (res *Status, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/v1/%s/sms/status.json",
	}
	r.setParam("messageid", s.messageID)

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return &Status{}, err
	}
	res = new(Status)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return &Status{}, err
	}
	return res, nil
}

type Status struct {
	Return struct {
		Status  int    `json:"status"`
		Message string `json:"message"`
	} `json:"return"`
	Entries []struct {
		MessageID  int    `json:"messageid"`
		Status     int    `json:"status"`
		StatusText string `json:"statustext"`
	} `json:"entries"`
}

type StatusByLocalIDService struct {
	c       *Client
	localID int64
}

func (s *StatusByLocalIDService) LocalID(localID int64) *StatusByLocalIDService {
	s.localID = localID
	return s
}

func (s *StatusByLocalIDService) Do(ctx context.Context, opts ...RequestOption) (res *StatusByLocalID, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/v1/%s/sms/status.json",
	}
	r.setParam("statuslocalmessageid", s.localID)

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return &StatusByLocalID{}, err
	}
	res = new(StatusByLocalID)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return &StatusByLocalID{}, err
	}
	return res, nil
}

type StatusByLocalID struct {
	Return struct {
		Status  int    `json:"status"`
		Message string `json:"message"`
	} `json:"return"`
	Entries []struct {
		Messageid  int    `json:"messageid"`
		Localid    string `json:"localid"`
		Status     int    `json:"status"`
		Statustext string `json:"statustext"`
	} `json:"entries"`
}
