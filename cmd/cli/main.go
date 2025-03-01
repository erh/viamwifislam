package main

import (
	"context"
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

	deps := resource.Dependencies{}
	// can load these from a remote moachine if you need

	cfg := viamwifislam.Config{}

	thing, err := viamwifislam.NewUnifi(ctx, deps, slam.Named("foo"), &cfg, logger)
	if err != nil {
		return err
	}
	defer thing.Close(ctx)

	_, err = viamwifislam.DoScan(ctx)
	if err != nil {
		return err
	}

	return nil
}
