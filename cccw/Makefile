SHELL := /usr/bin/env bash
TOP := .

PYTHON ?= python
VENV ?= $(TOP)/.venv

ACTIVATE = $(VENV)/bin/activate

# .SUFFIXES:
# .SUFFIXES: .py .whl

# .PRECIOUS: .py .json

.PHONY: all
all: venv

.PHONY: venv
venv: $(ACTIVATE)
$(ACTIVATE): requirements.txt
	@test -d $(VENV) || $(PYTHON) -m venv $(VENV) && \
  source $(ACTIVATE) && \
  $(PYTHON) -m pip install --upgrade pip setuptools wheel && \
  $(PYTHON) -m pip install --requirement requirements.txt && \
  touch $(ACTIVATE)

# The rule for venv_upgrade has no dependencies to force it to always run

.PHONY: venv_upgrade
venv_upgrade:
	@rm -rf $(VENV) && \
  $(PYTHON) -m venv $(VENV) && \
  source $(ACTIVATE) && \
  $(PYTHON) -m pip install --upgrade pip setuptools wheel && \
  if [ -f requirements_bare_private.txt ]; then $(PYTHON) -m pip install --upgrade --requirement requirements_bare_private.txt; fi && \
  $(PYTHON) -m pip install --upgrade --requirement requirements_bare.txt && \
  $(PYTHON) -m pip freeze > requirements.txt && \
  touch $(ACTIVATE)

.PHONY: clean
clean:
	@rm -rf $(VENV)
