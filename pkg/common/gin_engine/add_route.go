package gin_engine

// AddRoute is used for adding new route with existing method
func (engine *GinEngine) AddRoute(route Route) {
	engine.routes = append(engine.routes, route)
}
