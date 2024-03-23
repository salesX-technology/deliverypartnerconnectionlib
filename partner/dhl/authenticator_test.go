package dhl

import (
	"errors"

	"github.com/stretchr/testify/suite"
	gomock "go.uber.org/mock/gomock"
)

type AuthenticatorTestSuite struct {
	ctrl *gomock.Controller
	suite.Suite

	mDHLAuthenticationAPI *MockDHLAuthenticationAPI
	svc                   *dhlAuthenticator
}

func (s *AuthenticatorTestSuite) SetupTest() {
	s.ctrl = gomock.NewController(s.T())
	s.mDHLAuthenticationAPI = NewMockDHLAuthenticationAPI(s.ctrl)
	s.svc = &dhlAuthenticator{
		dhlAuthenticationAPI: s.mDHLAuthenticationAPI,
		clientID:             "clientID",
		password:             "password",
	}
}

func (s *AuthenticatorTestSuite) TearDownTest() {
	s.ctrl.Finish()
}

func (s *AuthenticatorTestSuite) TestAuthenticate() {
	request := DHLAuthenticationAPIRequest{}
	response := DHLAuthenticationAPIResponse{
		AccessTokenResponse: AccessTokenResponse{
			Token: "accessToken",
		},
	}
	s.mDHLAuthenticationAPI.EXPECT().Get(map[string]string{}, "?clientID=clientID&password=password", request).Return(response, nil)

	accessToken, err := s.svc.Authenticate()
	s.NoError(err)
	s.Equal("accessToken", accessToken)
}

func (s *AuthenticatorTestSuite) TestAuthenticateFailed() {
	request := DHLAuthenticationAPIRequest{}
	s.mDHLAuthenticationAPI.EXPECT().Get(map[string]string{}, "?clientID=clientID&password=password", request).Return(DHLAuthenticationAPIResponse{}, errors.New("error"))

	accessToken, err := s.svc.Authenticate()
	s.ErrorIs(err, errors.New("failed to authenticate with dhl: error"))
	s.Empty(accessToken)
}
