package config

import (
	"context"

	"golang.ngrok.com/ngrok"
	"golang.ngrok.com/ngrok/config"
)

func LoadNgrok(ctx context.Context) ngrok.Tunnel {
	l, err := ngrok.Listen(ctx,
		config.LabeledTunnel(
			config.WithLabel("edge", "edghts_2bPQhZVW5VersFF4wI3VaR8q1bl"),
		), ngrok.WithAuthtoken(ENV.NGROK_AUTHTOKEN))
	if err != nil {
		panic(err)
	}
	return l
}
