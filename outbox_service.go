package kavenegar

import (
	"context"
	"net/http"
)

type SelectOutboxService struct {
	c         *Client
	startDate int64
	endDate   *int64
	sender    *string
}

func (s *SelectOutboxService) StartDate(startDate int64) *SelectOutboxService {
	s.startDate = startDate
	return s
}

func (s *SelectOutboxService) EndDate(endDate int64) *SelectOutboxService {
	s.endDate = &endDate
	return s
}

func (s *SelectOutboxService) Sender(sender string) *SelectOutboxService {
	s.sender = &sender
	return s
}

func (s *SelectOutboxService) Do(ctx context.Context, opts ...RequestOption) (res *SelectOutbox, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/v1/%s/sms/selectoutbox.json",
	}
	r.setParam("messageid", s.startDate)
	if s.endDate != nil {
		r.setParam("messageid", s.endDate)
	}
	if s.sender != nil {
		r.setParam("sender", s.sender)
	}

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return &SelectOutbox{}, err
	}
	res = new(SelectOutbox)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return &SelectOutbox{}, err
	}
	return res, nil
}

type SelectOutbox struct {
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

type LatestOutBoxService struct {
	c        *Client
	pageSize *int64
	sender   *string
}

func (s *LatestOutBoxService) PageSize(pageSize int64) *LatestOutBoxService {
	s.pageSize = &pageSize
	return s
}

func (s *LatestOutBoxService) Sender(sender string) *LatestOutBoxService {
	s.sender = &sender
	return s
}

func (s *LatestOutBoxService) Do(ctx context.Context, opts ...RequestOption) (res *LatestOutBox, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/v1/%s/sms/latestoutbox.json",
	}
	if s.pageSize != nil {
		r.setParam("pageSize", *s.pageSize)
	}
	if s.sender != nil {
		r.setParam("sender", *s.sender)
	}

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return &LatestOutBox{}, err
	}
	res = new(LatestOutBox)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return &LatestOutBox{}, err
	}
	return res, nil
}

type LatestOutBox struct {
	Return struct {
		Status  int    `json:"status"`
		Message string `json:"message"`
	} `json:"return"`
	Entries []struct {
		Messageid  int    `json:"messageid"`
		Message    string `json:"message"`
		Status     int    `json:"status"`
		Statustext string `json:"statustext"`
		Sender     string `json:"sender"`
		Receptor   string `json:"receptor"`
		Date       int    `json:"date"`
		Cost       int    `json:"cost"`
	} `json:"entries"`
}

type GetCountOutboxService struct {
	c         *Client
	startDate int64
	endDate   *int64
	status    *string
}

func (s *GetCountOutboxService) StartDate(startDate int64) *GetCountOutboxService {
	s.startDate = startDate
	return s
}

func (s *GetCountOutboxService) EndDate(endDate int64) *GetCountOutboxService {
	s.endDate = &endDate
	return s
}

func (s *GetCountOutboxService) Status(status string) *GetCountOutboxService {
	s.status = &status
	return s
}

func (s *GetCountOutboxService) Do(ctx context.Context, opts ...RequestOption) (res *CountOutbox, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/v1/%s/sms/countoutbox.json",
	}
	r.setParam("messageid", s.startDate)
	if s.endDate != nil {
		r.setParam("messageid", s.endDate)
	}
	if s.status != nil {
		r.setParam("status", s.status)
	}

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return &CountOutbox{}, err
	}
	res = new(CountOutbox)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return &CountOutbox{}, err
	}
	return res, nil
}

type CountOutbox struct {
	Return struct {
		Status  int    `json:"status"`
		Message string `json:"message"`
	} `json:"return"`
	Entries []struct {
		Startdate int64 `json:"startdate"`
		Enddate   int64 `json:"enddate"`
		SumpPart  int   `json:"sumpart"`
		SumCount  int   `json:"sumcount"`
		Cost      int   `json:"cost"`
	} `json:"entries"`
}
