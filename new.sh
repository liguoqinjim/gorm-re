#!/bin/bash

set -e
# 生成model
go run main.go

# 复制到路径
cp model.go /Users/li/Workspace/weixin_tool/weixin_tool_server/server/datamodels