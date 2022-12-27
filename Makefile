

BUILD_FOLDER="./build/Bump version $(SEMVER)"
SOURCE_FOLDER := "./src"
LDFLAGS="-X 'main.BUMPVER=$(SEMVER)'"



all:
	@echo "Building All Files"
	@echo $(LDFLAGS)
	make macOS
	# make windows
	# make linux

macOS:
	@echo "Making Macos..."
	GOOS=darwin GOARCH=arm64 go build -ldflags=$(LDFLAGS) -o $(BUILD_FOLDER)/macOS-ARM/bump $(SOURCE_FOLDER)
	GOOS=darwin GOARCH=amd64 go build -o $(BUILD_FOLDER)/macOS-Intel/bump $(SOURCE_FOLDER)


windows:
	# @echo "Making Windows..."
	GOOS=windows GOARCH=386 go build -o $(BUILD_FOLDER)/window-x86/bump.exe $(SOURCE_FOLDER)
	GOOS=windows GOARCH=amd64 go build -o $(BUILD_FOLDER)/window-amd/bump.exe $(SOURCE_FOLDER)
	GOOS=windows GOARCH=arm64 go build -o $(BUILD_FOLDER)/window-arm/bump.exe $(SOURCE_FOLDER)


linux:
	# @echo "Making Linux..."
	GOOS=linux GOARCH=386 go build -o $(BUILD_FOLDER)/linux-x86/bump $(SOURCE_FOLDER)
	GOOS=linux GOARCH=arm64 go build -o $(BUILD_FOLDER)/linux-arm64/bump $(SOURCE_FOLDER)