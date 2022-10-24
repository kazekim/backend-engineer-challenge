build:
	TZ=Asia/Bangkok GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o ./dist/grchallengeapi-linux-x64 cmd/grchallengeapi/main.go
mock:
	sh scripts/generate_mock.sh