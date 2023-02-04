include .env
export

.PHONY: up
up:
	@docker compose --project-name ${APP_NAME} --file ./.docker/docker-compose.yaml up -d

.PHONY: down
down:
	@docker compose --project-name ${APP_NAME} down
