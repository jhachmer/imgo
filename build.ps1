# Write-Output "Building imgo.exe at ./bin/"
# $env:GOOS="windows";$env:GOARCH="amd64"; go build -o .\bin\imgo.exe .\cmd\imgo\main.go

Write-Output "Building imgo application at ./bin/"

$env:GOOS = "windows"
$env:GOARCH = "amd64"

$outputFile = ".\bin\imgo.exe"
$mainFile = ".\main.go"

# switch ($buildType) {
#     "cli" {
#         $outputFile = ".\bin\imgo-cli.exe"
#         $mainFile = ".\cmd\imgo\main.go"
#     }
#     "api" {
#         $outputFile = ".\bin\imgo-api.exe"
#         $mainFile = ".\cmd\imgo_server\main.go"
#     }
# }

go build -o $outputFile $mainFile

if ($LASTEXITCODE -eq 0) {
    Write-Output "Build successful: $outputFile"
} else {
    Write-Output "Build failed"
}
