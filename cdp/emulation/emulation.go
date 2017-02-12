// Package emulation provides the Chrome Debugging Protocol
// commands, types, and events for the Emulation domain.
//
// This domain emulates different environments for the page.
//
// Generated by the chromedp-gen command.
package emulation

// AUTOGENERATED. DO NOT EDIT.

import (
	"context"

	cdp "github.com/knq/chromedp/cdp"
	"github.com/mailru/easyjson"
)

// SetDeviceMetricsOverrideParams overrides the values of device screen
// dimensions (window.screen.width, window.screen.height, window.innerWidth,
// window.innerHeight, and "device-width"/"device-height"-related CSS media
// query results).
type SetDeviceMetricsOverrideParams struct {
	Width             int64              `json:"width"`                       // Overriding width value in pixels (minimum 0, maximum 10000000). 0 disables the override.
	Height            int64              `json:"height"`                      // Overriding height value in pixels (minimum 0, maximum 10000000). 0 disables the override.
	DeviceScaleFactor float64            `json:"deviceScaleFactor"`           // Overriding device scale factor value. 0 disables the override.
	Mobile            bool               `json:"mobile"`                      // Whether to emulate mobile device. This includes viewport meta tag, overlay scrollbars, text autosizing and more.
	FitWindow         bool               `json:"fitWindow"`                   // Whether a view that exceeds the available browser window area should be scaled down to fit.
	Scale             float64            `json:"scale,omitempty"`             // Scale to apply to resulting view image. Ignored in |fitWindow| mode.
	ScreenWidth       int64              `json:"screenWidth,omitempty"`       // Overriding screen width value in pixels (minimum 0, maximum 10000000). Only used for |mobile==true|.
	ScreenHeight      int64              `json:"screenHeight,omitempty"`      // Overriding screen height value in pixels (minimum 0, maximum 10000000). Only used for |mobile==true|.
	PositionX         int64              `json:"positionX,omitempty"`         // Overriding view X position on screen in pixels (minimum 0, maximum 10000000). Only used for |mobile==true|.
	PositionY         int64              `json:"positionY,omitempty"`         // Overriding view Y position on screen in pixels (minimum 0, maximum 10000000). Only used for |mobile==true|.
	ScreenOrientation *ScreenOrientation `json:"screenOrientation,omitempty"` // Screen orientation override.
}

// SetDeviceMetricsOverride overrides the values of device screen dimensions
// (window.screen.width, window.screen.height, window.innerWidth,
// window.innerHeight, and "device-width"/"device-height"-related CSS media
// query results).
//
// parameters:
//   width - Overriding width value in pixels (minimum 0, maximum 10000000). 0 disables the override.
//   height - Overriding height value in pixels (minimum 0, maximum 10000000). 0 disables the override.
//   deviceScaleFactor - Overriding device scale factor value. 0 disables the override.
//   mobile - Whether to emulate mobile device. This includes viewport meta tag, overlay scrollbars, text autosizing and more.
//   fitWindow - Whether a view that exceeds the available browser window area should be scaled down to fit.
func SetDeviceMetricsOverride(width int64, height int64, deviceScaleFactor float64, mobile bool, fitWindow bool) *SetDeviceMetricsOverrideParams {
	return &SetDeviceMetricsOverrideParams{
		Width:             width,
		Height:            height,
		DeviceScaleFactor: deviceScaleFactor,
		Mobile:            mobile,
		FitWindow:         fitWindow,
	}
}

// WithScale scale to apply to resulting view image. Ignored in |fitWindow|
// mode.
func (p SetDeviceMetricsOverrideParams) WithScale(scale float64) *SetDeviceMetricsOverrideParams {
	p.Scale = scale
	return &p
}

// WithScreenWidth overriding screen width value in pixels (minimum 0,
// maximum 10000000). Only used for |mobile==true|.
func (p SetDeviceMetricsOverrideParams) WithScreenWidth(screenWidth int64) *SetDeviceMetricsOverrideParams {
	p.ScreenWidth = screenWidth
	return &p
}

