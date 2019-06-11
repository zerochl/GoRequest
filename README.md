# GoRequest(专为Android与IOS提供的网络请求框架)
An android-ios-based network request framework implemented by golang
# 实现框架的初衷
1. 4.4一下的android机型不支持TLS1.2版本的部分加解密算法，导致无法使用https协议访问接口，使用java也能解决此问题，只是资料太少，耗时过多。
2. 网络框架实现跨平台，Android与IOS甚至服务端只需要一套代码
# 网络框架支持的功能
1. 线程池
可设置MaxIdle与CoreSize，如果线程数量超过MaxIdle会默认使用golang的携程，所以MaxIdle可能会形同虚设，可自行修改程序
2. header
可在初始化传入base header，也可根据不同接口传入特定header
# 使用方式
1. 获取gomobile程序并编译，参见：https://www.cnblogs.com/liuliu-word/p/9463859.html ，网上随便搜搜，多得很
2. cd到go src目录下，执行Android命令：gomobile bind -target=android GoRequest;IOS命令：gomobile bind -target=ios GoRequest;会生成GoRequest.aar和GoRequest.framework文件，对应移动端开发就知道怎么引入到项目中了

### 作者联系方式：QQ：975804495
### 疯狂的程序员群：186305789，没准你能遇到绝影大神
