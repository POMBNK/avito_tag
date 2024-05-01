generate:
	go generate ./...
.PHONY: generate

mup:
	goose -dir migrations postgres postgres://postgres:postgres@localhost:5432/postgres up
	goose -dir migrations postgres postgres://postgres:postgres@localhost:5432/postgres status
.PHONY: mup

mdown:
	goose -dir migrations postgres postgres://postgres:postgres@localhost:5432/postgres down
	goose -dir migrations postgres postgres://postgres:postgres@localhost:5432/postgres status
.PHONY: mdown