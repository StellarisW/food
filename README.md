# ð é¤é¥®æå¡åç«¯

## ð¡ é¡¹ç®ç®ä»

å®ç°äºç±»ä¼¼äºç¾å¢çä¸äºåè½ï¼å¦æç´¢é¤åï¼æç´¢èè°±ç­api

## ð¼ ä»£ç æ¶æ

### é¡¹ç®ç»æ

<details>
<summary>å±å¼æ¥ç</summary>
<pre>
<code>
    âââ app ----------------------------- (é¡¹ç®æä»¶)
    	âââ api ------------------------- (apiå±)
    		âââ recipe ------------------ (å³äºé£è°±çapi)
    		âââ restaurant -------------- (å³äºé¤åçapi)
    		âââ user -------------------- (å³äºç¨æ·çapi)
    	âââ global ---------------------- (å¨å±ç»ä»¶)
    	âââ internal -------------------- (åé¨å)
    		âââ middleware -------------- (ä¸­é´ä»¶)
    		âââ model ------------------- (æ¨¡å)
    		âââ service ----------------- (æå¡å±)
    	âââ router ---------------------- (è·¯ç±å±)
    âââ boot ---------------------------- (é¡¹ç®å¯å¨å)
    âââ manifest ------------------------ (äº¤ä»æ¸å)
    	âââ config ---------------------- (é¡¹ç®éç½®)
		âââ sql ------------------------- (sqlæä»¶)
			âââ mongodb ----------------- (mongodbæ°æ®é)
			âââ mysql ------------------- (mysqlè¡¨ç»æ)
    âââ utils --------------------------- (å·¥å·å)
    âââ build.sh ------------------------ (é¡¹ç®å¯å¨shellèæ¬)
    âââ docker-compose.yml -------------- (dockerå®¹å¨)
</code>
</pre>
</details>

### ææ¯æ 

<img src="https://raw.githubusercontent.com/gin-gonic/logo/master/color.png" width="15%">

