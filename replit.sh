go mod download
make
rm -r data
rm postgresql.log
bash replit_pg.sh
./build/inv server
