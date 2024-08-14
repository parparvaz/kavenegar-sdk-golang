package kavenegar

import (
	"context"
	"net/http"
)

type GetReceiveService struct {
	c          *Client
	lineNumber string
	isRead     int64
}

func (s *GetReceiveService) LineNumber(lineNumber string) *GetReceiveService {
	s.lineNumber = lineNumber
	return s
}

func (s *GetReceiveService) IsRead(isRead int64) *GetReceiveService {
	s.isRead = isRead
	return s
}

func (s *GetReceiveService) Do(ctx context.Context, opts ...RequestOption) (res *Receive, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/v1/%s/sms/receive.json",
	}
	r.setParam("linenumber", s.lineNumber)
	r.setParam("isread", s.isRead)

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return &Receive{}, err
	}
	res = new(Receive)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return &Receive{}, err
	}
	return res, nil
}

type Receive struct {
	Return struct {
		Status  int    `json:"status"`
		Message string `json:"message"`
	} `json:"return"`
	Entries []struct {
		Messageid int    `json:"messageid"`
		Message   string `json:"message"`
		Sender    string `json:"sender"`
		Receptor  string `json:"receptor"`
		Date      int    `json:"date"`
	} `json:"entries"`
}