// WithScreenHeight overriding screen height value in pixels (minimum 0,
// maximum 10000000). Only used for |mobile==true|.
func (p SetDeviceMetricsOverrideParams) WithScreenHeight(screenHeight int64) *SetDeviceMetricsOverrideParams {
	p.ScreenHeight = screenHeight
	return &p
}

// WithPositionX overriding view X position on screen in pixels (minimum 0,
// maximum 10000000). Only used for |mobile==true|.
func (p SetDeviceMetricsOverrideParams) WithPositionX(positionX int64) *SetDeviceMetricsOverrideParams {
	p.PositionX = positionX
	return &p
}

// WithPositionY overriding view Y position on screen in pixels (minimum 0,
// maximum 10000000). Only used for |mobile==true|.
func (p SetDeviceMetricsOverrideParams) WithPositionY(positionY int64) *SetDeviceMetricsOverrideParams {
	p.PositionY = positionY
	return &p
}

// WithScreenOrientation screen orientation override.
func (p SetDeviceMetricsOverrideParams) WithScreenOrientation(screenOrientation *ScreenOrientation) *SetDeviceMetricsOverrideParams {
	p.ScreenOrientation = screenOrientation
	return &p
}

// Do executes Emulation.setDeviceMetricsOverride against the provided context and
// target handler.
func (p *SetDeviceMetricsOverrideParams) Do(ctxt context.Context, h cdp.Handler) (err error) {
	if ctxt == nil {
		ctxt = context.Background()
	}

	// marshal
	buf, err := easyjson.Marshal(p)
	if err != nil {
		return err
	}

	// execute
	ch := h.Execute(ctxt, cdp.CommandEmulationSetDeviceMetricsOverride, easyjson.RawMessage(buf))

	// read response
	select {
	case res := <-ch:
		if res == nil {
			return cdp.ErrChannelClosed
		}

		switch v := res.(type) {
		case easyjson.RawMessage:
			return nil

		case error:
			return v
		}

	case <-ctxt.Done():
		return ctxt.Err()
	}

	return cdp.ErrUnknownResult
}

// ClearDeviceMetricsOverrideParams clears the overriden device metrics.
type ClearDeviceMetricsOverrideParams struct{}

// ClearDeviceMetricsOverride clears the overriden device metrics.
func ClearDeviceMetricsOverride() *ClearDeviceMetricsOverrideParams {
	return &ClearDeviceMetricsOverrideParams{}
}

// Do executes Emulation.clearDeviceMetricsOverride against the provided context and
// target handler.
func (p *ClearDeviceMetricsOverrideParams) Do(ctxt context.Context, h cdp.Handler) (err error) {
	if ctxt == nil {
		ctxt = context.Background()
	}

	// execute
	ch := h.Execute(ctxt, cdp.CommandEmulationClearDeviceMetricsOverride, cdp.Empty)

	// read response
	select {
	case res := <-ch:
		if res == nil {
			return cdp.ErrChannelClosed
		}

		switch v := res.(type) {
		case easyjson.RawMessage:
			return nil

		case error:
			return v
		}

	case <-ctxt.Done():
		return ctxt.Err()
	}

	return cdp.ErrUnknownResult
}

// ForceViewportParams overrides the visible area of the page. The change is
// hidden from the page, i.e. the observable scroll position and page scale does
// not change. In effect, the command moves the specified area of the page into
// the top-left corner of the frame.
type ForceViewportParams struct {
	X     float64 `json:"x"`     // X coordinate of top-left corner of the area (CSS pixels).
	Y     float64 `json:"y"`     // Y coordinate of top-left corner of the area (CSS pixels).
	Scale float64 `json:"scale"` // Scale to apply to the area (relative to a page scale of 1.0).
}

// ForceViewport overrides the visible area of the page. The change is hidden
// from the page, i.e. the observable scroll position and page scale does not
// change. In effect, the command moves the specified area of the page into the
// top-left corner of the frame.
//
// parameters:
//   x - X coordinate of top-left corner of the area (CSS pixels).
//   y - Y coordinate of top-left corner of the area (CSS pixels).
//   scale - Scale to apply to the area (relative to a page scale of 1.0).
func ForceViewport(x float64, y float64, scale float64) *ForceViewportParams {
	return &ForceViewportParams{
		X:     x,
		Y:     y,
		Scale: scale,
	}
}

