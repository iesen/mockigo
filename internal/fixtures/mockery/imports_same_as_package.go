package mockery

import (
	"github.com/iesen/mockigo/internal/fixtures/mockery/mockery"
)

type C int

type ImportsSameAsPackage interface {
	A() mockery.B
	B() KeyManager
	C(C)
}
