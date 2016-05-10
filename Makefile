usage:
	@echo "make server"

env:
	@go get -u github.com/jteeuwen/go-bindata/...

	@glide install

clean:
	@rm -rf public/bindata.go
	@rm -rf templates/bindata.go

bindata:
	@cd public && go-bindata -o bindata.go -pkg public -ignore="\\.DS_Store|\\.gitignore|\\.gitkeep" $(BINDATA_OPTS) ./...
	@cd templates && go-bindata -o bindata.go -pkg templates -ignore="\\.DS_Store|\\.gitignore|\\.gitkeep" $(BINDATA_OPTS) ./...

dev-bindata:
	@$(MAKE) --no-print-directory bindata BINDATA_OPTS="-debug"

server: clean bindata
	@go run server.go

dev-server: clean dev-bindata
	@go run server.go
