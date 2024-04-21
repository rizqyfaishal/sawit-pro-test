package services

import (
	"context"
	"errors"
	"fmt"
	"github.com/SawitProRecruitment/UserService/consts"
	"github.com/SawitProRecruitment/UserService/forms"
	"github.com/SawitProRecruitment/UserService/modules"
	"github.com/SawitProRecruitment/UserService/pojos"
	"github.com/SawitProRecruitment/UserService/repository"
	"github.com/SawitProRecruitment/UserService/utils"
	"github.com/SawitProRecruitment/UserService/validators"
	"github.com/go-playground/validator/v10"
)

const GenerateHashedPasswordCost int = 10

type UserService struct {
	repository   repository.UserRepositoryInterface
	passwordAuth modules.PasswordAuthInterface
}

func (u UserService) Register(form forms.UserRegisterForm) (*RegisterResult, error) {

	ctx := context.Background()

	result := &RegisterResult{
		ValidationErrors:    nil,
		HasValidationErrors: false,
	}

	validate := validator.New(validator.WithRequiredStructEnabled())
	err := validate.RegisterValidation(validators.AtLeastXCapitalCharValidationTag, validators.AtLeastXCapitalCharValidation)

	if err != nil {
		return nil, err
	}

	err = validate.RegisterValidation(validators.AtLeastXSpecialCharValidationTag, validators.AtLeastXSpecialCharValidation)

	if err != nil {
		return nil, err
	}

	err = validate.Struct(form)

	if err != nil {

		var validationErrors validator.ValidationErrors

		errors.As(err, &validationErrors)

		validationErrorMessages := utils.CollectValidationErrorMessages(form, validationErrors)

		result.HasValidationErrors = true
		result.ValidationErrors = validationErrorMessages

		return result, nil
	}

	getByPhoneNumberInput := repository.GetUserByPhoneNumberInput{
		PhoneNumber: form.PhoneNumber,
	}

	existedUser, err := u.repository.GetByPhoneNumberIncludePassword(ctx, getByPhoneNumberInput)

	if err != nil {
		return nil, err
	}

	if existedUser != nil {
		result.HasValidationErrors = true
		result.ValidationErrors = map[string]string{
			"phone_number": fmt.Sprintf("Phone number %s is unavailable for registering new user", existedUser.PhoneNumber),
		}

		return result, nil
	}

	hashedPassword, err := u.passwordAuth.GenerateHashedPassword(form.Password)

	if err != nil {
		return nil, err
	}

	hashedPasswordString := string(hashedPassword)

	insertUserInput := repository.InsertUserInput{
		PhoneNumber: form.PhoneNumber,
		FullName:    form.FullName,
		Password:    hashedPasswordString,
	}

	output, err := u.repository.Insert(ctx, insertUserInput)

	if err != nil {
		return nil, err
	}

	getUserByIdInput := repository.GetUserByIdInput{
		Id: output.Id,
	}

	getUserOutPut, err := u.repository.GetById(ctx, getUserByIdInput)

	if err != nil {
		return nil, err
	}

	if getUserOutPut == nil {
		return nil, errors.New("Unexpected error. After insert return nil user.")
	}

	result.User = u.buildRegisterUserResponse(*getUserOutPut)

	return result, nil
}

func (u UserService) Update(userId int64, form forms.UserUpdateForm) (*UpdateResult, error) {

	ctx := context.Background()
	result := UpdateResult{}

	if form.PhoneNumber == "" && form.FullName == "" {

		result.HasValidationErrors = true
		result.ValidationErrors = map[string]string{
			consts.GeneralErrorMessageKey: "Phone number or Full Name is required",
		}

		return &result, nil
	}

	validate := validator.New(validator.WithRequiredStructEnabled())
	err := validate.Struct(form)

	if err != nil {

		var validationErrors validator.ValidationErrors

		errors.As(err, &validationErrors)

		validationErrorMessages := utils.CollectValidationErrorMessages(form, validationErrors)

		result.HasValidationErrors = true
		result.ValidationErrors = validationErrorMessages

		return &result, nil
	}

	user, err := u.GetById(userId)

	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, errors.New("User not found")
	}

	if utils.StringIsEmpty(form.PhoneNumber) == false {
		user.PhoneNumber = form.PhoneNumber
	}

	if utils.StringIsEmpty(form.FullName) == false {
		user.FullName = form.FullName
	}

	updateUserInput := repository.UpdateUserInput{
		Id:                user.Id,
		PhoneNumber:       user.PhoneNumber,
		FullName:          user.FullName,
		LoginSuccessCount: user.LoginSuccessCount,
	}

	updateOutput, err := u.repository.Update(ctx, updateUserInput)

	if err != nil {
		return nil, err
	}

	if updateOutput.IsSuccessUpdate == false {
		return nil, errors.New("Failed to update the record")
	}

	user, err = u.GetById(userId)

	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, errors.New("User not found")
	}

	result.HasValidationErrors = false
	result.ValidationErrors = map[string]string{}
	result.User = *user

	return &result, nil
}

func (u UserService) GetById(userId int64) (*pojos.User, error) {

	ctx := context.Background()

	getUserByIdInput := repository.GetUserByIdInput{
		Id: userId,
	}

	output, err := u.repository.GetById(ctx, getUserByIdInput)

	if err != nil {
		return nil, err
	}

	if output == nil {
		return nil, nil
	}

	user := &pojos.User{
		Id:                output.Id,
		PhoneNumber:       output.PhoneNumber,
		FullName:          output.FullName,
		LoginSuccessCount: output.LoginSuccessCount,
		CreatedAt:         output.CreatedAt,
		UpdatedAt:         output.UpdatedAt,
	}

	return user, nil
}

func (u UserService) buildRegisterUserResponse(output repository.GetUserByIdOutput) pojos.User {

	return pojos.User{
		Id:                output.Id,
		PhoneNumber:       output.PhoneNumber,
		FullName:          output.FullName,
		LoginSuccessCount: output.LoginSuccessCount,
		CreatedAt:         output.CreatedAt,
		UpdatedAt:         output.UpdatedAt,
	}
}

func (u UserService) GetByPhoneNumber(phoneNumber string) (*pojos.User, error) {

	ctx := context.Background()

	getByPhoneNumberInput := repository.GetUserByPhoneNumberInput{
		PhoneNumber: phoneNumber,
	}

	output, err := u.repository.GetByPhoneNumberIncludePassword(ctx, getByPhoneNumberInput)

	if err != nil {
		return nil, err
	}

	if output == nil {
		return nil, nil
	}

	user := &pojos.User{
		Id:                output.Id,
		PhoneNumber:       output.PhoneNumber,
		FullName:          output.FullName,
		LoginSuccessCount: output.LoginSuccessCount,
		CreatedAt:         output.CreatedAt,
		UpdatedAt:         output.UpdatedAt,
	}

	return user, nil
}

func NewUserService(repository repository.UserRepositoryInterface, passwordAuth modules.PasswordAuthInterface) UserServiceInterface {

	return UserService{
		repository:   repository,
		passwordAuth: passwordAuth,
	}
}
