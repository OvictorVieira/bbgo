package fmaker

import (
	"fmt"
	"math"
	"time"

	"github.com/OvictorVieira/bbgo/pkg/datatype/floats"
	"github.com/OvictorVieira/bbgo/pkg/indicator"
	"github.com/OvictorVieira/bbgo/pkg/types"
)

//go:generate callbackgen -type S1
type S1 struct {
	types.IntervalWindow
	Values  floats.Slice
	EndTime time.Time

	UpdateCallbacks []func(value float64)
}

func (inc *S1) Last(int) float64 {
	if len(inc.Values) == 0 {
		return 0.0
	}
	return inc.Values[len(inc.Values)-1]
}

func (inc *S1) CalculateAndUpdate(klines []types.KLine) {
	if len(klines) < inc.Window {
		return
	}

	var end = len(klines) - 1
	var lastKLine = klines[end]

	if inc.EndTime != zeroTime && lastKLine.GetEndTime().Before(inc.EndTime) {
		return
	}

	var recentT = klines[end-(inc.Window-1) : end+1]

	correlation, err := calculateS1(recentT, inc.Window, KLineAmplitudeMapper, types.KLineVolumeMapper)
	if err != nil {
		log.WithError(err).Error("can not calculate correlation")
		return
	}
	inc.Values.Push(correlation)

	if len(inc.Values) > indicator.MaxNumOfVOL {
		inc.Values = inc.Values[indicator.MaxNumOfVOLTruncateSize-1:]
	}

	inc.EndTime = klines[end].GetEndTime().Time()

	inc.EmitUpdate(correlation)
}

func (inc *S1) handleKLineWindowUpdate(interval types.Interval, window types.KLineWindow) {
	if inc.Interval != interval {
		return
	}

	inc.CalculateAndUpdate(window)
}

func (inc *S1) Bind(updater indicator.KLineWindowUpdater) {
	updater.OnKLineWindowUpdate(inc.handleKLineWindowUpdate)
}

func calculateS1(klines []types.KLine, window int, valA KLineValueMapper, valB KLineValueMapper) (float64, error) {
	length := len(klines)
	if length == 0 || length < window {
		return 0.0, fmt.Errorf("insufficient elements for calculating VOL with window = %d", window)
	}

	sumA, sumB, sumAB, squareSumA, squareSumB := 0., 0., 0., 0., 0.
	for _, k := range klines {
		// sum of elements of array A
		sumA += valA(k)
		// sum of elements of array B
		sumB += valB(k)

		// sum of A[i] * B[i].
		sumAB = sumAB + valA(k)*valB(k)

		// sum of square of array elements.
		squareSumA = squareSumA + valA(k)*valA(k)
		squareSumB = squareSumB + valB(k)*valB(k)
	}
	// use formula for calculating correlation coefficient.
	corr := (float64(window)*sumAB - sumA*sumB) /
		math.Sqrt((float64(window)*squareSumA-sumA*sumA)*(float64(window)*squareSumB-sumB*sumB))

	return -corr, nil
}

func KLineAmplitudeMapper(k types.KLine) float64 {
	return k.High.Div(k.Low).Float64()
}
