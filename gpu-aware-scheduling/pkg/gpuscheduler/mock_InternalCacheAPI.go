// Code generated by mockery v1.0.0. DO NOT EDIT.

// +build !validation
// re-generate with 'make mock'

package gpuscheduler

import (
	mock "github.com/stretchr/testify/mock"
	cache "k8s.io/client-go/tools/cache"
)

// MockInternalCacheAPI is an autogenerated mock type for the InternalCacheAPI type
type MockInternalCacheAPI struct {
	mock.Mock
}

// WaitForCacheSync provides a mock function with given fields: stopCh, cacheSyncs
func (_m *MockInternalCacheAPI) WaitForCacheSync(stopCh <-chan struct{}, cacheSyncs ...cache.InformerSynced) bool {
	_va := make([]interface{}, len(cacheSyncs))
	for _i := range cacheSyncs {
		_va[_i] = cacheSyncs[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, stopCh)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 bool
	if rf, ok := ret.Get(0).(func(<-chan struct{}, ...cache.InformerSynced) bool); ok {
		r0 = rf(stopCh, cacheSyncs...)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}