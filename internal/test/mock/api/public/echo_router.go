// Code generated by mockery v2.16.0. DO NOT EDIT.

package public

import (
	echo "github.com/labstack/echo/v4"
	mock "github.com/stretchr/testify/mock"
)

// EchoRouter is an autogenerated mock type for the EchoRouter type
type EchoRouter struct {
	mock.Mock
}

// CONNECT provides a mock function with given fields: path, h, m
func (_m *EchoRouter) CONNECT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route {
	_va := make([]interface{}, len(m))
	for _i := range m {
		_va[_i] = m[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, path, h)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *echo.Route
	if rf, ok := ret.Get(0).(func(string, echo.HandlerFunc, ...echo.MiddlewareFunc) *echo.Route); ok {
		r0 = rf(path, h, m...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*echo.Route)
		}
	}

	return r0
}

// DELETE provides a mock function with given fields: path, h, m
func (_m *EchoRouter) DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route {
	_va := make([]interface{}, len(m))
	for _i := range m {
		_va[_i] = m[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, path, h)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *echo.Route
	if rf, ok := ret.Get(0).(func(string, echo.HandlerFunc, ...echo.MiddlewareFunc) *echo.Route); ok {
		r0 = rf(path, h, m...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*echo.Route)
		}
	}

	return r0
}

// GET provides a mock function with given fields: path, h, m
func (_m *EchoRouter) GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route {
	_va := make([]interface{}, len(m))
	for _i := range m {
		_va[_i] = m[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, path, h)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *echo.Route
	if rf, ok := ret.Get(0).(func(string, echo.HandlerFunc, ...echo.MiddlewareFunc) *echo.Route); ok {
		r0 = rf(path, h, m...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*echo.Route)
		}
	}

	return r0
}

// HEAD provides a mock function with given fields: path, h, m
func (_m *EchoRouter) HEAD(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route {
	_va := make([]interface{}, len(m))
	for _i := range m {
		_va[_i] = m[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, path, h)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *echo.Route
	if rf, ok := ret.Get(0).(func(string, echo.HandlerFunc, ...echo.MiddlewareFunc) *echo.Route); ok {
		r0 = rf(path, h, m...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*echo.Route)
		}
	}

	return r0
}

// OPTIONS provides a mock function with given fields: path, h, m
func (_m *EchoRouter) OPTIONS(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route {
	_va := make([]interface{}, len(m))
	for _i := range m {
		_va[_i] = m[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, path, h)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *echo.Route
	if rf, ok := ret.Get(0).(func(string, echo.HandlerFunc, ...echo.MiddlewareFunc) *echo.Route); ok {
		r0 = rf(path, h, m...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*echo.Route)
		}
	}

	return r0
}

// PATCH provides a mock function with given fields: path, h, m
func (_m *EchoRouter) PATCH(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route {
	_va := make([]interface{}, len(m))
	for _i := range m {
		_va[_i] = m[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, path, h)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *echo.Route
	if rf, ok := ret.Get(0).(func(string, echo.HandlerFunc, ...echo.MiddlewareFunc) *echo.Route); ok {
		r0 = rf(path, h, m...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*echo.Route)
		}
	}

	return r0
}

// POST provides a mock function with given fields: path, h, m
func (_m *EchoRouter) POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route {
	_va := make([]interface{}, len(m))
	for _i := range m {
		_va[_i] = m[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, path, h)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *echo.Route
	if rf, ok := ret.Get(0).(func(string, echo.HandlerFunc, ...echo.MiddlewareFunc) *echo.Route); ok {
		r0 = rf(path, h, m...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*echo.Route)
		}
	}

	return r0
}

// PUT provides a mock function with given fields: path, h, m
func (_m *EchoRouter) PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route {
	_va := make([]interface{}, len(m))
	for _i := range m {
		_va[_i] = m[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, path, h)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *echo.Route
	if rf, ok := ret.Get(0).(func(string, echo.HandlerFunc, ...echo.MiddlewareFunc) *echo.Route); ok {
		r0 = rf(path, h, m...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*echo.Route)
		}
	}

	return r0
}

// TRACE provides a mock function with given fields: path, h, m
func (_m *EchoRouter) TRACE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route {
	_va := make([]interface{}, len(m))
	for _i := range m {
		_va[_i] = m[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, path, h)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *echo.Route
	if rf, ok := ret.Get(0).(func(string, echo.HandlerFunc, ...echo.MiddlewareFunc) *echo.Route); ok {
		r0 = rf(path, h, m...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*echo.Route)
		}
	}

	return r0
}

type mockConstructorTestingTNewEchoRouter interface {
	mock.TestingT
	Cleanup(func())
}

// NewEchoRouter creates a new instance of EchoRouter. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewEchoRouter(t mockConstructorTestingTNewEchoRouter) *EchoRouter {
	mock := &EchoRouter{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
