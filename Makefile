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
PYENV := conda activate "./parser/venv" && cd parser
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
	@conda create --prefix ./parser/venv python=3.12 -y
	@$(PYENV) && pip install -r "requirements.txt"

py-dev:
	@$(ENVSET)DEV="True"; $(PYENV) && python app.py

py-freeze:
	@$(PYENV) && pip freeze > "requirements.txt"

# Clean
clean:
	@conda env remove -p ./parser/venv -y
	@$(RMDIR) "./parser/venv"
	@$(RMDIR) "./parser/application/__pycache__"
	@$(RMDIR) "./frontend/node_modules"