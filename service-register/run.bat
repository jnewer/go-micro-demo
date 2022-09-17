@echo off
start "p1" go run main.go --server_address :8001 &
start "p2" go run main.go --server_address :8002 &
start "p3" go run main.go --server_address :8003 &
pause