// Do executes Emulation.forceViewport against the provided context and
// target handler.
func (p *ForceViewportParams) Do(ctxt context.Context, h cdp.Handler) (err error) {
	if ctxt == nil {
		ctxt = context.Background()
	}

	// marshal
	buf, err := easyjson.Marshal(p)
	if err != nil {
		return err
	}

	// execute
	ch := h.Execute(ctxt, cdp.CommandEmulationForceViewport, easyjson.RawMessage(buf))

	// read response
	select {
	case res := <-ch:
		if res == nil {
			return cdp.ErrChannelClosed
		}

		switch v := res.(type) {
		case easyjson.RawMessage:
			return nil

		case error:
			return v
		}

	case <-ctxt.Done():
		return ctxt.Err()
	}

	return cdp.ErrUnknownResult
}

// ResetViewportParams resets the visible area of the page to the original
// viewport, undoing any effects of the forceViewport command.
type ResetViewportParams struct{}

// ResetViewport resets the visible area of the page to the original
// viewport, undoing any effects of the forceViewport command.
func ResetViewport() *ResetViewportParams {
	return &ResetViewportParams{}
}

// Do executes Emulation.resetViewport against the provided context and
// target handler.
func (p *ResetViewportParams) Do(ctxt context.Context, h cdp.Handler) (err error) {
	if ctxt == nil {
		ctxt = context.Background()
	}

	// execute
	ch := h.Execute(ctxt, cdp.CommandEmulationResetViewport, cdp.Empty)

	// read response
	select {
	case res := <-ch:
		if res == nil {
			return cdp.ErrChannelClosed
		}

		switch v := res.(type) {
		case easyjson.RawMessage:
			return nil

		case error:
			return v
		}

	case <-ctxt.Done():
		return ctxt.Err()
	}

	return cdp.ErrUnknownResult
}

// ResetPageScaleFactorParams requests that page scale factor is reset to
// initial values.
type ResetPageScaleFactorParams struct{}

// ResetPageScaleFactor requests that page scale factor is reset to initial
// values.
func ResetPageScaleFactor() *ResetPageScaleFactorParams {
	return &ResetPageScaleFactorParams{}
}

// Do executes Emulation.resetPageScaleFactor against the provided context and
// target handler.
func (p *ResetPageScaleFactorParams) Do(ctxt context.Context, h cdp.Handler) (err error) {
	if ctxt == nil {
		ctxt = context.Background()
	}

	// execute
	ch := h.Execute(ctxt, cdp.CommandEmulationResetPageScaleFactor, cdp.Empty)

	// read response
	select {
	case res := <-ch:
		if res == nil {
			return cdp.ErrChannelClosed
		}

		switch v := res.(type) {
		case easyjson.RawMessage:
			return nil

		case error:
			return v
		}

	case <-ctxt.Done():
		return ctxt.Err()
	}

	return cdp.ErrUnknownResult
}

// SetPageScaleFactorParams sets a specified page scale factor.
type SetPageScaleFactorParams struct {
	PageScaleFactor float64 `json:"pageScaleFactor"` // Page scale factor.
}

// SetPageScaleFactor sets a specified page scale factor.
//
// parameters:
//   pageScaleFactor - Page scale factor.
func SetPageScaleFactor(pageScaleFactor float64) *SetPageScaleFactorParams {
	return &SetPageScaleFactorParams{
		PageScaleFactor: pageScaleFactor,
	}
}

// Do executes Emulation.setPageScaleFactor against the provided context and
// target handler.
func (p *SetPageScaleFactorParams) Do(ctxt context.Context, h cdp.Handler) (err error) {
	if ctxt == nil {
		ctxt = context.Background()
	}

	// marshal
	buf, err := easyjson.Marshal(p)
	if err != nil {
		return err
	}

	// execute
	ch := h.Execute(ctxt, cdp.CommandEmulationSetPageScaleFactor, easyjson.RawMessage(buf))

	// read response
	select {
	case res := <-ch:
		if res == nil {
			return cdp.ErrChannelClosed
		}

		switch v := res.(type) {
		case easyjson.RawMessage:
			return nil

		case error:
			return v
		}

	case <-ctxt.Done():
		return ctxt.Err()
	}

	return cdp.ErrUnknownResult
}

