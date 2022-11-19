package recipe

type Group struct{}

func (g *Group) Recipe() *Api {
	return &insRecipe
}
