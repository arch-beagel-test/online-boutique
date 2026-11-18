package main

import (
	pb "github.com/GoogleCloudPlatform/microservices-demo/src/frontend/genproto"
	"github.com/GoogleCloudPlatform/microservices-demo/src/frontend/money"
)

// addMoney adds two amounts of one currency.
func addMoney(l, r pb.Money) pb.Money { return money.Must(money.Sum(l, r)) }

// multiplyMoney multiplies an amount by a quantity.
func multiplyMoney(m pb.Money, n uint32) pb.Money { return money.MultiplySlow(m, n) }