// SetVisibleSizeParams resizes the frame/viewport of the page. Note that
// this does not affect the frame's container (e.g. browser window). Can be used
// to produce screenshots of the specified size. Not supported on Android.
type SetVisibleSizeParams struct {
	Width  int64 `json:"width"`  // Frame width (DIP).
	Height int64 `json:"height"` // Frame height (DIP).
}

// SetVisibleSize resizes the frame/viewport of the page. Note that this does
// not affect the frame's container (e.g. browser window). Can be used to
// produce screenshots of the specified size. Not supported on Android.
//
// parameters:
//   width - Frame width (DIP).
//   height - Frame height (DIP).
func SetVisibleSize(width int64, height int64) *SetVisibleSizeParams {
	return &SetVisibleSizeParams{
		Width:  width,
		Height: height,
	}
}

// Do executes Emulation.setVisibleSize against the provided context and
// target handler.
func (p *SetVisibleSizeParams) Do(ctxt context.Context, h cdp.Handler) (err error) {
	if ctxt == nil {
		ctxt = context.Background()
	}

	// marshal
	buf, err := easyjson.Marshal(p)
	if err != nil {
		return err
	}

	// execute
	ch := h.Execute(ctxt, cdp.CommandEmulationSetVisibleSize, easyjson.RawMessage(buf))

	// read response
	select {
	case res := <-ch:
		if res == nil {
			return cdp.ErrChannelClosed
		}

		switch v := res.(type) {
		case easyjson.RawMessage:
			return nil

		case error:
			return v
		}

	case <-ctxt.Done():
		return ctxt.Err()
	}

	return cdp.ErrUnknownResult
}

// SetScriptExecutionDisabledParams switches script execution in the page.
type SetScriptExecutionDisabledParams struct {
	Value bool `json:"value"` // Whether script execution should be disabled in the page.
}

// SetScriptExecutionDisabled switches script execution in the page.
//
// parameters:
//   value - Whether script execution should be disabled in the page.
func SetScriptExecutionDisabled(value bool) *SetScriptExecutionDisabledParams {
	return &SetScriptExecutionDisabledParams{
		Value: value,
	}
}

// Do executes Emulation.setScriptExecutionDisabled against the provided context and
// target handler.
func (p *SetScriptExecutionDisabledParams) Do(ctxt context.Context, h cdp.Handler) (err error) {
	if ctxt == nil {
		ctxt = context.Background()
	}

	// marshal
	buf, err := easyjson.Marshal(p)
	if err != nil {
		return err
	}

	// execute
	ch := h.Execute(ctxt, cdp.CommandEmulationSetScriptExecutionDisabled, easyjson.RawMessage(buf))

	// read response
	select {
	case res := <-ch:
		if res == nil {
			return cdp.ErrChannelClosed
		}

		switch v := res.(type) {
		case easyjson.RawMessage:
			return nil

		case error:
			return v
		}

	case <-ctxt.Done():
		return ctxt.Err()
	}

	return cdp.ErrUnknownResult
}

// SetGeolocationOverrideParams overrides the Geolocation Position or Error.
// Omitting any of the parameters emulates position unavailable.
type SetGeolocationOverrideParams struct {
	Latitude  float64 `json:"latitude,omitempty"`  // Mock latitude
	Longitude float64 `json:"longitude,omitempty"` // Mock longitude
	Accuracy  float64 `json:"accuracy,omitempty"`  // Mock accuracy
}

// SetGeolocationOverride overrides the Geolocation Position or Error.
// Omitting any of the parameters emulates position unavailable.
//
// parameters:
func SetGeolocationOverride() *SetGeolocationOverrideParams {
	return &SetGeolocationOverrideParams{}
}

