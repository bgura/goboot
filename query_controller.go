package goboot

// An extended controller with the ability to read query parmeters
type QueryController struct {
	Controller
	ParamReader paramReader
}
