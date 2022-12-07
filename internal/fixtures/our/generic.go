package fixtures

import someinterface "github.com/iesen/mockigo/internal/fixtures/our/some-interface"

type GenericInterface[T any, B someinterface.SomeInterface, G ~int | float32] interface {
	SomeMethod(B) T
}

type GenericFunc[Y someinterface.SomeInterface] func() Y

type GenericComparable[T comparable] func() T
