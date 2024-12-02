.PHONY: run rerun test

run:
	@go run main.go -d $(filter-out $@,$(MAKECMDGOALS))

rerun:
	@rerun -watch . -run sh -c 'go run main.go -d $(filter-out $@,$(MAKECMDGOALS))'

test:
	@go test days/d$$(printf "%02d" $(filter-out $@,$(MAKECMDGOALS)))/main_test.go

retest:
	@rerun -watch . -run sh -c 'go test days/d$$(printf "%02d" $(filter-out $@,$(MAKECMDGOALS)))/main_test.go'

test-all:
	@go test ./...

generate-days:
	@./days/generate_days.sh
