// Code generated by "callbackgen -type PositionRiskControl"; DO NOT EDIT.

package riskcontrol

import (
	"github.com/OvictorVieira/promeheux.api/pkg/fixedpoint"
	"github.com/OvictorVieira/promeheux.api/pkg/types"
)

func (p *PositionRiskControl) OnReleasePosition(cb func(quantity fixedpoint.Value, side types.SideType)) {
	p.releasePositionCallbacks = append(p.releasePositionCallbacks, cb)
}

func (p *PositionRiskControl) EmitReleasePosition(quantity fixedpoint.Value, side types.SideType) {
	for _, cb := range p.releasePositionCallbacks {
		cb(quantity, side)
	}
}
