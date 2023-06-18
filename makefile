init:
	git submodule update --init --recursive
	docker compose up -d

pull:
	git fetch --all
	git pull
	git pull origin main --recurse-submodules
	git submodule update --recursive --remote