mig-create DESC:
  migrate create -seq -ext=.sql -dir=./migrations {{DESC}}
mig-run:
  migrate -path ./migrations -database $GREENLIGHT_DB_DSN up
mig-goto VERSION:
  migrate -path ./migrations -database $GREENLIGHT_DB_DSN goto {{VERSION}}
mig-version VERSION:
  migrate -path ./migrations -database $GREENLIGHT_DB_DSN {{VERSION}}
mig-down MANY:
  migrate -path ./migrations -database $GREENLIGHT_DB_DSN down {{MANY}}
mig-force VERSION:
  migrate -path ./migrations -database $GREENLIGHT_DB_DSN force {{VERSION}}

kill-api PORT="4000":
  lsof -ti :{{PORT}} | xargs -r kill 2>/dev/null
