# I don't really understand Makefile and docker. Have a lot of things to learn
all: docker-clean docker-hub docker-browser

vendor-prepare:
	@go get github.com/DATA-DOG/godog/cmd/godog
	@go get -t -d github.com/tebeka/selenium
	@go get github.com/joho/godotenv
	@go get -u github.com/logrusorgru/aurora
	@echo "Package installed"

kill-port:
	@kill -HUP $$(lsof -t -i:4545)
	@echo "Port 4545 is killed"

docker-clean:
	@docker container rm $$(docker ps -aq) -f

docker-hub:
	docker run -d -p 4545:4444 --name selenium-hub selenium/hub

docker-browser:
	docker run -d --link selenium-hub:hub selenium/node-chrome
	docker run -d --link selenium-hub:hub selenium/node-firefox
	@echo "Docker is running"
