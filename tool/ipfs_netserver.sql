-- MySQL dump 10.13  Distrib 8.0.16, for Win64 (x86_64)
--
-- Host: localhost    Database: ipfs_netsever
-- ------------------------------------------------------
-- Server version	8.0.15

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
 SET NAMES utf8 ;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `cloud`
--

DROP TABLE IF EXISTS `cloud`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
 SET character_set_client = utf8mb4 ;
CREATE TABLE `cloud` (
  `Dhash` varchar(45) NOT NULL COMMENT 'ipfs网络下，用户的id哈希',
  `status` int(11) NOT NULL DEFAULT '3' COMMENT '身份：\\n1. edge\\n2. gateway\\n3. cloud\\n4. other',
  `ip` varchar(45) NOT NULL COMMENT 'IP地址\\ntcp下包含端口号\\nxxx.xxx.xxx.xxx:port',
  `capacity` int(11) NOT NULL DEFAULT '500' COMMENT '总容量 （MB）',
  `remain` int(11) NOT NULL DEFAULT '400' COMMENT '剩余容量(MB)',
  PRIMARY KEY (`Dhash`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='云层';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `cloud`
--

LOCK TABLES `cloud` WRITE;
/*!40000 ALTER TABLE `cloud` DISABLE KEYS */;
INSERT INTO `cloud` VALUES ('11',1,'1',1,1);
/*!40000 ALTER TABLE `cloud` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `edge`
--

DROP TABLE IF EXISTS `edge`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
 SET character_set_client = utf8mb4 ;
CREATE TABLE `edge` (
  `Dhash` varchar(45) NOT NULL COMMENT 'ipfs网络下，用户的id哈希',
  `status` int(11) NOT NULL DEFAULT '1' COMMENT '身份：\\n1. edge\\n2. gateway\\n3. cloud\\n4. other',
  `ip` varchar(45) NOT NULL COMMENT 'IP地址\\ntcp下包含端口号\\nxxx.xxx.xxx.xxx:port',
  `capacity` int(11) NOT NULL DEFAULT '500' COMMENT '总容量 （MB）',
  `remain` int(11) NOT NULL DEFAULT '400' COMMENT '剩余容量(MB)',
  PRIMARY KEY (`Dhash`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='边缘层';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `edge`
--

LOCK TABLES `edge` WRITE;
/*!40000 ALTER TABLE `edge` DISABLE KEYS */;
/*!40000 ALTER TABLE `edge` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `gateway`
--

DROP TABLE IF EXISTS `gateway`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
 SET character_set_client = utf8mb4 ;
CREATE TABLE `gateway` (
  `Dhash` varchar(45) NOT NULL COMMENT 'ipfs网络下，用户的id哈希',
  `status` int(11) NOT NULL DEFAULT '2' COMMENT '身份：\\n1. edge\\n2. gateway\\n3. cloud\\n4. other',
  `ip` varchar(45) NOT NULL COMMENT 'IP地址\\ntcp下包含端口号\\nxxx.xxx.xxx.xxx:port',
  `capacity` int(11) NOT NULL DEFAULT '500' COMMENT '总容量 （MB）',
  `remain` int(11) NOT NULL DEFAULT '400' COMMENT '剩余容量(MB)',
  PRIMARY KEY (`Dhash`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='网关层';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `gateway`
--

LOCK TABLES `gateway` WRITE;
/*!40000 ALTER TABLE `gateway` DISABLE KEYS */;
/*!40000 ALTER TABLE `gateway` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `ownership`
--

DROP TABLE IF EXISTS `ownership`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
 SET character_set_client = utf8mb4 ;
CREATE TABLE `ownership` (
  `Dhash` varchar(45) NOT NULL COMMENT '用户哈希值',
  `Fhash` varchar(45) NOT NULL COMMENT '文件哈希值',
  PRIMARY KEY (`Dhash`,`Fhash`),
  KEY `FDASH_idx` (`Fhash`),
  CONSTRAINT `Dhash` FOREIGN KEY (`Dhash`) REFERENCES `cloud` (`Dhash`),
  CONSTRAINT `FDASH` FOREIGN KEY (`Fhash`) REFERENCES `video` (`Fhash`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='视频存储情况';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `ownership`
--

LOCK TABLES `ownership` WRITE;
/*!40000 ALTER TABLE `ownership` DISABLE KEYS */;
INSERT INTO `ownership` VALUES ('11','44');
/*!40000 ALTER TABLE `ownership` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `video`
--

DROP TABLE IF EXISTS `video`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
 SET character_set_client = utf8mb4 ;
CREATE TABLE `video` (
  `Uhash` int(11) NOT NULL COMMENT 'IPFS网络中，up主的id哈希值',
  `Fhash` varchar(45) NOT NULL COMMENT 'IPFS网络中，文件哈希值',
  `title` varchar(45) DEFAULT NULL COMMENT '文件标题',
  `time` int(11) NOT NULL COMMENT '上传时间戳',
  PRIMARY KEY (`Fhash`),
  UNIQUE KEY `ipfshash_UNIQUE` (`Fhash`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='视频文件登记表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `video`
--

LOCK TABLES `video` WRITE;
/*!40000 ALTER TABLE `video` DISABLE KEYS */;
INSERT INTO `video` VALUES (44,'44','1',1);
/*!40000 ALTER TABLE `video` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2020-04-30 16:19:49
