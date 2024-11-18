# Makefile

VENV_DIR := venv
PYTHON := python3
PIP := $(VENV_DIR)/bin/pip
MKDOCS := $(VENV_DIR)/bin/mkdocs

.PHONY: all venv install build serve clean

all: venv install build

venv:
	$(PYTHON) -m venv $(VENV_DIR)

install: venv
	$(PIP) install mkdocs mkdocs-material

build: install
	$(MKDOCS) build

serve: install
	$(MKDOCS) serve

clean:
	rm -rf $(VENV_DIR) site