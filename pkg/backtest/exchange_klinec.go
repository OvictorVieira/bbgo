package backtest

import (
	"github.com/OvictorVieira/bbgo/pkg/bbgo"
	"github.com/OvictorVieira/bbgo/pkg/types"
)

type ExchangeDataSource struct {
	C         chan types.KLine
	Exchange  *Exchange
	Session   *bbgo.ExchangeSession
	Callbacks []func(types.KLine, *ExchangeDataSource)
}
