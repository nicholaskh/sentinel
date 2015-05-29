package plugin

import (
	"net"
	"time"

	log "github.com/nicholaskh/log4go"
	"github.com/nicholaskh/sentinel/config"
	"github.com/nicholaskh/sentinel/engine"
)

const (
	PACKET_PING = "PING"
	PACKET_PONG = "PONG"
)

func init() {
	engine.RegisterPlugin("cmd_ping", func() engine.Plugin {
		return new(PingCommand)
	})
}

type PingCommand struct {
	target     string
	interval   time.Duration
	listenPort int
}

func (this *PingCommand) Init(config *config.ServiceConfig) {
	this.target = config.Target
	this.interval = config.Interval
	this.listenPort = config.ListenPort
}

func (this *PingCommand) Start() {
	socket, err := net.ListenUDP("udp4", &net.UDPAddr{
		IP:   net.IPv4(0, 0, 0, 0),
		Port: this.listenPort,
	})
	if err != nil {
		panic(err)
	}

	for {
		select {
		case <-time.Tick(this.interval):
			conn, err := net.DialTimeout("udp", this.target, time.Second*10)
			if err != nil {
				log.Warn("dial target[%s] error: %s", this.target, err.Error())
			}
			conn.Write([]byte(PACKET_PING))

			data := make([]byte, 256)
			read, _, err := socket.ReadFromUDP(data)
			if err != nil {
				log.Warn("read from target[%s] error, %s", this.target, err.Error())
				// TODO retry
			}
			pong := string(data[:read])
			if pong != PACKET_PONG {
				log.Warn("pong packet error: give %s, expected %s", pong, PACKET_PONG)
				// TODO retry
			}
		}
	}
}
