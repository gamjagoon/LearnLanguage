package broker

import (
	"errors"
	"log"
	"net"
	"time"

	"github.com/eclipse/paho.mqtt.golang/packets"
	"go.uber.org/zap"
)

const (
	CLIENT  = 0
	ROUTER  = 1
	REMOTE  = 2
	CLUSTER = 3
)

type Broker struct {
	id     string
	config *Config
}

func NewBroker(config *Config) (*Broker, error) {
	if config == nil {
		return nil, errors.New("config error")
	}

	b := &Broker{
		id:     GenUniqueId(),
		config: config,
	}

	return b, nil
}

func (b *Broker) Start() {
	if b == nil {
		log.Println("Broker is null")
		return
	}

	if b.config.Port != "" {
		go b.StartClientListening(false)
	}
}

func (b *Broker) StartClientListening(Tls bool) {
	var err error
	var l net.Listener

	for {
		if Tls {
			log.Println("Start TLS Listening clinet on ", zap.String("hp", hp))
		} else {
			hp := b.config.Host + ":" + b.config.Port
			l, err = net.Listen("tcp", hp)
			log.Println("Start Listening clinet on ", zap.String("hp", hp))
		}

		if err == nil {
			break
		}

		log.Println("Error Listening on", zap.Error(err))
		time.Sleep(1 * time.Second)
	}

	tmpDelay := 10

	for {
		conn, err := l.Accept()
		if err != nil {
			if ne, ok := err.(net.Error); ok && ne.Temporary() {
				log.Println("Temporary Client Accept Error")
				time.Sleep(time.Duration(tmpDelay))
				tmpDelay *= 2
				if tmpDelay > 100 {
					tmpDelay = 100
				}
			} else {
				log.Println("Accept error")
			}
			continue
		}

		tmpDelay = 10

		go b.handleConnection(CLIENT, conn)
	}
}

func (b *Broker) handleConnection(typ int, conn net.Conn) {
	packet, err := packets.ReadPacket(conn)
	if err != nil {
		log.Println("read connect packet error")
		conn.Close()
		return
	}
	if packet == nil {
		log.Println("received nil packet")
		return
	}
	msg, ok := packet.(*packets.ConnackPacket)
	if !ok {
		log.Println("received msg that was not Connect")
		return
	}

	connack := packets.NewControlPacket(packets.Connack).(*packets.ConnackPacket)
}
