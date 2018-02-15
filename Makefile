linter:
	@ echo "-> Running linters ..."
	@ gometalinter --vendor ./...
.PHONY: linter

mockgen:
	@ echo "-> Generate mocks for tests ..."
	@ mockgen -source interface.go -package mock_data -destination mock_data/interface_mock.go
.PHONY: linter