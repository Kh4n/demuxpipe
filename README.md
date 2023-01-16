# demuxpipe

A tool which multiplexes incoming connections over a TCP connection to a demultiplexer which forwards the data to another source.

This is useful for when you need to setup a proxy but want exactly 1 TCP connection used between your computer and the proxy.

It may also have other uses.

```
Multiplex incoming connections to another address

Usage:
  demuxpipe mux [flags]

Flags:
  -b, --bind string     address to send messages from (default ":9999")
  -h, --help            help for mux
  -l, --listen string   address to listen on (default ":8889")
  -w, --write string    address to mux to

Global Flags:
  -v, --verbose   enable all debug output
```
```
Demuxes incoming messages to another address

Usage:
  demuxpipe demux [flags]

Flags:
  -h, --help            help for demux
  -l, --listen string   address to receive mux (default ":8889")
  -w, --write string    address to demux to (eg. localhost proxy)

Global Flags:
  -v, --verbose   enable all debug output
```