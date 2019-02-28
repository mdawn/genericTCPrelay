Terminal 1: go run relay.go
Terminal 2: go run echo.go
Terminal 3: telnet localhost 8081

Why 3 separate terminals?  Because go hates that func main repeat, and the go routine in relay could use a channel to speed things up. As it stands, my version of telnet fails with much speed. 

TODO
- Works with one server (& many clients), & only reads strings
- Ports are hard coded


**************
# Generic TCP 

A little system with a single relay server that supports multiple echos, and allows each echo to have multiple clients(telnet). 

[Evans](https://github.com/ktr0731/evans)


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
- `xxx`

**STEP 2**: We open another terminal and 
- Install gRPC: `xxx`

**STEP 3**: We open our last terminal
- `xxx`

Then connect:
- `xxx`


## What would make it better

- Stress testing & security audit, among other measures