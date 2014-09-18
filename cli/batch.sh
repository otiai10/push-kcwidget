pkill push-batch
echo "batch killed"
sleep 1
nohup go run batch/push-batch.go &
echo "batch started"
