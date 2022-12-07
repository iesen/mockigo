package foo

import "github.com/iesen/mockigo/internal/fixtures/mockery/example_project/bar/foo"

type PackageNameSameAsImport interface {
	NewClient() foo.Client
}
