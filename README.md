# ib-auth

> refer:
>
> https://interactivebrokers.github.io/cpwebapi/
>
> https://interactivebrokers.com/api/doc.html

## 描述

管理账户登录状态

1. Keeping Session Alive

```http request
GET /v1/api/sso/validate
```

## 使用

```bash
docker pull ibgambler/ib-auth
docker run -e HOST=https://{url to your host}:5000 ibgambler/ib-auth
```