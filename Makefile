init-start:
	docker network create guardrails
	service-start api-build-start
service-start:
	docker-compose -f docker-compose-service.yml up -d
service-stop:
	docker-compose -f docker-compose-service.yml down
api-start:
	docker-compose -f docker-compose-api.yml up -d
api-build-start:
	docker-compose -f docker-compose-api.yml up -d --build
api-stop:
	docker-compose -f docker-compose-api.yml down
start:
	service-start api-build-start
stop: service-stop api-stop
prepare-config:
	cp docker_configs/grchallengeapi/config.yml.sample docker_configs/grchallengeapi/config.yml
	cp docker_configs/grscanningworker/config.yml.sample docker_configs/grscanningworker/config.yml
	cp grchallengeapi/configs/config.yml.sample grchallengeapi/configs/config.yml
	cp grscanningworker/configs/config.yml.sample grscanningworker/configs/config.yml
log-kafka:
	docker logs -f backend-engineer-challenge_kafka_1
log-api:
	docker logs -f backend-engineer-challenge_grchallengeapi_1
log-worker:
	docker logs -f backend-engineer-challenge_grscanningworker_1