build-echo:
	cd echo && go mod vendor && docker build -t echo:v0.1.0 .
up:
	docker-compose up -d
stop:
	docker-compose stop
down:
	docker-compose down