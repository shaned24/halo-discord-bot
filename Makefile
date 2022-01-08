# Tokens
DISCORD_TOKEN ?="discord_token"
AUTOCODE_TOKEN ?="autocode_token"

# .env
ifneq (,$(wildcard ./.env))
    include .env
    export
endif

# Docker
IMAGE         ?=shaned24/halo-discord-bot
TAG           ?=latest

# Helm
NAME          ?=halo-discord-bot
RELEASE_NAME  ?=crabbot-$(NAME)
CHART 		  ?=./deploy/crabbot
NAMESPACE     ?=halo-discord-bot
COMMAND       ?=/halo-discord-bot

.PHONY: deploy deploy-dry-run
docker-build:
	@docker build . -t $(IMAGE):$(TAG)

docker-publish: docker-build
	@docker push $(IMAGE):$(TAG)

docker-build-multi:
	docker buildx build --platform linux/amd64,linux/arm64 --push -t $(IMAGE):$(TAG) .

docker-run: docker-build
	docker run -it --rm \
 		-e DISCORD_TOKEN=$(DISCORD_TOKEN) \
 		-e AUTOCODE_TOKEN=$(AUTOCODE_TOKEN) \
 		$(IMAGE):$(TAG)

deploy:
	helm upgrade --install \
		$(RELEASE_NAME) $(CHART) \
		--namespace $(NAMESPACE) --create-namespace \
		--set "image.repository=$(IMAGE)" \
		--set "image.tag=$(TAG)" \
		--set "secrets.discordToken=$(DISCORD_TOKEN)" \
		--set "secrets.autoCodeToken=$(AUTOCODE_TOKEN)"

deploy-dry-run:
	helm upgrade $(RELEASE_NAME) $(CHART) \
		--install --dry-run --debug  \
		--namespace $(NAMESPACE) --create-namespace \
		--set "image.repository=$(IMAGE)" \
		--set "image.tag=$(TAG)" \
		--set "secrets.discordToken=$(DISCORD_TOKEN)" \
		--set "secrets.autoCodeToken=$(AUTOCODE_TOKEN)"
