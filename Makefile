MYSQL_HOST := db
MYSQL_PORT := 3306
MYSQL_USER := isucon
MYSQL_PASSWORD := isucon

.PHONY: migrate
migrate:
	mysql -h $(MYSQL_HOST) -P $(MYSQL_PORT) -u $(MYSQL_USER) -p$(MYSQL_PASSWORD) -f isubata < sql/index.sql

