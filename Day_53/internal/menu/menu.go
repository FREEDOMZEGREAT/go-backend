package menu

type Dish struct {
	Name        string
	PriceSantim int
	IsFasting   bool
}

func Menu() []Dish {
	return []Dish{
		{Name: "Doro Wot", PriceSantim: 1500, IsFasting: false},
		{Name: "Shiro", PriceSantim: 1200, IsFasting: true},
		{Name: "Kitfo", PriceSantim: 2000, IsFasting: false},
		{Name: "Gomen", PriceSantim: 1000, IsFasting: true},
	}
}