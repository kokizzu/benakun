
setup:
	go get -u -v github.com/kokizzu/gotro@latest
	go install github.com/cosmtrek/air@latest
	go install github.com/fatih/gomodifytags@latest
	go install github.com/kokizzu/replacer@latest
	go install github.com/akbarfa49/farify@latest
	go install golang.org/x/tools/cmd/goimports@latest
	curl -fsSL https://bun.sh/install | bash

local-tarantool:
	docker exec -it `docker ps | grep tarantool | cut -d ' ' -f 1` tarantoolctl connect benakunT:benakunPT@127.0.0.1:3302
	# box.space -- list all tables
	# box.execute [[ SELECT * FROM "users" LIMIT 1 ]]
	# \set language sql
	# \set delimiter ;

local-clickhouse:
	docker exec -it `docker ps | grep clickhouse | cut -d ' ' -f 1` clickhouse-client --host 127.0.0.1 --port 9001 --user benakunC --password benakunPC
	# SHOW TABLES -- list all tables
	# SELECT * FROM "actionLogs" LIMIT 1;