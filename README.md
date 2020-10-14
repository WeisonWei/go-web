# go-web

## Framework

Golang Web Framework: Gin

## Config

`GOPROXY=https://goproxy.cn`

## Api

[GIN-debug] GET    /sd/health                --> go-web/src/handler/sd.HealthCheck (5 handlers)  
[GIN-debug] GET    /sd/disk                  --> go-web/src/handler/sd.DiskCheck (5 handlers)  
[GIN-debug] GET    /sd/cpu                   --> go-web/src/handler/sd.CPUCheck (5 handlers)  
[GIN-debug] GET    /sd/ram                   --> go-web/src/handler/sd.RAMCheck (5 handlers)  

## DB
```sql
CREATE DATABASE /*!32312 IF NOT EXISTS*/ `user` /*!40100 DEFAULT CHARACTER SET utf8 */;
USE `user`;
DROP TABLE IF EXISTS `tb_users`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `tb_users` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `username` varchar(255) NOT NULL,
  `password` varchar(255) NOT NULL,
  `createdAt` timestamp NULL DEFAULT NULL,
  `updatedAt` timestamp NULL DEFAULT NULL,
  `deletedAt` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `username` (`username`),
  KEY `idx_tb_users_deletedAt` (`deletedAt`)
) ENGINE=MyISAM AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

-- Dumping data for table `tb_users`
LOCK TABLES `tb_users` WRITE;
/*!40000 ALTER TABLE `tb_users` DISABLE KEYS */;
INSERT INTO `tb_users` VALUES (0,'admin','$2a$10$veGcArz47VGj7l9xN7g2iuT9TF21jLI1YGXarGzvARNdnt4inC9PG','2018-05-27 16:25:33','2018-05-27 16:25:33',NULL);
/*!40000 ALTER TABLE `tb_users` ENABLE KEYS */;
UNLOCK TABLES;
```
## Quote

[Gin1](https://github.com/WeisonWei/apiserver_demos)  
[Gin2](https://www.jianshu.com/p/a31e4ee25305)  
[Gin3](https://gitbook.cn/gitchat/column/5dab061e7d66831b22aa0b44)  
[Gin4](https://juejin.im/book/6844733730678898702)  