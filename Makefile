GENMOCK := make gen-mock

gen-mock-all:
	$(GENMOCK) source=internal/core/ports/repository.go folder=mock_repository pkg=mockrepository
	$(GENMOCK) source=internal/core/ports/service.go folder=mock_service pkg=mockservice

gen-mock:
	mockgen -source=$(source) -package=$(pkg) -destination=pkg/testings/$(folder)/$(pkg).go