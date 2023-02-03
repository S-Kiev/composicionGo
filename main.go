package main

import (
	"fmt"

	"github.com/S-Kiev/composicionGo/pkg/cliente"
	"github.com/S-Kiev/composicionGo/pkg/factura"
	itemfactura "github.com/S-Kiev/composicionGo/pkg/itemFactura"
)

func main() {
	f := factura.New(
		"Uruguay",
		"Salto",
		cliente.New("Ezequiel", "Calle 123", "151365"),
		[]itemfactura.Item{
			itemfactura.New(1, "Pan", 12.34),
			itemfactura.New(1, "Leche", 39.14),
			itemfactura.New(1, "Carne", 75.21),
		},
	)

	f.SetTotal()

	fmt.Printf("%+v", f)
}
