# fc2录播机说明
##前端vue 后端golang


###使用说明
## 运行方法
### linux
需要 GLIBC_2.34 或以上 可以直接运行 录播机这个文件  版本不够需要对GLIBC进行升级  或自己编译golang代码运行  
操作指令如下
下载全部代码和文件  
打开并修改conf文件夹中的conf.ini文件的内容
服务器使用需要将ip地址改为 0.0.0.0  本地使用为127.0.0.1      默认为 ip=127.0.0.1  
端口号根据需求更改                              默认为 port=8080   
上传所有文件到服务器中  切换路径到目录文件夹下  
输入 chmod 777 录播机         
给予他使用限权  
输入 nohup ./录播机 &  
进行后台启动   
输入 exit    
等待5-10秒 重新连接服务器 进入logs文件夹内 查看最新的日志文件(文件格式为 xxxx-xx 年份 当前月份) 查看里面生成的登录密码  
示例  ----创建管理员账号----admin_Create----管理员密码为:123456789  
然后访问 服务器ip+设置的端口号 即可  
管理员账号为 admin     

### win10和以上
下载win版录播机.rar  
解压到文件夹中  
打开并修改conf文件夹中的conf.ini文件的内容  
服务器使用需要将ip地址改为 0.0.0.0  本地使用为127.0.0.1      默认为 ip=127.0.0.1  
端口号根据需求更改                                              默认为 port=8080  
然后双击win版录播机.exe  
等待5-10秒 进入logs文件夹内 查看最新的日志文件(文件格式为 xxxx-xx 年份 当前月份) 查看里面生成的登录密码   
示例  ----创建管理员账号----admin_Create----管理员密码为:123456789  
然后访问 服务器ip+设置的端口号 即可  
管理员账号为 admin     

### 自行编译
#### golang版本1.19.3
下载全部代码文件 解压到文件夹中  
打开并修改conf文件夹中的conf.ini文件的内容  
切换路径到文件夹内  
///////////////////////////////////////////////////////////////  
直接运行  
go run main.go  
后台运行  
输入 go run main.go &   
///////////////////////////////////////////////////////////////  
编译后运行  
go build main.go  
输入 chmod 777 main     
编译完成后    
输入 nohup ./main &   
///////////////////////////////////////////////////////////////  
等待5-10秒 进入logs文件夹内 查看最新的日志文件(文件格式为 xxxx-xx 年份 当前月份) 查看里面生成的登录密码    
示例  ----创建管理员账号----admin_Create----管理员密码为:123456789    
然后访问 服务器ip+设置的端口号 即可   
管理员账号为 admin       



## 具体使用方法  
打开网站后 进入设置  
可以直连的服务器或地区 直接点击检测连通性 通过即可      
不能直连的地区   请设置代理  然后点击检测连通性 通过 点击保存 即可   

点击up主列表页面   
点击添加新UP主按钮   
up主名称随便填写  
房间号  打开up主页面 查看浏览器网址栏中的id 然后复制粘贴即可  
示例  https://xxx.xxx.xxx/1748145052/  中的  1748145052  
点击提交  
添加弹窗消失 并更新了信息 代表添加成功  
添加弹窗未消失 代表 up名称或房间号 已存在  添加失败  

添加成功后  自动开始检测直播状态  1-2分钟检测一次开播状态    
直播完成后  就可以去视频列表中 下载对应的直播视频了  


##出现bug或使用问题 请在Issues中反馈