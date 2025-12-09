# Controller TCP/HTTP Server

![Go](https://img.shields.io/badge/Language-Go-00ADD8) ![License](https://img.shields.io/badge/License-GPL--3.0-blue)

## Projektübersicht

Dieses Projekt implementiert einen **minimalen Controller-Server** über TCP Sockets, der HTTP-ähnliche Requests verarbeiten kann.
Ziel ist es, einen **eigenen HTTP-Parser, Response-Handler und Routing-Mechanismus** zu demonstrieren, der sowohl Dummy-Daten als auch Timeout-Handling unterstützt.
Das Projekt ist Teil meiner Software- und Systementwicklungs-Portfolioarbeiten.

---

## Features

* TCP Socket Server mit eigenem HTTP-Parser
* Unterstützung für GET- und POST-Requests
* JSON Dummy-Antworten für verschiedene Endpoints:

  * `/health` → Service-Status
  * `/status` → Units & Health
  * `/captain` → Controller-Status
  * `/elect` → Election-Start (POST)
* Automatisches Handling von Browser-Requests wie `/favicon.ico`
* Request Timeout Handling (HTTP 408)
* Flexible Header-Struktur (`map[string][]string`) für Mehrfach-Header
* Eigenes `SerializeHTTPResponse` für browserkompatible Antworten
* Keep-Alive und Connection-Handling

---

## Projektstruktur

```
├── build.sh
├── cmd
│   ├── controller
│   │   ├── http
│   │   │   ├── request.go      # HTTP Request Parser
│   │   │   ├── response.go     # HTTP Response + Status Map + Serializer
│   │   │   └── routes.go       # RouteHandler & Dummy Responses
│   │   ├── main.go             # TCP Server & Connection Handling
│   │   └── rpc
│   │       └── rpcHandler.go   # (Zukünftige RPC-Integration)
│   └── robot
│       └── robot.go            # (Separates Robot-Modul)
├── docker
│   ├── controller.Dockerfile
│   ├── docker-compose.yml
│   └── robot.Dockerfile
├── go.mod
├── LICENSES
│   ├── CC0-1.0.txt
│   └── GPL-3.0-or-later.txt
├── README.md
└── target
    ├── controller
    └── robot
```

---

## Installation & Build

1. Repository klonen:

```bash
git clone <repo-url>
cd <repo-folder>
```

2. Build-Skript ausführen:

```bash
./build.sh
```

3. Server starten:

```bash
go run cmd/controller/main.go
```

* Standardmäßig lauscht der Server auf TCP-Port (konfigurierbar in `main.go`).

---

## API Endpoints (Dummy-Daten)

| Methode | Pfad           | Beschreibung               |
| ------- | -------------- | -------------------------- |
| GET     | `/`            | Basis-Status (200 OK)      |
| GET     | `/health`      | Service Health JSON        |
| GET     | `/status`      | Units & Controller Health  |
| GET     | `/captain`     | Controller Alive & Uptime  |
| POST    | `/elect`       | Election gestartet (JSON)  |
| GET     | `/favicon.ico` | Automatisch 204 No Content |

* Alle Antworten sind **HTTP/1.1-konform** und enthalten passende `Content-Length` & `Content-Type` Header.

---

## Timeout Handling

* Für inaktive oder fehlerhafte Requests wird automatisch ein **HTTP 408 Request Timeout** gesendet.
* Konfigurierbare Deadline über `conn.SetDeadline`.

---

## Technologien

* **Programmiersprache:** Go
* **Kommunikation:** TCP Sockets (Custom HTTP Parser)
* **Datenformat:** JSON
* **Build/Deployment:** Docker & Bash Scripts

---

## Lizenz

Dieses Projekt ist unter der **GPL-3.0-or-later** lizenziert.
Siehe LICENSES/ für Details.

---

## Fazit

Dieses Projekt demonstriert die Fähigkeit, **einen eigenen, minimalen HTTP-Server auf TCP-Ebene** zu implementieren, inklusive:

* Routing
* JSON Responses
* Timeout Handling
* Browser-kompatible Headers

Ideal für ein Portfolio, um **Systemverständnis, Go-Kenntnisse und Netzwerkprogrammierung** zu zeigen.
