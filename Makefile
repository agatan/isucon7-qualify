MYSQL_HOST := 127.0.0.1
MYSQL_PORT := 3306
MYSQL_USER := root

.PHONY: migrate
migrate:
	mysql -h $(MYSQL_HOST) -P $(MYSQL_PORT) -u $(MYSQL_USER) -f isubata < sql/index.sql

