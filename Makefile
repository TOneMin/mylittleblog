all-littlecat: run-littlecat-db run-littlecat

run-littlecat:
	docker-compose up littlecat
run-littlecat-db:
	docker-compose up -d postgres