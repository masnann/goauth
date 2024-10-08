// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import (
	echo "github.com/labstack/echo/v4"

	jwt "github.com/golang-jwt/jwt"

	mock "github.com/stretchr/testify/mock"
)

// GeneratorInterface is an autogenerated mock type for the GeneratorInterface type
type GeneratorInterface struct {
	mock.Mock
}

// CompareHash provides a mock function with given fields: hash, input
func (_m *GeneratorInterface) CompareHash(hash string, input string) (bool, error) {
	ret := _m.Called(hash, input)

	if len(ret) == 0 {
		panic("no return value specified for CompareHash")
	}

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string) (bool, error)); ok {
		return rf(hash, input)
	}
	if rf, ok := ret.Get(0).(func(string, string) bool); ok {
		r0 = rf(hash, input)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(hash, input)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GenerateHash provides a mock function with given fields: input
func (_m *GeneratorInterface) GenerateHash(input string) (string, error) {
	ret := _m.Called(input)

	if len(ret) == 0 {
		panic("no return value specified for GenerateHash")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (string, error)); ok {
		return rf(input)
	}
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(input)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(input)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GenerateJWT provides a mock function with given fields: userID, email, role
func (_m *GeneratorInterface) GenerateJWT(userID int64, email string, role string) (string, error) {
	ret := _m.Called(userID, email, role)

	if len(ret) == 0 {
		panic("no return value specified for GenerateJWT")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(int64, string, string) (string, error)); ok {
		return rf(userID, email, role)
	}
	if rf, ok := ret.Get(0).(func(int64, string, string) string); ok {
		r0 = rf(userID, email, role)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(int64, string, string) error); ok {
		r1 = rf(userID, email, role)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GenerateOTP provides a mock function with given fields: length
func (_m *GeneratorInterface) GenerateOTP(length int) (string, error) {
	ret := _m.Called(length)

	if len(ret) == 0 {
		panic("no return value specified for GenerateOTP")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(int) (string, error)); ok {
		return rf(length)
	}
	if rf, ok := ret.Get(0).(func(int) string); ok {
		r0 = rf(length)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(length)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GenerateRefreshToken provides a mock function with given fields: userID
func (_m *GeneratorInterface) GenerateRefreshToken(userID int64) (string, error) {
	ret := _m.Called(userID)

	if len(ret) == 0 {
		panic("no return value specified for GenerateRefreshToken")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(int64) (string, error)); ok {
		return rf(userID)
	}
	if rf, ok := ret.Get(0).(func(int64) string); ok {
		r0 = rf(userID)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(int64) error); ok {
		r1 = rf(userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// HandlerErr provides a mock function with given fields: ctx, errorType, serviceName, userEmail, msg, err
func (_m *GeneratorInterface) HandlerErr(ctx echo.Context, errorType string, serviceName string, userEmail string, msg string, err error) error {
	ret := _m.Called(ctx, errorType, serviceName, userEmail, msg, err)

	if len(ret) == 0 {
		panic("no return value specified for HandlerErr")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(echo.Context, string, string, string, string, error) error); ok {
		r0 = rf(ctx, errorType, serviceName, userEmail, msg, err)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ValidateRefreshToken provides a mock function with given fields: tokenString
func (_m *GeneratorInterface) ValidateRefreshToken(tokenString string) (int64, error) {
	ret := _m.Called(tokenString)

	if len(ret) == 0 {
		panic("no return value specified for ValidateRefreshToken")
	}

	var r0 int64
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (int64, error)); ok {
		return rf(tokenString)
	}
	if rf, ok := ret.Get(0).(func(string) int64); ok {
		r0 = rf(tokenString)
	} else {
		r0 = ret.Get(0).(int64)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(tokenString)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ValidateToken provides a mock function with given fields: tokenString
func (_m *GeneratorInterface) ValidateToken(tokenString string) (*jwt.Token, error) {
	ret := _m.Called(tokenString)

	if len(ret) == 0 {
		panic("no return value specified for ValidateToken")
	}

	var r0 *jwt.Token
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*jwt.Token, error)); ok {
		return rf(tokenString)
	}
	if rf, ok := ret.Get(0).(func(string) *jwt.Token); ok {
		r0 = rf(tokenString)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*jwt.Token)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(tokenString)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewGeneratorInterface creates a new instance of GeneratorInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewGeneratorInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *GeneratorInterface {
	mock := &GeneratorInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
