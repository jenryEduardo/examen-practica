package application

import "examen/serverUno/domain"


type GetProduct struct {
	repo domain.Iproduct
}	

//nos sirve para cuando sea invocado en infraestructure (controllers) creee el caso de uso 
//en pocas palabras como decirle a infra lo que puede hacer
func NewGetProduct(repo domain.Iproduct) *GetProduct{
		return &GetProduct{repo: repo}
}
//metodo que sirve para llamar a un metodo que realice una operacion desde x lugar 
//ya que no depende de nadie
func (cp *GetProduct) Execute() ([]domain.Products, error) {
	return cp.repo.GetAll()
}
