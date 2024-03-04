ifeq ($(OS),Windows_NT)
SHELL := pwsh.exe
.SHELLFLAGS := -NoProfile -Command
RMDIR := rm -Recurse -ErrorAction Ignore
PYENV := conda activate "./pmparser/venv" && cd pmparser
ENVSET := $$env:
ENVGET := $$env:
else
SHELL := /bin/bash
CONDA_BASE := $(shell conda info --base)
PYENV := source "$(CONDA_BASE)/etc/profile.d/conda.sh" && conda activate "./pmparser/venv" && cd pmparser
RMDIR := rm -rf
ENVSET := export
ENVGET := $
endif
JSENV := cd frontend
# Hints for Makefile
# $(RMDIR) "./folder"
# $(ENVSET)FOO="bar"; echo $(ENVGET)FOO
# $(JSENV) && npm run dev

# Frontend
js-install:
	@$(JSENV) && npm install

js-dev:
	@$(ENVSET) DEV="True"; $(JSENV) && npm run dev

# Parser
py-venv:
	@conda create --prefix ./pmparser/venv python=3.12 -y
	@$(PYENV) && python -m pip install -r "requirements.txt"

py-dev:
	$(ENVSET) DEV="True"; $(PYENV) && python app.py

py-run:
	@$(PYENV) && python app.py

py-freeze:
	@$(PYENV) && python -m pip freeze > "requirements_freeze.txt"

# Clean
clean:
	@git clean -Xdf

remove: clean
	@$(RMDIR) "./frontend/node_modules"
	@conda env remove -p ./pmparser/venv -y

proto:
	@$(PYENV)
	@python -m grpc_tools.protoc -I ./protos/parser -I ./protos/thirdParty \ --python_out=./pmparser --grpc_python_out=./pmparser --pyi_out=./pmparser \ ./protos/parser/app/protos/*.proto
# @python -m grpc_tools.protoc -I ./protos/parser -I ./protos/thirdParty --go_out=./backend/itemsWorker --go_opt=Mapp/protos/types.proto=app/protos/types ./protos/parser/app/protos/*.proto

compose-rm:
	docker-compose stop \
	&& docker-compose rm \
	&& sudo rm -rf ./db/pgdata

compose-up:
	docker-compose -f docker-compose.yml up --force-recreate