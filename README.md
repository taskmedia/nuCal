# nuCal

[![releases](https://img.shields.io/github/v/release/taskmedia/nuCal?style=flat-square)](https://github.com/taskmedia/nuCal/releases/latest)
[![docs](https://img.shields.io/badge/docs-pkg.go.dev-blue?style=flat-square)](https://pkg.go.dev/github.com/taskmedia/nuCal)
[![golang version](https://img.shields.io/github/go-mod/go-version/taskmedia/nuCal?style=flat-square)](https://golang.org/dl/#stable)
<br />
[![codecoverage](https://img.shields.io/codecov/c/github/taskmedia/nuCal?style=flat-square)](https://app.codecov.io/gh/taskmedia/nuCal)
![code size](https://img.shields.io/github/languages/code-size/taskmedia/nuCal?style=flat-square)
<br />
[![issues](https://img.shields.io/github/issues/taskmedia/nuCal?style=flat-square)](https://github.com/taskmedia/nuCal/issues)
[![pull requests](https://img.shields.io/github/issues-pr/taskmedia/nuCal?style=flat-square)](https://github.com/taskmedia/nuCal/pulls)
<br />
[![twitter](https://img.shields.io/twitter/follow/taskmediaDE?style=social)](https://twitter.com/taskmediaDE)
<br />
[![Open in Gitpod](https://gitpod.io/button/open-in-gitpod.svg)](https://gitpod.io/#https://github.com/taskmedia/nuCal)

This application will generate a REST endpoint to generate [ICS (iCalendar)](https://datatracker.ietf.org/doc/html/rfc5545) from [nuLiga](https://bhv-handball.liga.nu/).
The service will consume JSON content from [taskmedia/nuScrape](https://github.com/taskmedia/nuScrape).

## Start application

You will be able to start the application directly with golang:

```bash
go run cmd/nuCal/nuCal.go
```

Another option would be running the application in a [Docker container](https://hub.docker.com/r/taskmedia/nucal):

```bash
docker run \
  --name nucal \
  -p 8080:8080 \
  taskmedia/nucal:latest
```

