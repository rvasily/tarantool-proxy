package main

import (
	"github.com/tarantool/go-tarantool"
)

func (self *ProxyConnection) executeRequestPing(requestType uint32, requestId uint32,
	reader IprotoReader) (flags uint32, response *tarantool.Response, err error) {
	//Ping body is empty, so body_length == 0 and there's no body
	//|--------------- header ----------------|
	// <request_type><body_length><request_id>
	flags = FlagPing

	tnt16 := self.getTnt16Master("ping")
	response, err = tnt16.Ping()
	return
}
