package route

type BusRouteler interface {
	checkStop()
}

type BusRoute struct {
	name string
	startLocation string
	endLocation string
	locations []string
}

var busAStops = []string{"Mall Terminal\n", "Main Street\n", "Second Street\n", "City Hall\n", "Union Terminal\n"}
var busA = &BusRoute{"Temp", "Mall Terminal", "Union Terminal", busAStops}

//Checking if destination is in the bus route
func (bus BusRoute) checkStop(destination string) bool {
	for _, place := range bus.locations {
		if place == destination {
			return true
		}
	}
	return false
}
