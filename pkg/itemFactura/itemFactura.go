package itemfactura

type ItemFactura struct {
	id       uint
	producto string
	valor    float64
}

func New(id uint, producto string, valor float64) ItemFactura {
	return ItemFactura{id, producto, valor}
}

// Gettter del valor del item
func (i ItemFactura) Valor() float64 {
	return i.valor
}
