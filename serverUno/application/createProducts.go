package application

import "examen/serverUno/domain"

type CreateProduct struct {
	repo domain.Iproduct
}

func NewCreateProduct(repo domain.Iproduct)*CreateProduct{
	return &CreateProduct{repo: repo}
}

func (cp *CreateProduct) Execute(product domain.Products)error{
	return cp.repo.Save(&product)
}