// Code generated by "callbackgen -type MarketDataStore"; DO NOT EDIT.

package bbgo

import (
	"github.com/OvictorVieira/promeheux.api/pkg/types"
)

func (store *MarketDataStore) OnKLineWindowUpdate(cb func(interval types.Interval, klines types.KLineWindow)) {
	store.kLineWindowUpdateCallbacks = append(store.kLineWindowUpdateCallbacks, cb)
}

func (store *MarketDataStore) EmitKLineWindowUpdate(interval types.Interval, klines types.KLineWindow) {
	for _, cb := range store.kLineWindowUpdateCallbacks {
		cb(interval, klines)
	}
}

func (store *MarketDataStore) OnKLineClosed(cb func(k types.KLine)) {
	store.kLineClosedCallbacks = append(store.kLineClosedCallbacks, cb)
}

func (store *MarketDataStore) EmitKLineClosed(k types.KLine) {
	for _, cb := range store.kLineClosedCallbacks {
		cb(k)
	}
}
