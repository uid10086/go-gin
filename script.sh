#bin bash
go build main.go
killall main
nohup ./main &
exit
