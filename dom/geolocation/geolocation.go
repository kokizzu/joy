package geolocation

import (
	"github.com/matthewmueller/joy/dom/position"
	"github.com/matthewmueller/joy/dom/positionerror"
	"github.com/matthewmueller/joy/dom/positionoptions"
	"github.com/matthewmueller/joy/js"
)

// Geolocation struct
// js:"Geolocation,omit"
type Geolocation struct {
}

// ClearWatch fn
// js:"clearWatch"
func (*Geolocation) ClearWatch(watchId int) {
	js.Rewrite("$_.clearWatch($1)", watchId)
}

// GetCurrentPosition fn
// js:"getCurrentPosition"
func (*Geolocation) GetCurrentPosition(successCallback func(position *position.Position), errorCallback *func(err *positionerror.PositionError), options *positionoptions.PositionOptions) {
	js.Rewrite("$_.getCurrentPosition($1, $2, $3)", successCallback, errorCallback, options)
}

// WatchPosition fn
// js:"watchPosition"
func (*Geolocation) WatchPosition(successCallback func(position *position.Position), errorCallback *func(err *positionerror.PositionError), options *positionoptions.PositionOptions) (i int) {
	js.Rewrite("$_.watchPosition($1, $2, $3)", successCallback, errorCallback, options)
	return i
}
