.PHONY: migration-up
migration-up:
	dbmate --env POSTGRES_URL up

.PHONY: migration-down
migration-down:
	dbmate --env POSTGRES_URL down

.PHONY: migration-new
migration-new:
	dbmate --env POSTGRES_URL new $(NAME)
