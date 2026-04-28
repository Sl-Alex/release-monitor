@echo off
setlocal enabledelayedexpansion

REM --- paths ---
set ROOT=%~dp0
set TOOLS=%ROOT%tools
set TMP=%ROOT%_tmp_install

if exist "%TOOLS%\deno.exe" (
	echo.
	echo Deno is already installed, nothing to do
	exit /b 0
)

mkdir "%TMP%" 2>nul

cd /d "%TMP%"

echo.
echo Downloading Deno...

gh release download ^
  --repo denoland/deno ^
  --pattern "deno-x86_64-pc-windows-msvc.zip" ^
  --clobber

echo Extracting Deno...

powershell -Command ^
  "Expand-Archive -Force 'deno-x86_64-pc-windows-msvc.zip' '%TOOLS%'"

echo.
echo Cleaning temporary files...

cd /d "%ROOT%"
rmdir /s /q "%TMP%"

echo.
echo Done.
