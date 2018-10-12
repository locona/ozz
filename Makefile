.DEFAULT_GOAL := run

.PHONY: run
run:
	@go install
	@ozz
