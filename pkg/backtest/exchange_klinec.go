package backtest

import (
	"github.com/OvictorVieira/promeheux.api/pkg/bbgo"
	"github.com/OvictorVieira/promeheux.api/pkg/types"
)

type ExchangeDataSource struct {
	C         chan types.KLine
	Exchange  *Exchange
	Session   *bbgo.ExchangeSession
	Callbacks []func(types.KLine, *ExchangeDataSource)
}