// WithLatitude mock latitude.
func (p SetGeolocationOverrideParams) WithLatitude(latitude float64) *SetGeolocationOverrideParams {
	p.Latitude = latitude
	return &p
}

// WithLongitude mock longitude.
func (p SetGeolocationOverrideParams) WithLongitude(longitude float64) *SetGeolocationOverrideParams {
	p.Longitude = longitude
	return &p
}

// WithAccuracy mock accuracy.
func (p SetGeolocationOverrideParams) WithAccuracy(accuracy float64) *SetGeolocationOverrideParams {
	p.Accuracy = accuracy
	return &p
}

// Do executes Emulation.setGeolocationOverride against the provided context and
// target handler.
func (p *SetGeolocationOverrideParams) Do(ctxt context.Context, h cdp.Handler) (err error) {
	if ctxt == nil {
		ctxt = context.Background()
	}

	// marshal
	buf, err := easyjson.Marshal(p)
	if err != nil {
		return err
	}

	// execute
	ch := h.Execute(ctxt, cdp.CommandEmulationSetGeolocationOverride, easyjson.RawMessage(buf))

	// read response
	select {
	case res := <-ch:
		if res == nil {
			return cdp.ErrChannelClosed
		}

		switch v := res.(type) {
		case easyjson.RawMessage:
			return nil

		case error:
			return v
		}

	case <-ctxt.Done():
		return ctxt.Err()
	}

	return cdp.ErrUnknownResult
}

// ClearGeolocationOverrideParams clears the overriden Geolocation Position
// and Error.
type ClearGeolocationOverrideParams struct{}

// ClearGeolocationOverride clears the overriden Geolocation Position and
// Error.
func ClearGeolocationOverride() *ClearGeolocationOverrideParams {
	return &ClearGeolocationOverrideParams{}
}

// Do executes Emulation.clearGeolocationOverride against the provided context and
// target handler.
func (p *ClearGeolocationOverrideParams) Do(ctxt context.Context, h cdp.Handler) (err error) {
	if ctxt == nil {
		ctxt = context.Background()
	}

	// execute
	ch := h.Execute(ctxt, cdp.CommandEmulationClearGeolocationOverride, cdp.Empty)

	// read response
	select {
	case res := <-ch:
		if res == nil {
			return cdp.ErrChannelClosed
		}

		switch v := res.(type) {
		case easyjson.RawMessage:
			return nil

		case error:
			return v
		}

	case <-ctxt.Done():
		return ctxt.Err()
	}

	return cdp.ErrUnknownResult
}

// SetTouchEmulationEnabledParams toggles mouse event-based touch event
// emulation.
type SetTouchEmulationEnabledParams struct {
	Enabled       bool                 `json:"enabled"`                 // Whether the touch event emulation should be enabled.
	Configuration EnabledConfiguration `json:"configuration,omitempty"` // Touch/gesture events configuration. Default: current platform.
}

// SetTouchEmulationEnabled toggles mouse event-based touch event emulation.
//
// parameters:
//   enabled - Whether the touch event emulation should be enabled.
func SetTouchEmulationEnabled(enabled bool) *SetTouchEmulationEnabledParams {
	return &SetTouchEmulationEnabledParams{
		Enabled: enabled,
	}
}

// WithConfiguration touch/gesture events configuration. Default: current
// platform.
func (p SetTouchEmulationEnabledParams) WithConfiguration(configuration EnabledConfiguration) *SetTouchEmulationEnabledParams {
	p.Configuration = configuration
	return &p
}

// Do executes Emulation.setTouchEmulationEnabled against the provided context and
// target handler.
func (p *SetTouchEmulationEnabledParams) Do(ctxt context.Context, h cdp.Handler) (err error) {
	if ctxt == nil {
		ctxt = context.Background()
	}

	// marshal
	buf, err := easyjson.Marshal(p)
	if err != nil {
		return err
	}

	// execute
	ch := h.Execute(ctxt, cdp.CommandEmulationSetTouchEmulationEnabled, easyjson.RawMessage(buf))

	// read response
	select {
	case res := <-ch:
		if res == nil {
			return cdp.ErrChannelClosed
		}

		switch v := res.(type) {
		case easyjson.RawMessage:
			return nil

		case error:
			return v
		}

	case <-ctxt.Done():
		return ctxt.Err()
	}

	return cdp.ErrUnknownResult
}

