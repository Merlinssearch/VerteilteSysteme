// SPDX-FileCopyrightText: 2025 Maxim Ott <maxim.ott@informatik.hs-fulda.de>
// SPDX-FileCopyrightText: 2025 Maximilian Krönung <maximilian.kroenung@informatik.hs-fulda.de>
//
// SPDX-License-Identifier: GPL-3.0-or-later

package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"reflect"
	"strconv"
	"strings"
	"time"

	"group11/cmd/controller/http"
)

// bufio.NewScanner(),Erstellt einen Scanner aus einem io.Reader (z. B. net.Conn),Einmal pro Verbindung
// bufio:Reader() ist besser als scanner weil er Scanner liest zeilen weise und reader nicht
// scanner.Scan(),"Liest die nächste Zeile bis \n → gibt true wenn erfolgreich, false wenn Ende/Fehler",In deiner for-Schleife
// scanner.Text(),Gibt die zuletzt gelesene Zeile als string zurück (ohne \n),Direkt nach Scan()
// scanner.Err(),"Gibt den Fehler zurück, falls Scan() false wurde (z. B. EOF oder Netzwerkfehler)","Nach der Schleife, um zu wissen warum es aufgehört hat"

///////////////////////////////////////////////////////////
//                  Logs
///////////////////////////////////////////////////////////

func printHTTPRequest(req *http.Request) {
	if req == nil {
		fmt.Printf("no Valid http Object")
		return
	}
	v := reflect.ValueOf(*req) // Wert des Structs
	t := reflect.TypeOf(*req)  // Typinformationen

	for i := 0; i < v.NumField(); i++ {
		fieldName := t.Field(i).Name
		fieldValue := v.Field(i).Interface()

		switch val := fieldValue.(type) {
		case map[string][]string:
			fmt.Printf("%s:\n", fieldName)
			for k, values := range val {
				fmt.Printf("  %s:\n", k)
				for _, v := range values {
					fmt.Printf("    %s\n", v)
				}
			}
		default:
			fmt.Printf("%s: %v\n", fieldName, val)
		}
	}
}

// /////////////////////////////////////////////////////////
//              Connection Stuff
// /////////////////////////////////////////////////////////

func createTCPSocket(port int) (net.Listener, error) {
	addr := fmt.Sprintf(":%d", port)
	ln, err := net.Listen("tcp", addr)
	if err != nil {
		return nil, err
	}
	log.Printf("Listening on port %d", port)
	return ln, err
}

func handleConnection(conn net.Conn) {
	defer func() {
		fmt.Println("Client getrennt:", conn.RemoteAddr())
		conn.Close()
	}()
	fmt.Println("Client verbunden:", conn.RemoteAddr())
	reader := bufio.NewReader(conn)
	for {
		conn.SetDeadline(time.Now().Add(120 * time.Second))
		req := http.ParseHTTPRequestFromReader(reader)
		if req == nil {
			resp := &http.Response{
				StatusCode:      408,
				StatusText:      http.StatusTextMap[408],
				Proto:           "HTTP/1.1",
				CloseConnection: false,
				Body:            []byte("Request Timeout"),
				Headers: map[string][]string{
					"Content-Length": {strconv.Itoa(len("Request Timeout"))},
					"Content-Type":   {"text/plain"},
				},
			}
			conn.Write(http.SerializeHTTPResponse(resp))
			return
		}
		printHTTPRequest(req)
		resp := http.RouteHandler(req)
		if _, err := conn.Write(http.SerializeHTTPResponse(resp)); err != nil {
			log.Printf("Error writing response: %v", err)
			return // if error : connection closed
		}
		if values, ok := req.Header["Connection"]; ok {
			for _, v := range values {
				if strings.ToLower(v) == "close" {
					return
				}
			}
		}
	}
}

// /////////////////////////////////////////////////////////
//	             Die gute alte Main
// /////////////////////////////////////////////////////////

func main() {
	ln, err := createTCPSocket(8080)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Listening on :8080")
	// forloop für mehrer Verbindung
	for {
		conn, err := ln.Accept() // wenn ein Stream kommt wird accept
		if err != nil {
			log.Println("Error while Accept Connection ", err)
		}
		go handleConnection(conn)
		// ist wie fork() und macht schedueling für mehrer verbindung
		// conn.Write([]byte(newmessage + "\n"))
	}
}
