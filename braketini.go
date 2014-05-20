package braketini

import (
	"github.com/airbrake/gobrake"
	"github.com/go-martini/martini"
	"net/http"
)

const SecureConnection = true

type Config struct {
	ProjectId   int
	Key         string
	Environment string
}

func Middleware(config Config) martini.Handler {
	transport := gobrake.NewJSONTransport(config.ProjectId, config.Key, SecureConnection)
	notifier := gobrake.NewNotifier(transport)
	notifier.SetContext("environment", config.Environment)
	notifier.SetContext("version", "1.0")

	return func(context martini.Context, request *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				notifier.NotifyPanic(err, request, nil)
			}
		}()
		context.Next()
	}
}
