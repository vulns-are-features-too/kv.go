alias r := run
alias l := lint
alias f := fixup
alias t := test
alias tu := unit_test
alias ti := integration_test
alias ts := simulation_test
alias ta := api_test

[working-directory: './src']
run:
  go run .

[working-directory: './src']
lint:
  golangci-lint run

[working-directory: './src']
fixup:
  golangci-lint fmt
  golangci-lint run --fix

test: unit_test integration_test simulation_test api_test

[working-directory: './src/tests/unit_test']
@unit_test:
  echo "Running unit tests"
  go test

[working-directory: './src/tests/integration_test']
@integration_test:
  echo "Running integration tests"
  go test

[working-directory: './src/tests/simulation_test']
@simulation_test:
  echo "Running simulation tests"
  go test

[working-directory: './src/tests/api_test']
@api_test:
  echo "Running api tests"
  ./run_tests.sh
