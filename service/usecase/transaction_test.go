package usecase

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/saipulmuiz/krplus/models"
	"github.com/saipulmuiz/krplus/service/repository/mocks"
	"github.com/stretchr/testify/assert"
)

func Test_TransactionUsecase_RecordTransaction(t *testing.T) {
	type testCase struct {
		name                string
		wantError           bool
		request             *models.RecordTransactionRequest
		onGetUserByNIK      func(mock *mocks.MockUserRepository)
		onGetCredits        func(mock *mocks.MockCreditRepository)
		onCreateTransaction func(mock *mocks.MockTransactionRepository)
		onUpdateCredit      func(mock *mocks.MockCreditRepository)
	}

	var testTable []testCase
	testTable = append(testTable, testCase{
		name:      "success",
		wantError: false,
		request: &models.RecordTransactionRequest{
			NIK:            "1234567890",
			ContractNumber: "CN123",
			OTR:            500000,
			Tenor:          12,
			AdminFee:       5000,
			Interest:       10000,
			AssetName:      "Mobil",
		},
		onGetUserByNIK: func(mock *mocks.MockUserRepository) {
			mock.EXPECT().GetUserByNIK("1234567890").Return(&models.User{UserID: 1}, nil)
		},
		onGetCredits: func(mock *mocks.MockCreditRepository) {
			mock.EXPECT().GetCredits(gomock.Any()).Return(&[]models.CreditLimit{
				{
					CreditID:             1,
					RemainingLimitAmount: 1000000,
					UsedLimitAmount:      0,
				},
			}, int64(1), nil)
		},
		onCreateTransaction: func(mock *mocks.MockTransactionRepository) {
			mock.EXPECT().CreateTransaction(gomock.Any()).Return(nil)
		},
		onUpdateCredit: func(mock *mocks.MockCreditRepository) {
			mock.EXPECT().UpdateCredit(gomock.Any(), gomock.Any()).Return(nil)
		},
	})

	testTable = append(testTable, testCase{
		name:      "user not found",
		wantError: true,
		request: &models.RecordTransactionRequest{
			NIK: "1234567890",
		},
		onGetUserByNIK: func(mock *mocks.MockUserRepository) {
			mock.EXPECT().GetUserByNIK("1234567890").Return(nil, errors.New("user not found"))
		},
	})

	testTable = append(testTable, testCase{
		name:      "insufficient credit limit",
		wantError: true,
		request: &models.RecordTransactionRequest{
			NIK:   "1234567890",
			OTR:   2000000,
			Tenor: 12,
		},
		onGetUserByNIK: func(mock *mocks.MockUserRepository) {
			mock.EXPECT().GetUserByNIK("1234567890").Return(&models.User{UserID: 1}, nil)
		},
		onGetCredits: func(mock *mocks.MockCreditRepository) {
			mock.EXPECT().GetCredits(gomock.Any()).Return(&[]models.CreditLimit{
				{
					CreditID:             1,
					RemainingLimitAmount: 1000000,
					UsedLimitAmount:      0,
				},
			}, int64(1), nil)
		},
	})

	for _, tc := range testTable {
		t.Run(tc.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()

			userRepo := mocks.NewMockUserRepository(mockCtrl)
			creditRepo := mocks.NewMockCreditRepository(mockCtrl)
			transactionRepo := mocks.NewMockTransactionRepository(mockCtrl)

			if tc.onGetUserByNIK != nil {
				tc.onGetUserByNIK(userRepo)
			}

			if tc.onGetCredits != nil {
				tc.onGetCredits(creditRepo)
			}

			if tc.onCreateTransaction != nil {
				tc.onCreateTransaction(transactionRepo)
			}

			if tc.onUpdateCredit != nil {
				tc.onUpdateCredit(creditRepo)
			}

			usecase := &TransactionUsecase{
				userRepo:        userRepo,
				creditRepo:      creditRepo,
				transactionRepo: transactionRepo,
			}

			err := usecase.RecordTransaction(*tc.request)

			if tc.wantError {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}
