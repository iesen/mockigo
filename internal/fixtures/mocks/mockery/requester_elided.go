// Code generated by mockigo. DO NOT EDIT.

package test

import match "github.com/subtle-byte/mockigo/match"
import mock "github.com/subtle-byte/mockigo/mock"

var _ = match.Any[int]

type RequesterElided struct {
	mock *mock.Mock
}

func NewRequesterElided(t mock.Testing) *RequesterElided {
	t.Helper()
	return &RequesterElided{mock: mock.NewMock(t)}
}

type _RequesterElided_Expecter struct {
	mock *mock.Mock
}

func (_mock *RequesterElided) EXPECT() _RequesterElided_Expecter {
	 return _RequesterElided_Expecter{mock: _mock.mock}
}

type _RequesterElided_Get_Call struct {
	*mock.Call
}

func (_mock *RequesterElided) Get(path string, url string) (error) {
	_mock.mock.T.Helper()
	_results := _mock.mock.Called("Get", path, url)
	_r0 := _results.Error(0)
	return _r0
}

func (_expecter _RequesterElided_Expecter) Get(path match.Arg[string], url match.Arg[string]) _RequesterElided_Get_Call {
	return _RequesterElided_Get_Call{Call: _expecter.mock.ExpectCall("Get", path.Arg, url.Arg)}
}

func (_call _RequesterElided_Get_Call) Return(_r0 error) _RequesterElided_Get_Call {
	_call.Call.Return(_r0)
	return _call
}

func (_call _RequesterElided_Get_Call) RunReturn(f func(path string, url string) (error)) _RequesterElided_Get_Call {
	_call.Call.RunReturn(f)
	return _call
}