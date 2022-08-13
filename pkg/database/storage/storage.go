package storage

import (
	storage "bot/pkg/models"
)



func GetStorage() storage.IStorage

func GetMemStorage() storage.IMemStorage