package application

import "examen/serverUno/domain"


type GetProduct struct {
	repo domain.Iproduct
}	


func NewGetProduct(repo domain.Iproduct) *GetProduct{
		return &GetProduct{repo: repo}
}

func (cp *GetProduct) Execute() ([]domain.Products, error) {
	return cp.repo.GetAll()
}
