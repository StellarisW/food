package user

type Group struct{}

func (g *Group) Sign() *SignApi {
	return &insSign
}

func (g *Group) Collect() *CollectApi {
	return &insCollect
}
