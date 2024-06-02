
setup:
	go get -u -v github.com/kokizzu/gotro@latest
	go install github.com/cosmtrek/air@latest
	go install github.com/fatih/gomodifytags@latest
	go install github.com/kokizzu/replacer@latest
	go install github.com/akbarfa49/farify@latest
	go install golang.org/x/tools/cmd/goimports@latest
	curl -fsSL https://bun.sh/install | bash

local-tarantool:
	#docker exec -it benakun-tarantool1-1 tarantoolctl connect benakunT:benakunPT@127.0.0.1:3301 # 2.11
	#docker exec -it benakun-tarantool1-1 tt connect benakunT:benakunPT@127.0.0.1:3301 # 3.1 lua mode
	docker exec -it benakun-tarantool1-1 tt connect benakunT:benakunPT@127.0.0.1:3301 -l sql # 3.1 sql mode
	# box.space -- list all tables
	# box.execute [[ SELECT * FROM "users" LIMIT 1 ]]
	# \set language sql
	# \set delimiter ;
    # SET SESSION "sql_seq_scan" = true; -- since 3.1 sequential scan no longer allowed

local-clickhouse:
	docker exec -it benakun-clickhouse1-1 clickhouse-client --host 127.0.0.1 --port 9000 --user benakunC --password benakunPC
	# SHOW TABLES -- list all tables
	# SELECT * FROM "actionLogs" LIMIT 1;

modtidy:
	sudo chmod -R a+rwx _tmpdb && go mod tidy

orm:
	# generate ORM
	./gen-orm.sh

views:
	# generate views and routes
	./gen-views.sh

migrate:
	# migrate table schema
	go run main.go migrate