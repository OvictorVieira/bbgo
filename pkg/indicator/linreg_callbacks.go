// Code generated by "callbackgen -type LinReg"; DO NOT EDIT.

package indicator

import ()

func (lr *LinReg) OnUpdate(cb func(value float64)) {
	lr.UpdateCallbacks = append(lr.UpdateCallbacks, cb)
}

func (lr *LinReg) EmitUpdate(value float64) {
	for _, cb := range lr.UpdateCallbacks {
		cb(value)
	}
}
