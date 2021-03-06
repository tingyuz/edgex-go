// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"
import models "github.com/edgexfoundry/go-mod-core-contracts/models"

// DeviceProfileLoader is an autogenerated mock type for the DeviceProfileLoader type
type DeviceProfileLoader struct {
	mock.Mock
}

// GetAllDeviceProfiles provides a mock function with given fields:
func (_m *DeviceProfileLoader) GetAllDeviceProfiles() ([]models.DeviceProfile, error) {
	ret := _m.Called()

	var r0 []models.DeviceProfile
	if rf, ok := ret.Get(0).(func() []models.DeviceProfile); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.DeviceProfile)
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

// GetDeviceProfileById provides a mock function with given fields: id
func (_m *DeviceProfileLoader) GetDeviceProfileById(id string) (models.DeviceProfile, error) {
	ret := _m.Called(id)

	var r0 models.DeviceProfile
	if rf, ok := ret.Get(0).(func(string) models.DeviceProfile); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(models.DeviceProfile)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetDeviceProfileByName provides a mock function with given fields: n
func (_m *DeviceProfileLoader) GetDeviceProfileByName(n string) (models.DeviceProfile, error) {
	ret := _m.Called(n)

	var r0 models.DeviceProfile
	if rf, ok := ret.Get(0).(func(string) models.DeviceProfile); ok {
		r0 = rf(n)
	} else {
		r0 = ret.Get(0).(models.DeviceProfile)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(n)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetDeviceProfilesByManufacturer provides a mock function with given fields: man
func (_m *DeviceProfileLoader) GetDeviceProfilesByManufacturer(man string) ([]models.DeviceProfile, error) {
	ret := _m.Called(man)

	var r0 []models.DeviceProfile
	if rf, ok := ret.Get(0).(func(string) []models.DeviceProfile); ok {
		r0 = rf(man)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.DeviceProfile)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(man)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetDeviceProfilesByManufacturerModel provides a mock function with given fields: man, mod
func (_m *DeviceProfileLoader) GetDeviceProfilesByManufacturerModel(man string, mod string) ([]models.DeviceProfile, error) {
	ret := _m.Called(man, mod)

	var r0 []models.DeviceProfile
	if rf, ok := ret.Get(0).(func(string, string) []models.DeviceProfile); ok {
		r0 = rf(man, mod)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.DeviceProfile)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(man, mod)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetDeviceProfilesByModel provides a mock function with given fields: m
func (_m *DeviceProfileLoader) GetDeviceProfilesByModel(m string) ([]models.DeviceProfile, error) {
	ret := _m.Called(m)

	var r0 []models.DeviceProfile
	if rf, ok := ret.Get(0).(func(string) []models.DeviceProfile); ok {
		r0 = rf(m)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.DeviceProfile)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(m)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetDeviceProfilesWithLabel provides a mock function with given fields: l
func (_m *DeviceProfileLoader) GetDeviceProfilesWithLabel(l string) ([]models.DeviceProfile, error) {
	ret := _m.Called(l)

	var r0 []models.DeviceProfile
	if rf, ok := ret.Get(0).(func(string) []models.DeviceProfile); ok {
		r0 = rf(l)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.DeviceProfile)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(l)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
