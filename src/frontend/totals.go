package main

import pb "github.com/GoogleCloudPlatform/microservices-demo/src/frontend/genproto"

const nanosPerUnit = 1000000000

// addMoney adds two amounts of one currency, carrying whole units out of the nanos.
func addMoney(l, r pb.Money) pb.Money {
	units, nanos := l.GetUnits()+r.GetUnits(), l.GetNanos()+r.GetNanos()
	units += int64(nanos / nanosPerUnit)
	nanos %= nanosPerUnit
	return pb.Money{Units: units, Nanos: nanos, CurrencyCode: l.GetCurrencyCode()}
}

// multiplyMoney adds an amount to itself n-1 times.
func multiplyMoney(m pb.Money, n uint32) pb.Money {
	out := m
	for ; n > 1; n-- {
		out = addMoney(out, m)
	}
	return out
}
