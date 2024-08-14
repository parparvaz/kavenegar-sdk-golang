package kavenegar

import (
	"context"
	"net/http"
)

type CountInboxService struct {
	c          *Client
	startDate  int64
	lineNumber *string
	isRead     *int64
	endDate    *int64
}

func (s *CountInboxService) StartDate(startDate int64) *CountInboxService {
	s.startDate = startDate
	return s
}

func (s *CountInboxService) EndDate(endDate int64) *CountInboxService {
	s.endDate = &endDate
	return s
}

func (s *CountInboxService) LineNumber(lineNumber string) *CountInboxService {
	s.lineNumber = &lineNumber
	return s
}

func (s *CountInboxService) IsRead(isRead int64) *CountInboxService {
	s.isRead = &isRead
	return s
}

func (s *CountInboxService) Do(ctx context.Context, opts ...RequestOption) (res *CountInbox, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/v1/%s/sms/countinbox.json",
	}
	r.setParam("messageid", s.startDate)
	if s.endDate != nil {
		r.setParam("messageid", s.endDate)
	}
	if s.lineNumber != nil {
		r.setParam("linenumber", s.lineNumber)
	}
	if s.isRead != nil {
		r.setParam("isread", s.isRead)
	}

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return &CountInbox{}, err
	}
	res = new(CountInbox)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return &CountInbox{}, err
	}
	return res, nil
}

type CountInbox struct {
	Return struct {
		Status  int    `json:"status"`
		Message string `json:"message"`
	} `json:"return"`
	Entries []struct {
		StartDate int `json:"startdate"`
		EndDate   int `json:"enddate"`
		SumCount  int `json:"sumcount"`
	} `json:"entries"`
}

type InfoService struct {
	c *Client
}

func (s *InfoService) Do(ctx context.Context, opts ...RequestOption) (res *Info, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/v1/%s/sms/info.json",
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return &Info{}, err
	}
	res = new(Info)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return &Info{}, err
	}
	return res, nil
}

type Info struct {
	Return struct {
		Status  int    `json:"status"`
		Message string `json:"message"`
	} `json:"return"`
	Entries struct {
		Remaincredit int    `json:"remaincredit"`
		Expiredate   int    `json:"expiredate"`
		Type         string `json:"type"`
	} `json:"entries"`
}
