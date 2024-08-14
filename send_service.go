package kavenegar

import (
	"context"
	"net/http"
	"strings"
)

type SendService struct {
	c        *Client
	receptor string
	message  string
	sender   *string
	date     *int64
	sendType *string
	localID  *int64
	hide     *[]byte
}

func (s *SendService) Receptor(receptor []string) *SendService {
	s.receptor = strings.Join(receptor, ",")
	return s
}

func (s *SendService) Message(message string) *SendService {
	s.message = message
	return s
}

func (s *SendService) Sender(sender string) *SendService {
	s.sender = &sender
	return s
}

func (s *SendService) Date(date int64) *SendService {
	s.date = &date
	return s
}

func (s *SendService) SendType(sendType string) *SendService {
	s.sendType = &sendType
	return s
}

func (s *SendService) LocalID(localID int64) *SendService {
	s.localID = &localID
	return s
}

func (s *SendService) Hide(hide []byte) *SendService {
	s.hide = &hide
	return s
}

func (s *SendService) Do(ctx context.Context, opts ...RequestOption) (res *Send, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/v1/%s/sms/send.json",
	}
	r.setParam("receptor", s.receptor)
	r.setParam("message", s.message)
	if s.sender != nil {
		r.setParam("sender", s.sender)
	}
	if s.date != nil {
		r.setParam("date", s.date)
	}
	if s.sendType != nil {
		r.setParam("type", s.sendType)
	}
	if s.localID != nil {
		r.setParam("localid", s.localID)
	}
	if s.hide != nil {
		r.setParam("hide", s.hide)
	}

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return &Send{}, err
	}
	res = new(Send)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return &Send{}, err
	}
	return res, nil
}

type Send struct {
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

type SendArrayService struct {
	c               *Client
	receptor        []string
	message         []string
	sender          []string
	date            *int64
	sendType        *[]int8
	localMessageIDs *[]int64
	hide            *[]byte
}

func (s *SendArrayService) Receptor(receptor []string) *SendArrayService {
	s.receptor = receptor
	return s
}

func (s *SendArrayService) Message(message []string) *SendArrayService {
	s.message = message
	return s
}

func (s *SendArrayService) Sender(sender []string) *SendArrayService {
	s.sender = sender
	return s
}

func (s *SendArrayService) Date(date int64) *SendArrayService {
	s.date = &date
	return s
}

func (s *SendArrayService) SendType(sendTypes []sendType) *SendArrayService {
	var st []int8
	for _, sendType := range sendTypes {
		st = append(st, sendType.int8())
	}
	s.sendType = &st
	return s
}

func (s *SendArrayService) LocalID(localmessageids []int64) *SendArrayService {
	s.localMessageIDs = &localmessageids
	return s
}

func (s *SendArrayService) Hide(hide []byte) *SendArrayService {
	s.hide = &hide
	return s
}

func (s *SendArrayService) Do(ctx context.Context, opts ...RequestOption) (res *SendArray, err error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: "/v1/%s/sms/sendarray.json",
	}
	r.setParam("receptor", s.receptor)
	r.setParam("message", s.message)
	r.setParam("sender", s.sender)
	if s.date != nil {
		r.setParam("date", s.date)
	}
	if s.sendType != nil {
		r.setParam("type", s.sendType)
	}
	if s.localMessageIDs != nil {
		r.setParam("localmessageids", s.localMessageIDs)
	}
	if s.hide != nil {
		r.setParam("hide", s.hide)
	}

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return &SendArray{}, err
	}
	res = new(SendArray)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return &SendArray{}, err
	}
	return res, nil
}

type SendArray struct {
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
