# oauth2 
oauth2 是什么


oauth2 的实现分两种， 
* 自己的网站获取资源网站的资源，这时候需要实现oauth2 client端。
* 自己的网站需要编程资源端，给其他oauth2 client端提供资源，需要实现oauth2的server端。



## 解释
* 第三方网站，资源请求站：资源的请求方，通常是你实现的网站，比如你想让用户用微信在你的网站上登录
* 资源提供网站: 比如微信
* 用户: 微信账号拥有者

* scope 你需要获得的用户信息的范围，会在授权的时候提示给用户

## oauth2 的工作机制:

`refresh_token` vs `access_token|auth token`:
两者都有一定的失效时间, refresh token 有效时间比 access token 时间长, 这意味着 access token 会首先失效掉。
这时候就需要拿着refresh token再去获得一次新的 access token。
我猜测使用 refresh token 的目的是为了不损失安全性上, 给用户以最好的体验，毕竟不用重复性的登录。


client端请求资源站:
* 去资源网站申请 appid, secret
没有获得过授权或者用户revoke掉了:
* 请求 refresh token, 携带`client_id scope ...` 信息
* 拿到 refresh token 发送auth请求, 获得 access token
* 资源站会通过url回调将 access token和refresh token, 还有用户信息传回。拿到 access token之后就可以通过 bearer token 方式去发请求了
  获得过授权但是在系统(你自己实现的)退出了:

server端实现步骤:
* 提供申请 `client_id, client_secret` 的地址, 需要用户提供 `redirect_uri` app名称, logo， 描述等信息
* 提供 refresh token 发放地址, 需要用户携带 `client_id` 这些信息, 然后随机生成一个 refresh token
  通过url回调的方式将refresh token携带，转入 `redirect_uri`。
  * 设置 refresh token 失效的时间，scope 。
  * 如果失效的 refresh token ，那么就会回到原点，通常是弹出用户授权页面。将 refresh token 对应的信息保存
  * 如果 refresh token 未失效，则直接以 302 转入 redirect url, 携带新的 refresh code。
* 提供 access token 的发放地址，用户需要提供 `refresh_token, client_id, client_secret, redirect_uri`, 
  * 匹配保存的 refresh token， 确定失效与否， 
  * 匹配提供的 redirect_uri 是否和申请 `client_id， client_secret`的时候一致,
  * 如果都ok， 提供 access token, 设置失效时间




## oauth2 的安全性:
* hacker即使获得了第三方网站的secret 和 id 也没有方式修改 redirect_uri 的规则



## oauth2 实现:

### oauth2 client端实现:
以[Github](https://developer.github.com/v3/oauth/)为例:

* 获得refresh token, 发送请求, 带上`client_id ...` 参数
  
    GET https://github.com/login/oauth/authorize

    如果用户没有授权过，那么就会打开一个tab页面, 这个页面是github提供的登录页面, 
    如果用户授权过，但是auth token没有失效, 那么返回一个302

* 获得access_token(auth token): 参数 `client_id, client_secret, code(第一步返回的), redirect_uri, ...`

    POST https://github.com/login/oauth/access_token

    返回信息: 
    Accept: application/json
    {"access_token":"e72e16c7e42f292c6912e7710c838347ae178b4a", "scope":"repo,gist", "token_type":"bearer"}

* 然后就可以用access_token来代替用户向资源网站发送请求了, 在 header 中添加 Authorization 字段，或者在url中附加 `access_token` 请求参数

    curl -H "Authorization: token OAUTH-TOKEN" https://api.github.com/user


### oauth2 client端代码实现:




# Links
* [passport-github](https://github.com/CrazyFork/passport-github)
* [passport-oauth2](https://github.com/CrazyFork/passport-oauth2)
* [passport-oauth2-client-password: 用户资源发放网站验证请求者身份](https://github.com/jaredhanson/passport-oauth2-client-password)
















