module main

go 1.26.2

require database v0.0.1

require (
	common v0.0.1
	server v0.0.1
)

require api v0.0.1 // indirect

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/stretchr/testify v1.11.1
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace database => ./database

replace server => ./server

replace api => ./server/api

replace common => ./common
