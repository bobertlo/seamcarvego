# seamcarvego

[![Build
Status](https://travis-ci.org/bobertlo/seamcarvego.svg?branch=master)](https://travis-ci.org/bobertlo/seamcarvego)
[![Go Report
Card](https://goreportcard.com/badge/github.com/bobertlo/seamcarvego)](https://goreportcard.com/report/github.com/bobertlo/seamcarvego)

*seamcarvego* is an implementation of a [seam
carving](https://en.m.wikipedia.org/wiki/Seam_carving) algorithm. For now, I
have simply ported an assignment from an algorithms class to go, in a test
driven manner. I plan to attempt optimizations in the future.

# Frontend

There is included a front-end with the following syntax:

```
seamcarvego [input.png] [output.png] [cols] [rows]
```

The frontend may be installed with `go get` using the following command:

```
go get github.com/bobertlo/seamcarvego/cmd/seamcarvego
```
