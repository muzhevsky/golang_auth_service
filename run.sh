sudo apt update
sudo apt install postgresql-client


sudo docker compose up -d --build

psql -h 127.0.0.1 -U postgres -f ./authorization/migrations/000001_create_user_table.up.sql