# It's simply not possible to run postgres in a docker image in replit like locally so we're here :((
pg_ctl stop

initdb
cp postgresql.conf.tpl data/postgresql.conf

socker_dir="\/home\/runner\/${REPL_SLUG}\/postgres"

sed -i "s/replace_unix_dir/${socker_dir}/" data/postgresql.conf

pg_ctl -l /home/runner/${REPL_SLUG}/postgresql.log start

createdb -h localhost -p 5432 -e "testDB"
psql -h localhost -p 5432 -d "testDB" -c "create user test with password 'testpass';"
psql -h localhost -p 5432 -d "testDB" -f db/migrations/*up.sql -U test

