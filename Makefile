BUILD_DIR := ./.build
BUILD_FILE := $(BUILD_DIR)/build.zip

.PHONY: trunk-serve
trunk:
	trunk serve --open --dist ./target/trunk-serve

.PHONY: create-build-dir
create-build-dir:
	@if [ ! -d $(BUILD_DIR) ]; then \
		echo "Creating directory $(BUILD_DIR)"; \
		mkdir -p $(BUILD_DIR); \
	fi

.PHONY: trunk-build
trunk-build: create-build-dir
	trunk build --release --dist ./target/trunk-build
	zip --recurse-paths $(BUILD_FILE) ./target/trunk-build
