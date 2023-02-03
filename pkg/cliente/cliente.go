package cliente

type Cliente struct {
	nombre    string
	direccion string
	celular   string
}

func New(nombre, direccion, celular string) Cliente {
	return Cliente{nombre, direccion, celular}
}
