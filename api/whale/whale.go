package whale

type (
	// Whale 建構子
	Whale struct {
	}
)

func New() *Whale {
	whaleRoute := &Whale{}
	return whaleRoute
}
