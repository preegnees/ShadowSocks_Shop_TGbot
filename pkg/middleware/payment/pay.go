package payment

import (
	models "bot/pkg/models"
)

var _ models.IPayment = (*Payment)(nil)

type Payment struct{}

func (*Payment) GetSatus(storage models.IStorage, id models.Id) (models.PaymentStatus, error)

func (*Payment) GetPaymentLink(storage models.IStorage, id models.Id) (models.PaymentLink, error)

func (*Payment) GetHistory(storage models.IStorage, id models.Id) (models.PaymentHistory, error)