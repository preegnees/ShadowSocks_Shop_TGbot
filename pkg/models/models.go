package models

// Id
type Id string
// Label. Строка, которая генерируется при выставления щета, чтобы можно было потом узнать был ли платеж
type Label string
// Username
type Username string
// TimeCreated. Время оплаты
type TimeCreated int
// Amount. Сумма оплаты
type Amount int

// User. Описание пользователя
type User struct {
	Id
	Username
}

// Transaction. Описание транзакции
type Transaction struct {
	Label
	Id
	TimeCreated
	Amount
}

// IStorage. Интерфейс для постоянного хранилища
type IStorage interface {
	GetUser(Id) (User, error)
	SaveUser(User) error

	SaveTransaction(Label) error
	GetAllTransactions(Id) ([]Transaction, error)
	GetLastTransaction(Id) (Transaction, error)
	ConfirmTransaction(Label) error
}

// IMemStorage. Интерфейс для временного хранилища
type IMemStorage interface {
	SaveLabel(Label, Id) error
	GetLabel(Id) (Label, error)
	ConfirmTransaction(Label) error
}