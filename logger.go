package archiver

import "go.uber.org/zap"

var lg *zap.Logger

func init() {
	var err error
	lg, err = zap.NewProduction()
	if err != nil {
		panic(err)
	}
}
