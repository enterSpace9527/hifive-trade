@echo off

REM 设置目标系统和目标平台
set GOOS=linux
set GOARCH=amd64

rem 获取当前脚本文件的路径
set "scriptPath=%~dp0"

rem 拼接相对路径成绝对路径
set "absolutePath=%scriptPath%trade_service"

echo "pwd: %absolutePath%"
cd %absolutePath%
REM 运行交叉编译命令，修改为你的项目路径和输出文件名
go build -o hifive-trade

REM 恢复本地环境变量
set GOOS=
set GOARCH=