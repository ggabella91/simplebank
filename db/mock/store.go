// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/ggabella91/simplebank/db/sqlc (interfaces: Store)

// Package mockdb is a generated GoMock package.
package mockdb

import (
	context "context"
	reflect "reflect"

	db "github.com/ggabella91/simplebank/db/sqlc"
	uuid "github.com/google/uuid"
	gomock "go.uber.org/mock/gomock"
)

// MockStore is a mock of Store interface.
type MockStore struct {
	ctrl     *gomock.Controller
	recorder *MockStoreMockRecorder
}

// MockStoreMockRecorder is the mock recorder for MockStore.
type MockStoreMockRecorder struct {
	mock *MockStore
}

// NewMockStore creates a new mock instance.
func NewMockStore(ctrl *gomock.Controller) *MockStore {
	mock := &MockStore{ctrl: ctrl}
	mock.recorder = &MockStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStore) EXPECT() *MockStoreMockRecorder {
	return m.recorder
}

// AddAccountBalance mocks base method.
func (m *MockStore) AddAccountBalance(arg0 context.Context, arg1 db.AddAccountBalanceParams) (db.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddAccountBalance", arg0, arg1)
	ret0, _ := ret[0].(db.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddAccountBalance indicates an expected call of AddAccountBalance.
func (mr *MockStoreMockRecorder) AddAccountBalance(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddAccountBalance", reflect.TypeOf((*MockStore)(nil).AddAccountBalance), arg0, arg1)
}

// CreateAccount mocks base method.
func (m *MockStore) CreateAccount(arg0 context.Context, arg1 db.CreateAccountParams) (db.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateAccount", arg0, arg1)
	ret0, _ := ret[0].(db.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateAccount indicates an expected call of CreateAccount.
func (mr *MockStoreMockRecorder) CreateAccount(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateAccount", reflect.TypeOf((*MockStore)(nil).CreateAccount), arg0, arg1)
}

// CreateEntry mocks base method.
func (m *MockStore) CreateEntry(arg0 context.Context, arg1 db.CreateEntryParams) (db.Entry, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateEntry", arg0, arg1)
	ret0, _ := ret[0].(db.Entry)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateEntry indicates an expected call of CreateEntry.
func (mr *MockStoreMockRecorder) CreateEntry(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateEntry", reflect.TypeOf((*MockStore)(nil).CreateEntry), arg0, arg1)
}

// CreateSession mocks base method.
func (m *MockStore) CreateSession(arg0 context.Context, arg1 db.CreateSessionParams) (db.Session, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateSession", arg0, arg1)
	ret0, _ := ret[0].(db.Session)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateSession indicates an expected call of CreateSession.
func (mr *MockStoreMockRecorder) CreateSession(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateSession", reflect.TypeOf((*MockStore)(nil).CreateSession), arg0, arg1)
}

// CreateTransfer mocks base method.
func (m *MockStore) CreateTransfer(arg0 context.Context, arg1 db.CreateTransferParams) (db.Transfer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTransfer", arg0, arg1)
	ret0, _ := ret[0].(db.Transfer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateTransfer indicates an expected call of CreateTransfer.
func (mr *MockStoreMockRecorder) CreateTransfer(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTransfer", reflect.TypeOf((*MockStore)(nil).CreateTransfer), arg0, arg1)
}

// CreateUser mocks base method.
func (m *MockStore) CreateUser(arg0 context.Context, arg1 db.CreateUserParams) (db.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", arg0, arg1)
	ret0, _ := ret[0].(db.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockStoreMockRecorder) CreateUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockStore)(nil).CreateUser), arg0, arg1)
}

// CreateUserTx mocks base method.
func (m *MockStore) CreateUserTx(arg0 context.Context, arg1 db.CreateUserTxParams) (db.CreateUserTxResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUserTx", arg0, arg1)
	ret0, _ := ret[0].(db.CreateUserTxResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUserTx indicates an expected call of CreateUserTx.
func (mr *MockStoreMockRecorder) CreateUserTx(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUserTx", reflect.TypeOf((*MockStore)(nil).CreateUserTx), arg0, arg1)
}

// CreateVerifyEmail mocks base method.
func (m *MockStore) CreateVerifyEmail(arg0 context.Context, arg1 db.CreateVerifyEmailParams) (db.VerifyEmail, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateVerifyEmail", arg0, arg1)
	ret0, _ := ret[0].(db.VerifyEmail)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateVerifyEmail indicates an expected call of CreateVerifyEmail.
func (mr *MockStoreMockRecorder) CreateVerifyEmail(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateVerifyEmail", reflect.TypeOf((*MockStore)(nil).CreateVerifyEmail), arg0, arg1)
}

// DeleteAccount mocks base method.
func (m *MockStore) DeleteAccount(arg0 context.Context, arg1 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAccount", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteAccount indicates an expected call of DeleteAccount.
func (mr *MockStoreMockRecorder) DeleteAccount(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAccount", reflect.TypeOf((*MockStore)(nil).DeleteAccount), arg0, arg1)
}

// GetAccount mocks base method.
func (m *MockStore) GetAccount(arg0 context.Context, arg1 int64) (db.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAccount", arg0, arg1)
	ret0, _ := ret[0].(db.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAccount indicates an expected call of GetAccount.
func (mr *MockStoreMockRecorder) GetAccount(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccount", reflect.TypeOf((*MockStore)(nil).GetAccount), arg0, arg1)
}

// GetAccountForUpdate mocks base method.
func (m *MockStore) GetAccountForUpdate(arg0 context.Context, arg1 int64) (db.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAccountForUpdate", arg0, arg1)
	ret0, _ := ret[0].(db.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAccountForUpdate indicates an expected call of GetAccountForUpdate.
func (mr *MockStoreMockRecorder) GetAccountForUpdate(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccountForUpdate", reflect.TypeOf((*MockStore)(nil).GetAccountForUpdate), arg0, arg1)
}

// GetEntry mocks base method.
func (m *MockStore) GetEntry(arg0 context.Context, arg1 int64) (db.Entry, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEntry", arg0, arg1)
	ret0, _ := ret[0].(db.Entry)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetEntry indicates an expected call of GetEntry.
func (mr *MockStoreMockRecorder) GetEntry(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEntry", reflect.TypeOf((*MockStore)(nil).GetEntry), arg0, arg1)
}

// GetSession mocks base method.
func (m *MockStore) GetSession(arg0 context.Context, arg1 uuid.UUID) (db.Session, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSession", arg0, arg1)
	ret0, _ := ret[0].(db.Session)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSession indicates an expected call of GetSession.
func (mr *MockStoreMockRecorder) GetSession(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSession", reflect.TypeOf((*MockStore)(nil).GetSession), arg0, arg1)
}

// GetTransfer mocks base method.
func (m *MockStore) GetTransfer(arg0 context.Context, arg1 int64) (db.Transfer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTransfer", arg0, arg1)
	ret0, _ := ret[0].(db.Transfer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTransfer indicates an expected call of GetTransfer.
func (mr *MockStoreMockRecorder) GetTransfer(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTransfer", reflect.TypeOf((*MockStore)(nil).GetTransfer), arg0, arg1)
}

// GetUser mocks base method.
func (m *MockStore) GetUser(arg0 context.Context, arg1 string) (db.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUser", arg0, arg1)
	ret0, _ := ret[0].(db.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUser indicates an expected call of GetUser.
func (mr *MockStoreMockRecorder) GetUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUser", reflect.TypeOf((*MockStore)(nil).GetUser), arg0, arg1)
}

// ListAccounts mocks base method.
func (m *MockStore) ListAccounts(arg0 context.Context, arg1 db.ListAccountsParams) ([]db.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAccounts", arg0, arg1)
	ret0, _ := ret[0].([]db.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAccounts indicates an expected call of ListAccounts.
func (mr *MockStoreMockRecorder) ListAccounts(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAccounts", reflect.TypeOf((*MockStore)(nil).ListAccounts), arg0, arg1)
}

// ListEntries mocks base method.
func (m *MockStore) ListEntries(arg0 context.Context, arg1 db.ListEntriesParams) ([]db.Entry, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListEntries", arg0, arg1)
	ret0, _ := ret[0].([]db.Entry)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListEntries indicates an expected call of ListEntries.
func (mr *MockStoreMockRecorder) ListEntries(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListEntries", reflect.TypeOf((*MockStore)(nil).ListEntries), arg0, arg1)
}

// ListTransfers mocks base method.
func (m *MockStore) ListTransfers(arg0 context.Context, arg1 db.ListTransfersParams) ([]db.Transfer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListTransfers", arg0, arg1)
	ret0, _ := ret[0].([]db.Transfer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListTransfers indicates an expected call of ListTransfers.
func (mr *MockStoreMockRecorder) ListTransfers(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTransfers", reflect.TypeOf((*MockStore)(nil).ListTransfers), arg0, arg1)
}

// TransferTx mocks base method.
func (m *MockStore) TransferTx(arg0 context.Context, arg1 db.TransferTxParams) (db.TransferTxResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TransferTx", arg0, arg1)
	ret0, _ := ret[0].(db.TransferTxResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TransferTx indicates an expected call of TransferTx.
func (mr *MockStoreMockRecorder) TransferTx(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TransferTx", reflect.TypeOf((*MockStore)(nil).TransferTx), arg0, arg1)
}

// UpdateAccount mocks base method.
func (m *MockStore) UpdateAccount(arg0 context.Context, arg1 db.UpdateAccountParams) (db.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAccount", arg0, arg1)
	ret0, _ := ret[0].(db.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateAccount indicates an expected call of UpdateAccount.
func (mr *MockStoreMockRecorder) UpdateAccount(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAccount", reflect.TypeOf((*MockStore)(nil).UpdateAccount), arg0, arg1)
}

// UpdateUser mocks base method.
func (m *MockStore) UpdateUser(arg0 context.Context, arg1 db.UpdateUserParams) (db.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUser", arg0, arg1)
	ret0, _ := ret[0].(db.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateUser indicates an expected call of UpdateUser.
func (mr *MockStoreMockRecorder) UpdateUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUser", reflect.TypeOf((*MockStore)(nil).UpdateUser), arg0, arg1)
}

// UpdateVerifyEmail mocks base method.
func (m *MockStore) UpdateVerifyEmail(arg0 context.Context, arg1 db.UpdateVerifyEmailParams) (db.VerifyEmail, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateVerifyEmail", arg0, arg1)
	ret0, _ := ret[0].(db.VerifyEmail)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateVerifyEmail indicates an expected call of UpdateVerifyEmail.
func (mr *MockStoreMockRecorder) UpdateVerifyEmail(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateVerifyEmail", reflect.TypeOf((*MockStore)(nil).UpdateVerifyEmail), arg0, arg1)
}

// VerifyEmailTx mocks base method.
func (m *MockStore) VerifyEmailTx(arg0 context.Context, arg1 db.VerifyEmailTxParams) (db.VerifyEmailTxResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VerifyEmailTx", arg0, arg1)
	ret0, _ := ret[0].(db.VerifyEmailTxResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// VerifyEmailTx indicates an expected call of VerifyEmailTx.
func (mr *MockStoreMockRecorder) VerifyEmailTx(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VerifyEmailTx", reflect.TypeOf((*MockStore)(nil).VerifyEmailTx), arg0, arg1)
}
