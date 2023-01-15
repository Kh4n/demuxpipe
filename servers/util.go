package servers

import (
	"encoding/gob"
	"log"
	"net"
	"sync"
)

type lockMap struct {
	table map[string]net.Conn
	lock  *sync.RWMutex
}

func newLockMap() *lockMap {
	return &lockMap{
		table: make(map[string]net.Conn),
		lock:  &sync.RWMutex{},
	}
}

func (m *lockMap) get(key string) (net.Conn, bool) {
	m.lock.RLock()
	ret, exists := m.table[key]
	m.lock.RUnlock()
	return ret, exists
}
func (m *lockMap) put(key string, val net.Conn) {
	m.lock.Lock()
	m.table[key] = val
	m.lock.Unlock()
}
func (m *lockMap) del(key string) {
	m.lock.Lock()
	delete(m.table, key)
	m.lock.Unlock()
}

func handlePipeConn(conn net.Conn, conns *lockMap, addr string, enc *gob.Encoder) {
	conns.put(addr, conn)
	buf := make([]byte, 1024)
	for {
		size, err := conn.Read(buf)
		if err != nil {
			enc.Encode(newCloseMessage(addr))
			log.Println("unable to read from conn:", addr, err)
			conns.del(addr)
			err = conn.Close()
			if err != nil {
				log.Println("unable to close conn:", addr, err)
				return
			}
			log.Println("closed conn:", addr)
			return
		}
		log.Println("read", size, "bytes from", addr)
		msg := newDataMessage(addr, buf[:size])
		err = enc.Encode(msg)
		if err != nil {
			log.Println("unable to write to mux:", err)
		}
		log.Println("sent", size, "bytes to pipe")
	}
}

func handleMessages(enc *gob.Encoder, dec *gob.Decoder, conns *lockMap, writeAddr *string) {
	for {
		var msg message
		err := dec.Decode(&msg)
		if err != nil {
			log.Fatalln("unable to decode from pipe:", err)
		}
		log.Println("received", len(msg.Data), "bytes from pipe with addr:", msg.Addr)
		conn, exists := conns.get(msg.Addr)
		if !exists {
			if writeAddr == nil {
				log.Println("non existent addr received from pipe:", msg.Addr)
				continue
			}
			conn, err = net.Dial("tcp", *writeAddr)
			if err != nil {
				log.Fatalln("unable to dial:", *writeAddr)
			}
			go handlePipeConn(conn, conns, msg.Addr, enc)
		}

		switch msg.Type {
		case CLOSE:
			{
				conns.del(msg.Addr)
				err = conn.Close()
				if err != nil {
					log.Println("unable to close conn:", msg.Addr, err)
					continue
				}
				log.Println("close instruction received:", msg.Addr)
			}
		case DATA:
			{
				size, err := conn.Write(msg.Data)
				if err != nil {
					log.Println("unable to write to connection:", msg.Addr, err)
					continue
				}
				if writeAddr != nil {
					log.Println("sent", size, "bytes to", *writeAddr, "on behalf of", msg.Addr)
				} else {
					log.Println("sent", size, "bytes to", msg.Addr)
				}
			}
		default:
			log.Println("invalid msg type received:", msg.Type)
		}
	}
}
