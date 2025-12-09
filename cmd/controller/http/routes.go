// SPDX-FileCopyrightText: 2025 Maxim Ott <maxim.ott@informatik.hs-fulda.de>
// SPDX-FileCopyrightText: 2025 Maximilian Krönung <maximilian.kroenung@informatik.hs-fulda.de>
//
// SPDX-License-Identifier: GPL-3.0-or-later

package http

import (
	"fmt"
	"strings"
)

// Dummy answers for test

func HandelHealth(req *Request) *Response {
	jsonText := `{"service": "controller","status": "ok","health": 100, "uptime": 1234}`
	return &Response{
		StatusCode: 200,
		StatusText: StatusTextMap[200],
		Proto:      "HTTP/1.1",
		Headers: map[string][]string{
			"Content-Type":   {"application/json"},
			"Content-Length": {fmt.Sprintf("%d", len(jsonText))},
		},
		Body:            []byte(jsonText),
		CloseConnection: false,
	}
}

func HandleStatus(req *Request) *Response {
	jsonText := `{"Units Alive": 10, "Health": 100, "service": "controller"}`
	return &Response{
		StatusCode: 200,
		StatusText: StatusTextMap[200],
		Proto:      "HTTP/1.1",
		Headers: map[string][]string{
			"Content-Type":   {"application/json"},
			"Content-Length": {fmt.Sprintf("%d", len(jsonText))},
		},
		Body:            []byte(jsonText),
		CloseConnection: false,
	}
}

func HandelCaptain(req *Request) *Response {
	jsonText := `{"alive": true, "uptime": 1234, "service": "controller"}`
	return &Response{
		StatusCode: 200,
		StatusText: StatusTextMap[200],
		Proto:      "HTTP/1.1",
		Headers: map[string][]string{
			"Content-Type":   {"application/json"},
			"Content-Length": {fmt.Sprintf("%d", len(jsonText))},
		},
		Body:            []byte(jsonText),
		CloseConnection: false,
	}
}

func HandelElect(req *Request) *Response {
	jsonText := `{"election":"started","timestamp": 1733740000}`
	return &Response{
		StatusCode: 200,
		StatusText: StatusTextMap[200],
		Proto:      "HTTP/1.1",
		Headers: map[string][]string{
			"Content-Type":   {"application/json"},
			"Content-Length": {fmt.Sprintf("%d", len(jsonText))},
		},
		Body:            []byte(jsonText),
		CloseConnection: false,
	}
}

func RouteHandler(req *Request) *Response {
	path := strings.ToLower(req.Path)     // can be bad if we work with files
	method := strings.ToUpper(req.Method) // safer tolowerCase ?
	switch true {
	case path == "/":
		return &Response{
			StatusCode:      200,
			StatusText:      StatusTextMap[200],
			Proto:           "HTTP/1.1",
			CloseConnection: false,
		}

	case path == "/health" && method == "GET":
		return HandelHealth(req)
	case path == "/status" && method == "GET":
		return HandleStatus(req)
	case path == "/captain" && method == "GET":
		return HandelCaptain(req)
	case path == "/elect" && method == "POST":
		return HandelElect(req)
	case path == "/favicon.ico" && method == "GET":
		return &Response{
			StatusCode:      204,
			StatusText:      StatusTextMap[204],
			Proto:           "HTTP/1.1",
			CloseConnection: false,
		}
	default:
		fmt.Println("Pfad nicht gültig : ", path)
		return &Response{
			StatusCode:      400,
			StatusText:      StatusTextMap[400],
			Proto:           "HTTP/1.1",
			CloseConnection: false,
		}
	}
}
