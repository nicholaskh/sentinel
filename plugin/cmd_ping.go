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
	engine.RegisterPlugin("ping", func() engine.Plugin {
		return new(PingCommand)
	})
}

type PingCommand struct {
	target        string
	interval      time.Duration
	retryTimes    int
	retryInterval time.Duration
	readTimeout   time.Duration
	listenPort    int
}

func (this *PingCommand) Init(config *config.ServiceConfig) {
	this.target = config.Target
	this.interval = config.Interval
	this.retryTimes = config.Retry
	this.retryInterval = config.RetryInterval
	this.readTimeout = config.ReadTimeout
}

func (this *PingCommand) Start() {
	for {
		select {
		case <-time.Tick(this.interval):
			var success bool
			log.Info("ping %s", this.target)
			if success = this.ping(); !success {
				time.Sleep(this.retryInterval)
				success = this.retry(0)
			}

			if success {
				log.Info("ping %s success", this.target)
			} else {
				log.Info("ping %s fail", this.target)
				// TODO -- alarm
			}
		}
	}
}

func (this *PingCommand) ping() (success bool) {
	conn, err := net.DialTimeout("udp", this.target, time.Second*10)
	if err != nil {
		log.Warn("dial target[%s] error: %s", this.target, err.Error())
		return false
	}
	conn.Write([]byte(PACKET_PING))

	conn.SetReadDeadline(time.Now().Add(this.readTimeout))
	data := make([]byte, 256)
	read, err := conn.Read(data)
	if err != nil {
		log.Warn("read from target[%s] error, %s", this.target, err.Error())
		return false
	}
	pong := string(data[:read])
	if pong != PACKET_PONG {
		log.Warn("pong packet error: give %s, expected %s", pong, PACKET_PONG)
		return false
	}
	return true
}

func (this *PingCommand) retry(retried int) bool {
	retried++
	log.Info("retry %d", retried)
	if retried >= this.retryTimes {
		return false
	}
	if !this.ping() {
		time.Sleep(this.retryInterval)
		return this.retry(retried)
	}
	return true
}
