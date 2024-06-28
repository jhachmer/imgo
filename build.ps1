Write-Output "Building imgo.exe at ./bin/"
$env:GOOS="windows";$env:GOARCH="amd64"; go build -o .\bin\imgo.exe .\cmd\imgo\main.go
