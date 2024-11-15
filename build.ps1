# Write-Output "Building imgo_cli.exe at ./bin/"
# $env:GOOS="windows";$env:GOARCH="amd64"; go build -o .\bin\imgo_cli.exe .\cmd\imgo_cli\main.go

param(
    [ValidateSet("test", "api", "cli")]
    [string]$buildType = "test"
)

Write-Output "Building $buildType application at ./bin/"

$env:GOOS = "windows"
$env:GOARCH = "amd64"

$outputFile = ".\bin\imgo.exe"
$mainFile = ".\main.go"

switch ($buildType) {
    "api" {
        $outputFile = ".\bin\imgo-api.exe"
        $mainFile = ".\cmd\imgo_server\main.go"
        go build -o $outputFile $mainFile
    }
    "cli" {
        $outputFile = ".\bin\imgo-cli.exe"
        $mainFile = ".\cmd\imgo_cli\main.go"
        go build -o $outputFile $mainFile
    }
    "test" {
        go test -v ./...
    }
}


if ($LASTEXITCODE -eq 0) {
    Write-Output "Build successful: $outputFile"
} else {
    Write-Output "Build failed"
}
