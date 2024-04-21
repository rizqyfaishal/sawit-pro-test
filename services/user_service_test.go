package services

import (
	"errors"
	"github.com/SawitProRecruitment/UserService/consts"
	"github.com/SawitProRecruitment/UserService/forms"
	"github.com/SawitProRecruitment/UserService/modules"
	"github.com/SawitProRecruitment/UserService/pojos"
	"github.com/SawitProRecruitment/UserService/repository"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"
	"reflect"
	"testing"
)

type UserServiceTestSuite struct {
	suite.Suite

	repository   *repository.MockUserRepositoryInterface
	passwordAuth *modules.MockPasswordAuthInterface

	MockController *gomock.Controller
}

func TestUserServiceTestSuite(t *testing.T) {
	suite.Run(t, new(UserServiceTestSuite))
}

func (ts *UserServiceTestSuite) SetupSuite() {

	mockCtrl := gomock.NewController(ts.T())

	ts.MockController = mockCtrl

	defer mockCtrl.Finish()

	ts.repository = repository.NewMockUserRepositoryInterface(mockCtrl)
	ts.passwordAuth = modules.NewMockPasswordAuthInterface(mockCtrl)
}

func (ts *UserServiceTestSuite) TestNewUserService() {
	type args struct {
		repository   repository.UserRepositoryInterface
		passwordAuth modules.PasswordAuthInterface
	}
	tests := []struct {
		name string
		args args
		mock func()
		want UserServiceInterface
	}{
		{
			name: "When instantiate correctly it will return implementation of UserService",
			args: args{
				repository:   ts.repository,
				passwordAuth: ts.passwordAuth,
			},
			mock: func() {

			},
			want: UserService{
				repository:   ts.repository,
				passwordAuth: ts.passwordAuth,
			},
		},
	}
	for _, tt := range tests {
		ts.T().Run(tt.name, func(t *testing.T) {
			if got := NewUserService(tt.args.repository, tt.args.passwordAuth); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUserService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func (ts *UserServiceTestSuite) TestUserService_GetById() {
	type fields struct {
		repository   repository.UserRepositoryInterface
		passwordAuth modules.PasswordAuthInterface
	}
	type args struct {
		userId int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *pojos.User
		mock    func()
		wantErr bool
	}{
		{
			name: "Given valid user id, but repository return error, then service will return error",
			fields: fields{
				repository:   ts.repository,
				passwordAuth: ts.passwordAuth,
			},
			args: args{
				userId: 123,
			},
			want: nil,
			mock: func() {
				ts.repository.EXPECT().GetById(gomock.Any(), gomock.Any()).Return(nil, errors.New("Unexpected error"))
			},
			wantErr: true,
		},

		{
			name: "Given valid user id, repository return no error, but the record actually not found, then service will return empty result",
			fields: fields{
				repository:   ts.repository,
				passwordAuth: ts.passwordAuth,
			},
			args: args{
				userId: 123,
			},
			want: nil,
			mock: func() {
				ts.repository.EXPECT().GetById(gomock.Any(), gomock.Any()).Return(nil, nil)
			},
			wantErr: false,
		},

		{
			name: "Given valid user id, repository return no error, but the record actually found, then service will return user result",
			fields: fields{
				repository:   ts.repository,
				passwordAuth: ts.passwordAuth,
			},
			args: args{
				userId: 123,
			},
			want: &pojos.User{
				Id:                123,
				PhoneNumber:       "+6285773801038",
				FullName:          "Rizqy Faishal Tanjung",
				LoginSuccessCount: 2,
			},
			mock: func() {
				ts.repository.EXPECT().GetById(gomock.Any(), gomock.Any()).Return(&repository.GetUserByIdOutput{
					Id:                123,
					PhoneNumber:       "+6285773801038",
					FullName:          "Rizqy Faishal Tanjung",
					LoginSuccessCount: 2,
				}, nil)
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		ts.T().Run(tt.name, func(t *testing.T) {
			tt.mock()
			u := UserService{
				repository:   tt.fields.repository,
				passwordAuth: tt.fields.passwordAuth,
			}
			got, err := u.GetById(tt.args.userId)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetById() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func (ts *UserServiceTestSuite) TestUserService_GetByPhoneNumber() {
	type fields struct {
		repository   repository.UserRepositoryInterface
		passwordAuth modules.PasswordAuthInterface
	}
	type args struct {
		phoneNumber string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *pojos.User
		wantErr bool
		mock    func()
	}{
		{
			name: "Given valid user id, but repository return error, then service will return error",
			fields: fields{
				repository:   ts.repository,
				passwordAuth: ts.passwordAuth,
			},
			args: args{
				phoneNumber: "+6285773801038",
			},
			want: nil,
			mock: func() {
				ts.repository.EXPECT().GetByPhoneNumberIncludePassword(gomock.Any(), gomock.Any()).Return(nil, errors.New("Unexpected error"))
			},
			wantErr: true,
		},

		{
			name: "Given valid user id, repository return no error, but the record actually not found, then service will return empty result",
			fields: fields{
				repository:   ts.repository,
				passwordAuth: ts.passwordAuth,
			},
			args: args{
				phoneNumber: "+6285773801038",
			},
			want: nil,
			mock: func() {
				ts.repository.EXPECT().GetByPhoneNumberIncludePassword(gomock.Any(), gomock.Any()).Return(nil, nil)
			},
			wantErr: false,
		},

		{
			name: "Given valid user id, repository return no error, but the record actually found, then service will return user result",
			fields: fields{
				repository:   ts.repository,
				passwordAuth: ts.passwordAuth,
			},
			args: args{
				phoneNumber: "+6285773801038",
			},
			want: &pojos.User{
				Id:                123,
				PhoneNumber:       "+6285773801038",
				FullName:          "Rizqy Faishal Tanjung",
				LoginSuccessCount: 2,
			},
			mock: func() {
				ts.repository.EXPECT().GetByPhoneNumberIncludePassword(gomock.Any(), gomock.Any()).Return(&repository.GetUserByPhoneNumberOutput{
					Id:                123,
					PhoneNumber:       "+6285773801038",
					FullName:          "Rizqy Faishal Tanjung",
					LoginSuccessCount: 2,
				}, nil)
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		ts.T().Run(tt.name, func(t *testing.T) {
			tt.mock()
			u := UserService{
				repository:   tt.fields.repository,
				passwordAuth: tt.fields.passwordAuth,
			}
			got, err := u.GetByPhoneNumber(tt.args.phoneNumber)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetByPhoneNumber() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetByPhoneNumber() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func (ts *UserServiceTestSuite) TestUserService_Register() {
	type fields struct {
		repository   repository.UserRepositoryInterface
		passwordAuth modules.PasswordAuthInterface
	}
	type args struct {
		form forms.UserRegisterForm
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *RegisterResult
		wantErr bool
		mock    func()
	}{
		{
			name: "When user register without phone number, password, full name, it will return validation error",
			fields: fields{
				repository:   ts.repository,
				passwordAuth: ts.passwordAuth,
			},
			args: args{
				form: forms.UserRegisterForm{},
			},
			want: &RegisterResult{
				ValidationErrors: map[string]string{
					"full_name": "Full name is required", "password": "Password is required", "phone_number": "Phone number is required",
				},
				HasValidationErrors: true,
			},
			wantErr: false,
			mock: func() {

			},
		},

		{
			name: "When user register without phone number, it will return validation error",
			fields: fields{
				repository:   ts.repository,
				passwordAuth: ts.passwordAuth,
			},
			args: args{
				form: forms.UserRegisterForm{
					FullName: "Rizqy Faishal Tanjung",
					Password: "Asdasd123#",
				},
			},
			want: &RegisterResult{
				ValidationErrors: map[string]string{
					"phone_number": "Phone number is required",
				},
				HasValidationErrors: true,
			},
			wantErr: false,
			mock: func() {

			},
		},

		{
			name: "When user register without full name,  it will return validation error",
			fields: fields{
				repository:   ts.repository,
				passwordAuth: ts.passwordAuth,
			},
			args: args{
				form: forms.UserRegisterForm{
					Password:    "Asdasd123#",
					PhoneNumber: "+62857738010",
				},
			},
			want: &RegisterResult{
				ValidationErrors: map[string]string{
					"full_name": "Full name is required",
				},
				HasValidationErrors: true,
			},
			wantErr: false,
			mock: func() {

			},
		},

		{
			name: "When user register but password full name did not contains capital character, it will return validation error",
			fields: fields{
				repository:   ts.repository,
				passwordAuth: ts.passwordAuth,
			},
			args: args{
				form: forms.UserRegisterForm{
					FullName:    "Rizqy Faishal Tanjung",
					Password:    "asdasd123#",
					PhoneNumber: "+62857738010",
				},
			},
			want: &RegisterResult{
				ValidationErrors: map[string]string{
					"password": "Password must contains at least 1 captial characters",
				},
				HasValidationErrors: true,
			},
			wantErr: false,
			mock: func() {

			},
		},

		{
			name: "When user register but password did not contains special character, it will return validation error",
			fields: fields{
				repository:   ts.repository,
				passwordAuth: ts.passwordAuth,
			},
			args: args{
				form: forms.UserRegisterForm{
					FullName:    "Rizqy Faishal Tanjung",
					Password:    "Asdasd123",
					PhoneNumber: "+62857738010",
				},
			},
			want: &RegisterResult{
				ValidationErrors: map[string]string{
					"password": "Password must contains at least 1 special characters",
				},
				HasValidationErrors: true,
			},
			wantErr: false,
			mock: func() {

			},
		},

		{
			name: "When user register but password it less than 6 characters, it will return validation error",
			fields: fields{
				repository:   ts.repository,
				passwordAuth: ts.passwordAuth,
			},
			args: args{
				form: forms.UserRegisterForm{
					FullName:    "Rizqy Faishal Tanjung",
					Password:    "Adasd",
					PhoneNumber: "+62857738010",
				},
			},
			want: &RegisterResult{
				ValidationErrors: map[string]string{
					"password": "Password must have minimum 6 characters long",
				},
				HasValidationErrors: true,
			},
			wantErr: false,
			mock: func() {

			},
		},

		{
			name: "When user register but password it more than than 64 characters, it will return validation error",
			fields: fields{
				repository:   ts.repository,
				passwordAuth: ts.passwordAuth,
			},
			args: args{
				form: forms.UserRegisterForm{
					FullName:    "Rizqy Faishal Tanjung",
					Password:    "AdasdAdasdAdasdAdasdAdasdAdasdAdasdAdasdAdasdAdasdAdasdAdasdAdasdAdasdAdasdAdasdAdasdAdasdAdasdAdasd",
					PhoneNumber: "+62857738010",
				},
			},
			want: &RegisterResult{
				ValidationErrors: map[string]string{
					"password": "Password must have maximum 64 characters long",
				},
				HasValidationErrors: true,
			},
			wantErr: false,
			mock: func() {

			},
		},

		{
			name: "When user register but full name it more than than 60 characters, it will return validation error",
			fields: fields{
				repository:   ts.repository,
				passwordAuth: ts.passwordAuth,
			},
			args: args{
				form: forms.UserRegisterForm{
					FullName:    "Rizqy Faishal Tanjung Rizqy Faishal Tanjung Rizqy Faishal Tanjung Rizqy Faishal Tanjung",
					Password:    "Asdasd12#",
					PhoneNumber: "+62857738010",
				},
			},
			want: &RegisterResult{
				ValidationErrors: map[string]string{
					"full_name": "Full name must have maximum 60 characters long",
				},
				HasValidationErrors: true,
			},
			wantErr: false,
			mock: func() {

			},
		},

		{
			name: "When user register but full name it less than than 3 characters, it will return validation error",
			fields: fields{
				repository:   ts.repository,
				passwordAuth: ts.passwordAuth,
			},
			args: args{
				form: forms.UserRegisterForm{
					FullName:    "Ri",
					Password:    "Asdasd12#",
					PhoneNumber: "+62857738010",
				},
			},
			want: &RegisterResult{
				ValidationErrors: map[string]string{
					"full_name": "Full name must have minimum 3 characters long",
				},
				HasValidationErrors: true,
			},
			wantErr: false,
			mock: func() {

			},
		},

		{
			name: "When user register but phone number it less than than 3 characters, it will return validation error",
			fields: fields{
				repository:   ts.repository,
				passwordAuth: ts.passwordAuth,
			},
			args: args{
				form: forms.UserRegisterForm{
					FullName:    "Rizqy Faishal Tanjung",
					Password:    "Asdasd12#",
					PhoneNumber: "+6",
				},
			},
			want: &RegisterResult{
				ValidationErrors: map[string]string{
					"phone_number": "Phone number must have minimum 10 characters long",
				},
				HasValidationErrors: true,
			},
			wantErr: false,
			mock: func() {

			},
		},

		{
			name: "When user register but phone number it more than 13 characters, it will return validation error",
			fields: fields{
				repository:   ts.repository,
				passwordAuth: ts.passwordAuth,
			},
			args: args{
				form: forms.UserRegisterForm{
					FullName:    "Rizqy Faishal Tanjung",
					Password:    "Asdasd12#",
					PhoneNumber: "+6232424242424424",
				},
			},
			want: &RegisterResult{
				ValidationErrors: map[string]string{
					"phone_number": "Phone number must have maximum 13 characters long",
				},
				HasValidationErrors: true,
			},
			wantErr: false,
			mock: func() {

			},
		},

		{
			name: "When user register but phone number did not starts with +62, it will return validation error",
			fields: fields{
				repository:   ts.repository,
				passwordAuth: ts.passwordAuth,
			},
			args: args{
				form: forms.UserRegisterForm{
					FullName:    "Rizqy Faishal Tanjung",
					Password:    "Asdasd12#",
					PhoneNumber: "+68242424424",
				},
			},
			want: &RegisterResult{
				ValidationErrors: map[string]string{
					"phone_number": "Phone number must starts with +62",
				},
				HasValidationErrors: true,
			},
			wantErr: false,
			mock: func() {

			},
		},

		{
			name: "When user register but the repo return error when get user by phone number, it will return error",
			fields: fields{
				repository:   ts.repository,
				passwordAuth: ts.passwordAuth,
			},
			args: args{
				form: forms.UserRegisterForm{
					FullName:    "Rizqy Faishal Tanjung",
					Password:    "Asdasd12#",
					PhoneNumber: "+62242424424",
				},
			},
			want:    nil,
			wantErr: true,
			mock: func() {
				ts.repository.EXPECT().GetByPhoneNumberIncludePassword(gomock.Any(), gomock.Any()).Return(nil, errors.New("Unepxected error"))
			},
		},

		{
			name: "When user register but the repo return user found when get user by phone number, it will return validatioin error",
			fields: fields{
				repository:   ts.repository,
				passwordAuth: ts.passwordAuth,
			},
			args: args{
				form: forms.UserRegisterForm{
					FullName:    "Rizqy Faishal Tanjung",
					Password:    "Asdasd12#",
					PhoneNumber: "+62242424424",
				},
			},
			want: &RegisterResult{
				ValidationErrors: map[string]string{
					"phone_number": "Phone number +62242424424 is unavailable for registering new user",
				},
				HasValidationErrors: true,
			},
			wantErr: false,
			mock: func() {
				ts.repository.EXPECT().GetByPhoneNumberIncludePassword(gomock.Any(), gomock.Any()).Return(&repository.GetUserByPhoneNumberOutput{
					Id:                123,
					PhoneNumber:       "+62242424424",
					FullName:          "Rizqy Faishal Tanjung",
					LoginSuccessCount: 1,
					Password:          "Asdasd123%",
				}, nil)
			},
		},

		{
			name: "When user register hashing password got error, it will return error",
			fields: fields{
				repository:   ts.repository,
				passwordAuth: ts.passwordAuth,
			},
			args: args{
				form: forms.UserRegisterForm{
					FullName:    "Rizqy Faishal Tanjung",
					Password:    "Asdasd12#",
					PhoneNumber: "+62242424424",
				},
			},
			want:    nil,
			wantErr: true,
			mock: func() {
				ts.repository.EXPECT().GetByPhoneNumberIncludePassword(gomock.Any(), gomock.Any()).Return(nil, nil)
				ts.passwordAuth.EXPECT().GenerateHashedPassword(gomock.Any()).Return("", errors.New("unexpected error"))
			},
		},

		{
			name: "When user register repo got error when inserting record, it will return error",
			fields: fields{
				repository:   ts.repository,
				passwordAuth: ts.passwordAuth,
			},
			args: args{
				form: forms.UserRegisterForm{
					FullName:    "Rizqy Faishal Tanjung",
					Password:    "Asdasd12#",
					PhoneNumber: "+62242424424",
				},
			},
			want:    nil,
			wantErr: true,
			mock: func() {
				ts.repository.EXPECT().GetByPhoneNumberIncludePassword(gomock.Any(), gomock.Any()).Return(nil, nil)
				ts.passwordAuth.EXPECT().GenerateHashedPassword(gomock.Any()).Return("asdasdsdsada", nil)
				ts.repository.EXPECT().Insert(gomock.Any(), gomock.Any()).Return(nil, errors.New("unexpected error"))
			},
		},

		{
			name: "When user register repo got no error, it will return register result",
			fields: fields{
				repository:   ts.repository,
				passwordAuth: ts.passwordAuth,
			},
			args: args{
				form: forms.UserRegisterForm{
					FullName:    "Rizqy Faishal Tanjung",
					Password:    "Asdasd12#",
					PhoneNumber: "+62242424424",
				},
			},
			want: &RegisterResult{
				User: pojos.User{
					Id:          123,
					FullName:    "Rizqy Faishal Tanjung",
					PhoneNumber: "+62242424424",
				},
				HasValidationErrors: false,
			},
			wantErr: false,
			mock: func() {
				ts.repository.EXPECT().GetByPhoneNumberIncludePassword(gomock.Any(), gomock.Any()).Return(nil, nil)
				ts.passwordAuth.EXPECT().GenerateHashedPassword(gomock.Any()).Return("asdasdsdsada", nil)
				ts.repository.EXPECT().Insert(gomock.Any(), gomock.Any()).Return(&repository.InsertUserOutput{
					Id: 123,
				}, nil)

				ts.repository.EXPECT().GetById(gomock.Any(), gomock.Any()).Return(&repository.GetUserByIdOutput{
					Id:                123,
					PhoneNumber:       "+62242424424",
					FullName:          "Rizqy Faishal Tanjung",
					LoginSuccessCount: 0,
				}, nil)

			},
		},
	}
	for _, tt := range tests {
		ts.T().Run(tt.name, func(t *testing.T) {
			tt.mock()
			u := UserService{
				repository:   tt.fields.repository,
				passwordAuth: tt.fields.passwordAuth,
			}
			got, err := u.Register(tt.args.form)
			if (err != nil) != tt.wantErr {
				t.Errorf("Register() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Register() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func (ts *UserServiceTestSuite) TestUserService_Update() {
	type fields struct {
		repository   repository.UserRepositoryInterface
		passwordAuth modules.PasswordAuthInterface
	}
	type args struct {
		userId int64
		form   forms.UserUpdateForm
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *UpdateResult
		wantErr bool
		mock    func()
	}{
		{
			name: "When the phone number and full name is empty, then return validation errors",
			fields: fields{
				repository:   ts.repository,
				passwordAuth: ts.passwordAuth,
			},
			args: args{
				userId: 123,
				form:   forms.UserUpdateForm{},
			},
			want: &UpdateResult{
				User:                pojos.User{},
				HasValidationErrors: true,
				ValidationErrors: map[string]string{
					consts.GeneralErrorMessageKey: "Phone number or Full Name is required",
				},
			},
			wantErr: false,
			mock: func() {

			},
		},

		{
			name: "When the fullname is only 2 characters long, then return validation errors",
			fields: fields{
				repository:   ts.repository,
				passwordAuth: ts.passwordAuth,
			},
			args: args{
				userId: 123,
				form: forms.UserUpdateForm{
					FullName: "as",
				},
			},
			want: &UpdateResult{
				User:                pojos.User{},
				HasValidationErrors: true,
				ValidationErrors: map[string]string{
					"full_name": "Full name must have minimum 3 characters long",
				},
			},
			wantErr: false,
			mock: func() {

			},
		},

		{
			name: "When the fullname is at 82 characters long, then return validation errors",
			fields: fields{
				repository:   ts.repository,
				passwordAuth: ts.passwordAuth,
			},
			args: args{
				userId: 123,
				form: forms.UserUpdateForm{
					FullName: "Character Counter is a free online character count calculator that's simple to use",
				},
			},
			want: &UpdateResult{
				User:                pojos.User{},
				HasValidationErrors: true,
				ValidationErrors: map[string]string{
					"full_name": "Full name must have maximum 60 characters long",
				},
			},
			wantErr: false,
			mock: func() {

			},
		},

		{
			name: "When the phone number is at more than 13 characters long, then return validation errors",
			fields: fields{
				repository:   ts.repository,
				passwordAuth: ts.passwordAuth,
			},
			args: args{
				userId: 123,
				form: forms.UserUpdateForm{
					PhoneNumber: "+6285732932932792",
				},
			},
			want: &UpdateResult{
				User:                pojos.User{},
				HasValidationErrors: true,
				ValidationErrors: map[string]string{
					"phone_number": "Phone number must have maximum 13 characters long",
				},
			},
			wantErr: false,
			mock: func() {

			},
		},

		{
			name: "When the phone number is at less than 10 characters long, then return validation errors",
			fields: fields{
				repository:   ts.repository,
				passwordAuth: ts.passwordAuth,
			},
			args: args{
				userId: 123,
				form: forms.UserUpdateForm{
					PhoneNumber: "+62857",
				},
			},
			want: &UpdateResult{
				User:                pojos.User{},
				HasValidationErrors: true,
				ValidationErrors: map[string]string{
					"phone_number": "Phone number must have minimum 10 characters long",
				},
			},
			wantErr: false,
			mock: func() {

			},
		},

		{
			name: "When the phone number is not started with +62, then return validation errors",
			fields: fields{
				repository:   ts.repository,
				passwordAuth: ts.passwordAuth,
			},
			args: args{
				userId: 123,
				form: forms.UserUpdateForm{
					PhoneNumber: "+82857382923",
				},
			},
			want: &UpdateResult{
				User:                pojos.User{},
				HasValidationErrors: true,
				ValidationErrors: map[string]string{
					"phone_number": "Phone number must starts with +62",
				},
			},
			wantErr: false,
			mock: func() {

			},
		},

		{
			name: "When the form is valid, but the user returned by repo is not found, then return error",
			fields: fields{
				repository:   ts.repository,
				passwordAuth: ts.passwordAuth,
			},
			args: args{
				userId: 123,
				form: forms.UserUpdateForm{
					PhoneNumber: "+62857382923",
				},
			},
			want:    nil,
			wantErr: true,
			mock: func() {
				ts.repository.EXPECT().GetById(gomock.Any(), gomock.Any()).Return(nil, nil)
			},
		},

		{
			name: "When the form is valid, but the user returned by repo error, then return error",
			fields: fields{
				repository:   ts.repository,
				passwordAuth: ts.passwordAuth,
			},
			args: args{
				userId: 123,
				form: forms.UserUpdateForm{
					PhoneNumber: "+62857382923",
				},
			},
			want:    nil,
			wantErr: true,
			mock: func() {
				ts.repository.EXPECT().GetById(gomock.Any(), gomock.Any()).Return(nil, errors.New("unexpected error"))
			},
		},

		{
			name: "When the form is valid, but the user returned by repo error when update, then return error",
			fields: fields{
				repository:   ts.repository,
				passwordAuth: ts.passwordAuth,
			},
			args: args{
				userId: 123,
				form: forms.UserUpdateForm{
					PhoneNumber: "+62857382923",
				},
			},
			want:    nil,
			wantErr: true,
			mock: func() {
				ts.repository.EXPECT().GetById(gomock.Any(), gomock.Any()).Return(&repository.GetUserByIdOutput{
					Id:                123,
					PhoneNumber:       "62857382923",
					FullName:          "Rizqy",
					LoginSuccessCount: 0,
				}, nil)
				ts.repository.EXPECT().Update(gomock.Any(), gomock.Any()).Return(nil, errors.New("unexpected error"))
			},
		},

		{
			name: "When the form is valid, successfully update, then return updated user",
			fields: fields{
				repository:   ts.repository,
				passwordAuth: ts.passwordAuth,
			},
			args: args{
				userId: 123,
				form: forms.UserUpdateForm{
					PhoneNumber: "+62857382923",
				},
			},
			want: &UpdateResult{
				User: pojos.User{
					Id:                123,
					PhoneNumber:       "62857382923",
					FullName:          "Rizqy",
					LoginSuccessCount: 0,
				},
				HasValidationErrors: false,
				ValidationErrors:    map[string]string{},
			},
			wantErr: false,
			mock: func() {
				ts.repository.EXPECT().GetById(gomock.Any(), gomock.Any()).AnyTimes().Return(&repository.GetUserByIdOutput{
					Id:                123,
					PhoneNumber:       "62857382923",
					FullName:          "Rizqy",
					LoginSuccessCount: 0,
				}, nil)
				ts.repository.EXPECT().Update(gomock.Any(), gomock.Any()).Return(&repository.UpdateUserOutput{
					IsSuccessUpdate: true,
				}, nil)
			},
		},
	}
	for _, tt := range tests {
		ts.T().Run(tt.name, func(t *testing.T) {
			tt.mock()
			u := UserService{
				repository:   tt.fields.repository,
				passwordAuth: tt.fields.passwordAuth,
			}
			got, err := u.Update(tt.args.userId, tt.args.form)
			if (err != nil) != tt.wantErr {
				t.Errorf("Update() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Update() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserService_buildRegisterUserResponse(t *testing.T) {
	type fields struct {
		repository   repository.UserRepositoryInterface
		passwordAuth modules.PasswordAuthInterface
	}
	type args struct {
		output repository.GetUserByIdOutput
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   pojos.User
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := UserService{
				repository:   tt.fields.repository,
				passwordAuth: tt.fields.passwordAuth,
			}
			if got := u.buildRegisterUserResponse(tt.args.output); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("buildRegisterUserResponse() = %v, want %v", got, tt.want)
			}
		})
	}
}
