// Code generated by mockigo. DO NOT EDIT.

package test

import json "encoding/json"
import match "github.com/subtle-byte/mockigo/match"
import mock "github.com/subtle-byte/mockigo/mock"

var _ = match.Any[int]

type RequesterArgSameAsNamedImport struct {
	mock *mock.Mock
}

func NewRequesterArgSameAsNamedImport(t mock.Testing) *RequesterArgSameAsNamedImport {
	t.Helper()
	return &RequesterArgSameAsNamedImport{mock: mock.NewMock(t)}
}

type _RequesterArgSameAsNamedImport_Expecter struct {
	mock *mock.Mock
}

func (_mock *RequesterArgSameAsNamedImport) EXPECT() _RequesterArgSameAsNamedImport_Expecter {
	 return _RequesterArgSameAsNamedImport_Expecter{mock: _mock.mock}
}

type _RequesterArgSameAsNamedImport_Get_Call struct {
	*mock.Call
}

func (_mock *RequesterArgSameAsNamedImport) Get(_a0 string) (*json.RawMessage) {
	_mock.mock.T.Helper()
	_results := _mock.mock.Called("Get", _a0)
	var _r0 *json.RawMessage
	if _got := _results.Get(0); _got != nil {
		_r0 = _got.(*json.RawMessage)
	}
	return _r0
}

func (_expecter _RequesterArgSameAsNamedImport_Expecter) Get(_a0 match.Arg[string]) _RequesterArgSameAsNamedImport_Get_Call {
	return _RequesterArgSameAsNamedImport_Get_Call{Call: _expecter.mock.ExpectCall("Get", _a0.Arg)}
}

func (_call _RequesterArgSameAsNamedImport_Get_Call) Return(_r0 *json.RawMessage) _RequesterArgSameAsNamedImport_Get_Call {
	_call.Call.Return(_r0)
	return _call
}

func (_call _RequesterArgSameAsNamedImport_Get_Call) RunReturn(f func(json string) (*json.RawMessage)) _RequesterArgSameAsNamedImport_Get_Call {
	_call.Call.RunReturn(f)
	return _call
}
