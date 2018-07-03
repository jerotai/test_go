package lobster

type (
	// Lobster 建構子
	Lobster struct {
	}
)

func New() *Lobster {
	route := &Lobster{}
	return route
}
