# XYZUA

A Union Authentication System, covering SSO services and other features.

## Features

- HMAC Login with SHA256/MD5
- Time sacrifice defense
- Separated Front-End (**Vue**-based) and Back-End (**Golang**-based)
- SSO Services
- Temperary Authentication **Ticket**
- Various **Mark** Operation

## Installation

1. XYZUA use MySQL as database, please create a database named ```xyzua```

2. run ```git clone https://github.com/phantomlsh/xyzua``` and ```go get -u```

3. create ```./conf/app.conf``` and write the config:
```
appname = xyzua
httpport = [port]
runmode = pro
autorender = false
copyrequestbody = true
sessionon = true
SessionGCMaxLifetime = 60
sqlconn = [username]:[password]@([mysql host]:3306)/[database]
```

4. run ```bee run``` or ```bee build``` or ```bee pack```

> Note: When running Outside China, please do the following:

change ```/web/home.html <head>``` :
```
<head>
  <title>HOME | XYZUA</title>
  <meta charset="utf-8">
  <meta http-equiv="X-UA-Compatible" content="IE=edge">
  <meta name="viewport" content="width=device-width,initial-scale=1.0">
  <link href="https://cdnjs.cloudflare.com/ajax/libs/material-design-icons/3.0.1/iconfont/material-icons.min.css" rel="stylesheet">
  <link rel="stylesheet" type="text/css" href="./home.css">
  <script src="https://cdnjs.cloudflare.com/ajax/libs/vue/2.6.10/vue.min.js"></script>
  <script src="https://cdnjs.cloudflare.com/ajax/libs/axios/0.19.0/axios.min.js"></script>
</head>
```

change ```/web/index.html <head>``` :
```
<head>
  <title>LOGIN | XYZUA</title>
  <meta charset="utf-8">
  <meta http-equiv="X-UA-Compatible" content="IE=edge">
  <meta name="viewport" content="width=device-width,initial-scale=1.0">
  <link rel="stylesheet" type="text/css" href="./index.css">
  <script src="https://cdnjs.cloudflare.com/ajax/libs/vue/2.6.10/vue.min.js"></script>
  <script src="https://cdnjs.cloudflare.com/ajax/libs/axios/0.19.0/axios.min.js"></script>
  <script src="https://cdnjs.cloudflare.com/ajax/libs/jshashes/1.0.7/hashes.min.js"></script>
</head>
```

> change ```/web/changepassword.html <head>``` as ```index.html``` except the ```<title>```
> change ```/web/ticket.html <head>``` as ```home.html``` except the ```<title>```