dep:
	@cd src && go mod download

run:
	@cd cmd && go run *.go

ADAPTERS_DIR := $(patsubst %,%,$(notdir $(wildcard ./adapters/*)))

gen_mock_adapters:
	@for dir in $(ADAPTERS_DIR); do \
  		mockery -case=underscore -dir=./adapters/$$dir -output ./mocks/adapters/$$dir -all ; \
    done
gen_mock_adapter2:
	@for dir in $(ADAPTERS_DIR); do \
  		mockery --case=underscore --dir=./adapters/$$dir --output ./mocks/adapters/$$dir --all ; \
    done
