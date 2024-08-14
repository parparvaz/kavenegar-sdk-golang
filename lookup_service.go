package kavenegar

import (
	"context"
	"net/http"
)

type GetLookupService struct {
	c          *Client
	receptor   string
	token      string
	template   string
	token2     *string
	token3     *string
	lookupType *string
}

func (s *GetLookupService) Receptor(receptor string) *GetLookupService {
	s.receptor = receptor
	return s
}

func (s *GetLookupService) Token(token string) *GetLookupService {
	s.token = token
	return s
}

func (s *GetLookupService) Template(template string) *GetLookupService {
	s.template = template
	return s
}

func (s *GetLookupService) Token2(token2 string) *GetLookupService {
	s.token2 = &token2
	return s
}

func (s *GetLookupService) Token3(token3 string) *GetLookupService {
	s.token3 = &token3
	return s
}

func (s *GetLookupService) Type(lookupType string) *GetLookupService {
	s.lookupType = &lookupType
	return s
}

func (s *GetLookupService) Do(ctx context.Context, opts ...RequestOption) (res *Lookup, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/v1/%s/verify/lookup.json",
	}
	r.setParam("receptor", s.receptor)
	r.setParam("token", s.token)
	r.setParam("template", s.template)
	if s.token2 != nil {
		r.setParam("token2", *s.token2)
	}
	if s.token3 != nil {
		r.setParam("token3", *s.token3)
	}
	if s.lookupType != nil {
		r.setParam("type", *s.lookupType)
	}

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return &Lookup{}, err
	}
	res = new(Lookup)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return &Lookup{}, err
	}
	return res, nil
}

type Lookup struct {
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
