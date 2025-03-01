package main

import (
	"context"
	"flag"
	"fmt"

	"go.viam.com/rdk/logging"
	"go.viam.com/rdk/resource"
	"go.viam.com/rdk/services/slam"
	"viamwifislam"
)

func main() {
	err := realMain()
	if err != nil {
		panic(err)
	}
}

func realMain() error {
	ctx := context.Background()
	logger := logging.NewLogger("cli")

	ssid := ""

	flag.StringVar(&ssid, "ssid", ssid, "what ssid to scan for")

	flag.Parse()

	deps := resource.Dependencies{}
	// can load these from a remote moachine if you need

	cfg := viamwifislam.Config{}

	thing, err := viamwifislam.NewUnifi(ctx, deps, slam.Named("foo"), &cfg, logger)
	if err != nil {
		return err
	}
	defer thing.Close(ctx)

	res, err := viamwifislam.DoScan(ctx, ssid)
	if err != nil {
		return err
	}

	for _, r := range res {
		fmt.Printf("%v\n", r)
	}

	return nil
}
