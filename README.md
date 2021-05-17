# oook


## 安装
### 下载安装
https://github.com/Joeyscat/oook/releases

### 使用 go get 命令安装
```bash
> go get github.com/joeyscat/oook
```

## 使用
```bash
# 静态服务器
❯ oook static-server -d=/images -p=8001
Static Server Running on http://127.0.0.1:8001/
```

## TODO
```bash
# 抓包
> oook proxy?? -h=127.0.0.1 -p=8000

# restful api
> oook rest --data=/api.json

# linux/windows 代码风格检查
> oook style-check --path=/code

# golang 项目模板生成
> oook gogen --cli=cobra
Generating cli application with cobra ...
> oook gogen --web=echo --db=mongodb
Generating web application with echo, mongodb ... 

```
