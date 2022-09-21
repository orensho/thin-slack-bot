SLACKBOT_REPO=myrepo
SLACKBOT_IMAGENAME=thin-slack-bot
SLACKBOT_VERSION=0.0.1

build:
	go test .
	GOOS=linux GOARCH=amd64 go build -o thin-slack-bot
	docker build --platform=linux/amd64 -t ${SLACKBOT_REPO}/${SLACKBOT_IMAGENAME}:${SLACKBOT_VERSION} .