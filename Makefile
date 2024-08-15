TARGET_TRUNK_BUILD_DIR := ./target/trunk-build
BUILD_DIR := ./.build
BUILD_FILE := $(BUILD_DIR)/build.zip

.PHONY: trunk-serve
trunk:
	trunk serve --open --dist ./target/trunk-serve

.PHONY: create-target-trunk-build-dir
create-target-trunk-build-dir:
	@if [ ! -d $(TARGET_TRUNK_BUILD_DIR) ]; then \
		echo "Creating directory $(TARGET_TRUNK_BUILD_DIR)"; \
		mkdir -p $(TARGET_TRUNK_BUILD_DIR); \
	fi

.PHONY: create-build-dir
create-build-dir:
	@if [ ! -d $(BUILD_DIR) ]; then \
		echo "Creating directory $(BUILD_DIR)"; \
		mkdir -p $(BUILD_DIR); \
	fi

.PHONY: trunk-build
trunk-build: create-target-trunk-build-dir create-build-dir
	trunk build --release --dist $(TARGET_TRUNK_BUILD_DIR)
	(cd $(TARGET_TRUNK_BUILD_DIR) && zip --recurse-paths ../../$(BUILD_FILE) .)
