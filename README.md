
mockigo
=======
[![go.dev reference](https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white&style=flat-square)](https://pkg.go.dev/github.com/subtle-byte/mockigo?tab=overview)

`mockigo` provides the ability to easily generate mocks for golang interfaces.

Caveat: `mockigo` uses Go 1.18 generics in the generated mocks.

Table of Contents
-----------------

- [Installation](#installation)
- [Usage in command line](#usage-in-command-line)
- [Generated mocks usage](#generated-mocks-usage)
- [Comparison with `mockery` and `gomock`](#comparison-with-mockery-and-gomock)


Installation
------------

You can use the `go install`:

```
go install github.com/subtle-byte/mockigo/cmd/mockigo@latest
```

Usage in command line
------------

You can configurate either using command line flags or using `mockigo.yaml` file. Command line flags has priority over `mockigo.yaml`.

| flag | `mockigo.yaml` key | default value |description  |
|---|---|---|---|
| `--root-dir` | `root-dir` | `internal` | Directory from which the recursive search for go files begins. |
| `--mocks-dir` | `mocks-dir` | `internal/mocks` | Directory where are the generated mocks saved. |
| `--walk` | `walk` | `""` | Rules of walking in `root-dir` finding interfaces. |

`mockigo` will walk recursively in the `root-dir`. You can configure to avoid some dirs inside `root-dir` using `walk` setting which works similar to `.gitignore` file (but `!` _excludes_ and `@` let to specify interfaces).

For example you can create `mockigo.yaml`
```yaml
root-dir: pkg
mocks-dir: pkg/mocks
walk:
  - "!." # exclude everything in `pkg`
  - "service" # includes dir `pkg/service` back
  - "server@SomeInterface1,SomeInterface2" # includes interfaces SomeInterface1 and SomeInterface2 in `pkg/server` back
```
and run `mockigo`.

Alternatively you can do the same without `mockigo.yaml` running command `mockigo --root-dir=pkg --mocks-dir=pkg/mocks --walk="!.;service;server@SomeInterface1,SomeInterface2`.


Generated mocks usage
------------

For example you had the file `internal/service/service.go`:

```go
package service

type FooBar interface {
	Foo(n int) (string, error)
	Bar()
}

type MyFunc func(a, b int) int
```

If you run `mockigo` you will get mocks in the dir `mocks`.

Then you can create file `internal/service/service_test.go` and use the generated mocks:

```go
package service

import (
	"testing"

	"github.com/stretchr/testify/require"
	service_mocks "github.com/someuser/somerepo/internal/mocks/service"
	"github.com/subtle-byte/mockigo/match"
)

func TestMocks(t *testing.T) {
	fooBar := service_mocks.NewFooBar(t)
	fooBar.EXPECT().Foo(match.Eq(7)).Return("7", nil)
	r1, err := fooBar.Foo(7)
	require.Equal(t, "7", r1)
	require.NoError(t, err)

	myFunc := service_mocks.NewMyFunc(t)
	myFunc.EXPECT().Execute(match.Any[int](), match.Any[int]()).Return(10)
	r2 := myFunc.Execute(1, 2)
	require.Equal(t, 10, r2)
}
```

`mathc.Eq` and `match.Any` check the arguments of mock method when it is called. Everything in this testing code is typechecked in compile time using Go 1.18 generics.

More powerful usage (also everything is typechecked in compile time):

```go
package service

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
	service_mocks "github.com/someuser/somerepo/internal/mocks/service"
	"github.com/subtle-byte/mockigo/match"
	"github.com/subtle-byte/mockigo/mock"
)

func TestMocks(t *testing.T) {
	fooBar := service_mocks.NewFooBar(t)
	mock.InOrder( // Foo can be called only after at least one call of Bar
		fooBar.EXPECT().Bar(),
		fooBar.EXPECT().Foo(match.MatchedBy(func(n int) bool {
			return n > 0
		})).RunReturn(func(n int) (string, error) {
			return strconv.Itoa(n), nil
		}),
	)
	fooBar.Bar()
	r1, err := fooBar.Foo(9)
	require.Equal(t, "9", r1)
	require.NoError(t, err)
}
```

Comparison with `mockery` and `gomock`
------------

Neither `mockery` nor `gomock` generate type-safe mocks.

For example with `gomock` the following code (which generated using `FooBar` interface defined above) is valid during compilation, but it will fail in runtime:

```go
fooBar.EXPECT().Foo("not int").Return([]byte("not string"), "not error", "not allowed third return")
```

Also with `mockery` the following code is valid during compilation, but it will fail in runtime:

```go
fooBar.EXPECT().Foo("not int").Call.Return(func(notInt string) {
	// returns nothing instead of 2 expected returns
})
```

Another benefit of `mockigo` - it is faster. For example lets generate mocks for ![CoreDNS](https://github.com/coredns/coredns) project using commands `/usr/bin/time --verbose mockigo --root-dir . --mocks-dir mocks` and `/usr/bin/time --verbose mockery --all --dir . --output mocks2 --keeptree --with-expecter`. The results:

| |Time of generation|Max RAM usage|
|---|---|---|
|mockery|34.29 sec|483 MB|
|mockigo|0.76 sec|105 MB|

Also `mockigo` allows to recursively walk the dirs with flexible rules.
