package domain


type Iproduct interface {
	Save(product *Products)error
	GetAll()([]Products,error)
	Delete(id string)error
	Update(id int,product *Products)error
}