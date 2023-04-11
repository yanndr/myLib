VERSION=`cat ../.version`
LDFLAGS=-ldflags "-X main.Version=${VERSION}"

.PHONY: api
api:

	cd api && go install  ${LDFLAGS}  ./cmd/myLibApi
	echo "api installed. exec: myLibApi"

.PHONY: client
client:
	cd myLib && go install ${LDFLAGS}
	echo "client installed. exec: myLib"