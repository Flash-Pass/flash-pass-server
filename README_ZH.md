# Flash Pass Server

## 本地开发须知
1. 环境变量：在本地开发，编译项目之前需要设置环境变量 `FLASH_PASS_USER`，对应的值为用户名（xinzhengfei、yucong），如需设置其他值请事先与管理员沟通。


2. 配置中心：项目运行前，通过 SSH 客户端首先完成 server 连接，再编译运行项目，否则项目将无法连接配置中心（原因：未开放 nacos 直连端口）。
```
ssh -L 18848:localhost:18848 CbwYYmlU@82.156.171.8
输入密码：b%b=C]'OPuSH
```

## 项目部署须知
1. test & pro 环境部署时需要设置环境变量：`FLASH_PASS_NAMESPACE`。


2. 环境变量 `FLASH_PASS_USER` 指针对于本地开发时生效，用于 db 连接时匹配数据库，实现用户隔离。test & pro 环境不使用该变量。

