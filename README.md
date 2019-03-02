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
2. Run multiple servers concurrently (I've provided an echo & a ping for your enjoyment) 
3. Run multiple clients concurrently per server (we use telnet)

## Example Scenario

You wrote a program that, tragically, sits behind a firewall. You desperately want to expose your server! Fortunately we can bypass the firewall with a relay. 

**STEP 1**: 
We open a terminal and run the relay server, initiating a connection
- Run `go run relay.go`
- The default port is `7` if none is specified. However, we're free to use flag `-p` to designate a port for compatibility with concurrent relays

**STEP 2**: 
Now we connect our server.

We open another terminal and 
- Run `go run echo.go` (our server we want to connect, in this case)
- Running this returns the `established relay address` 
- Tell your clients to connect to this `established relay address`

**STEP 3**: We open our last terminal and
- Run `telnet localhost 8081` 
(where `8081` is the `established relay address`)

Now you and your client are both being hosted by the relay via TCP channel. It's party time!

We can type our little hearts out:
- ex. `Hello, world` echoes back `Hello, world`


## What would make it better

- Stress testing, security audit, read binary, etc.  