// Code generated by mockery v2.32.0. DO NOT EDIT.

package mocks

import (
	context "context"
	mongo "khanhanhtr/sample2/mongo"

	mock "github.com/stretchr/testify/mock"

	mongo_drivermongo "go.mongodb.org/mongo-driver/mongo"

	options "go.mongodb.org/mongo-driver/mongo/options"
)

// Collection is an autogenerated mock type for the Collection type
type Collection struct {
	mock.Mock
}

// Aggregate provides a mock function with given fields: ctx, pipeline
func (_m *Collection) Aggregate(ctx context.Context, pipeline interface{}) (mongo.Cursor, error) {
	ret := _m.Called(ctx, pipeline)

	var r0 mongo.Cursor
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, interface{}) (mongo.Cursor, error)); ok {
		return rf(ctx, pipeline)
	}
	if rf, ok := ret.Get(0).(func(context.Context, interface{}) mongo.Cursor); ok {
		r0 = rf(ctx, pipeline)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(mongo.Cursor)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, interface{}) error); ok {
		r1 = rf(ctx, pipeline)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CountDocuments provides a mock function with given fields: ctx, filter, _a2
func (_m *Collection) CountDocuments(ctx context.Context, filter interface{}, _a2 ...*options.CountOptions) (int64, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, filter)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 int64
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, interface{}, ...*options.CountOptions) (int64, error)); ok {
		return rf(ctx, filter, _a2...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, interface{}, ...*options.CountOptions) int64); ok {
		r0 = rf(ctx, filter, _a2...)
	} else {
		r0 = ret.Get(0).(int64)
	}

	if rf, ok := ret.Get(1).(func(context.Context, interface{}, ...*options.CountOptions) error); ok {
		r1 = rf(ctx, filter, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteMany provides a mock function with given fields: ctx, filter
func (_m *Collection) DeleteMany(ctx context.Context, filter interface{}) (int64, error) {
	ret := _m.Called(ctx, filter)

	var r0 int64
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, interface{}) (int64, error)); ok {
		return rf(ctx, filter)
	}
	if rf, ok := ret.Get(0).(func(context.Context, interface{}) int64); ok {
		r0 = rf(ctx, filter)
	} else {
		r0 = ret.Get(0).(int64)
	}

	if rf, ok := ret.Get(1).(func(context.Context, interface{}) error); ok {
		r1 = rf(ctx, filter)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteOne provides a mock function with given fields: ctx, filter
func (_m *Collection) DeleteOne(ctx context.Context, filter interface{}) (int64, error) {
	ret := _m.Called(ctx, filter)

	var r0 int64
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, interface{}) (int64, error)); ok {
		return rf(ctx, filter)
	}
	if rf, ok := ret.Get(0).(func(context.Context, interface{}) int64); ok {
		r0 = rf(ctx, filter)
	} else {
		r0 = ret.Get(0).(int64)
	}

	if rf, ok := ret.Get(1).(func(context.Context, interface{}) error); ok {
		r1 = rf(ctx, filter)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Find provides a mock function with given fields: ctx, filter, _a2
func (_m *Collection) Find(ctx context.Context, filter interface{}, _a2 ...*options.FindOptions) (mongo.Cursor, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, filter)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 mongo.Cursor
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, interface{}, ...*options.FindOptions) (mongo.Cursor, error)); ok {
		return rf(ctx, filter, _a2...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, interface{}, ...*options.FindOptions) mongo.Cursor); ok {
		r0 = rf(ctx, filter, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(mongo.Cursor)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, interface{}, ...*options.FindOptions) error); ok {
		r1 = rf(ctx, filter, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindOne provides a mock function with given fields: ctx, filter
func (_m *Collection) FindOne(ctx context.Context, filter interface{}) mongo.SingleResult {
	ret := _m.Called(ctx, filter)

	var r0 mongo.SingleResult
	if rf, ok := ret.Get(0).(func(context.Context, interface{}) mongo.SingleResult); ok {
		r0 = rf(ctx, filter)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(mongo.SingleResult)
		}
	}

	return r0
}

// InsertMany provides a mock function with given fields: ctx, documents
func (_m *Collection) InsertMany(ctx context.Context, documents []interface{}) ([]interface{}, error) {
	ret := _m.Called(ctx, documents)

	var r0 []interface{}
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, []interface{}) ([]interface{}, error)); ok {
		return rf(ctx, documents)
	}
	if rf, ok := ret.Get(0).(func(context.Context, []interface{}) []interface{}); ok {
		r0 = rf(ctx, documents)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]interface{})
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, []interface{}) error); ok {
		r1 = rf(ctx, documents)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// InsertOne provides a mock function with given fields: ctx, document
func (_m *Collection) InsertOne(ctx context.Context, document interface{}) (interface{}, error) {
	ret := _m.Called(ctx, document)

	var r0 interface{}
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, interface{}) (interface{}, error)); ok {
		return rf(ctx, document)
	}
	if rf, ok := ret.Get(0).(func(context.Context, interface{}) interface{}); ok {
		r0 = rf(ctx, document)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, interface{}) error); ok {
		r1 = rf(ctx, document)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateMany provides a mock function with given fields: ctx, filter, update, _a3
func (_m *Collection) UpdateMany(ctx context.Context, filter interface{}, update interface{}, _a3 ...*options.UpdateOptions) (*mongo_drivermongo.UpdateResult, error) {
	_va := make([]interface{}, len(_a3))
	for _i := range _a3 {
		_va[_i] = _a3[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, filter, update)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *mongo_drivermongo.UpdateResult
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, interface{}, interface{}, ...*options.UpdateOptions) (*mongo_drivermongo.UpdateResult, error)); ok {
		return rf(ctx, filter, update, _a3...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, interface{}, interface{}, ...*options.UpdateOptions) *mongo_drivermongo.UpdateResult); ok {
		r0 = rf(ctx, filter, update, _a3...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*mongo_drivermongo.UpdateResult)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, interface{}, interface{}, ...*options.UpdateOptions) error); ok {
		r1 = rf(ctx, filter, update, _a3...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateOne provides a mock function with given fields: ctx, filter, update, _a3
func (_m *Collection) UpdateOne(ctx context.Context, filter interface{}, update interface{}, _a3 ...*options.UpdateOptions) (*mongo_drivermongo.UpdateResult, error) {
	_va := make([]interface{}, len(_a3))
	for _i := range _a3 {
		_va[_i] = _a3[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, filter, update)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *mongo_drivermongo.UpdateResult
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, interface{}, interface{}, ...*options.UpdateOptions) (*mongo_drivermongo.UpdateResult, error)); ok {
		return rf(ctx, filter, update, _a3...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, interface{}, interface{}, ...*options.UpdateOptions) *mongo_drivermongo.UpdateResult); ok {
		r0 = rf(ctx, filter, update, _a3...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*mongo_drivermongo.UpdateResult)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, interface{}, interface{}, ...*options.UpdateOptions) error); ok {
		r1 = rf(ctx, filter, update, _a3...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewCollection creates a new instance of Collection. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewCollection(t interface {
	mock.TestingT
	Cleanup(func())
}) *Collection {
	mock := &Collection{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}