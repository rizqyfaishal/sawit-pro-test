package services

import (
	"errors"
	"github.com/SawitProRecruitment/UserService/forms"
	"github.com/SawitProRecruitment/UserService/modules"
	"github.com/SawitProRecruitment/UserService/repository"
	"github.com/golang-jwt/jwt/v5"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"
	"os"
	"reflect"
	"testing"
	"time"
)

type AuthenticationServiceTestSuite struct {
	suite.Suite

	repository   *repository.MockUserRepositoryInterface
	passwordAuth *modules.MockPasswordAuthInterface
	jwtAuth      *modules.MockJsonWebTokenUtilInterface

	MockController *gomock.Controller
}

func TestAuthenticationServiceTestSuite(t *testing.T) {
	suite.Run(t, new(AuthenticationServiceTestSuite))
}

func (ts *AuthenticationServiceTestSuite) SetupSuite() {

	mockCtrl := gomock.NewController(ts.T())

	ts.MockController = mockCtrl

	defer mockCtrl.Finish()

	ts.repository = repository.NewMockUserRepositoryInterface(mockCtrl)
	ts.passwordAuth = modules.NewMockPasswordAuthInterface(mockCtrl)
	ts.jwtAuth = modules.NewMockJsonWebTokenUtilInterface(mockCtrl)

}

