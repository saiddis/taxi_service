package types

type NewRider struct {
	Username string
}

type FindDrivers struct {
	Lat   float64
	Lng   float64
	Range float64
}
