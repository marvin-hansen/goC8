# Make will use bash instead of sh
SHELL := /usr/bin/env bash
CC=clang # required by bazel
#ENV=LOCAL # Required by auto-config.

# GNU make man page
# http://www.gnu.org/software/make/manual/make.html

# For some strange reasons, intends & blanks shift in bash when calling 'make' so the formatting below should align intend at least on Bash on OSX.
.PHONY: help
help:
	@echo ' '
	@echo 'Setup: '
	@echo '    make check        	    	Checks all requirements.'
	@echo ' '
	@echo 'Test: '
	@echo '    make test-all   		Runs all API tests in order.'
	@echo '    make test-collection   	Tests collection API. '
	@echo '    make test-document   	Tests document API. '
	@echo '    make test-index   		Tests index API. '
	@echo '    make test-kv   		Tests key-value API. '

# "---------------------------------------------------------"
# Setup
# "---------------------------------------------------------"
.PHONY: check
check:
	@source scripts/setup/check_requirements.sh

# "---------------------------------------------------------"
# Test
# "---------------------------------------------------------"

.PHONY: test-all
test-all:
	@source scripts/test/test-all.sh

.PHONY: test-collection
test-collection:
	@source scripts/test/test-collection.sh

.PHONY: test-document
test-document:
	@source scripts/test/test-document.sh

.PHONY: test-index
test-index:
	@source scripts/test/test-index.sh

.PHONY: test-kv
test-kv:
	@source scripts/test/test-kv.sh
