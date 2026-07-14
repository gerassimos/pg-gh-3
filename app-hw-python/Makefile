
# Environment variables with default values
# Can be overridden by exporting them before running make commands
# Example: export DEVICE1_WEB_SERVER_COUNT=1
DEVICE1_WEB_SERVER_COUNT ?= 3
START_PORT ?= 5001

## create python virtual environment
venv-create:
	python3 -m venv venv

## activate python virtual environment
venv-activate:	
	@echo "Run this command in your terminal:"
	@echo "source venv/bin/activate"

## deactivate python virtual environment
venv-deactivate:
	@echo "Run this command in your terminal:"
	@echo "deactivate"

## generate requirements.txt file for python dependencies
build-py-requirements:
	. venv/bin/activate && pip freeze > requirements.txt

## install dependencies from requirements.txt file
install-py-requirements:
	cat requirements.txt
	. venv/bin/activate && pip install -r requirements.txt

start-app-hw-python:
	. venv/bin/activate && python hello.py



