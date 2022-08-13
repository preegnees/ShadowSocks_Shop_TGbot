package redis

import (
	models "bot/pkg/models"
)

type Redis struct {} 

var _ models.IMemStorage = (*Redis)(nil)

func (*Redis) SaveLabel(label models.Label, id models.Id) error

func (*Redis) GetLabel(id models.Id) (models.Label, error)

func (*Redis) ConfirmTransaction(label models.Label) error