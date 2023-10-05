FORCE:

test:
	go test -coverprofile cicd/cover.out ./...

cover: test
	go tool cover -html cicd/cover.out

bench: FORCE
	go test -benchmem -bench=. ./...

license:
	addlicense -c "appkit Authors" -l apache -y 2023 .
