source ./files/env/env.sh
source ./files/dbmate_config/dbmate.sh

docker compose up -d

# Do database migration (TODO: Only do synchronous migration with app start in dev)
dbmate up

# Start the HTTP service locally
go run ./cmd/main.go