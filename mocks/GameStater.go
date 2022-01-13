// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import common "github.com/Amari-Mecheri/GoSnakeLogic/pkg/common"

import mock "github.com/stretchr/testify/mock"

// GameStater is an autogenerated mock type for the GameStater type
type GameStater struct {
	mock.Mock
}

// BoardSize provides a mock function with given fields:
func (_m *GameStater) BoardSize() common.Size {
	ret := _m.Called()

	var r0 common.Size
	if rf, ok := ret.Get(0).(func() common.Size); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(common.Size)
	}

	return r0
}

// CreateObjects provides a mock function with given fields:
func (_m *GameStater) CreateObjects() ([]common.Sprite, error) {
	ret := _m.Called()

	var r0 []common.Sprite
	if rf, ok := ret.Get(0).(func() []common.Sprite); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]common.Sprite)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Dirty provides a mock function with given fields:
func (_m *GameStater) Dirty() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// GameInProgress provides a mock function with given fields:
func (_m *GameStater) GameInProgress() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// HighScore provides a mock function with given fields:
func (_m *GameStater) HighScore() int {
	ret := _m.Called()

	var r0 int
	if rf, ok := ret.Get(0).(func() int); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}

// InitBoard provides a mock function with given fields: size
func (_m *GameStater) InitBoard(size common.Size) error {
	ret := _m.Called(size)

	var r0 error
	if rf, ok := ret.Get(0).(func(common.Size) error); ok {
		r0 = rf(size)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MoveDown provides a mock function with given fields:
func (_m *GameStater) MoveDown() {
	_m.Called()
}

// MoveLeft provides a mock function with given fields:
func (_m *GameStater) MoveLeft() {
	_m.Called()
}

// MoveRight provides a mock function with given fields:
func (_m *GameStater) MoveRight() {
	_m.Called()
}

// MoveUp provides a mock function with given fields:
func (_m *GameStater) MoveUp() {
	_m.Called()
}

// Play provides a mock function with given fields:
func (_m *GameStater) Play() ([]common.Sprite, error) {
	ret := _m.Called()

	var r0 []common.Sprite
	if rf, ok := ret.Get(0).(func() []common.Sprite); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]common.Sprite)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Round provides a mock function with given fields:
func (_m *GameStater) Round() int {
	ret := _m.Called()

	var r0 int
	if rf, ok := ret.Get(0).(func() int); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}

// Score provides a mock function with given fields:
func (_m *GameStater) Score() int {
	ret := _m.Called()

	var r0 int
	if rf, ok := ret.Get(0).(func() int); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}

// SetGameInProgress provides a mock function with given fields: _a0
func (_m *GameStater) SetGameInProgress(_a0 bool) {
	_m.Called(_a0)
}

// SnakePosition provides a mock function with given fields:
func (_m *GameStater) SnakePosition() (common.Position, error) {
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

// SnakeSize provides a mock function with given fields:
func (_m *GameStater) SnakeSize() (int, error) {
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

// Start provides a mock function with given fields:
func (_m *GameStater) Start() {
	_m.Called()
}
