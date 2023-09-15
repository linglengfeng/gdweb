@echo off

:: copy proto_go ..\web3Server\proto_go
 xcopy proto_go ..\web3Server\proto_go /s /e
 xcopy proto_go ..\webServer\proto_go /s /e
:: echo "copy successed..."

pause