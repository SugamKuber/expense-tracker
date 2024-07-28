SERVICES := auth:3000 tracker:3001 file-manager:3002
PIDS_FILE := server_pids.txt
LOG_DIR := logs

.PHONY: all start stop clean setup restart check

all: setup start

setup:
	@mkdir -p $(LOG_DIR)

start: setup
	@echo "Starting all services..."
	@> $(PIDS_FILE)
	@for service_port in $(SERVICES); do \
		service=$${service_port%%:*}; \
		port=$${service_port#*:}; \
		echo "Starting $$service on port $$port..."; \
		cd $$service && go mod tidy && \
		( go run cmd/main.go > $(PWD)/$(LOG_DIR)/$$service.log 2>&1 & ) && \
		cd $(PWD); \
		for i in $$(seq 1 10); do \
			sleep 1; \
			pid=$$(lsof -t -i :$$port); \
			if [ -n "$$pid" ]; then \
				echo "$$pid" >> $(PIDS_FILE); \
				echo "$$service started with PID $$pid"; \
				break; \
			fi; \
			if [ $$i -eq 10 ]; then \
				echo "Failed to start $$service or find its PID"; \
			fi; \
		done; \
	done
	@echo "All services started. PIDs saved in $(PIDS_FILE)"
	@echo "Logs are being written to $(LOG_DIR)/*.log"

stop:
	@echo "Stopping all services..."
	@if [ -f $(PIDS_FILE) ]; then \
		while read pid; do \
			if kill -0 $$pid 2>/dev/null; then \
				echo "Stopping process $$pid"; \
				kill $$pid; \
			else \
				echo "Process $$pid is not running"; \
			fi; \
		done < $(PIDS_FILE); \
		rm $(PIDS_FILE); \
	else \
		echo "No PIDs file found. Are the servers running?"; \
	fi

check:
	@echo "Checking all services..."
	@for service_port in $(SERVICES); do \
		port=$${service_port#*:}; \
		echo "Checking service on port $$port:"; \
		curl -s localhost:$$port/h || echo "Failed to connect to service on port $$port"; \
		echo; \
	done

restart: stop clean start check

clean:
	@echo "Cleaning up logs..."
	@rm -rf $(LOG_DIR)