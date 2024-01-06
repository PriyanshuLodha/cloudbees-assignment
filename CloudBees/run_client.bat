@echo off

rem Get the directory of the script
SET "SCRIPT_DIR=%~dp0"

rem Set the working directory to the project root
cd /d "%SCRIPT_DIR%\CLOUDBEES"

rem Run client tests
echo Running client tests...
go test -v "./client"

rem Check if tests passed
if %ERRORLEVEL% EQU 0 (
    echo Client tests passed. Running client main file...

    rem Run client main file
    go run "./client/main.go"

    echo Client main file executed successfully.
) else (
    echo Client tests failed. Not running client main file.
)
