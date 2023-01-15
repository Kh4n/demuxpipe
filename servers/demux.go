package servers

import (
	"encoding/gob"
	"log"
	"net"
	"time"
)

func PipeToDemux(listenAddr string, writeAddr string) {
	addr, err := net.ResolveTCPAddr("tcp", listenAddr)
	if err != nil {
		log.Fatalln("unable to resolve:", listenAddr, err)
	}
	ln, err := net.ListenTCP("tcp", addr)
	if err != nil {
		log.Fatalln("unable to listen on:", listenAddr)
	}
	pipe, err := ln.AcceptTCP()
	if err != nil {
		log.Fatalln("unable to connect to mux:", err)
	}
	pipe.SetKeepAlivePeriod(time.Second * 10)
	pipe.SetKeepAlive(true)
	log.Println("connected to mux on:", pipe.RemoteAddr().String())

	handlePipeToConns(pipe, writeAddr)
}

func handlePipeToConns(pipe net.Conn, writeAddr string) {
	conns := newLockMap()
	enc := gob.NewEncoder(pipe)
	dec := gob.NewDecoder(pipe)
	handleMessages(enc, dec, conns, &writeAddr)
}
