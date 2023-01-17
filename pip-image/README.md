
## 什么是apk

Alpine Linux 的包管理工具，主要包含两部分内容：

* 包管理：提供安装/更新/删除包的功能 --- apk
* 系统备份恢复：将系统恢复至之前安装和配置的备份状态 --- lbu

命令类别	子命令	命令说明
包管理	add	添加指定包并确认所有依赖满足要求
包管理	del	删除指定包
系统维护	fix	在不改动主要的依赖的情况下进行包的修复或者升级
系统维护	update	从远程仓库获取信息更新本地仓库索引
系统维护	upgrade	对已安装了的包进行更新
系统维护	cache	对缓存进行操作，比如对缺失的包进行缓存或者对于不需要的包进行缓存删除
信息查询	info	对于指定的包进行包或者仓库的详细信息进行显示
信息查询	list	按照指定条件进行包的列表信息显示
信息查询	search	查询相关的包的详细信息
信息查询	dot	生成依赖之间的关联关系图（使用箭头描述）
信息查询	policy	显示包的仓库策略信息
信息查询	stats	显示仓库和包的安装相关的统计信息
仓库管理	index	使用文件生成仓库索引文件
仓库管理	fetch	从全局仓库下载包到本地目录
仓库管理	verify	验证包的完整性和签名信息
仓库管理	manifest	显示package各组成部分的checksum




配置 pip 镜像源

pip config set global.index-url https://mirrors.aliyun.com/pypi/simple/
pip config set install.trusted-host mirrors.aliyun.com


