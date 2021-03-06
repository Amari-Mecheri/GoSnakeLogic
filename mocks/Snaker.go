// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import common "github.com/Amari-Mecheri/GoSnakeLogic/pkg/common"
import mock "github.com/stretchr/testify/mock"

// Snaker is an autogenerated mock type for the Snaker type
type Snaker struct {
	mock.Mock
}

// Direction provides a mock function with given fields:
func (_m *Snaker) Direction() (common.Direction, error) {
	ret := _m.Called()

	var r0 common.Direction
	if rf, ok := ret.Get(0).(func() common.Direction); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(common.Direction)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GrowTo provides a mock function with given fields: newPosition
func (_m *Snaker) GrowTo(newPosition common.Position) error {
	ret := _m.Called(newPosition)

	var r0 error
	if rf, ok := ret.Get(0).(func(common.Position) error); ok {
		r0 = rf(newPosition)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MoveTo provides a mock function with given fields: newPosition
func (_m *Snaker) MoveTo(newPosition common.Position) (common.Position, error) {
	ret := _m.Called(newPosition)

	var r0 common.Position
	if rf, ok := ret.Get(0).(func(common.Position) common.Position); ok {
		r0 = rf(newPosition)
	} else {
		r0 = ret.Get(0).(common.Position)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(common.Position) error); ok {
		r1 = rf(newPosition)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NextMove provides a mock function with given fields:
func (_m *Snaker) NextMove() (common.Position, error) {
	ret := _m.Called()

	var r0 common.Position
	if rf, ok := ret.Get(0).(func() common.Position); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(common.Position)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Position provides a mock function with given fields:
func (_m *Snaker) Position() (common.Position, error) {
	ret := _m.Called()

	var r0 common.Position
	if rf, ok := ret.Get(0).(func() common.Position); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(common.Position)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SetDirection provides a mock function with given fields: direction
func (_m *Snaker) SetDirection(direction common.Direction) {
	_m.Called(direction)
}

// Size provides a mock function with given fields:
func (_m *Snaker) Size() (int, error) {
	ret := _m.Called()

	var r0 int
	if rf, ok := ret.Get(0).(func() int); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Tail provides a mock function with given fields:
func (_m *Snaker) Tail() (common.Position, error) {
	ret := _m.Called()

	var r0 common.Position
	if rf, ok := ret.Get(0).(func() common.Position); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(common.Position)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
