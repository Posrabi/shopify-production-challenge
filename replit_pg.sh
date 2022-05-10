# It's simply not possible to run postgres in a docker image in replit like locally so we're here :((
pg_ctl stop

initdb
cp postgresql.conf.tpl data/postgresql.conf

socker_dir="\/home\/runner\/${REPL_SLUG}\/postgres"

sed -i "s/replace_unix_dir/${socker_dir}/" data/postgresql.conf

pg_ctl -l /home/runner/${REPL_SLUG}/postgresql.log start

createdb -h localhost -p 5432 -U test -W testpass
psql -h localhost -c "create database testDB;" -U test -W testpass
psql -h localhost -d testDB -f db/migrations/*up.sql -U test -W testpass
