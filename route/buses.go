package route

type BusRouteler interface {
	checkStop()
}

type BusRoute struct {
	name string
	startLocation string
	endLocation string
	locations []Place
}

//Checking if destination is in the bus route
func (bus BusRoute) checkStop(destination string) bool {
	for _, place := range bus.locations {
		if place.name == destination {
			return true
		}
	}
	return false
}