func (ts *AuthenticationServiceTestSuite) TestAuthenticationService_Authenticate() {
	os.Setenv("LOGIN_EXPIRATION_DURATION", "24h")

	type fields struct {
		repository   repository.UserRepositoryInterface
		passwordAuth modules.PasswordAuthInterface
		jwtAuth      modules.JsonWebTokenUtilInterface
	}
	type args struct {
		form forms.UserLoginForm
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *AuthenticationResult
		wantErr bool
		mock    func()
	}{
		{
			name: "When the form is invalid, the phone number and password is not found, then return validation errors",
			fields: fields{
				repository:   ts.repository,
				passwordAuth: ts.passwordAuth,
				jwtAuth:      ts.jwtAuth,
			},
			args: args{
				form: forms.UserLoginForm{},
			},
			want: &AuthenticationResult{
				IsSuccess:           false,
				IsUserNotFound:      false,
				HasValidationErrors: true,
				ValidationErrors: map[string]string{
					"password":     "Password is required",
					"phone_number": "Phone number is required",
				},
				Credential: nil,
			},
			mock: func() {

			},
			wantErr: false,
		},

		{
			name: "When the form is invalid, the password is empty, then return validation errors",
			fields: fields{
				repository:   ts.repository,
				passwordAuth: ts.passwordAuth,
				jwtAuth:      ts.jwtAuth,
			},
			args: args{
				form: forms.UserLoginForm{
					PhoneNumber: "+628577480121",
				},
			},
			want: &AuthenticationResult{
				IsSuccess:           false,
				IsUserNotFound:      false,
				HasValidationErrors: true,
				ValidationErrors: map[string]string{
					"password": "Password is required",
				},
				Credential: nil,
			},
			mock: func() {

			},
			wantErr: false,
		},

		{
			name: "When the form is invalid, the phone number is empty, then return validation errors",
			fields: fields{
				repository:   ts.repository,
				passwordAuth: ts.passwordAuth,
				jwtAuth:      ts.jwtAuth,
			},
			args: args{
				form: forms.UserLoginForm{
					Password: "asdasd123",
				},
			},
			want: &AuthenticationResult{
				IsSuccess:           false,
				IsUserNotFound:      false,
				HasValidationErrors: true,
				ValidationErrors: map[string]string{
					"phone_number": "Phone number is required",
				},
				Credential: nil,
			},
			mock: func() {

			},
			wantErr: false,
		},

		{
			name: "When the form is valid, the phone number and password are not empty, but the user is not found, then return validation errors",
			fields: fields{
				repository:   ts.repository,
				passwordAuth: ts.passwordAuth,
				jwtAuth:      ts.jwtAuth,
			},
			args: args{
				form: forms.UserLoginForm{
					Password:    "asdasd123",
					PhoneNumber: "+628329328932",
				},
			},
			want: &AuthenticationResult{
				IsSuccess:      false,
				IsUserNotFound: true,
			},
			mock: func() {
				ts.repository.EXPECT().GetByPhoneNumberIncludePassword(gomock.Any(), gomock.Any()).Return(nil, nil)
			},
			wantErr: false,
		},

		{
			name: "When the form is valid, the phone number and password are not empty, but the repository return error, then return errors",
			fields: fields{
				repository:   ts.repository,
				passwordAuth: ts.passwordAuth,
				jwtAuth:      ts.jwtAuth,
			},
			args: args{
				form: forms.UserLoginForm{
					Password:    "asdasd123",
					PhoneNumber: "+628329328932",
				},
			},
			want: nil,
			mock: func() {
				ts.repository.EXPECT().GetByPhoneNumberIncludePassword(gomock.Any(), gomock.Any()).Return(nil, errors.New("unexpected error"))
			},
			wantErr: true,
		},

		{
			name: "When the form is valid, the phone number and password are not empty, but the password is invalid by the bcrypt, then return is not authenticated",
			fields: fields{
				repository:   ts.repository,
				passwordAuth: ts.passwordAuth,
				jwtAuth:      ts.jwtAuth,
			},
			args: args{
				form: forms.UserLoginForm{
					Password:    "asdasd123",
					PhoneNumber: "+628329328932",
				},
			},
			want: &AuthenticationResult{
				IsSuccess:           false,
				IsUserNotFound:      false,
				HasValidationErrors: false,
			},
			mock: func() {
				ts.repository.EXPECT().GetByPhoneNumberIncludePassword(gomock.Any(), gomock.Any()).Return(&repository.GetUserByPhoneNumberOutput{
					Id:                123,
					PhoneNumber:       "+628329328932",
					FullName:          "Rizqy Faishal",
					LoginSuccessCount: 0,
					Password:          "asdasd123",
					CreatedAt:         time.Time{},
					UpdatedAt:         time.Time{},
				}, nil)

				ts.passwordAuth.EXPECT().CompareHashedPassword(gomock.Any(), gomock.Any()).Return(false, errors.New("password not match"))
			},
			wantErr: false,
		},

		{
			name: "When the form is valid, the phone number and password are not empty, but the password is valid, but jwt generator is failed, then return errors",
			fields: fields{
				repository:   ts.repository,
				passwordAuth: ts.passwordAuth,
				jwtAuth:      ts.jwtAuth,
			},
			args: args{
				form: forms.UserLoginForm{
					Password:    "asdasd123",
					PhoneNumber: "+628329328932",
				},
			},
			want: nil,
			mock: func() {
				ts.repository.EXPECT().GetByPhoneNumberIncludePassword(gomock.Any(), gomock.Any()).Return(&repository.GetUserByPhoneNumberOutput{
					Id:                123,
					PhoneNumber:       "+628329328932",
					FullName:          "Rizqy Faishal",
					LoginSuccessCount: 0,
					Password:          "asdasd123",
					CreatedAt:         time.Time{},
					UpdatedAt:         time.Time{},
				}, nil)

				ts.passwordAuth.EXPECT().CompareHashedPassword(gomock.Any(), gomock.Any()).Return(true, nil)
				ts.jwtAuth.EXPECT().GenerateJwt(gomock.Any()).Return(nil, errors.New("token not generated"))

			},
			wantErr: true,
		},

		{
			name: "When the form is valid, the phone number and password are not empty, but the password is valid, but has repo error, then return errors",
			fields: fields{
				repository:   ts.repository,
				passwordAuth: ts.passwordAuth,
				jwtAuth:      ts.jwtAuth,
			},
			args: args{
				form: forms.UserLoginForm{
					Password:    "asdasd123",
					PhoneNumber: "+628329328932",
				},
			},
			want: nil,
			mock: func() {
				ts.repository.EXPECT().GetByPhoneNumberIncludePassword(gomock.Any(), gomock.Any()).Return(&repository.GetUserByPhoneNumberOutput{
					Id:                123,
					PhoneNumber:       "+628329328932",
					FullName:          "Rizqy Faishal",
					LoginSuccessCount: 0,
					Password:          "asdasd123",
					CreatedAt:         time.Time{},
					UpdatedAt:         time.Time{},
				}, nil)

				ts.passwordAuth.EXPECT().CompareHashedPassword(gomock.Any(), gomock.Any()).Return(true, nil)
				token := "jwt token"
				ts.jwtAuth.EXPECT().GenerateJwt(gomock.Any()).Return(&token, nil)
				ts.repository.EXPECT().Update(gomock.Any(), gomock.Any()).Return(nil, errors.New("Error at update"))
			},
			wantErr: true,
		},

		{
			name: "Positive case, When the form is valid, the phone number and password are not empty, but the password is valid, return success authentication result",
			fields: fields{
				repository:   ts.repository,
				passwordAuth: ts.passwordAuth,
				jwtAuth:      ts.jwtAuth,
			},
			args: args{
				form: forms.UserLoginForm{
					Password:    "asdasd123",
					PhoneNumber: "+628329328932",
				},
			},
			want: &AuthenticationResult{
				IsSuccess:           true,
				IsUserNotFound:      false,
				HasValidationErrors: false,
				ValidationErrors:    map[string]string{},
				Credential: &AuthenticationCredential{
					Token: "jwt token",
				},
			},
			mock: func() {
				ts.repository.EXPECT().GetByPhoneNumberIncludePassword(gomock.Any(), gomock.Any()).Return(&repository.GetUserByPhoneNumberOutput{
					Id:                123,
					PhoneNumber:       "+628329328932",
					FullName:          "Rizqy Faishal",
					LoginSuccessCount: 0,
					Password:          "asdasd123",
					CreatedAt:         time.Time{},
					UpdatedAt:         time.Time{},
				}, nil)

				ts.passwordAuth.EXPECT().CompareHashedPassword(gomock.Any(), gomock.Any()).Return(true, nil)
				token := "jwt token"
				ts.jwtAuth.EXPECT().GenerateJwt(gomock.Any()).Return(&token, nil)
				ts.repository.EXPECT().Update(gomock.Any(), gomock.Any()).Return(&repository.UpdateUserOutput{
					IsSuccessUpdate: true,
				}, nil)
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		ts.T().Run(tt.name, func(t *testing.T) {
			tt.mock()
			a := AuthenticationService{
				repository:   tt.fields.repository,
				passwordAuth: tt.fields.passwordAuth,
				jwtAuth:      tt.fields.jwtAuth,
			}
			got, err := a.Authenticate(tt.args.form)
			if (err != nil) != tt.wantErr {
				t.Errorf("Authenticate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Authenticate() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func (ts *AuthenticationServiceTestSuite) TestAuthenticationService_Authorize() {
	type fields struct {
		repository   repository.UserRepositoryInterface
		passwordAuth modules.PasswordAuthInterface
		jwtAuth      modules.JsonWebTokenUtilInterface
	}
	type args struct {
		tokenString string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *AuthorizationResult
		mock    func()
		wantErr bool
	}{
		{
			name: "When the token is invalid and should be unauthorized, then return error",
			fields: fields{
				repository:   ts.repository,
				passwordAuth: ts.passwordAuth,
				jwtAuth:      ts.jwtAuth,
			},
			args: args{
				tokenString: "a json web token",
			},
			want:    nil,
			wantErr: true,
			mock: func() {
				ts.jwtAuth.EXPECT().VerifyJwt(gomock.Any()).Return(nil, errors.New("invalid token"))
			},
		},

		{
			name: "When the token is valid and should be authorized, then return authorize result",
			fields: fields{
				repository:   ts.repository,
				passwordAuth: ts.passwordAuth,
				jwtAuth:      ts.jwtAuth,
			},
			args: args{
				tokenString: "a json web token",
			},
			want: &AuthorizationResult{
				IsAuthorized: true,
				UserId:       123,
			},
			wantErr: false,
			mock: func() {
				ts.jwtAuth.EXPECT().VerifyJwt(gomock.Any()).Return(&modules.CustomClaims{
					RegisteredClaims: jwt.RegisteredClaims{},
					UserId:           123,
				}, nil)
			},
		},
	}
	for _, tt := range tests {
		ts.T().Run(tt.name, func(t *testing.T) {
			tt.mock()

			a := AuthenticationService{
				repository:   tt.fields.repository,
				passwordAuth: tt.fields.passwordAuth,
				jwtAuth:      tt.fields.jwtAuth,
			}
			got, err := a.Authorize(tt.args.tokenString)
			if (err != nil) != tt.wantErr {
				t.Errorf("Authorize() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Authorize() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func (ts *AuthenticationServiceTestSuite) TestNewAuthenticationService() {
	type args struct {
		repository   repository.UserRepositoryInterface
		passwordAuth modules.PasswordAuthInterface
		jwtAuth      modules.JsonWebTokenUtilInterface
	}
	tests := []struct {
		name string
		args args
		want AuthenticationServiceInterface
	}{
		{
			name: "When given valid dependencies module, it will return authentication service",
			args: args{
				repository:   ts.repository,
				passwordAuth: ts.passwordAuth,
				jwtAuth:      ts.jwtAuth,
			},
			want: AuthenticationService{
				repository:   ts.repository,
				passwordAuth: ts.passwordAuth,
				jwtAuth:      ts.jwtAuth,
			},
		},
	}
	for _, tt := range tests {
		ts.T().Run(tt.name, func(t *testing.T) {
			if got := NewAuthenticationService(tt.args.repository, tt.args.passwordAuth, tt.args.jwtAuth); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewAuthenticationService() = %v, want %v", got, tt.want)
			}
		})
	}
}
