package servers

import (
	"encoding/gob"
	"log"
	"net"
	"time"
)

func MuxToPipe(listenAddrStr string, writeAddrStr string, bindAddrStr string) {
	writeAddr, err := net.ResolveTCPAddr("tcp", writeAddrStr)
	if err != nil {
		log.Fatalln("unable to resolve writeAddr:", writeAddrStr, err)
	}
	bindAddr, err := net.ResolveTCPAddr("tcp", bindAddrStr)
	if err != nil {
		log.Fatalln("unable to resolve bindAddr:", bindAddrStr, err)
	}

	ln, err := net.Listen("tcp", listenAddrStr)
	if err != nil {
		log.Fatalln("unable to listen:", listenAddrStr, err)
	}

	pipe, err := net.DialTCP("tcp", bindAddr, writeAddr)
	if err != nil {
		log.Fatalln("unable dial:", bindAddr, writeAddr, err)
	}
	pipe.SetKeepAlivePeriod(time.Second * 10)
	pipe.SetKeepAlive(true)
	log.Println("connected to", writeAddr)

	handleConnsToPipe(pipe, ln)
}

func handleConnsToPipe(pipe net.Conn, ln net.Listener) {
	conns := newLockMap()
	enc := gob.NewEncoder(pipe)
	dec := gob.NewDecoder(pipe)
	go handleNewConns(ln, conns, enc)

	handleMessages(enc, dec, conns, nil)
}

func handleNewConns(ln net.Listener, conns *lockMap, enc *gob.Encoder) {
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println("unable to accept connection:", err)
			continue
		}
		raddr := conn.RemoteAddr()
		if raddr == nil {
			log.Println("unable to get remote address")
			continue
		}
		go handlePipeConn(conn, conns, raddr.String(), enc)
	}
}
