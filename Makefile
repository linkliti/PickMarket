ifeq ($(OS),Windows_NT)
SHELL := pwsh.exe
.SHELLFLAGS := -NoProfile -Command
RMDIR := rm -Recurse -ErrorAction Ignore
ENVSET := $$env:
ENVGET := $$env:
else
RMDIR := rm -rf
ENVSET := export
ENVGET := $
endif
PYENV := conda activate "./pmparser/venv" && cd pmparser
JSENV := cd frontend
# Hints for Makefile
# $(RMDIR) "./folder"
# $(ENVSET)FOO="bar"; echo $(ENVGET)FOO
# $(JSENV) && npm run dev

# Frontend
js-install:
	@$(JSENV) && npm install

js-dev:
	@$(ENVSET)DEV="True"; $(JSENV) && npm run dev

# Parser
py-venv:
	@conda create --prefix ./pmparser/venv python=3.12 -y
	@$(PYENV) && pip install -r "requirements.txt"

py-dev:
	@$(ENVSET)DEV="True"; $(PYENV) && python app.py

py-run:
	@$(PYENV) && python app.py

py-freeze:
	@$(PYENV) && pip freeze > "requirements.txt"

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