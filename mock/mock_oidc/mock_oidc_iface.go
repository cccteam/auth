// Code generated by MockGen. DO NOT EDIT.
// Source: ../oidc/oidc_iface.go
//
// Generated by this command:
//
//	mockgen -source ../oidc/oidc_iface.go -destination mock_oidc/mock_oidc_iface.go
//

// Package mock_oidc is a generated GoMock package.
package mock_oidc

import (
	context "context"
	http "net/http"
	reflect "reflect"

	oidc "github.com/coreos/go-oidc/v3/oidc"
	gomock "go.uber.org/mock/gomock"
	oauth2 "golang.org/x/oauth2"
)

// MockAuthenticator is a mock of Authenticator interface.
type MockAuthenticator struct {
	ctrl     *gomock.Controller
	recorder *MockAuthenticatorMockRecorder
	isgomock struct{}
}

// MockAuthenticatorMockRecorder is the mock recorder for MockAuthenticator.
type MockAuthenticatorMockRecorder struct {
	mock *MockAuthenticator
}

// NewMockAuthenticator creates a new mock instance.
func NewMockAuthenticator(ctrl *gomock.Controller) *MockAuthenticator {
	mock := &MockAuthenticator{ctrl: ctrl}
	mock.recorder = &MockAuthenticatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAuthenticator) EXPECT() *MockAuthenticatorMockRecorder {
	return m.recorder
}

// AuthCodeURL mocks base method.
func (m *MockAuthenticator) AuthCodeURL(w http.ResponseWriter, returnURL string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AuthCodeURL", w, returnURL)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AuthCodeURL indicates an expected call of AuthCodeURL.
func (mr *MockAuthenticatorMockRecorder) AuthCodeURL(w, returnURL any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AuthCodeURL", reflect.TypeOf((*MockAuthenticator)(nil).AuthCodeURL), w, returnURL)
}

// LoginURL mocks base method.
func (m *MockAuthenticator) LoginURL() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LoginURL")
	ret0, _ := ret[0].(string)
	return ret0
}

// LoginURL indicates an expected call of LoginURL.
func (mr *MockAuthenticatorMockRecorder) LoginURL() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LoginURL", reflect.TypeOf((*MockAuthenticator)(nil).LoginURL))
}

// Verify mocks base method.
func (m *MockAuthenticator) Verify(ctx context.Context, w http.ResponseWriter, r *http.Request, claims any) (string, string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Verify", ctx, w, r, claims)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(string)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Verify indicates an expected call of Verify.
func (mr *MockAuthenticatorMockRecorder) Verify(ctx, w, r, claims any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Verify", reflect.TypeOf((*MockAuthenticator)(nil).Verify), ctx, w, r, claims)
}

// Mockprovider is a mock of provider interface.
type Mockprovider struct {
	ctrl     *gomock.Controller
	recorder *MockproviderMockRecorder
	isgomock struct{}
}

// MockproviderMockRecorder is the mock recorder for Mockprovider.
type MockproviderMockRecorder struct {
	mock *Mockprovider
}

// NewMockprovider creates a new mock instance.
func NewMockprovider(ctrl *gomock.Controller) *Mockprovider {
	mock := &Mockprovider{ctrl: ctrl}
	mock.recorder = &MockproviderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *Mockprovider) EXPECT() *MockproviderMockRecorder {
	return m.recorder
}

// Endpoint mocks base method.
func (m *Mockprovider) Endpoint() oauth2.Endpoint {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Endpoint")
	ret0, _ := ret[0].(oauth2.Endpoint)
	return ret0
}

// Endpoint indicates an expected call of Endpoint.
func (mr *MockproviderMockRecorder) Endpoint() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Endpoint", reflect.TypeOf((*Mockprovider)(nil).Endpoint))
}

// Verifier mocks base method.
func (m *Mockprovider) Verifier(config *oidc.Config) *oidc.IDTokenVerifier {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Verifier", config)
	ret0, _ := ret[0].(*oidc.IDTokenVerifier)
	return ret0
}

// Verifier indicates an expected call of Verifier.
func (mr *MockproviderMockRecorder) Verifier(config any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Verifier", reflect.TypeOf((*Mockprovider)(nil).Verifier), config)
}

// Mockconfig is a mock of config interface.
type Mockconfig struct {
	ctrl     *gomock.Controller
	recorder *MockconfigMockRecorder
	isgomock struct{}
}

// MockconfigMockRecorder is the mock recorder for Mockconfig.
type MockconfigMockRecorder struct {
	mock *Mockconfig
}

// NewMockconfig creates a new mock instance.
func NewMockconfig(ctrl *gomock.Controller) *Mockconfig {
	mock := &Mockconfig{ctrl: ctrl}
	mock.recorder = &MockconfigMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *Mockconfig) EXPECT() *MockconfigMockRecorder {
	return m.recorder
}

// AuthCodeURL mocks base method.
func (m *Mockconfig) AuthCodeURL(state string, opts ...oauth2.AuthCodeOption) string {
	m.ctrl.T.Helper()
	varargs := []any{state}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "AuthCodeURL", varargs...)
	ret0, _ := ret[0].(string)
	return ret0
}

// AuthCodeURL indicates an expected call of AuthCodeURL.
func (mr *MockconfigMockRecorder) AuthCodeURL(state any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{state}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AuthCodeURL", reflect.TypeOf((*Mockconfig)(nil).AuthCodeURL), varargs...)
}

// ClientID mocks base method.
func (m *Mockconfig) ClientID() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClientID")
	ret0, _ := ret[0].(string)
	return ret0
}

// ClientID indicates an expected call of ClientID.
func (mr *MockconfigMockRecorder) ClientID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClientID", reflect.TypeOf((*Mockconfig)(nil).ClientID))
}

// Exchange mocks base method.
func (m *Mockconfig) Exchange(ctx context.Context, code string, opts ...oauth2.AuthCodeOption) (*oauth2.Token, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, code}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Exchange", varargs...)
	ret0, _ := ret[0].(*oauth2.Token)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Exchange indicates an expected call of Exchange.
func (mr *MockconfigMockRecorder) Exchange(ctx, code any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, code}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Exchange", reflect.TypeOf((*Mockconfig)(nil).Exchange), varargs...)
}