// SetEmulatedMediaParams emulates the given media for CSS media queries.
type SetEmulatedMediaParams struct {
	Media string `json:"media"` // Media type to emulate. Empty string disables the override.
}

// SetEmulatedMedia emulates the given media for CSS media queries.
//
// parameters:
//   media - Media type to emulate. Empty string disables the override.
func SetEmulatedMedia(media string) *SetEmulatedMediaParams {
	return &SetEmulatedMediaParams{
		Media: media,
	}
}

// Do executes Emulation.setEmulatedMedia against the provided context and
// target handler.
func (p *SetEmulatedMediaParams) Do(ctxt context.Context, h cdp.Handler) (err error) {
	if ctxt == nil {
		ctxt = context.Background()
	}

	// marshal
	buf, err := easyjson.Marshal(p)
	if err != nil {
		return err
	}

	// execute
	ch := h.Execute(ctxt, cdp.CommandEmulationSetEmulatedMedia, easyjson.RawMessage(buf))

	// read response
	select {
	case res := <-ch:
		if res == nil {
			return cdp.ErrChannelClosed
		}

		switch v := res.(type) {
		case easyjson.RawMessage:
			return nil

		case error:
			return v
		}

	case <-ctxt.Done():
		return ctxt.Err()
	}

	return cdp.ErrUnknownResult
}

// SetCPUThrottlingRateParams enables CPU throttling to emulate slow CPUs.
type SetCPUThrottlingRateParams struct {
	Rate float64 `json:"rate"` // Throttling rate as a slowdown factor (1 is no throttle, 2 is 2x slowdown, etc).
}

// SetCPUThrottlingRate enables CPU throttling to emulate slow CPUs.
//
// parameters:
//   rate - Throttling rate as a slowdown factor (1 is no throttle, 2 is 2x slowdown, etc).
func SetCPUThrottlingRate(rate float64) *SetCPUThrottlingRateParams {
	return &SetCPUThrottlingRateParams{
		Rate: rate,
	}
}

// Do executes Emulation.setCPUThrottlingRate against the provided context and
// target handler.
func (p *SetCPUThrottlingRateParams) Do(ctxt context.Context, h cdp.Handler) (err error) {
	if ctxt == nil {
		ctxt = context.Background()
	}

	// marshal
	buf, err := easyjson.Marshal(p)
	if err != nil {
		return err
	}

	// execute
	ch := h.Execute(ctxt, cdp.CommandEmulationSetCPUThrottlingRate, easyjson.RawMessage(buf))

	// read response
	select {
	case res := <-ch:
		if res == nil {
			return cdp.ErrChannelClosed
		}

		switch v := res.(type) {
		case easyjson.RawMessage:
			return nil

		case error:
			return v
		}

	case <-ctxt.Done():
		return ctxt.Err()
	}

	return cdp.ErrUnknownResult
}

// CanEmulateParams tells whether emulation is supported.
type CanEmulateParams struct{}

// CanEmulate tells whether emulation is supported.
func CanEmulate() *CanEmulateParams {
	return &CanEmulateParams{}
}

// CanEmulateReturns return values.
type CanEmulateReturns struct {
	Result bool `json:"result,omitempty"` // True if emulation is supported.
}

// Do executes Emulation.canEmulate against the provided context and
// target handler.
//
// returns:
//   result - True if emulation is supported.
func (p *CanEmulateParams) Do(ctxt context.Context, h cdp.Handler) (result bool, err error) {
	if ctxt == nil {
		ctxt = context.Background()
	}

	// execute
	ch := h.Execute(ctxt, cdp.CommandEmulationCanEmulate, cdp.Empty)

	// read response
	select {
	case res := <-ch:
		if res == nil {
			return false, cdp.ErrChannelClosed
		}

		switch v := res.(type) {
		case easyjson.RawMessage:
			// unmarshal
			var r CanEmulateReturns
			err = easyjson.Unmarshal(v, &r)
			if err != nil {
				return false, cdp.ErrInvalidResult
			}

			return r.Result, nil

		case error:
			return false, v
		}

	case <-ctxt.Done():
		return false, ctxt.Err()
	}

	return false, cdp.ErrUnknownResult
}

