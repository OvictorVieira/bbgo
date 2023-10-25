package glassnode

import (
	"time"

	"github.com/OvictorVieira/bbgo/pkg/datasource/glassnode/glassnodeapi"
)

type QueryOptions struct {
	Since    *time.Time
	Until    *time.Time
	Interval *glassnodeapi.Interval
	Currency *string
}
