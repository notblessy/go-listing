run:
	go run main.go httpsrv

migration:
	go run main.go migrate

model/mock/mock_product_repository.go:
	mockgen -destination=model/mock/mock_product_repository.go -package=mock github.com/notblessy/go-listing/model ProductRepository

model/mock/mock_product_usecase.go:
	mockgen -destination=model/mock/mock_product_usecase.go -package=mock github.com/notblessy/go-listing/model ProductUsecase

mockgen: model/mock/mock_product_repository.go \
	model/mock/mock_product_usecase.go

test: unit-test
unit-test: mockgen
	SVC_ENV=test SVC_DISABLE_CACHING=true go test ./... -v --cover