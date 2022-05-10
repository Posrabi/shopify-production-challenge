go mod download
make
rm -r data
rm postgresql.log
rm logfile
bash replit_pg.sh
./build/inv server
