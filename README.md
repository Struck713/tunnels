# tunnels

A Go application for self-hosting tunnels, similar to cloudflared.

### How does it work?

This program basically creates a man-in-the-middle which brokers requests between the client and the Internet.
We just forward all the traffic through. This allows the client to have an application facing the Internet without forwarding a port.

## Getting started

You can install the program using `go install nstruck.dev/tunnels@latest`.

### Server

Now, you can host a server (somewhere open to the Internet) using `tunnels server`. Make sure there is `config.json` file in the folder the program is ran in.

There is an example server config in `example_server_config.json`. In this config,

- `address` is the address you want to bind the proxy to
- `subdomain` is the subdomain to issue certificates to
- `email` is the email used for Let's Encrypt.

```JSON
{
  "address": "0.0.0.0:9999",
  "subdomain": "proxy.example.org",
  "email": "mail@example.org"
}
```

The server will start and provide an authentication key which you can use to establish a connection with a client.

### Client

Once the server is hosted, you can connect to it and proxy a server using the client. Use `tunnels client <auth key>`. Make sure there is `config.json` file in the folder the program is ran in.

There is an example server config in `example_client_config.json`. In this config,

- `service` is the service you want to forward
- `server` is the address and port of the server you hosted.

```JSON
{
  "service": "http://localhost:8085",
  "server": "0.0.0.0:9999"
}
```
