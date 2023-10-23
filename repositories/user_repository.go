package repositories

type UserRepository interface {
	Create()
	Update()
	Delete()
	FindById()
	FindAll()
}
