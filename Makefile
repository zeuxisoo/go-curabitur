usage:
	@echo "make server"

env:
	@go get -u github.com/jteeuwen/go-bindata/...
	@go get github.com/Unknwon/bra

	@glide install

clean:
	@rm -rf public/bindata.go
	@rm -rf templates/bindata.go

	@rm -rf go-curabitur

bindata:
	@cd public && go-bindata -o bindata.go -pkg public $(BINDATA_OPTS) -ignore="\\.DS_Store|\\.gitignore|\\.gitkeep" ./...
	@cd templates && go-bindata -o bindata.go -pkg templates $(BINDATA_OPTS) -ignore="\\.DS_Store|\\.gitignore|\\.gitkeep" ./...

dev-bindata:
	@$(MAKE) --no-print-directory bindata BINDATA_OPTS="-debug"

server: clean bindata
	@bra run

dev-server: clean dev-bindata
	@bra run
