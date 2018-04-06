# hunter
movie hunter


## 功能愿景
1. 用于树莓派新从电影天堂获取科幻片磁力链接
1. 新片通过邮件推送给用户。邮件里介绍新片的标题，评分，简介等信息，并提供可点击的连接。用户点击链接即可下发下载任务

## 关键实现
1. 下载的功能使用qtorrent这个软件实现
1. 远程下载使用的时qtorrent的[http API](https://github.com/qbittorrent/qBittorrent/wiki/Web-API-Documentation)实现，相当丰富
1. 使用go语言及第三方辅助包开发

## 其他舒适建议（不听你就不会舒适。。。）
1. 建议使用树莓派，功耗低，可以全年运行。
1. (大坑，慎入啊兄台)舍弃SD卡，因为长时间高负荷运行，寿命一般不超过3个月。
建议按照[官网教程](https://www.raspberrypi.org/documentation/hardware/raspberrypi/bootmodes/msd.md)使用移动硬盘代替(需要外接个5w电源)
1. 开发一个服务端，搭建于私有VPS。实时监控树莓派的状态，列举可下载的电影列表，并提供触发下载的链接。（代码后续看情况开放）
1. 使用Supervisor来启动你的app。
1. 使用Samba来共享树莓派资源。使用VLC这个APP（IOS）来播放电影。