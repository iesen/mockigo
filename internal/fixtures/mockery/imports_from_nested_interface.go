package mockery

import (
	"github.com/iesen/mockigo/internal/fixtures/mockery/http"
)

type HasConflictingNestedImports interface {
	RequesterNS
	Z() http.MyStruct
}
