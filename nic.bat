@echo off
go build -o "./build/nic.exe" && ( "./build/nic" %* )