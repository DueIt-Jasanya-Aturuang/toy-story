package _usecase

import (
	"errors"
)

var NamePaymentExist = errors.New("nama payment sudah tersedia")
var PaymentNotExist = errors.New("payment tidak tersedia")
var TitleIconExist = errors.New("title icon sudah tersedia")
var IconNotExist = errors.New("icon tidak tersedia")
