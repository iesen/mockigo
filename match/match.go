package match

import (
	"github.com/iesen/mockigo/internal/util"
	"github.com/iesen/mockigo/mock"
)

type Arg[T any] struct {
	Arg mock.Matcher
}

// ArgsToMatchers is useful in mock objects while working with variadic arguments
func ArgsToMatchers[T any](args []Arg[T]) []mock.Matcher {
	return util.MapSlice(args, func(arg Arg[T]) mock.Matcher {
		return arg.Arg
	})
}

func Eq[T any](expectedArg T) Arg[T] {
	return Arg[T]{mock.Eq(expectedArg)}
}

func Any[T any]() Arg[T] {
	return Arg[T]{mock.Any()}
}

func Not[T any](arg Arg[T]) Arg[T] {
	return Arg[T]{
		func(x interface{}) bool {
			return !arg.Arg(x)
		},
	}
}

func MatchedBy[T any](matcher func(T) bool) Arg[T] {
	return Arg[T]{
		func(x interface{}) bool {
			t, ok := x.(T)
			if !ok {
				return false
			}
			return matcher(t)
		},
	}
}