// SetVirtualTimePolicyParams turns on virtual time for all frames (replacing
// real-time with a synthetic time source) and sets the current virtual time
// policy. Note this supersedes any previous time budget.
type SetVirtualTimePolicyParams struct {
	Policy VirtualTimePolicy `json:"policy"`
	Budget int64             `json:"budget,omitempty"` // If set, after this many virtual milliseconds have elapsed virtual time will be paused and a virtualTimeBudgetExpired event is sent.
}

// SetVirtualTimePolicy turns on virtual time for all frames (replacing
// real-time with a synthetic time source) and sets the current virtual time
// policy. Note this supersedes any previous time budget.
//
// parameters:
//   policy
func SetVirtualTimePolicy(policy VirtualTimePolicy) *SetVirtualTimePolicyParams {
	return &SetVirtualTimePolicyParams{
		Policy: policy,
	}
}

// WithBudget if set, after this many virtual milliseconds have elapsed
// virtual time will be paused and a virtualTimeBudgetExpired event is sent.
func (p SetVirtualTimePolicyParams) WithBudget(budget int64) *SetVirtualTimePolicyParams {
	p.Budget = budget
	return &p
}

// Do executes Emulation.setVirtualTimePolicy against the provided context and
// target handler.
func (p *SetVirtualTimePolicyParams) Do(ctxt context.Context, h cdp.Handler) (err error) {
	if ctxt == nil {
		ctxt = context.Background()
	}

	// marshal
	buf, err := easyjson.Marshal(p)
	if err != nil {
		return err
	}

	// execute
	ch := h.Execute(ctxt, cdp.CommandEmulationSetVirtualTimePolicy, easyjson.RawMessage(buf))

	// read response
	select {
	case res := <-ch:
		if res == nil {
			return cdp.ErrChannelClosed
		}

		switch v := res.(type) {
		case easyjson.RawMessage:
			return nil

		case error:
			return v
		}

	case <-ctxt.Done():
		return ctxt.Err()
	}

	return cdp.ErrUnknownResult
}

// SetDefaultBackgroundColorOverrideParams sets or clears an override of the
// default background color of the frame. This override is used if the content
// does not specify one.
type SetDefaultBackgroundColorOverrideParams struct {
	Color *cdp.RGBA `json:"color,omitempty"` // RGBA of the default background color. If not specified, any existing override will be cleared.
}

// SetDefaultBackgroundColorOverride sets or clears an override of the
// default background color of the frame. This override is used if the content
// does not specify one.
//
// parameters:
func SetDefaultBackgroundColorOverride() *SetDefaultBackgroundColorOverrideParams {
	return &SetDefaultBackgroundColorOverrideParams{}
}

// WithColor rGBA of the default background color. If not specified, any
// existing override will be cleared.
func (p SetDefaultBackgroundColorOverrideParams) WithColor(color *cdp.RGBA) *SetDefaultBackgroundColorOverrideParams {
	p.Color = color
	return &p
}

// Do executes Emulation.setDefaultBackgroundColorOverride against the provided context and
// target handler.
func (p *SetDefaultBackgroundColorOverrideParams) Do(ctxt context.Context, h cdp.Handler) (err error) {
	if ctxt == nil {
		ctxt = context.Background()
	}

	// marshal
	buf, err := easyjson.Marshal(p)
	if err != nil {
		return err
	}

	// execute
	ch := h.Execute(ctxt, cdp.CommandEmulationSetDefaultBackgroundColorOverride, easyjson.RawMessage(buf))

	// read response
	select {
	case res := <-ch:
		if res == nil {
			return cdp.ErrChannelClosed
		}

		switch v := res.(type) {
		case easyjson.RawMessage:
			return nil

		case error:
			return v
		}

	case <-ctxt.Done():
		return ctxt.Err()
	}

	return cdp.ErrUnknownResult
}
