package types

import "github.com/OvictorVieira/promeheux.api/pkg/datatype/floats"

var _ Series = floats.Slice([]float64{}).Addr()
