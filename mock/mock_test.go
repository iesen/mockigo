package mock

import (
	"os"
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
)

func Example_simple() {
	runWithT(func(t *t) {
		m := newMockWithT(t)
		m.ExpectCall("Foo", Eq("hello"))
		m.Called("Foo", "hello")
		m.Called("Foo", "bye")
	})
	//Output:
	// runWithT:4: Unexpected call of method "Foo" because:
	// expected call runWithT:2 doesn't match the argument "bye" at index 0
}

func TestMock_withReturn(t *testing.T) {
	m := NewMock(t)
	m.ExpectCall("Foo", Eq("hello")).Return(43)
	ret := m.Called("Foo", "hello")
	ret0 := ret.Get(0).(int)
	require.Equal(t, 43, ret0)
}

func TestMock_withRunReturn(t *testing.T) {
	m := NewMock(t)
	m.ExpectCall("Foo", Eq("hello")).RunReturn(func(s string) int {
		return 43
	})
	ret := m.Called("Foo", "hello")
	ret0 := ret.Get(0).(int)
	require.Equal(t, 43, ret0)
}

func Example_withTimes() {
	runWithT(func(t *t) {
		m := newMockWithT(t)
		m.ExpectCall("BarBar").Times(0, 1)
		m.ExpectCall("Bar").Times(1, -1)
		m.Called("Bar")
		m.Called("Bar")
		m.Called("Bar")
		m.ExpectCall("Foo", Eq("hello")).Times(1, -100)
		m.Called("Foo", "hello")
		m.Called("Foo", "hello")
	})
	//Output:
	// runWithT:9: Unexpected call of method "Foo" because:
	// expected call runWithT:7 has already been called the max number of times
}

func Example_withAfter() {
	runWithT(func(t *t) {
		m := newMockWithT(t)
		fooCall := m.ExpectCall("Foo")
		m.ExpectCall("Bar").After(1, -1, fooCall)
		m.Called("Bar")
	})
	//Output:
	// runWithT:4: Unexpected call of method "Bar" because:
	// expected call runWithT:3 should be called after call runWithT:2
}

func Example_withInOrder() {
	runWithT(func(t *t) {
		m := newMockWithT(t)
		InOrder(1, -1,
			m.ExpectCall("Foo"),
			m.ExpectCall("Bar"),
		)
		m.Called("Bar")
	})
	//Output:
	// runWithT:6: Unexpected call of method "Bar" because:
	// expected call runWithT:4 should be called after call runWithT:3
}

func TestMock_withMatchers(t *testing.T) {
	m := NewMock(t)
	m.ExpectCall("Foo", Any(), Eq("hello"), Not(Eq(0)))
	m.Called("Foo", 100, "hello", 1)
}

func TestMock_withMany(t *testing.T) {
	m := NewMock(t)
	InOrder(1, 1,
		m.ExpectCall("Foo", Any(), Eq("hello")).Return(45),
		m.ExpectCall("Bar", Any()).RunReturn(func(n int) string {
			return strconv.Itoa(n)
		}),
	)
	fooRet := m.Called("Foo", 100, "hello").Get(0).(int) // == 45
	barRet := m.Called("Bar", 200).Get(0).(string)       // == "200"

	require.Equal(t, 45, fooRet)
	require.Equal(t, "200", barRet)
}

func Example_unusedMethods() {
	runWithT(func(t *t) {
		m := newMockWithT(t)
		m.ExpectCall("Foo")
	})
	//Output:
	// runWithT:1: clean up phase: missing call(s) to expected call runWithT:2
}

func Example_noReturn() {
	runWithT(func(t *t) {
		m := newMockWithT(t)
		m.ExpectCall("Foo")
		rets := m.Called("Foo")
		rets.Get(1)
	})
	//Output:
	// runWithT:4: Call runWithT:2 does not have return value at index 1
}

func Example_errReturn() {
	runWithT(func(t *t) {
		m := newMockWithT(t)
		m.ExpectCall("Foo").Return(os.ErrClosed, nil, 7)
		rets := m.Called("Foo")
		e := rets.Error(0)
		require.Equal(t, os.ErrClosed, e)
		e = rets.Error(1)
		require.True(t, e == nil)
		rets.Error(2)
	})
	//Output:
	// runWithT:8: Call runWithT:2 does not have return value of the error type at index 2
}

func Test_unknownFile(t *testing.T) {
	m := NewMock(t)
	m.getCaller = func(skip int) (pc uintptr, file string, line int, ok bool) {
		return 0, "", 0, false
	}
	info := m.callerInfo(1)
	require.Equal(t, "unknown file", info)
}

func Example_wrongNumberOfArguments() {
	runWithT(func(t *t) {
		m := newMockWithT(t)
		m.ExpectCall("Foo")
		m.Called("Foo", 3, 4)
	})
	//Output:
	// runWithT:3: Unexpected call of method "Foo" because:
	// expected call runWithT:2 expects 0 arguments, got 2
}

func Example_wrongNumberOfRunReturnArguments() {
	runWithT(func(t *t) {
		m := newMockWithT(t)
		m.ExpectCall("Foo").RunReturn(func(int) {})
	})
	//Output:
	// runWithT:2: Wrong number of arguments in RunReturn func for call runWithT:2: got 0, want 1
}

func Example_unexpectedMethod() {
	runWithT(func(t *t) {
		m := newMockWithT(t)
		m.Called("Foo")
	})
	//Output:
	// runWithT:2: Unexpected call of method "Foo" because: there are no expected calls of the method "Foo"
}
