# yeego


关于测试: 为了避免有时候将敏感信息传导上面，所以，建了一个conf文件夹，将自己的敏感信息放到conf里面，

然后同时修改
```
func Init() {
	yeego.MustInitConfig("/Users/vannnnish/go_project/project_sh/yeego/conf", "conf")
	appId = yeego.Config.GetString("wechat.AppId")
	appSecret = yeego.Config.GetString("wechat.AppSecret")
	mchId = yeego.Config.GetString("wechat.MchId")
	key = yeego.Config.GetString("wechat.Key")
}
```

## 对标准库的封装，

### 例如: http, strings ,strconv, 加密库的一些封装。

### 对第三方包的封装，

### 第三方服务的封装，例如，微信支付 ，阿里云短信，submail等