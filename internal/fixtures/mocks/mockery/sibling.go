// Code generated by mockigo. DO NOT EDIT.

package test

import match "github.com/subtle-byte/mockigo/match"
import mock "github.com/subtle-byte/mockigo/mock"

var _ = match.Any[int]

type Sibling struct {
	mock *mock.Mock
}

func NewSibling(t mock.Testing) *Sibling {
	t.Helper()
	return &Sibling{mock: mock.NewMock(t)}
}

type _Sibling_Expecter struct {
	mock *mock.Mock
}

func (_mock *Sibling) EXPECT() _Sibling_Expecter {
	 return _Sibling_Expecter{mock: _mock.mock}
}

type _Sibling_DoSomething_Call struct {
	*mock.Call
}

func (_mock *Sibling) DoSomething() () {
	_mock.mock.T.Helper()
	_mock.mock.Called("DoSomething", )
}

func (_expecter _Sibling_Expecter) DoSomething() _Sibling_DoSomething_Call {
	return _Sibling_DoSomething_Call{Call: _expecter.mock.ExpectCall("DoSomething", )}
}

func (_call _Sibling_DoSomething_Call) Return() _Sibling_DoSomething_Call {
	_call.Call.Return()
	return _call
}

func (_call _Sibling_DoSomething_Call) RunReturn(f func() ()) _Sibling_DoSomething_Call {
	_call.Call.RunReturn(f)
	return _call
}