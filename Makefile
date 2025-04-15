.PHONY: migration-up
migration-up:
	dbmate \
		--env MARY_POSTGRES_URL \
		--migrations-dir "./dependency/postgres/migrations" \
		--schema-file "./dependency/postgres/schema.sql" \
		up

.PHONY: migration-down
migration-down:
	dbmate \
		--env MARY_POSTGRES_URL \
		--migrations-dir "./dependency/postgres/migrations" \
		--schema-file "./dependency/postgres/schema.sql" \
		down

.PHONY: migration-new
migration-new:
	dbmate \
		--env MARY_POSTGRES_URL \
		--migrations-dir "./dependency/postgres/migrations" \
		new $(NAME)

.PHONY: sqlc-gen
sqlc-gen:
	sqlc generate

.PHONY: buf-gen
buf-gen:
	buf generate

.PHONY: mock
mock:
	mockery
