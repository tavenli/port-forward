# port-forward
Go语言开发的端口转发工具  for port data forward

```
开发语言：GO
控制台框架：beego
```

# 功能介绍

> 支持Web控制台添加端口映射

> 支持对每条端口映射进行开启和关闭控制

> 支持API接口，方便与其它系统集成

# 已编译好的程序包

[port-forward.win64.v1.0.zip](http://files.git.oschina.net/group1/M00/01/4E/PaAvDFktmFmADf8FAFiePV2ray4682.zip?token=95b2ff6047714107a757f930495c9530&ts=1496192518&attname=port-forward.win64.v1.0.zip)

[port-forward.linux64.v1.0.zip](http://files.git.oschina.net/group1/M00/01/4E/PaAvDFktmK6Ac0A1AFDn6l67Jso205.zip?token=06b1de3e885eec779e2c9dbd965a5239&ts=1496192669&attname=port-forward.linux64.v1.0.zip)

# 使用交流群

> 欢迎大家就使用问题或个性化需求在QQ群中讨论，QQ群号：99134862

# 快速安装说明
1. 下载编译好的程序包，并解压程序包
2. 在执行程序包目录下找 data/PortForwardDb.sql 的数据库创建文件，创建好数据库 PortForwardDb
3. 修改 conf/data.conf 中的数据库连接串，主要是修改连接MySQL的用户名和密码
4. 执行 start.sh （Linux）或 start.bat （Win）命令
5. 打开浏览器，进入控制台，打开 http://127.0.0.1:8000/login
6. 输入用户 admin  密码 123456 进入控制台


# 控制台UI
![登录](http://git.oschina.net/tavenli/port-forward/raw/master/screenshot/Login.png "在这里输入图片标题")


![转发列表](http://git.oschina.net/tavenli/port-forward/raw/master/screenshot/List.png "在这里输入图片标题")


![端口转发配置](http://git.oschina.net/tavenli/port-forward/raw/master/screenshot/edit.png "在这里输入图片标题")


![方便与其它平台集成接口](http://git.oschina.net/tavenli/port-forward/raw/master/screenshot/ApiDoc.png "在这里输入图片标题")

