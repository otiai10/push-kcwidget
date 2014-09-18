pkill push-batch
echo "batch killed"
sleep 1
# TODO: date
nohup go run batch/push-batch.go >> log/batch.log &
echo "batch started"
