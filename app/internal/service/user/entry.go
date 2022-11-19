package user

type Group struct{}

func (g *Group) User() *sUser {
	return &insUser
}

func (g *Group) Collect() *sCollect {
	return &insCollect
}
