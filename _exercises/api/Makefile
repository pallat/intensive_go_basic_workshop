.PHONY: init create load
init:
	pre-commit install --config ./githook/pre-commit-config.yaml
create:
	sqlite3 languages.sqlite < ./create_table.sql
load:
	sqlite3 languages.sqlite < ./load_data.sql
