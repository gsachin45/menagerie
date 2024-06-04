package utils

import (
	"github.com/go-playground/validator"
	"github.com/sirupsen/logrus"
)

var Validate *validator.Validate

func init() {
	Validate = validator.New()
}

var Logger *logrus.Logger
