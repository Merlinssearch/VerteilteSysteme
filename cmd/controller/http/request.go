// SPDX-FileCopyrightText: 2025 Maxim Ott <maxim.ott@informatik.hs-fulda.de>
// SPDX-FileCopyrightText: 2025 Maximilian Krönung <maximilian.kroenung@informatik.hs-fulda.de>
//
// SPDX-License-Identifier: GPL-3.0-or-later

package http

import (
	"bufio"
	"fmt"
	"log"
	"strconv"
	"strings"
)

type Request struct {
	Method  string
	Path    string
	Version string
	Header  map[string][]string
	Body    string
}

// conn is now in handleconnection because , if this function is done the connection is closed
// we just give the pointer from the reader

func ParseHTTPRequestFromReader(reader *bufio.Reader) *Request {
	// reader := bufio.NewReader(conn) reader now in handelconnection so we can handle multible http request
	parsedHTTP := &Request{
		Header: make(map[string][]string),
	}
	///////////////////////////////////////////////////////////
	//             Erste Line von HTTP
	///////////////////////////////////////////////////////////
	line, err := reader.ReadString('\n')
	if err != nil {
		log.Println("Failed to Readline from bufio reader ", err)
		return nil
	}
	line = strings.TrimSpace(line) // entfern äußere leerzeichen
	// parts := strings.Split(line, " ")  // keine gute idee , kann sein das zu viele leerzeichen da sind
	parts := strings.Fields(line)
	if len(parts) < 3 {
		log.Println("Invalid Http Request")
		return nil
	}
	parsedHTTP.Method = parts[0]
	parsedHTTP.Path = parts[1]
	parsedHTTP.Version = parts[2]
	///////////////////////////////////////////////////////////
	//             Header von HTTP
	///////////////////////////////////////////////////////////
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			log.Println("Failed to Readline from bufio reade ", err)
			return nil
		}
		line = strings.TrimSpace(line) // entfernt äußere leerzeichen
		if line == "" {
			break
		}
		slices := strings.SplitN(line, ":", 2)
		key := strings.TrimSpace(slices[0])
		value := strings.TrimSpace(slices[1])
		parsedHTTP.Header[key] = append(parsedHTTP.Header[key], value)
	}
	///////////////////////////////////////////////////////////
	//              Body von HTTP
	///////////////////////////////////////////////////////////
	if contentLengthStrs, ok := parsedHTTP.Header["Content-Length"]; ok &&
		len(contentLengthStrs) > 0 {
		intLength, err := strconv.Atoi(contentLengthStrs[0])
		if err != nil {
			fmt.Println("Couldn't convert ASCII to INT", err)
		}
		bufferBody := make([]byte, intLength)
		// Allocate buffer from length Bodylength in bytes
		total := 0
		for total < intLength {
			// read from stream and write in bufferBody , n is the return in Int (how many bytes were read)
			n, err := reader.Read(bufferBody[total:])
			if err != nil {
				fmt.Println("Writing in Buff or Reading from Stream :fail", err)
			}
			// add the read bytes so buffer doesn't overwrite stuff in bufferBody
			total += n
		}
		BodyContent := string(bufferBody)
		parsedHTTP.Body = BodyContent
	}
	return parsedHTTP
}
