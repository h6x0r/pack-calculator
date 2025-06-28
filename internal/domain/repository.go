package domain

type PackRepository interface {
	List() ([]Pack, error)
	Create(size int) (Pack, error)
	Delete(size int) error
	Update(oldSize, newSize int) error
}

type OrderRepository interface {
	Save(order *Order) error
}
