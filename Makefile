.PHONY: run-tests
run-tests:
	@go test -v -failfast `go list ./...` -cover

.PHONY: run-tests-report
run-tests-report:
	@go test -v -failfast `go list ./...` -cover -coverprofile=coverage.out -json > test-report.out