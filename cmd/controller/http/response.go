// SPDX-FileCopyrightText: 2025 Maxim Ott <maxim.ott@informatik.hs-fulda.de>
// SPDX-FileCopyrightText: 2025 Maximilian Krönung <maximilian.kroenung@informatik.hs-fulda.de>
//
// SPDX-License-Identifier: GPL-3.0-or-later

// Package http implements a simple HTTP request parser and serializer.
package http

import (
	"fmt"
	"strconv"
	"strings"
)

var StatusTextMap = map[int]string{
	200: "OK",
	201: "Created",
	204: "No Content",
	400: "Bad Request",
	401: "Unauthorized",
	403: "Forbidden",
	404: "Not Found",
	408: "Request Timeout",
	500: "Internal Server Error",
}

type Response struct {
	StatusCode      int
	StatusText      string
	Proto           string
	Headers         map[string][]string //
	Body            []byte
	CloseConnection bool
}

func SerializeHTTPResponse(resp *Response) []byte {
	proto := resp.Proto
	if proto == "" {
		proto = "HTTP/1.1"
	}

	statusText := resp.StatusText
	if statusText == "" {
		if txt, ok := StatusTextMap[resp.StatusCode]; ok {
			statusText = txt
		} else {
			statusText = "Unknown"
		}
	}

	var builder strings.Builder
	builder.WriteString(
		fmt.Sprintf("%s %d %s\r\n", proto, resp.StatusCode, statusText),
	)

	if resp.Headers == nil {
		resp.Headers = make(map[string][]string)
	}
	// Set Content-Length
	resp.Headers["Content-Length"] = []string{strconv.Itoa(len(resp.Body))}
	// Add Connection: close if needed
	if resp.CloseConnection {
		resp.Headers["Connection"] = []string{"close"}
	}
	// Add headers
	for key, values := range resp.Headers {
		for _, v := range values {
			builder.WriteString(fmt.Sprintf("%s: %s\r\n", key, v))
		}
	} // End of headers
	builder.WriteString("\r\n")

	// Add body
	builder.Write(resp.Body)

	return []byte(builder.String())
}
