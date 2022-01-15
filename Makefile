APP_NAME=hlsconv
VERSION=$(shell head -n 1 VERSION)
APP_FOLDER=$(APP_NAME)-$(VERSION)

build:
	go build -o ./bin/

run:
	bin/$(APP_NAME) -i video/ -o video/

pack-win:
	rm -rf $(APP_FOLDER) bin
	mkdir -p $(APP_FOLDER)
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 make build;
	cp -r conf bin lib README.md $(APP_FOLDER)
	zip -r $(APP_FOLDER)-win.zip $(APP_FOLDER) 
	rm -rf $(APP_FOLDER) bin

pack-mac:
	rm -rf $(APP_FOLDER) bin
	mkdir -p $(APP_FOLDER)
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 make build;
	cp -r conf bin lib README.md $(APP_FOLDER)
	zip -r $(APP_FOLDER)-macos.zip $(APP_FOLDER) 
	rm -rf $(APP_FOLDER) bin

pack-linux:
	rm -rf $(APP_FOLDER) bin
	mkdir -p $(APP_FOLDER)
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 make build;
	cp -r conf bin lib README.md $(APP_FOLDER)
	zip -r $(APP_FOLDER)-linux.zip $(APP_FOLDER) 
	rm -rf $(APP_FOLDER)