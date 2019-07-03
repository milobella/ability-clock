# clock
Ability to get the time whenever you are !

## Features
- [x] Get the time according to the device's timezone or default.
- [ ] Get the time in a city.
- [ ] Get the time in a country.

## Prerequisites

- Having access to [gitlab.milobella.com](https://gitlab.milobella.com/milobella)
- Having ``golang`` installed [instructions](https://golang.org/doc/install)
- Having ``go dep`` installed [instructions](https://golang.github.io/dep/docs/installation.html)

## Build

```bash
$ dep ensure
$ go build -o bin/ability cmd/ability/main.go
```

## Run

```bash
$ bin/ability
```

## Requests example

#### Get the time in default timezone
```bash
$ curl -i -X POST http://localhost:4444/resolve -d '{"nlu":{"BestIntent": "GET_TIME"}}'             130 ↵
HTTP/1.1 200 OK
Date: Wed, 03 Jul 2019 06:33:51 GMT
Content-Length: 126
Content-Type: text/plain; charset=utf-8

{"nlg":{"sentence":"It is {{time}}","params":[{"name":"time","value":"8 h 33","type":"time"}]},"context":{"slot_filling":{}}}
```

#### Get the time in shanghai
```bash
$ curl -i -X POST http://localhost:4444/resolve -d '{"nlu":{"BestIntent": "GET_TIME"}, "device": {"state": {"timezone": "Asia/Shanghai"}}}'
HTTP/1.1 200 OK
Date: Wed, 03 Jul 2019 06:34:38 GMT
Content-Length: 127
Content-Type: text/plain; charset=utf-8

{"nlg":{"sentence":"It is {{time}}","params":[{"name":"time","value":"14 h 34","type":"time"}]},"context":{"slot_filling":{}}}
```

## CHANGELOGS
- [Application changelog](CHANGELOG.md)
- [Helm chart changelog](clock-ability/helm/ability-cinema/CHANGELOG.md)