- [gin](https://github.com/gin-gonic/gin)

> `Gin`æ¯ä¸ä¸ªç¨Goè¯­è¨ç¼åçwebæ¡æ¶ãå®æ¯ä¸ä¸ªç±»ä¼¼äº`martini`ä½æ¥ææ´å¥½æ§è½çAPIæ¡æ¶, ç±äºä½¿ç¨äº`httprouter`ï¼éåº¦æé«äºè¿40åã å¦æä½ æ¯æ§è½åé«æçè¿½æ±è, ä½ ä¼ç±ä¸`Gin`ã

[Ginæ¡æ¶ä»ç»åä½¿ç¨-ææå¨çåå®¢](https://www.liwenzhou.com/posts/Go/gin/)

[è§é¢æç¨](https://www.bilibili.com/video/BV1gJ411p7xC/)

<img src="http://jwt.io/img/logo-asset.svg">

- jwt

> SON Web Token (JWT)æ¯ä¸ä¸ªå¼æ¾æ å(RFC 7519)ï¼å®å®ä¹äºä¸ç§ç´§åçãèªåå«çæ¹å¼ï¼ç¨äºä½ä¸ºJSONå¯¹è±¡å¨åæ¹ä¹é´å®å¨å°ä¼ è¾ä¿¡æ¯ãè¯¥ä¿¡æ¯å¯ä»¥è¢«éªè¯åä¿¡ä»»ï¼å ä¸ºå®æ¯æ°å­ç­¾åçã

[äºåéå¸¦ä½ äºè§£å¥æ¯JWT](https://zhuanlan.zhihu.com/p/86937325)

[JSON Web Token å¥é¨æç¨](https://www.ruanyifeng.com/blog/2018/07/json_web_token-tutorial.html)

[jwt.io](https://jwt.io/) å¯ä»¥å¨è¿ä¸ªç½ç«æ ¡éª jwt

<img src="https://miro.medium.com/max/700/1*YcVBLTidq861sJhIlVby5w.png" width="50%">

- [zap](https://github.com/uber-go/zap)

> `zap`æ¯`Uber`å¼åçéå¸¸å¿«çãç»æåçï¼åæ¥å¿çº§å«çGoæ¥å¿åºãæ ¹æ®Uber-go Zapçææ¡£ï¼å®çæ§è½æ¯ç±»ä¼¼çç»æåæ¥å¿åæ´å¥½ï¼ä¹æ¯æ ååºæ´å¿«ãå·ä½çæ§è½æµè¯å¯ä»¥å»`github`ä¸çå°ã

[ä½¿ç¨zapæ¥æ¶ginæ¡æ¶é»è®¤çæ¥å¿å¹¶éç½®æ¥å¿å½æ¡£](https://www.liwenzhou.com/posts/Go/use_zap_in_gin/)

[æ·±å¥æµægolang zap æ¥å¿åºä½¿ç¨ï¼å«æä»¶åå²ãåçº§å«å­å¨åå¨å±ä½¿ç¨ç­ï¼](https://www.yisu.com/zixun/154695.html)

<img src="https://github.com/spf13/viper/raw/master/.github/logo.png?raw=true" width="50%">

- [viper](https://github.com/spf13/viper)

> Viperæ¯éç¨äºGoåºç¨ç¨åºçå®æ´éç½®è§£å³æ¹æ¡ãå®è¢«è®¾è®¡ç¨äºå¨åºç¨ç¨åºä¸­å·¥ä½ï¼å¹¶ä¸å¯ä»¥å¤çææç±»åçéç½®éæ±åæ ¼å¼ã

[Goè¯­è¨éç½®ç®¡çç¥å¨ââViperä¸­ææç¨](https://www.liwenzhou.com/posts/Go/viper_tutorial/)

<img src="https://upload.wikimedia.org/wikipedia/zh/thumb/6/62/MySQL.svg/1200px-MySQL.svg.png" width="30%">

- [mysql](https://www.mysql.com/)

> ä¸ä¸ªå³ç³»åæ°æ®åºç®¡çç³»ç»ï¼ç±çå¸MySQL AB å¬å¸å¼åï¼å±äº Oracle æä¸äº§åãMySQL æ¯ææµè¡çå³ç³»åæ°æ®åºç®¡çç³»ç»å³ç³»åæ°æ®åºç®¡çç³»ç»ä¹ä¸ï¼å¨ WEB åºç¨æ¹é¢ï¼MySQLæ¯æå¥½ç RDBMS (Relational Database Management Systemï¼å³ç³»æ°æ®åºç®¡çç³»ç») åºç¨è½¯ä»¶ä¹ä¸

[Goæä½MySQL](https://www.liwenzhou.com/posts/Go/go_mysql/)

[sqlxåºä½¿ç¨æå](https://www.liwenzhou.com/posts/Go/sqlx/)

[GORMå¥é¨æå](https://www.liwenzhou.com/posts/Go/gorm/)

[GORMä¸­æææ¡£](https://gorm.io/zh_CN/docs/)

<img src="https://upload.wikimedia.org/wikipedia/en/thumb/6/6b/Redis_Logo.svg/1200px-Redis_Logo.svg.png" width="40%">

- [redis](https://redis.io/)

> ä¸ä¸ªå¼æºçãä½¿ç¨Cè¯­è¨ç¼åçãæ¯æç½ç»äº¤äºçãå¯åºäºåå­ä¹å¯æä¹åçKey-Valueæ°æ®åº

[Goè¯­è¨æä½Redis](https://www.liwenzhou.com/posts/Go/redis/)

<img src="https://upload.wikimedia.org/wikipedia/commons/thumb/0/00/Mongodb.png/1200px-Mongodb.png" width="40%">

- [MongoDB](https://www.mongodb.com/)

> ææ¡£åæ°æ®åºï¼æå´è¶£çå¯ä»¥èªå·±å»äºè§£

[Goè¯­è¨æä½mongoDB](https://www.liwenzhou.com/posts/Go/go_mongodb/)

<img src="https://developers.redhat.com/sites/default/files/styles/article_feature/public/blog/2014/05/homepage-docker-logo.png?itok=zx0e-vcP" width="30%">

- [docker](https://www.docker.com/)

> Google å¬å¸æ¨åºç Go è¯­è¨ è¿è¡å¼åå®ç°ï¼åºäº Linux åæ ¸ç cgroupï¼namespaceï¼ä»¥å AUFS ç±»ç Union FS ç­ææ¯çä¸ä¸ªå®¹å¨æå¡

â	å®¹å¨ç¨docker-composeé¨ç½²

## ð åè½æ¨¡å

### APIææ¡£

[apipost](https://console-docs.apipost.cn/preview/75cf02d1bc40f846/9a5d4e23d3fc1ea7)
