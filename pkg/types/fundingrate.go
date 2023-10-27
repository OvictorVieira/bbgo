package types

import (
	"time"

	"github.com/OvictorVieira/promeheux.api/pkg/fixedpoint"
)

type FundingRate struct {
	FundingRate fixedpoint.Value
	FundingTime time.Time
	Time        time.Time
}
