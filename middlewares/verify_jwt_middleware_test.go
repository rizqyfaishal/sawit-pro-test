package middlewares

import (
	"github.com/SawitProRecruitment/UserService/services"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"
	"reflect"
	"testing"
)

type VerifyJWTMiddlewareTestSuite struct {
	suite.Suite

	authenticationService *services.MockAuthenticationServiceInterface
	MockController        *gomock.Controller
}

func TestVerifyJWTMiddlewareTestSuite(t *testing.T) {
	suite.Run(t, new(VerifyJWTMiddlewareTestSuite))
}

func (ts *VerifyJWTMiddlewareTestSuite) SetupSuite() {

	mockCtrl := gomock.NewController(ts.T())

	ts.MockController = mockCtrl

	defer mockCtrl.Finish()

	ts.authenticationService = services.NewMockAuthenticationServiceInterface(mockCtrl)
}

func (ts *VerifyJWTMiddlewareTestSuite) TestNewVerifyJwtMiddleware() {
	type args struct {
		svc services.Services
	}
	tests := []struct {
		name string
		args args
		want VerifyJwtMiddleware
	}{
		{
			name: "When given valid authentication service, then it will return Verify JWT middleware",
			args: args{
				svc: struct {
					Authentication services.AuthenticationServiceInterface
					User           services.UserServiceInterface
				}{Authentication: ts.authenticationService, User: nil},
			},
			want: VerifyJwtMiddleware{
				authenticationService: ts.authenticationService,
			},
		},
	}
	for _, tt := range tests {
		ts.T().Run(tt.name, func(t *testing.T) {
			if got := NewVerifyJwtMiddleware(tt.args.svc); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewVerifyJwtMiddleware() = %v, want %v", got, tt.want)
			}
		})
	}
}

func (ts *VerifyJWTMiddlewareTestSuite) TestVerifyJwtMiddleware_isUrlAllowed() {
	type fields struct {
		authenticationService services.AuthenticationServiceInterface
	}
	type args struct {
		requestMethod string
		url           string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "When the url is POST /users/login, then it bypassed and return true",
			fields: fields{
				authenticationService: ts.authenticationService,
			},
			args: args{
				requestMethod: "POST",
				url:           "/users/login",
			},
			want: true,
		},
		{
			name: "When the url is POST /users/me, then it not bypassed and return false",
			fields: fields{
				authenticationService: ts.authenticationService,
			},
			args: args{
				requestMethod: "POST",
				url:           "/users/me",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		ts.T().Run(tt.name, func(t *testing.T) {
			v := &VerifyJwtMiddleware{
				authenticationService: tt.fields.authenticationService,
			}
			if got := v.isUrlAllowed(tt.args.requestMethod, tt.args.url); got != tt.want {
				t.Errorf("isUrlAllowed() = %v, want %v", got, tt.want)
			}
		})
	}
}

func (ts *VerifyJWTMiddlewareTestSuite) TestVerifyJwtMiddleware_isTokenAllowed() {

	validUsedId := int64(123)
	type fields struct {
		authenticationService services.AuthenticationServiceInterface
	}
	type args struct {
		jwtToken string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		mock   func() error
		want   bool
		want1  *int64
	}{
		{
			name: "When the authorization service return not authorized given the token, then itu return false and nil",
			fields: fields{
				authenticationService: ts.authenticationService,
			},
			args: args{
				jwtToken: "dummy json web token",
			},
			want:  false,
			want1: nil,
			mock: func() error {

				ts.authenticationService.EXPECT().Authorize(gomock.Any()).Return(&services.AuthorizationResult{
					IsAuthorized: false,
					UserId:       0,
				}, nil)

				return nil
			},
		},

		{
			name: "When the authorization service return authorized given the token, then itu return true and userId",
			fields: fields{
				authenticationService: ts.authenticationService,
			},
			args: args{
				jwtToken: "dummy json web token",
			},
			want:  true,
			want1: &validUsedId,
			mock: func() error {

				ts.authenticationService.EXPECT().Authorize(gomock.Any()).Return(&services.AuthorizationResult{
					IsAuthorized: true,
					UserId:       123,
				}, nil)

				return nil
			},
		},
	}
	for _, tt := range tests {
		ts.T().Run(tt.name, func(t *testing.T) {
			v := &VerifyJwtMiddleware{
				authenticationService: tt.fields.authenticationService,
			}

			tt.mock()

			got, got1 := v.isTokenAllowed(tt.args.jwtToken)
			if got != tt.want {
				t.Errorf("isTokenAllowed() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("isTokenAllowed() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
