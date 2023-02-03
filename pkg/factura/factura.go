package factura

import (
	"github.com/S-Kiev/composicionGo/pkg/cliente"
)

type Factura struct {
	pais    string
	ciudad  string
	total   float64
	cliente cliente.Cliente
	//En este slice estoy usando la composicion de Go
	//porque la catura se compone de muchos items
	items []itemfactura.ItemFactura
}

func New(pais, ciudad string, cliente cliente.Cliente, items []itemfactura.ItemFactura) Factura {
	return Factura{
		pais:    pais,
		ciudad:  ciudad,
		cliente: cliente,
		items:   items,
	}
}

// Setter del total de la factura
// para actualizar valores se usa puntero
func (i *Factura) SetTotal() {
	for _, item := range i.items {
		i.total += item.Valor()
	}
}
