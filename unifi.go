package viamwifislam

import (
	"context"

	"github.com/pkg/errors"
	"go.viam.com/rdk/logging"
	"go.viam.com/rdk/resource"
	"go.viam.com/rdk/services/slam"
	"go.viam.com/rdk/spatialmath"
	"go.viam.com/utils/rpc"
)

var (
	Unifi            = resource.NewModel("erh", "viamwifislam", "unifi")
	errUnimplemented = errors.New("unimplemented")
)

func init() {
	resource.RegisterService(slam.API, Unifi,
		resource.Registration[slam.Service, *Config]{
			Constructor: newViamwifislamUnifi,
		},
	)
}

type Config struct {
	/*
		Put config attributes here. There should be public/exported fields
		with a `json` parameter at the end of each attribute.

		Example config struct:
			type Config struct {
				Pin   string `json:"pin"`
				Board string `json:"board"`
				MinDeg *float64 `json:"min_angle_deg,omitempty"`
			}

		If your model does not need a config, replace *Config in the init
		function with resource.NoNativeConfig
	*/
}

// Validate ensures all parts of the config are valid and important fields exist.
// Returns implicit dependencies based on the config.
// The path is the JSON path in your robot's config (not the `Config` struct) to the
// resource being validated; e.g. "components.0".
func (cfg *Config) Validate(path string) ([]string, error) {
	// Add config validation code here
	return nil, nil
}

type viamwifislamUnifi struct {
	resource.AlwaysRebuild

	name resource.Name

	logger logging.Logger
	cfg    *Config

	cancelCtx  context.Context
	cancelFunc func()
}

func newViamwifislamUnifi(ctx context.Context, deps resource.Dependencies, rawConf resource.Config, logger logging.Logger) (slam.Service, error) {
	conf, err := resource.NativeConfig[*Config](rawConf)
	if err != nil {
		return nil, err
	}

	return NewUnifi(ctx, deps, rawConf.ResourceName(), conf, logger)

}

func NewUnifi(ctx context.Context, deps resource.Dependencies, name resource.Name, conf *Config, logger logging.Logger) (slam.Service, error) {

	cancelCtx, cancelFunc := context.WithCancel(context.Background())

	s := &viamwifislamUnifi{
		name:       name,
		logger:     logger,
		cfg:        conf,
		cancelCtx:  cancelCtx,
		cancelFunc: cancelFunc,
	}
	return s, nil
}

func (s *viamwifislamUnifi) Name() resource.Name {
	return s.name
}

func (s *viamwifislamUnifi) NewClientFromConn(ctx context.Context, conn rpc.ClientConn, remoteName string, name resource.Name, logger logging.Logger) (slam.Service, error) {
	panic("not implemented")
}

func (s *viamwifislamUnifi) Position(ctx context.Context) (spatialmath.Pose, error) {
	panic("not implemented")
}

func (s *viamwifislamUnifi) PointCloudMap(ctx context.Context, returnEditedMap bool) (func() ([]byte, error), error) {
	panic("not implemented")
}

func (s *viamwifislamUnifi) InternalState(ctx context.Context) (func() ([]byte, error), error) {
	panic("not implemented")
}

func (s *viamwifislamUnifi) Properties(ctx context.Context) (slam.Properties, error) {
	panic("not implemented")
}

func (s *viamwifislamUnifi) DoCommand(ctx context.Context, cmd map[string]interface{}) (map[string]interface{}, error) {
	panic("not implemented")
}

func (s *viamwifislamUnifi) Close(context.Context) error {
	// Put close code here
	s.cancelFunc()
	return nil
}
