start=`date +%s`

set -eu -o pipefail

pids=()

cur_dir=$PWD

echo "Building DB"

cd infra/postgres && docker-compose down --rmi all --volumes --remove-orphans && docker-compose up -d &
pids+=( $! )

for pid in ${pids[*]}; do
  wait $pid
done

unset pids

end=`date +%s`

runtime=$((end-start))
echo "Duration: ${runtime}s"
