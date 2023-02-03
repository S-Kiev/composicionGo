package itemfactura

type Item struct {
	id       uint
	producto string
	valor    float64
}

func New(id uint, producto string, valor float64) Item {
	return Item{id, producto, valor}
}

// Gettter del valor del item
func (i Item) Valor() float64 {
	return i.valor
}
