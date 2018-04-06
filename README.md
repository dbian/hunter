# hunter
movie hunter


## 功能愿景
1. 用于树莓派从电影天堂获取科幻片磁力链接
1. 新片通过邮件推送给用户。邮件里介绍新片的标题，评分，简介等信息，并提供可点击的连接。用户点击链接即可下发下载任务
1. 移动端使用VLC播放树莓派通过samba共享的电影。


## 配置文件
1. 配置文件在根目录，config.json，顾名思义是json格式
1. 下载需要使用qtorrent，将它的web ui功能打开，输入账号密码，并且填入本工程的config.json
1. 填入邮箱的账号密码，用于自动发送邮件通知


## 关键实现
1. 下载的功能使用qtorrent这个软件实现
1. 远程下载使用的时qtorrent的[v3 http API](https://github.com/qbittorrent/qBittorrent/wiki/WebUI-API-Documentation)实现，相当丰富
1. 使用go语言及第三方辅助包开发

## 其他舒适建议（不然很蛋疼。。。）
1. 使用树莓派，功耗低，可以全年运行。其他的低功耗服务器亦可。
1. (大坑，慎入啊兄台)舍弃SD卡，因为长时间高负荷运行，寿命一般不超过3个月。
建议按照[官网教程](https://www.raspberrypi.org/documentation/hardware/raspberrypi/bootmodes/msd.md)使用移动硬盘代替(需要外接个5w电源)
1. 开发一个服务端，搭建于私有VPS。实时监控树莓派的状态，列举可下载的电影列表，并提供触发下载的链接。（代码后续放出）
1. 使用Supervisor来启动你的app。
1. 使用Samba来共享树莓派资源。IOS使用VLC这个APP来播放电影。

## TODO
1. 适配qtorrent其他版本的API
1. 功能更加开放，可定制化功能更多
1. 欢迎提交PR(带上测试代码。。。)