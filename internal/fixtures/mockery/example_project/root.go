package example_project

import "github.com/iesen/mockigo/internal/fixtures/mockery/example_project/foo"

type Root interface {
	TakesBaz(*foo.Baz)
	ReturnsFoo() (foo.Foo, error)
}
