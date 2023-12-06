@echo off


rem 获取当前脚本文件的路径
set "scriptPath=%~dp0"

rem 拼接相对路径成绝对路径
set "absolutePath=%scriptPath%"

echo "pwd: %absolutePath%"
cd %scriptPath%
REM 生成或更新api
goctl api go -dir .\trade_service -api .\trade_api\trade.api
