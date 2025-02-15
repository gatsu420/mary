.PHONY: migrate
migrate:
	dbmate --env POSTGRES_URL migrate

.PHONY: rollback
rollback:
	dbmate --env POSTGRES_URL rollback
