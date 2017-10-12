package goboot

// Interface defines an endpoint in the API in which
// a user can connect and make queries. All models which
// are modifiable must export a list of routes
type Api interface {
	GetRoutes() []Route
}
