package mock

import (
	models "github.com/odpf/optimus/models"
	mock "github.com/stretchr/testify/mock"

	optimus "github.com/odpf/optimus/api/proto/odpf/optimus/core/v1beta1"

	tree "github.com/odpf/optimus/core/tree"
)

// ProtoAdapter is an autogenerated mock type for the ProtoAdapter type
type ProtoAdapter struct {
	mock.Mock
}

// FromJobProto provides a mock function with given fields: _a0
func (_m *ProtoAdapter) FromJobProto(_a0 *optimus.JobSpecification) (models.JobSpec, error) {
	ret := _m.Called(_a0)

	var r0 models.JobSpec
	if rf, ok := ret.Get(0).(func(*optimus.JobSpecification) models.JobSpec); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(models.JobSpec)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*optimus.JobSpecification) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FromNamespaceProto provides a mock function with given fields: specification
func (_m *ProtoAdapter) FromNamespaceProto(specification *optimus.NamespaceSpecification) models.NamespaceSpec {
	ret := _m.Called(specification)

	var r0 models.NamespaceSpec
	if rf, ok := ret.Get(0).(func(*optimus.NamespaceSpecification) models.NamespaceSpec); ok {
		r0 = rf(specification)
	} else {
		r0 = ret.Get(0).(models.NamespaceSpec)
	}

	return r0
}

// FromProjectProto provides a mock function with given fields: _a0
func (_m *ProtoAdapter) FromProjectProto(_a0 *optimus.ProjectSpecification) models.ProjectSpec {
	ret := _m.Called(_a0)

	var r0 models.ProjectSpec
	if rf, ok := ret.Get(0).(func(*optimus.ProjectSpecification) models.ProjectSpec); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(models.ProjectSpec)
	}

	return r0
}

// FromResourceProto provides a mock function with given fields: res, storeName
func (_m *ProtoAdapter) FromResourceProto(res *optimus.ResourceSpecification, storeName string) (models.ResourceSpec, error) {
	ret := _m.Called(res, storeName)

	var r0 models.ResourceSpec
	if rf, ok := ret.Get(0).(func(*optimus.ResourceSpecification, string) models.ResourceSpec); ok {
		r0 = rf(res, storeName)
	} else {
		r0 = ret.Get(0).(models.ResourceSpec)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*optimus.ResourceSpecification, string) error); ok {
		r1 = rf(res, storeName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ToInstanceProto provides a mock function with given fields: _a0
func (_m *ProtoAdapter) ToInstanceProto(_a0 models.InstanceSpec) (*optimus.InstanceSpec, error) {
	ret := _m.Called(_a0)

	var r0 *optimus.InstanceSpec
	if rf, ok := ret.Get(0).(func(models.InstanceSpec) *optimus.InstanceSpec); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*optimus.InstanceSpec)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(models.InstanceSpec) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ToJobProto provides a mock function with given fields: _a0
func (_m *ProtoAdapter) ToJobProto(_a0 models.JobSpec) (*optimus.JobSpecification, error) {
	ret := _m.Called(_a0)

	var r0 *optimus.JobSpecification
	if rf, ok := ret.Get(0).(func(models.JobSpec) *optimus.JobSpecification); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*optimus.JobSpecification)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(models.JobSpec) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ToNamespaceProto provides a mock function with given fields: spec
func (_m *ProtoAdapter) ToNamespaceProto(spec models.NamespaceSpec) *optimus.NamespaceSpecification {
	ret := _m.Called(spec)

	var r0 *optimus.NamespaceSpecification
	if rf, ok := ret.Get(0).(func(models.NamespaceSpec) *optimus.NamespaceSpecification); ok {
		r0 = rf(spec)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*optimus.NamespaceSpecification)
		}
	}

	return r0
}

// ToProjectProto provides a mock function with given fields: _a0
func (_m *ProtoAdapter) ToProjectProto(_a0 models.ProjectSpec) *optimus.ProjectSpecification {
	ret := _m.Called(_a0)

	var r0 *optimus.ProjectSpecification
	if rf, ok := ret.Get(0).(func(models.ProjectSpec) *optimus.ProjectSpecification); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*optimus.ProjectSpecification)
		}
	}

	return r0
}

// ToReplayExecutionTreeNode provides a mock function with given fields: res
func (_m *ProtoAdapter) ToReplayExecutionTreeNode(res *tree.TreeNode) (*optimus.ReplayExecutionTreeNode, error) {
	ret := _m.Called(res)

	var r0 *optimus.ReplayExecutionTreeNode
	if rf, ok := ret.Get(0).(func(*tree.TreeNode) *optimus.ReplayExecutionTreeNode); ok {
		r0 = rf(res)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*optimus.ReplayExecutionTreeNode)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*tree.TreeNode) error); ok {
		r1 = rf(res)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ToReplayStatusTreeNode provides a mock function with given fields: res
func (_m *ProtoAdapter) ToReplayStatusTreeNode(res *tree.TreeNode) (*optimus.ReplayStatusTreeNode, error) {
	ret := _m.Called(res)

	var r0 *optimus.ReplayStatusTreeNode
	if rf, ok := ret.Get(0).(func(*tree.TreeNode) *optimus.ReplayStatusTreeNode); ok {
		r0 = rf(res)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*optimus.ReplayStatusTreeNode)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*tree.TreeNode) error); ok {
		r1 = rf(res)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ToResourceProto provides a mock function with given fields: res
func (_m *ProtoAdapter) ToResourceProto(res models.ResourceSpec) (*optimus.ResourceSpecification, error) {
	ret := _m.Called(res)

	var r0 *optimus.ResourceSpecification
	if rf, ok := ret.Get(0).(func(models.ResourceSpec) *optimus.ResourceSpecification); ok {
		r0 = rf(res)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*optimus.ResourceSpecification)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(models.ResourceSpec) error); ok {
		r1 = rf(res)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
