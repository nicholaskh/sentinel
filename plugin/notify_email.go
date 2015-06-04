package plugin

import (
	"fmt"
	"net/smtp"
	"strings"

	log "github.com/nicholaskh/log4go"
	"github.com/nicholaskh/sentinel/config"
	"github.com/nicholaskh/sentinel/engine"
)

func init() {
	engine.RegisterPlugin("email", func() engine.Plugin {
		return new(EmailNotification)
	})
}

type EmailNotification struct {
	config        *config.NotificationConfig
	serviceConfig *config.ServiceConfig
	notifyQueue   chan interface{}
}

func (this *EmailNotification) Init(config *config.NotificationConfig, serviceConfig *config.ServiceConfig) {
	this.config = config
	this.serviceConfig = serviceConfig
	this.notifyQueue = make(chan interface{})
}

func (this *EmailNotification) Start() {
	for _ = range this.notifyQueue {
		err := this.sendMail()
		if err != nil {
			log.Warn("Send mail error: %s", err.Error())
		}
	}
}

func (this *EmailNotification) sendMail() error {
	log.Info("send mail")
	user := "zhangkh.3@163.com"
	password := "19871013aA"
	host := "smtp.163.com:25"

	subject := fmt.Sprintf("【ALERT】%s[%s] Down", this.serviceConfig.Name, this.serviceConfig.Target)

	body := fmt.Sprintf(`
		<html>
		<body>
		<h3>
		%s at %s is down, please contact administrator ASAP!
		</h3>
		<br />
		If you do not care about this message, please ignore.
		</body>
		</html>
		`, this.serviceConfig.Name, this.serviceConfig.Target)
	hp := strings.Split(host, ":")
	auth := smtp.PlainAuth("", user, password, hp[0])
	content_type := "Content-Type: text/html; charset=UTF-8"

	msg := []byte("To: " + strings.Join(this.config.Notifiers, ";") + "\r\nFrom: " + user + ">\r\nSubject: " + subject + "\r\n" + content_type + "\r\n\r\n" + body)
	err := smtp.SendMail(host, auth, user, this.config.Notifiers, msg)
	return err
}

func (this *EmailNotification) Notify() {
	this.notifyQueue <- 1
}
