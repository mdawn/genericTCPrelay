# Generic TCP 

## Description

A system with a single relay server which supports multiple servers and, in turn, multiple clients for each server. 

## What We Need

We're on a Mac, right? Of course we are. 

- The latest version of [Go](https://golang.org/doc/install#install)
- This repo

## What it Does

With this TCP we can:

1. Run our relay with a selected listen port
2. Run multiple servers concurrently (an echo & a ping)
3. Run multiple clients concurrently per server (we use telnet)

## Getting Operational

**STEP 1**: We open a terminal and  
- Run `go run relay.go`
- The default port is `8080` if none is specified. However, we're free to use flag `-p` to designate a port for compatibility with concurrent relays

**STEP 2**: We open another terminal and 
- Run `go run echo.go`
- Running this returns the `established relay address`

**STEP 3**: We open our last terminal and
- Run `telnet localhost 8081`

Then we can type our little hearts out:
- ex. `Hello, world` echoes back `Hello, world`

## Example Scenario

You wrote a program that, tragically, sits behind a firewall. You desperately want to expose your server! Fortunately we can use a TCP connection to bypass the firewall. 

**First:** Run the relay server and choose a host port for it, such as `8080`. Initiate a connection to that host port.

**Second:** Once connected, the relay server graciously allows you to read and use a different port for your clients to use, such as `8081`. Tell your clients to connect to that port so the relay can host them as well. 

**Second:** Now you and your client are both being hosted by the relay via TCP channel. It's party time. 


## What would make it better

- Stress testing, security audit, graceful shutdown, scripts, etc.