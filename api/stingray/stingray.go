package stingray

type (
	// Stingray 建構子
	Stingray struct {
		//whitebaitSerivce
		WhitebaitSerivce struct{
		}
	}
)

func New() *Stingray {
	stingrayRoute := &Stingray{}
	return stingrayRoute
}
