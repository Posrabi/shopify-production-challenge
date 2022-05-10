go mod download
make
sudo ln -s /usr/local/bin /usr/bin
bash ./db.sh
./build/inv server
