package mongo

import (
	models "bot/pkg/models"
)

type Mongo struct{}

var _ models.IStorage = (*Mongo)(nil)

func (*Mongo) GetUser(id models.Id) (models.User, error)

func (*Mongo) SaveUser(user models.User) error

func (*Mongo) SaveTransaction(label models.Label) error

func (*Mongo) GetAllTransactions(id models.Id) ([]models.Transaction, error)

func (*Mongo) GetLastTransaction(id models.Id) (models.Transaction, error)

func (*Mongo) ConfirmTransaction(label models.Label) error
