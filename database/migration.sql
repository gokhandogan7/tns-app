-- MySQL dump 10.13  Distrib 8.0.31, for Win64 (x86_64)
--
-- Host: 127.0.0.1    Database: firstdatabase
-- ------------------------------------------------------
-- Server version	8.0.31

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `articles`
--

DROP TABLE IF EXISTS `articles`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `articles` (
  `id` int NOT NULL,
  `title` varchar(45) NOT NULL,
  `description` varchar(145) DEFAULT NULL,
  `context` varchar(245) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `idarticles_UNIQUE` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `articles`
--

LOCK TABLES `articles` WRITE;
/*!40000 ALTER TABLE `articles` DISABLE KEYS */;
INSERT INTO `articles` VALUES (1,'title 1','description 1','context 1'),(3,'title 3','description 1','context 1'),(4,'title 4','description 4','content 4'),(7,'title 7','description 7','context 7'),(9,'Title 10','Description 10 new','new content with create method'),(10,'Title 10','Description 10 new','new content with create method'),(11,'Newly Created Post','The description for my new post','my articles content'),(14,'Title 10','Description 10 new','new content with create method'),(15,'yenisfsgsfafas','yenisdfasdfsfahod','yeni fadfasfmetfasfashod'),(18,'article2.Title','article2.Desc',''),(19,'article2.Title','article2.Desc','article2.Context'),(23,'title 3','description 1','context 1'),(24,'article2.Title','article2.Degfgdfhsc',''),(26,'article2.Title','article2.Degfgdfhsc','article2.Context'),(31,'Newly Created Post','The description for my new post','my articles content'),(45,'yeni title upgrade','yeni metdasdfsfahod','yeni metfasfashod'),(127,'article2.Title','Description degisti2','article2.Context'),(227,'article2.Title','Description degisti2 new title','article2.Context'),(345,'Newly Created Post','The description for my new post','my articles content');
/*!40000 ALTER TABLE `articles` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2022-10-20 13:44:02
