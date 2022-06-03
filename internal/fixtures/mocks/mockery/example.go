// Code generated by mockigo. DO NOT EDIT.

package test

import match "github.com/subtle-byte/mockigo/match"
import mock "github.com/subtle-byte/mockigo/mock"
import mockery_http "github.com/subtle-byte/mockigo/internal/fixtures/mockery/http"
import net_http "net/http"

var _ = match.Any[int]

type Example struct {
	mock *mock.Mock
}

func NewExample(t mock.Testing) *Example {
	t.Helper()
	return &Example{mock: mock.NewMock(t)}
}

type _Example_Expecter struct {
	mock *mock.Mock
}

func (_mock *Example) EXPECT() _Example_Expecter {
	 return _Example_Expecter{mock: _mock.mock}
}

type _Example_A_Call struct {
	*mock.Call
}

func (_mock *Example) A() (net_http.Flusher) {
	_mock.mock.T.Helper()
	_results := _mock.mock.Called("A", )
	var _r0 net_http.Flusher
	if _got := _results.Get(0); _got != nil {
		_r0 = _got.(net_http.Flusher)
	}
	return _r0
}

func (_expecter _Example_Expecter) A() _Example_A_Call {
	return _Example_A_Call{Call: _expecter.mock.ExpectCall("A", )}
}

func (_call _Example_A_Call) Return(_r0 net_http.Flusher) _Example_A_Call {
	_call.Call.Return(_r0)
	return _call
}

func (_call _Example_A_Call) RunReturn(f func() (net_http.Flusher)) _Example_A_Call {
	_call.Call.RunReturn(f)
	return _call
}

type _Example_B_Call struct {
	*mock.Call
}

func (_mock *Example) B(fixtureshttp string) (mockery_http.MyStruct) {
	_mock.mock.T.Helper()
	_results := _mock.mock.Called("B", fixtureshttp)
	_r0 := _results.Get(0).(mockery_http.MyStruct)
	return _r0
}

func (_expecter _Example_Expecter) B(fixtureshttp match.Arg[string]) _Example_B_Call {
	return _Example_B_Call{Call: _expecter.mock.ExpectCall("B", fixtureshttp.Arg)}
}

func (_call _Example_B_Call) Return(_r0 mockery_http.MyStruct) _Example_B_Call {
	_call.Call.Return(_r0)
	return _call
}

func (_call _Example_B_Call) RunReturn(f func(fixtureshttp string) (mockery_http.MyStruct)) _Example_B_Call {
	_call.Call.RunReturn(f)
	return _call
}