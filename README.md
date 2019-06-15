# XYZUA

> basic information about **XYZUA** and installation.

## database

> XYZUA use MySQL as database

SQL script for the database is below:

```
-- --------------------------------------------------------
-- Server version:               10.3.4-MariaDB - MariaDB Server
-- Server OS:                    Linux
-- HeidiSQL Version:             9.5.0.5196
-- --------------------------------------------------------

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET NAMES utf8 */;
/*!50503 SET NAMES utf8mb4 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;

-- Dumping database structure for xyzua
CREATE DATABASE IF NOT EXISTS `xyzua` /*!40100 DEFAULT CHARACTER SET utf8 */;
USE `xyzua`;

-- Dumping structure for table xyzua.app
CREATE TABLE IF NOT EXISTS `app` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT 'appid',
  `appname` varchar(64) NOT NULL COMMENT 'application name',
  `secret` varchar(64) NOT NULL COMMENT 'application secret',
  `callback` varchar(512) NOT NULL COMMENT 'callback url',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='application information';

-- Data exporting was unselected.
-- Dumping structure for table xyzua.code
CREATE TABLE IF NOT EXISTS `code` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT 'code id',
  `userid` int(11) unsigned NOT NULL COMMENT 'user id',
  `appid` int(11) unsigned NOT NULL COMMENT 'application id',
  `code` varchar(128) NOT NULL COMMENT 'the code',
  `expire` int(10) unsigned DEFAULT NULL COMMENT 'expire time',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='the trans-application code for authentication';

-- Data exporting was unselected.
-- Dumping structure for table xyzua.mark
CREATE TABLE IF NOT EXISTS `mark` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT 'markid',
  `type` varchar(32) NOT NULL COMMENT 'type of mark',
  `identity` varchar(64) NOT NULL COMMENT 'username or ticket',
  `timestamp` timestamp NOT NULL DEFAULT current_timestamp() COMMENT 'timestamp of this mark',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='mark of ticket/user';

-- Data exporting was unselected.
-- Dumping structure for table xyzua.ticket
CREATE TABLE IF NOT EXISTS `ticket` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT 'id of ticket',
  `ticket` varchar(64) DEFAULT NULL COMMENT 'ticket',
  `info` text DEFAULT NULL COMMENT 'information of ticket',
  `expire` int(10) unsigned DEFAULT NULL COMMENT 'expire time of ticket',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='a kind of temperary identity';

-- Data exporting was unselected.
-- Dumping structure for table xyzua.user
CREATE TABLE IF NOT EXISTS `user` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT 'userid',
  `username` varchar(64) DEFAULT NULL COMMENT 'username',
  `password` varchar(64) DEFAULT NULL COMMENT 'hash of password',
  `info` text DEFAULT NULL COMMENT 'information of user',
  `role` varchar(50) DEFAULT NULL COMMENT 'role of user',
  `authority` tinyint(3) unsigned DEFAULT NULL COMMENT 'authority of user',
  `token` varchar(64) DEFAULT NULL COMMENT 'token for user',
  `expire` int(10) unsigned DEFAULT NULL COMMENT 'token expire time',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='user info';

/*!40101 SET SQL_MODE=IFNULL(@OLD_SQL_MODE, '') */;
/*!40014 SET FOREIGN_KEY_CHECKS=IF(@OLD_FOREIGN_KEY_CHECKS IS NULL, 1, @OLD_FOREIGN_KEY_CHECKS) */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;

```

## web

### run outside China

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