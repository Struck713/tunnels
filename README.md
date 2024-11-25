# tunnels

A Go application for self-hosting tunnels, similar to cloudflared.

### How does it work?

This program basically creates a man-in-the-middle which brokers requests between the client and the Internet.
We just forward all the traffic through. This allows the client to have an application facing the Internet without forwarding a port.

### Getting started

You can build the program with `go build` and then run it with `./tunnels client` or `./tunnels server` depending on what you're doing.

### Some love for Go

Writing this program in Go made it super simple to actually get it working. The channels really make the asyncronous programming super simple and
are a great abstraction on a normally complex topic.
