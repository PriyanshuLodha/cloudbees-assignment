@echo off

rem Get the directory of the script
SET "SCRIPT_DIR=%~dp0"

rem Set the working directory to the project root
cd /d "%SCRIPT_DIR%\CLOUDBEES"

rem Run server tests
echo Running server tests...
go test -v "./server"

rem Check if tests passed
if %ERRORLEVEL% EQU 0 (
    echo Server tests passed. Running server main file...

    rem Run server main file
    go run "./server/main.go"

    echo Server main file executed successfully.
) else (
    echo Server tests failed. Not running server main file.
)
