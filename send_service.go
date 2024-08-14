package kavenegar

import (
	"context"
	"net/http"
	"strings"
)

type GetSendService struct {
	c        *Client
	receptor string
	message  string
	sender   *string
	date     *int64
	sendType *string
	localID  *int64
	hide     *[]byte
}

func (s *GetSendService) Receptor(receptor []string) *GetSendService {
	s.receptor = strings.Join(receptor, ",")
	return s
}

func (s *GetSendService) Message(message string) *GetSendService {
	s.message = message
	return s
}

func (s *GetSendService) Sender(sender string) *GetSendService {
	s.sender = &sender
	return s
}

func (s *GetSendService) Date(date int64) *GetSendService {
	s.date = &date
	return s
}

func (s *GetSendService) SendType(sendType string) *GetSendService {
	s.sendType = &sendType
	return s
}

func (s *GetSendService) LocalID(localID int64) *GetSendService {
	s.localID = &localID
	return s
}

func (s *GetSendService) Hide(hide []byte) *GetSendService {
	s.hide = &hide
	return s
}

func (s *GetSendService) Do(ctx context.Context, opts ...RequestOption) (res *Send, err error) {
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

type PostSendArrayService struct {
	c               *Client
	receptor        []string
	message         []string
	sender          []string
	date            *int64
	sendType        *[]int8
	localMessageIDs *[]int64
	hide            *[]byte
}

func (s *PostSendArrayService) Receptor(receptor []string) *PostSendArrayService {
	s.receptor = receptor
	return s
}

func (s *PostSendArrayService) Message(message []string) *PostSendArrayService {
	s.message = message
	return s
}

func (s *PostSendArrayService) Sender(sender []string) *PostSendArrayService {
	s.sender = sender
	return s
}

func (s *PostSendArrayService) Date(date int64) *PostSendArrayService {
	s.date = &date
	return s
}

func (s *PostSendArrayService) SendType(sendTypes []sendType) *PostSendArrayService {
	var st []int8
	for _, sendType := range sendTypes {
		st = append(st, sendType.int8())
	}
	s.sendType = &st
	return s
}

func (s *PostSendArrayService) LocalID(localmessageids []int64) *PostSendArrayService {
	s.localMessageIDs = &localmessageids
	return s
}

func (s *PostSendArrayService) Hide(hide []byte) *PostSendArrayService {
	s.hide = &hide
	return s
}

func (s *PostSendArrayService) Do(ctx context.Context, opts ...RequestOption) (res *SendArray, err error) {
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
