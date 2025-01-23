.PHONY: run-tests
run-tests:
	@go test -v -failfast `go list ./...` -cover

.PHONY: run-tests-report
run-tests-report:
	@go test -v -failfast `go list ./...` -cover -coverprofile=coverage.out -json > test-report.out

.PHONY: mock
mock:
	@`go env GOPATH`/bin/mockgen -source ./$(util)/$(subutil).go -destination ./tests/mock/$(util)/$(subutil).go

.PHONY: mock-all
mock-all:
	@make mock util=configreader subutil=configreader
	@make mock util=instrument subutil=instrument
	@make mock util=log subutil=log
	@make mock util=sql subutil=sql
	@make mock util=sql subutil=sql_tx
	@make mock util=sql subutil=sql_stmt
	@make mock util=sql subutil=sql_cmd
	@make mock util=query subutil=sql_builder
	@make mock util=redis subutil=redis
	@make mock util=parser subutil=parser
	@make mock util=parser subutil=json
	@make mock util=jwtAuth subutil=jwt
	@make mock util=mongo subutil=mongo
	@make mock util=mongo subutil=mongo_monitor
	@make mock util=timelib subutil=timelib
