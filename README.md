# sample-go-api
an go-api-based sample 

# feature
- url:          https://staging.p2shop.com.cn/sample/fruits
- swagger 2.0:  https://staging.p2shop.com.cn/sample/docs/index.html
- jwt url:      https://staging.p2shop.com.cn/sample/v1/fruits
- sign get:     https://staging.p2shop.com.cn/sample/sign?name=xiao.xinmiao&sign=3F7EC1885326B9D1FD078DB2276C84D6&authors=%5B%7B"age"%3A18%2C"name"%3A"eland8"%7D%2C%7B"name"%3A%20"eland2"%2C"age"%3A%2019%7D%5D
- sign post:    https://staging.p2shop.com.cn/sample/sign
``` 
{
    "name": "xiao.xinmiao", 
    "sign": "52357A730D11306FDEF190FD95A8D4E9", 
    "authors": [
        {
            "age": 18, 
            "name": "eland8"
        }, 
        {
            "name": "eland2", 
            "age": 19
        }
    ]
}
```


# tips

## Before submitting the code, please download the go-api code
```
git remote add a-upstream https://gitlab.p2shop.cn:8443/sample/go-api.git
git fetch a-upstream master
git rebase upstream/master a-upstream/master
```
