package engine

import "bili/backend/fetcher"

type ParseFunc func(contents []byte, url, dir string) ParseResult

type Request struct {
	Url           string
	Dir           string
	ParseFunction ParseFunc
	FetchFun      fetcher.FetchFun
}

func NewRequest(url, dir string, parseFunction ParseFunc, fetchFun fetcher.FetchFun) *Request {
	return &Request{Url: url, Dir: dir, ParseFunction: parseFunction, FetchFun: fetchFun}
}

type ParseResult struct {
	Requests []*Request
	Items    []*Item
}

type Item struct {
	Payload interface{}
}

func NewItem(payload interface{}) *Item {
	return &Item{Payload: payload}
}
