# It's simply not possible to run postgres in a docker image in replit like locally so we're here :((
pg_ctl stop

initdb
cp postgresql.conf.tpl data/postgresql.conf

socker_dir="\/home\/runner\/${REPL_SLUG}\/postgres"

sed -i "s/replace_unix_dir/${socker_dir}/" data/postgresql.conf

pg_ctl -l /home/runner/${REPL_SLUG}/postgresql.log start

createdb -h 127.0.0.1 -p 5432 -U test -W testpass
psql -h 127.0.0.1 -c "create database testDB;"
psql -h 127.0.0.1 -d testDB -f db/migrations/*up.sql
