## Create a virtual environment
python3 -m venv venv

## Activate the Environment:
source venv/bin/activate

## Install dependencies
pip install -r requirements.txt

## deactivate the virtual environment when done
deactivate

## example pip install commands
```shell
#!/bin/bash
pip install flask 
pip install requests
# safer to use python -m pip to ensure the correct pip is used
python -m pip install rich
```