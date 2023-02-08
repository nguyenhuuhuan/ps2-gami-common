package newrelic

import (
	"os"

	"gitlab.id.vin/gami/go-agent/v3/newrelic"
)

func NewNewRelic() (*newrelic.Application, error) {
	return newrelic.NewApplication(
		newrelic.ConfigAppName(os.Getenv("NEW_RELIC_APP_NAME")),
		newrelic.ConfigFromEnvironment(),
	)
}
