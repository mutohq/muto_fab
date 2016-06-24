#!/bin/bash
begin=$(date +%s)
echo "Building for mac"
env GOOS=darwin GOARCH=amd64 go build -v -o copy-mac sourceToDest.go &


echo "Building for linux"
env GOOS=linux GOARCH=amd64 go build -v -o copy-linux sourceToDest.go &


echo "Building for windows"
env GOOS=windows GOARCH=amd64 go build -v -o copy-win.exe sourceToDest.go &

wait

end=$(date +%s)

total=$((end-begin))

echo "Built-in "$total"s"
