.PHONY: init create load
init:
	pre-commit install --config ./githook/pre-commit-config.yaml
create:
	sqlite3 laguages.sqlite < ./create_table.sql
load:
	sqlite3 laguages.sqlite < ./load_data.sql