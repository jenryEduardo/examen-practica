package infraestructure


import (
	"examen/serverUno/domain"
	"examen/core"
)

type MySQLRepository struct {
	conn *core.Conn_MySQL
}

func NewMySQLRepository() *MySQLRepository {
	conn := core.GetDBPool()
	return &MySQLRepository{conn: conn}
}

func (r *MySQLRepository) Save(p *domain.Products) error {
	query := "INSERT INTO Products (nombre, precio,cantidad,codigoBarra) VALUES (?,?,?,?)"
	_, err := r.conn.DB.Exec(query, p.Nombre, p.Precio,p.Cantidad,p.CodigoBarra)
	return err
}

func (r *MySQLRepository) Delete(p string)error{
	nombre :=p
	query := "DELETE FROM Products WHERE nombre = ?"
	_,err :=r.conn.DB.Exec(query,nombre)
	return err
}

func (r *MySQLRepository) Update(id int,p *domain.Products)error{
	query := "UPDATE Products SET nombre = ?, precio = ? WHERE idProduct = ?"
    _, err := r.conn.DB.Exec(query, p.Nombre, p.Precio,id)
    if err != nil {
        return err
    }
	return err
}

func (r *MySQLRepository) GetAll() ([]domain.Products, error) {
	query := "SELECT nombre, precio, cantidad, codigoBarra FROM Products"
	rows, err := r.conn.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []domain.Products
	for rows.Next() {
		var product domain.Products
		if err := rows.Scan(&product.Nombre, &product.Precio, &product.Cantidad, &product.CodigoBarra); err != nil {
			return nil, err
		}
		
		products = append(products, product)
	}
	return products, nil
}