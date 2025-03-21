package usecase

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/saipulmuiz/krplus/models"
	"github.com/saipulmuiz/krplus/service/helper"
	"github.com/saipulmuiz/krplus/service/repository/mocks"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func Test_UserUsecase_Register(t *testing.T) {
	type testCase struct {
		name             string
		wantError        bool
		expectedResponse *models.User
		request          *models.RegisterUser
		onRegister       func(mock *mocks.MockUserRepository)
		onGetUserByEmail func(mock *mocks.MockUserRepository)
		onGetUserByNIK   func(mock *mocks.MockUserRepository)
	}

	var testTable []testCase
	testTable = append(testTable, testCase{
		name:      "success",
		wantError: false,
		request: &models.RegisterUser{
			FullName:  "John Doe",
			LegalName: "John Doe",
			Email:     "john@example.com",
			Password:  "password123",
			NIK:       "1234567890",
		},
		onGetUserByEmail: func(mock *mocks.MockUserRepository) {
			mock.EXPECT().GetUserByEmail("john@example.com").Return(&models.User{}, nil)
		},
		onGetUserByNIK: func(mock *mocks.MockUserRepository) {
			mock.EXPECT().GetUserByNIK("1234567890").Return(&models.User{}, nil)
		},
		onRegister: func(mock *mocks.MockUserRepository) {
			mock.EXPECT().Register(gomock.Any()).DoAndReturn(func(user *models.User) (*models.User, error) {
				return &models.User{
					UserID:    1,
					FullName:  user.FullName,
					LegalName: user.LegalName,
					Email:     user.Email,
					NIK:       user.NIK,
				}, nil
			})
		},
		expectedResponse: &models.User{
			UserID:    1,
			FullName:  "John Doe",
			LegalName: "John Doe",
			Email:     "john@example.com",
			NIK:       "1234567890",
		},
	})

	testTable = append(testTable, testCase{
		name:      "user already registered",
		wantError: true,
		request: &models.RegisterUser{
			FullName:  "John Doe",
			LegalName: "John Doe",
			Email:     "john@example.com",
			Password:  "password123",
			NIK:       "1234567890",
		},
		onGetUserByEmail: func(mock *mocks.MockUserRepository) {
			mock.EXPECT().GetUserByEmail("john@example.com").Return(&models.User{
				UserID:   1,
				FullName: "John Doe",
				Email:    "john@example.com",
				NIK:      "1234567890",
			}, nil)
		},
		expectedResponse: nil,
	})

	testTable = append(testTable, testCase{
		name:      "error checking by email",
		wantError: true,
		request: &models.RegisterUser{
			FullName:  "Jane Doe",
			LegalName: "Jane Doe",
			Email:     "error@example.com",
			Password:  "password123",
			NIK:       "0987654321",
		},
		onGetUserByEmail: func(mock *mocks.MockUserRepository) {
			mock.EXPECT().GetUserByEmail("error@example.com").Return(nil, errors.New("database error"))
		},
		expectedResponse: nil,
	})

	for _, tc := range testTable {
		t.Run(tc.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()

			userRepo := mocks.NewMockUserRepository(mockCtrl)

			if tc.onGetUserByEmail != nil {
				tc.onGetUserByEmail(userRepo)
			}

			if tc.onGetUserByNIK != nil {
				tc.onGetUserByNIK(userRepo)
			}

			if tc.onRegister != nil {
				tc.onRegister(userRepo)
			}

			usecase := &UserUsecase{userRepo: userRepo}

			resp, err := usecase.Register(tc.request)

			if tc.wantError {
				assert.NotNil(t, err)
				assert.Nil(t, resp)
			} else {
				assert.Nil(t, err)
				assert.Equal(t, tc.expectedResponse, resp)
			}
		})
	}
}

func Test_UserUsecase_Login(t *testing.T) {
	type testCase struct {
		name             string
		wantError        bool
		expectedResponse *models.LoginResponse
		request          *models.LoginUser
		onGetUserByEmail func(mock *mocks.MockUserRepository)
	}

	var testTable []testCase
	testTable = append(testTable, testCase{
		name:      "user not found",
		wantError: true,
		request: &models.LoginUser{
			Email:    "john@example.com",
			Password: "password",
		},
		onGetUserByEmail: func(mock *mocks.MockUserRepository) {
			mock.EXPECT().GetUserByEmail("john@example.com").Return(nil, gorm.ErrRecordNotFound)
		},
	})

	testTable = append(testTable, testCase{
		name:      "password does not match",
		wantError: true,
		request: &models.LoginUser{
			Email:    "john@example.com",
			Password: "wrongpassword",
		},
		onGetUserByEmail: func(mock *mocks.MockUserRepository) {
			mockUser := &models.User{
				UserID:   1,
				FullName: "John Doe",
				Email:    "john@example.com",
				Password: helper.HashPassword("password"),
			}
			mock.EXPECT().GetUserByEmail("john@example.com").Return(mockUser, nil)
		},
		expectedResponse: nil,
	})

	testTable = append(testTable, testCase{
		name:      "error checking user by email",
		wantError: true,
		request: &models.LoginUser{
			Email:    "error@example.com",
			Password: "password",
		},
		onGetUserByEmail: func(mock *mocks.MockUserRepository) {
			mock.EXPECT().GetUserByEmail("error@example.com").Return(nil, errors.New("database error"))
		},
		expectedResponse: nil,
	})

	for _, tc := range testTable {
		t.Run(tc.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()

			userRepo := mocks.NewMockUserRepository(mockCtrl)

			if tc.onGetUserByEmail != nil {
				tc.onGetUserByEmail(userRepo)
			}

			usecase := &UserUsecase{userRepo: userRepo}

			resp, err := usecase.Login(tc.request)

			if tc.wantError {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
				assert.Equal(t, tc.expectedResponse, resp)
			}
		})
	}
}
