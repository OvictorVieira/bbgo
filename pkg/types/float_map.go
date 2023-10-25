package types

import "github.com/OvictorVieira/bbgo/pkg/datatype/floats"

var _ Series = floats.Slice([]float64{}).Addr()
