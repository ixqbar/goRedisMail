package mailer

import (
	"errors"
	"github.com/jonnywang/go-kits/redis"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

type mailerRedisHandler struct {
	redis.RedisHandler
}

func (obj *mailerRedisHandler) Init() error {
	obj.Initiation(func() {
		GmailHandler.Init()
	})

	return nil
}

func (obj *mailerRedisHandler) Shutdown() {
	Logger.Print("mailer server will shutdown")
}

func (obj *mailerRedisHandler) Version() (string, error) {
	return VERSION, nil
}

func (obj *mailerRedisHandler) Ping(message string) (string, error) {
	if len(message) > 0 {
		return message, nil
	}

	return "PONG", nil
}

func (obj *mailerRedisHandler) Hmset(name string, values map[string][]byte) (error) {
	if _, ok := values["To"]; !ok {
		return errors.New("not found params `To`")
	}

	if _, ok := values["Subject"]; !ok {
		return errors.New("not found params `Subject`")
	}

	if _, ok := values["Content"]; !ok {
		return errors.New("not found params `Content`")
	}

	go GmailHandler.Sender(strings.Split(string(values["To"]), ","), string(values["Subject"]), string(values["Content"]))

	return nil
}

func Run() {
	mailRedisHandler := &mailerRedisHandler{}

	err := mailRedisHandler.Init()
	if err != nil {
		Logger.Print(err)
		return
	}

	mailRedisServer, err := redis.NewServer(GConfig.ListenServer, mailRedisHandler)
	if err != nil {
		Logger.Print(err)
		return
	}

	serverStop := make(chan bool)
	stopSignal := make(chan os.Signal)
	signal.Notify(stopSignal, syscall.SIGTERM, syscall.SIGINT)

	go func() {
		<-stopSignal
		Logger.Print("catch exit signal")
		mailRedisServer.Stop(10)
		GmailHandler.Stop()
		serverStop <- true
	}()

	err = mailRedisServer.Start()
	if err != nil {
		Logger.Print(err)
		stopSignal <- syscall.SIGTERM
	}

	<-serverStop

	close(serverStop)
	close(stopSignal)

	Logger.Print("all server shutdown")
}
