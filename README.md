# Generic TCP 

## Description

A little system with a single relay server that supports multiple echos, and allows each echo to have multiple clients(telnet). 

## What We Need

We're on a Mac, right? Of course we are. 

- The latest version of [Go](https://golang.org/doc/install#install)
- This repo

## What it Does

With this TCP we can:

1. Run two relays at the same time, given we tell the relays what port to listen on
2. Run multiple echo servers concurrently
3. Run multiple clients concurrently per echo server

## Getting Operational

**STEP 1**: We open a terminal and  
- Run `go run relay.go`
- No specified port defaults to `8080`. However, we're free to utilize the optional flag `-p` to designate a port for compatibility with concurrent relays

**STEP 2**: We open another terminal and 
- Run `go run echo.go`
- Running this returns the `established relay address`

**STEP 3**: We open our last terminal and
- Run `telnet localhost 8081`

Then we can type our little hearts out:
- ex. `Hello, world` echoes back `Hello, world`

## Example Scenario

You wrote a program that, tragically, sits behind a firewall. You desperately want to expose your server! 

Fortunately we can use a TCP connection to bypass the firewall. 
**First:** Run the relay server and choose a port. Have your program connect on that same port. Your server will read the relay address through that connection.
**Second:** The relay server provides a different port for your clients to use. Tell your clients to connect to the relay server on that port. 

The relay server acts as an intermediary between our program's server and our desired clients - here we used telnet. 


## What would make it better

- A script to reduce the annoyance of having to open 3 terminals
- Stress testing, security audit, among other measures