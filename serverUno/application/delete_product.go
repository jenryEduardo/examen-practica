package application

import "examen/serverUno/domain"

type DeleteProduct struct {
	rep domain.Iproduct
}

func NewDeleteProduct(rep domain.Iproduct) *DeleteProduct{
	return &DeleteProduct{
		rep: rep,
	}
}

func (cp *DeleteProduct) Execute(nombre string)error {
	return cp.rep.Delete(nombre)
}