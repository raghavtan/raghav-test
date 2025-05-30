// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/motain/of-catalog/internal/modules/component/repository (interfaces: RepositoryInterface)
//
// Generated by this command:
//
//	mockgen -destination=./mocks/mock_repository.go -package=repository github.com/motain/of-catalog/internal/modules/component/repository RepositoryInterface
//

// Package repository is a generated GoMock package.
package repository

import (
	context "context"
	reflect "reflect"
	time "time"

	resources "github.com/motain/of-catalog/internal/modules/component/resources"
	gomock "go.uber.org/mock/gomock"
)

// MockRepositoryInterface is a mock of RepositoryInterface interface.
type MockRepositoryInterface struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryInterfaceMockRecorder
	isgomock struct{}
}

// MockRepositoryInterfaceMockRecorder is the mock recorder for MockRepositoryInterface.
type MockRepositoryInterfaceMockRecorder struct {
	mock *MockRepositoryInterface
}

// NewMockRepositoryInterface creates a new mock instance.
func NewMockRepositoryInterface(ctrl *gomock.Controller) *MockRepositoryInterface {
	mock := &MockRepositoryInterface{ctrl: ctrl}
	mock.recorder = &MockRepositoryInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepositoryInterface) EXPECT() *MockRepositoryInterfaceMockRecorder {
	return m.recorder
}

// AddDocument mocks base method.
func (m *MockRepositoryInterface) AddDocument(ctx context.Context, component resources.Component, document resources.Document) (resources.Document, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddDocument", ctx, component, document)
	ret0, _ := ret[0].(resources.Document)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddDocument indicates an expected call of AddDocument.
func (mr *MockRepositoryInterfaceMockRecorder) AddDocument(ctx, component, document any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddDocument", reflect.TypeOf((*MockRepositoryInterface)(nil).AddDocument), ctx, component, document)
}

// BindMetric mocks base method.
func (m *MockRepositoryInterface) BindMetric(ctx context.Context, component resources.Component, metricID, identifier string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BindMetric", ctx, component, metricID, identifier)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BindMetric indicates an expected call of BindMetric.
func (mr *MockRepositoryInterfaceMockRecorder) BindMetric(ctx, component, metricID, identifier any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BindMetric", reflect.TypeOf((*MockRepositoryInterface)(nil).BindMetric), ctx, component, metricID, identifier)
}

// Create mocks base method.
func (m *MockRepositoryInterface) Create(ctx context.Context, component resources.Component) (resources.Component, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, component)
	ret0, _ := ret[0].(resources.Component)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockRepositoryInterfaceMockRecorder) Create(ctx, component any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockRepositoryInterface)(nil).Create), ctx, component)
}

// Delete mocks base method.
func (m *MockRepositoryInterface) Delete(ctx context.Context, component resources.Component) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, component)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockRepositoryInterfaceMockRecorder) Delete(ctx, component any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockRepositoryInterface)(nil).Delete), ctx, component)
}

// GetBySlug mocks base method.
func (m *MockRepositoryInterface) GetBySlug(ctx context.Context, component resources.Component) (*resources.Component, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBySlug", ctx, component)
	ret0, _ := ret[0].(*resources.Component)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBySlug indicates an expected call of GetBySlug.
func (mr *MockRepositoryInterfaceMockRecorder) GetBySlug(ctx, component any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBySlug", reflect.TypeOf((*MockRepositoryInterface)(nil).GetBySlug), ctx, component)
}

// Push mocks base method.
func (m *MockRepositoryInterface) Push(ctx context.Context, metricSource resources.MetricSource, value float64, recordedAt time.Time) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Push", ctx, metricSource, value, recordedAt)
	ret0, _ := ret[0].(error)
	return ret0
}

// Push indicates an expected call of Push.
func (mr *MockRepositoryInterfaceMockRecorder) Push(ctx, metricSource, value, recordedAt any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Push", reflect.TypeOf((*MockRepositoryInterface)(nil).Push), ctx, metricSource, value, recordedAt)
}

// RemoveDocument mocks base method.
func (m *MockRepositoryInterface) RemoveDocument(ctx context.Context, component resources.Component, document resources.Document) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveDocument", ctx, component, document)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveDocument indicates an expected call of RemoveDocument.
func (mr *MockRepositoryInterfaceMockRecorder) RemoveDocument(ctx, component, document any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveDocument", reflect.TypeOf((*MockRepositoryInterface)(nil).RemoveDocument), ctx, component, document)
}

// SetAPISpecifications mocks base method.
func (m *MockRepositoryInterface) SetAPISpecifications(ctx context.Context, component resources.Component, apiSpecs, apiSpecsFile string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetAPISpecifications", ctx, component, apiSpecs, apiSpecsFile)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetAPISpecifications indicates an expected call of SetAPISpecifications.
func (mr *MockRepositoryInterfaceMockRecorder) SetAPISpecifications(ctx, component, apiSpecs, apiSpecsFile any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetAPISpecifications", reflect.TypeOf((*MockRepositoryInterface)(nil).SetAPISpecifications), ctx, component, apiSpecs, apiSpecsFile)
}

// SetDependency mocks base method.
func (m *MockRepositoryInterface) SetDependency(ctx context.Context, dependent, provider resources.Component) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetDependency", ctx, dependent, provider)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetDependency indicates an expected call of SetDependency.
func (mr *MockRepositoryInterfaceMockRecorder) SetDependency(ctx, dependent, provider any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetDependency", reflect.TypeOf((*MockRepositoryInterface)(nil).SetDependency), ctx, dependent, provider)
}

// UnbindMetric mocks base method.
func (m *MockRepositoryInterface) UnbindMetric(ctx context.Context, metricSource resources.MetricSource) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UnbindMetric", ctx, metricSource)
	ret0, _ := ret[0].(error)
	return ret0
}

// UnbindMetric indicates an expected call of UnbindMetric.
func (mr *MockRepositoryInterfaceMockRecorder) UnbindMetric(ctx, metricSource any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnbindMetric", reflect.TypeOf((*MockRepositoryInterface)(nil).UnbindMetric), ctx, metricSource)
}

// UnsetDependency mocks base method.
func (m *MockRepositoryInterface) UnsetDependency(ctx context.Context, dependent, provider resources.Component) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UnsetDependency", ctx, dependent, provider)
	ret0, _ := ret[0].(error)
	return ret0
}

// UnsetDependency indicates an expected call of UnsetDependency.
func (mr *MockRepositoryInterfaceMockRecorder) UnsetDependency(ctx, dependent, provider any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnsetDependency", reflect.TypeOf((*MockRepositoryInterface)(nil).UnsetDependency), ctx, dependent, provider)
}

// Update mocks base method.
func (m *MockRepositoryInterface) Update(ctx context.Context, component resources.Component) (resources.Component, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, component)
	ret0, _ := ret[0].(resources.Component)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockRepositoryInterfaceMockRecorder) Update(ctx, component any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockRepositoryInterface)(nil).Update), ctx, component)
}

// UpdateDocument mocks base method.
func (m *MockRepositoryInterface) UpdateDocument(ctx context.Context, component resources.Component, document resources.Document) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateDocument", ctx, component, document)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateDocument indicates an expected call of UpdateDocument.
func (mr *MockRepositoryInterfaceMockRecorder) UpdateDocument(ctx, component, document any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateDocument", reflect.TypeOf((*MockRepositoryInterface)(nil).UpdateDocument), ctx, component, document)
}
