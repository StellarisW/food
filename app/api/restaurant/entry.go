package restaurant

type Group struct{}

func (g *Group) Restaurant() *Api {
	return &insRestaurant
}
