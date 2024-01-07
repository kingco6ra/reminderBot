package repos

type Object interface{}

type RepositoryInterface interface {
	Insert(interface{}) error
	Select() ([]interface{}, error)
	SelectByID(id uint64) (interface{}, error)
	Update(interface{}) error
	Delete(interface{}) error
}
