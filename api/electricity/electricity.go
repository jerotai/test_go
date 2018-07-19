package electricity

type (
	// electricity 建構子
	Electricity struct {
	}
)

func New() *Electricity {
	route := &Electricity{}
	return route
}

