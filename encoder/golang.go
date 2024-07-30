package encoder

import (
	"github.com/oodle-ai/grabana/encoder/golang"
	"github.com/oodle-ai/grafana-sdk"
	"go.uber.org/zap"
)

func ToGolang(logger *zap.Logger, dashboard sdk.Board) (string, error) {
	golangEncoder := golang.NewEncoder(logger)

	return golangEncoder.EncodeDashboard(dashboard)
}
