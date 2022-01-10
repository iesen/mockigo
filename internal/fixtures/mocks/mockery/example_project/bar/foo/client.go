// Code generated by mockigo. DO NOT EDIT.

package foo

import match "github.com/subtle-byte/mockigo/match"
import mock "github.com/subtle-byte/mockigo/mock"

var _ = match.Any[int]

type Client struct {
	mock *mock.Mock
}

func NewClient(t mock.Testing) *Client {
	t.Helper()
	return &Client{mock: mock.NewMock(t)}
}

type _Client_Expecter struct {
	mock *mock.Mock
}

func (_mock *Client) EXPECT() _Client_Expecter {
	 return _Client_Expecter{mock: _mock.mock}
}

type _Client_Search_Call struct {
	*mock.Call
}

func (_mock *Client) Search(query string) ([]string, error) {
	_mock.mock.T.Helper()
	_results := _mock.mock.Called("Search", query)
	var _r0 []string
	if _got := _results.Get(0); _got != nil {
		_r0 = _got.([]string)
	}
	_r1 := _results.Error(1)
	return _r0, _r1
}

func (_expecter _Client_Expecter) Search(query match.Arg[string]) _Client_Search_Call {
	return _Client_Search_Call{Call: _expecter.mock.ExpectCall("Search", query.Arg)}
}

func (_call _Client_Search_Call) Return(_r0 []string, _r1 error) _Client_Search_Call {
	_call.Call.Return(_r0, _r1)
	return _call
}

func (_call _Client_Search_Call) RunReturn(f func(query string) ([]string, error)) _Client_Search_Call {
	_call.Call.RunReturn(f)
	return _call
}
