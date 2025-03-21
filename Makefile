generate-mocks:
	# repositories
	@mockgen -destination=./service/repository/mocks/mock_user_repository.go -package=mocks github.com/saipulmuiz/krplus/service UserRepository
	@mockgen -destination=./service/repository/mocks/mock_credit_repository.go -package=mocks github.com/saipulmuiz/krplus/service CreditRepository
	@mockgen -destination=./service/repository/mocks/mock_payment_repository.go -package=mocks github.com/saipulmuiz/krplus/service PaymentRepository
	@mockgen -destination=./service/repository/mocks/mock_transaction_repository.go -package=mocks github.com/saipulmuiz/krplus/service TransactionRepository