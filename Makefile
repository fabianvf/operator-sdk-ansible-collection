build_modules:
	cd src && \
	go build -o ../plugins/modules/version_info cmd/version/version.go

build_collection:
	ansible-galaxy collection build --output-path build/ --force

build: build_modules build_collection

install: build
	ansible-galaxy collection install build/*.tar.gz --force

clean:
	rm -rf build
