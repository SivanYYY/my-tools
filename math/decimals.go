package main

import (
	"errors"
	"github.com/shopspring/decimal"
	"math/big"
)

func converNumber(n any) (num decimal.Decimal, err error) {
	switch n.(type) {
	case string:
		num, err = decimal.NewFromString(n.(string))
		if err != nil {
			return num, err
		}
	case big.Int:
		num = decimal.NewFromBigInt(n.(*big.Int), 0)
	case int8:
		num = decimal.NewFromInt(int64(n.(int8)))
	case int32:
		num = decimal.NewFromInt(int64(n.(int32)))
	case int64:
		num = decimal.NewFromInt(n.(int64))
	case int:
		num = decimal.NewFromInt(int64(n.(int)))
	default:
		return num, errors.New("not match")
	}
	return
}

func Add(n1, n2 any) (string, error) {
	num1, err := converNumber(n1)
	if err != nil {
		return "", err
	}
	num2, err := converNumber(n2)
	if err != nil {
		return "", err
	}
	return num1.Add(num2).String(), nil
}

func Sub(n1, n2 any) (string, error) {
	num1, err := converNumber(n1)
	if err != nil {
		return "", err
	}
	num2, err := converNumber(n2)
	if err != nil {
		return "", err
	}
	return num1.Sub(num2).String(), nil
}

func Mul(n1, n2 any) (string, error) {
	num1, err := converNumber(n1)
	if err != nil {
		return "", err
	}
	num2, err := converNumber(n2)
	if err != nil {
		return "", err
	}
	return num1.Mul(num2).String(), nil
}

func Div(n1, n2 any) (string, error) {
	num1, err := converNumber(n1)
	if err != nil {
		return "", err
	}
	num2, err := converNumber(n2)
	if err != nil {
		return "", err
	}
	return num1.Div(num2).String(), nil
}
