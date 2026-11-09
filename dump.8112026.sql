-- MySQL dump 10.13  Distrib 26.10.0, for Linux (x86_64)
--
-- Host: localhost    Database: gymapp2
-- ------------------------------------------------------
-- Server version	26.10.0

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;
SET @MYSQLDUMP_TEMP_LOG_BIN = @@SESSION.SQL_LOG_BIN;
SET @@SESSION.SQL_LOG_BIN= 0;

--
-- GTID state at the beginning of the backup 
--

SET @@GLOBAL.GTID_PURGED=/*!80000 '+'*/ '8deed65a-d6e5-11f1-8fde-0242ac110003:1-460';

--
-- Table structure for table `api_users`
--

DROP TABLE IF EXISTS `api_users`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `api_users` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `login` varchar(255) NOT NULL,
  `avatar_url` text,
  `url` text,
  `tipo` varchar(50) DEFAULT NULL,
  `nome` varchar(255) DEFAULT NULL,
  `empresa` varchar(255) DEFAULT NULL,
  `blog` text,
  `localizacao` varchar(255) DEFAULT NULL,
  `email` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `api_users`
--

LOCK TABLES `api_users` WRITE;
/*!40000 ALTER TABLE `api_users` DISABLE KEYS */;
INSERT INTO `api_users` VALUES (1,'KaduHod','https://avatars.githubusercontent.com/u/68855956?v=4','https://api.github.com/users/KaduHod','User','Carlos Alberto Ribas Junior','NEXCORE','https://kadu.tec.br/','Curitiba',NULL);
/*!40000 ALTER TABLE `api_users` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `articulation_movement`
--

DROP TABLE IF EXISTS `articulation_movement`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `articulation_movement` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `articulation_id` int unsigned NOT NULL,
  `movement_id` int unsigned NOT NULL,
  PRIMARY KEY (`id`),
  KEY `articulation_articulation_movement_foreign_key` (`articulation_id`),
  KEY `movement_articulation_movement_foreign_key` (`movement_id`),
  CONSTRAINT `articulation_articulation_movement_foreign_key` FOREIGN KEY (`articulation_id`) REFERENCES `articulations` (`id`),
  CONSTRAINT `movement_articulation_movement_foreign_key` FOREIGN KEY (`movement_id`) REFERENCES `movements` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=164 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `articulation_movement`
--

LOCK TABLES `articulation_movement` WRITE;
/*!40000 ALTER TABLE `articulation_movement` DISABLE KEYS */;
INSERT INTO `articulation_movement` VALUES (1,1,5),(20,1,6),(21,1,4),(22,1,3),(32,1,1),(33,1,2),(34,1,25),(35,2,7),(36,2,8),(37,2,9),(38,2,10),(39,2,15),(40,2,16),(41,2,31),(42,2,32),(43,2,30),(44,3,7),(45,3,8),(46,3,17),(47,3,18),(48,3,21),(49,3,20),(50,4,17),(51,4,18),(52,5,18),(53,5,2),(54,5,21),(55,5,20),(56,5,9),(57,5,10),(58,5,30),(59,6,7),(60,6,8),(61,6,9),(62,6,10),(63,7,7),(64,8,8),(65,8,7),(66,8,9),(67,8,10),(68,8,22),(69,9,7),(70,9,8),(71,9,9),(72,9,10),(73,10,10),(74,10,8),(75,11,8),(76,11,7),(77,11,23),(78,11,25),(79,12,7),(80,12,8),(81,12,23),(82,12,25),(83,13,7),(84,13,8),(85,13,23),(86,13,25),(87,14,7),(88,14,8),(89,14,9),(90,14,10),(91,14,15),(92,14,16),(93,15,7),(94,15,8),(95,16,27),(96,16,26),(97,16,28),(98,16,29),(99,17,28),(100,17,29),(101,17,27),(102,17,26),(103,18,7),(104,18,8),(105,19,7),(106,19,8),(107,20,4),(108,20,3),(109,20,1),(110,20,2),(111,20,25),(112,21,4),(113,21,3),(114,22,4),(115,22,3),(116,22,1),(117,22,2),(118,22,5),(119,22,6),(120,23,4),(121,23,3),(122,23,1),(123,23,2),(124,23,5),(125,23,6),(126,24,7),(127,24,8),(128,24,23),(129,24,25),(130,25,7),(131,25,8),(132,25,23),(133,25,25),(134,26,33),(135,26,34),(136,26,35),(137,26,25),(138,27,25),(139,27,23),(140,27,38),(141,27,37),(142,28,39),(143,29,7),(144,29,8),(145,29,23),(146,29,25),(147,30,4),(148,30,3),(149,30,16),(150,30,15),(151,31,7),(152,31,8),(153,31,23),(154,32,1),(155,32,2),(156,32,25),(157,32,40),(158,32,41),(159,5,7),(160,5,8),(161,15,42),(162,2,43),(163,2,32);
/*!40000 ALTER TABLE `articulation_movement` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `articulation_movement_muscle`
--

DROP TABLE IF EXISTS `articulation_movement_muscle`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `articulation_movement_muscle` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `muscle_portion_id` int unsigned NOT NULL,
  `articulation_id` int unsigned NOT NULL,
  `movement_id` int unsigned DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `muscle_portion_id` (`muscle_portion_id`,`articulation_id`,`movement_id`),
  KEY `articulation_foreign_key` (`articulation_id`),
  KEY `movement_foreign_key` (`movement_id`),
  CONSTRAINT `articulation_foreign_key` FOREIGN KEY (`articulation_id`) REFERENCES `articulations` (`id`),
  CONSTRAINT `movement_foreign_key` FOREIGN KEY (`movement_id`) REFERENCES `movements` (`id`),
  CONSTRAINT `muscle_portion_foreign_key` FOREIGN KEY (`muscle_portion_id`) REFERENCES `muscle_portion` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=273 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `articulation_movement_muscle`
--

LOCK TABLES `articulation_movement_muscle` WRITE;
/*!40000 ALTER TABLE `articulation_movement_muscle` DISABLE KEYS */;
INSERT INTO `articulation_movement_muscle` VALUES (1,1,2,7),(176,1,2,8),(2,1,2,9),(3,1,2,15),(183,2,2,7),(6,2,2,8),(4,2,2,9),(5,2,2,15),(184,3,2,7),(9,3,2,8),(7,3,2,9),(8,3,2,15),(185,4,2,7),(186,4,2,8),(10,4,23,1),(11,4,23,3),(12,4,23,6),(179,5,1,1),(181,5,1,5),(259,5,1,6),(187,5,2,7),(188,5,2,8),(14,5,23,1),(15,5,23,5),(256,5,23,6),(218,5,23,7),(219,5,23,8),(189,6,2,7),(190,6,2,8),(16,6,20,3),(191,7,2,7),(17,7,2,8),(18,7,2,9),(255,7,2,10),(19,7,2,15),(20,7,23,2),(192,8,2,7),(193,8,2,8),(21,8,23,4),(194,9,2,7),(195,9,2,8),(22,9,13,7),(23,9,13,8),(24,9,13,23),(25,9,13,25),(254,10,1,3),(180,10,1,4),(209,10,1,5),(196,10,2,7),(197,10,2,8),(28,10,12,8),(214,10,20,7),(215,10,20,8),(26,10,22,4),(216,10,22,7),(217,10,22,8),(212,10,23,4),(27,10,23,5),(29,11,1,2),(30,11,1,5),(198,11,2,7),(199,11,2,8),(31,12,1,2),(32,12,1,3),(33,12,1,5),(200,12,2,7),(201,12,2,8),(257,12,23,NULL),(210,13,1,2),(202,13,2,7),(203,13,2,8),(34,13,23,2),(35,13,23,6),(178,14,1,2),(182,14,1,6),(204,14,2,7),(205,14,2,8),(36,14,23,2),(37,14,23,6),(208,15,1,4),(206,15,2,7),(207,15,2,8),(40,15,12,23),(38,15,22,4),(213,15,23,4),(39,15,23,6),(50,16,2,7),(258,16,2,8),(48,16,3,7),(220,16,3,8),(247,16,4,17),(49,16,4,18),(51,17,2,7),(221,17,2,8),(52,17,3,7),(240,17,3,8),(248,17,4,17),(53,17,4,18),(54,18,3,7),(241,18,3,8),(55,19,2,8),(222,19,3,7),(56,19,3,8),(266,20,2,NULL),(177,20,3,7),(57,20,3,8),(225,21,2,7),(226,21,2,8),(211,21,3,7),(58,21,3,8),(60,22,3,7),(59,22,4,17),(64,23,5,7),(263,23,5,8),(61,23,5,21),(65,24,5,7),(264,24,5,8),(63,24,5,20),(66,25,5,7),(265,25,5,8),(77,26,14,10),(78,27,14,10),(267,28,14,7),(79,28,14,8),(80,28,14,16),(88,30,13,7),(87,30,13,25),(268,30,26,34),(81,31,13,8),(82,31,13,23),(83,31,13,25),(269,31,26,34),(84,32,13,7),(86,32,13,23),(85,32,13,25),(270,32,26,34),(89,33,13,7),(91,33,13,23),(90,33,13,25),(271,33,26,34),(92,34,13,7),(260,34,13,8),(272,34,26,34),(239,35,13,7),(93,35,13,8),(94,35,13,23),(95,35,13,25),(99,36,14,8),(101,36,14,16),(100,36,15,7),(249,36,15,8),(104,37,14,8),(102,37,14,15),(103,37,15,7),(250,37,15,8),(105,38,14,8),(107,38,14,15),(106,38,15,7),(251,38,15,8),(108,39,15,8),(109,40,15,8),(110,41,15,8),(112,42,14,7),(252,42,15,7),(111,42,15,8),(123,43,15,7),(121,43,16,26),(244,43,16,27),(124,44,16,26),(245,44,16,27),(126,45,2,7),(130,46,2,7),(127,46,2,10),(132,46,2,16),(131,46,2,43),(137,48,11,7),(139,48,11,23),(138,48,11,25),(140,49,11,8),(141,49,11,23),(142,49,11,25),(113,50,14,7),(114,50,14,10),(115,51,14,7),(116,51,15,7),(41,52,2,31),(67,53,5,7),(68,53,5,21),(69,54,5,7),(70,54,5,21),(71,55,5,8),(72,55,5,20),(144,56,14,7),(143,56,14,9),(145,56,14,15),(147,57,14,7),(146,57,14,9),(148,57,14,15),(151,58,14,7),(150,58,14,8),(149,58,14,9),(152,58,14,15),(154,59,14,7),(153,59,14,9),(156,59,14,15),(155,59,15,7),(238,60,2,15),(237,60,2,16),(43,61,2,16),(46,62,2,8),(44,62,2,9),(45,62,2,15),(96,63,14,7),(97,63,14,9),(98,63,14,15),(157,64,14,16),(158,65,14,16),(159,66,14,16),(160,67,14,16),(161,68,14,16),(246,69,16,NULL),(125,69,17,27),(47,70,2,15),(261,70,2,16),(163,71,13,8),(162,71,14,7),(253,71,14,8),(164,71,14,16),(117,72,15,7),(118,72,15,42),(128,73,2,8),(135,73,2,16),(129,73,2,32),(73,74,3,7),(242,74,3,8),(74,74,4,17),(75,74,4,18),(262,75,3,NULL),(243,75,4,17),(76,75,4,18),(167,76,14,7),(166,76,14,8),(165,76,14,9),(168,77,2,2),(169,77,2,9),(170,77,2,15),(171,77,23,2),(172,78,2,2),(173,78,2,9),(174,78,2,15),(175,78,23,2);
/*!40000 ALTER TABLE `articulation_movement_muscle` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `articulation_movement_muscle_portion`
--

DROP TABLE IF EXISTS `articulation_movement_muscle_portion`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `articulation_movement_muscle_portion` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `muscle_portion_id` int unsigned NOT NULL,
  `articulation_movement_id` int unsigned NOT NULL,
  `role` enum('agonist','synergist','antagonist') NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `unique_relation` (`muscle_portion_id`,`articulation_movement_id`,`role`),
  KEY `articulation_movement_id_muscle_portion_articulation_movement` (`articulation_movement_id`),
  CONSTRAINT `articulation_movement_id_muscle_portion_articulation_movement` FOREIGN KEY (`articulation_movement_id`) REFERENCES `articulation_movement` (`id`),
  CONSTRAINT `muscle_portion_articulation_movement_muscle_portion` FOREIGN KEY (`muscle_portion_id`) REFERENCES `muscle_portion` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=14 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `articulation_movement_muscle_portion`
--

LOCK TABLES `articulation_movement_muscle_portion` WRITE;
/*!40000 ALTER TABLE `articulation_movement_muscle_portion` DISABLE KEYS */;
INSERT INTO `articulation_movement_muscle_portion` VALUES (5,1,35,'agonist'),(6,1,37,'agonist'),(7,1,39,'agonist'),(10,2,36,'agonist'),(8,2,37,'agonist'),(9,2,39,'agonist'),(13,3,36,'agonist'),(11,3,37,'agonist'),(12,3,39,'agonist');
/*!40000 ALTER TABLE `articulation_movement_muscle_portion` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `articulations`
--

DROP TABLE IF EXISTS `articulations`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `articulations` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `createdAt` datetime DEFAULT CURRENT_TIMESTAMP,
  `updatedAt` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=33 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `articulations`
--

LOCK TABLES `articulations` WRITE;
/*!40000 ALTER TABLE `articulations` DISABLE KEYS */;
INSERT INTO `articulations` VALUES (1,'Escápula e Clavícula','2022-11-20 03:00:00','2022-11-20 03:00:00'),(2,'Ombro (Glenoumeral)','2022-11-20 03:00:00','2022-11-20 03:00:00'),(3,'Cotovelo','2022-11-20 03:00:00','2022-11-20 03:00:00'),(4,'Antebraço (Rádio-ulnar)','2022-11-20 03:00:00','2022-11-20 03:00:00'),(5,'Punho e Intercarpais','2022-11-20 03:00:00','2022-11-20 03:00:00'),(6,'Metacarpofalangiana dos Dedos','2022-11-20 03:00:00','2022-11-20 03:00:00'),(7,'Interfalangiana dos Dedos','2022-11-20 03:00:00','2022-11-20 03:00:00'),(8,'Carpometacarpiana do Polegar','2022-11-20 03:00:00','2022-11-20 03:00:00'),(9,'Metacarpofalangiana do Polegar','2022-11-20 03:00:00','2022-11-20 03:00:00'),(10,'Interfalangiana do Polegar','2022-11-20 03:00:00','2022-11-20 03:00:00'),(11,'Pescoço (Atlanto-occipital e Atlantoaxial)','2022-11-20 03:00:00','2022-11-20 03:00:00'),(12,'Coluna Cervical','2022-11-20 03:00:00','2022-11-20 03:00:00'),(13,'Coluna Torácica e Lombar','2022-11-20 03:00:00','2022-11-20 03:00:00'),(14,'Quadril','2022-11-20 03:00:00','2022-11-20 03:00:00'),(15,'Joelho','2022-11-20 03:00:00','2022-11-20 03:00:00'),(16,'Tornozelo','2022-11-20 03:00:00','2022-11-20 03:00:00'),(17,'Intertarsais do Pé','2022-11-20 03:00:00','2022-11-20 03:00:00'),(18,'Metatarsofalangiana dos Dedos do Pé','2022-11-20 03:00:00','2022-11-20 03:00:00'),(19,'Interfalangiana dos Dedos do Pé','2022-11-20 03:00:00','2022-11-20 03:00:00'),(20,'Esternoclavicular','2023-03-17 14:28:06',NULL),(21,'Costocondral','2023-03-17 14:28:38',NULL),(22,'Acromioclavicular','2023-03-17 14:29:16',NULL),(23,'Escapulotorácica','2023-03-17 14:30:11',NULL),(24,'Toracolombar','2023-03-17 14:30:38',NULL),(25,'Crista Ilíaca','2023-03-17 14:31:00',NULL),(26,'Pelve','2023-03-17 14:34:15',NULL),(27,'Sacroilíaca','2023-03-17 14:53:09',NULL),(28,'Crista Púbica','2023-03-17 14:54:29',NULL),(29,'Vértebras','2023-03-17 14:58:36',NULL),(30,'Costelas','2023-03-17 15:00:23',NULL),(31,'Processo Xifoide','2023-03-17 18:34:45','2023-03-17 18:34:52'),(32,'Temporomandibular','2023-03-17 18:48:35',NULL);
/*!40000 ALTER TABLE `articulations` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Temporary view structure for view `auxiliar`
--

DROP TABLE IF EXISTS `auxiliar`;
/*!50001 DROP VIEW IF EXISTS `auxiliar`*/;
SET @saved_cs_client     = @@character_set_client;
/*!50503 SET character_set_client = utf8mb4 */;
/*!50001 CREATE VIEW `auxiliar` AS SELECT 
 1 AS `id_exercicio`,
 1 AS `nome_exercicio`,
 1 AS `grupo_muscular`,
 1 AS `porcao_muscular`,
 1 AS `articulacao`,
 1 AS `movimento`*/;
SET character_set_client = @saved_cs_client;

--
-- Table structure for table `cache`
--

DROP TABLE IF EXISTS `cache`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `cache` (
  `key` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `value` mediumtext CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `expiration` int NOT NULL,
  PRIMARY KEY (`key`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `cache`
--

LOCK TABLES `cache` WRITE;
/*!40000 ALTER TABLE `cache` DISABLE KEYS */;
/*!40000 ALTER TABLE `cache` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `cache_locks`
--

DROP TABLE IF EXISTS `cache_locks`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `cache_locks` (
  `key` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `owner` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `expiration` int NOT NULL,
  PRIMARY KEY (`key`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `cache_locks`
--

LOCK TABLES `cache_locks` WRITE;
/*!40000 ALTER TABLE `cache_locks` DISABLE KEYS */;
/*!40000 ALTER TABLE `cache_locks` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `exercicios`
--

DROP TABLE IF EXISTS `exercicios`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `exercicios` (
  `id` int unsigned NOT NULL,
  `createdAt` datetime DEFAULT CURRENT_TIMESTAMP,
  `updatedAt` datetime DEFAULT NULL,
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `force` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `link` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `execution` text CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci,
  `mechanic` text CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `exercicios`
--

LOCK TABLES `exercicios` WRITE;
/*!40000 ALTER TABLE `exercicios` DISABLE KEYS */;
INSERT INTO `exercicios` VALUES (1,'2023-03-19 22:33:40',NULL,'Bench Press',NULL,NULL,NULL,NULL),(2,'2023-03-19 22:33:40',NULL,'Incline Bench Press',NULL,NULL,NULL,NULL),(3,'2023-03-19 22:33:40',NULL,'Dumbbell Fly',NULL,NULL,NULL,NULL),(4,'2023-03-19 22:33:40',NULL,'Incline Dumbbell Press',NULL,NULL,NULL,NULL),(5,'2023-03-19 22:33:40',NULL,'Chest Dip',NULL,NULL,NULL,NULL),(6,'2023-03-19 22:33:40',NULL,'Cable Crossover',NULL,NULL,NULL,NULL),(7,'2023-03-19 22:33:40',NULL,'Decline Bench Press',NULL,NULL,NULL,NULL),(8,'2023-03-19 22:33:40',NULL,'Decline Dumbbell Fly',NULL,NULL,NULL,NULL),(9,'2023-03-19 22:33:40',NULL,'Push-Up',NULL,NULL,NULL,NULL),(10,'2023-03-19 22:33:40',NULL,'Dumbbell Pullover',NULL,NULL,NULL,NULL),(11,'2023-03-19 22:33:40',NULL,'Incline Dumbbell Fly',NULL,NULL,NULL,NULL),(12,'2023-03-19 22:33:40',NULL,'Dumbbell Bench Press with Neutral Grip',NULL,NULL,NULL,NULL),(13,'2023-03-19 22:33:40',NULL,'Push-Up with Serratus Focus',NULL,NULL,NULL,NULL),(14,'2023-03-19 22:33:40',NULL,'Dumbbell Punches',NULL,NULL,NULL,NULL),(15,'2023-03-19 22:33:40',NULL,'Scapular Protraction on Cable Machine',NULL,NULL,NULL,NULL),(16,'2023-03-19 22:33:40',NULL,'Barbell Shrugs',NULL,NULL,NULL,NULL),(17,'2023-03-19 22:33:40',NULL,'Cable Cross-Over',NULL,NULL,NULL,NULL),(18,'2023-03-19 22:33:40',NULL,'Incline Dumbbell Fly',NULL,NULL,NULL,NULL),(19,'2023-03-19 22:33:40',NULL,'Pull-Up (Barra Fixa)',NULL,NULL,NULL,NULL),(20,'2023-03-19 22:33:40',NULL,'Remada Curvada com Barra',NULL,NULL,NULL,NULL),(21,'2023-03-19 22:33:40',NULL,'Remada Baixa no Cabo',NULL,NULL,NULL,NULL),(22,'2023-03-19 22:33:40',NULL,'Pull-Apart com Elástico',NULL,NULL,NULL,NULL),(23,'2023-03-19 22:33:40',NULL,'Crucifixo Inverso com Halteres',NULL,NULL,NULL,NULL),(24,'2023-03-19 22:33:40',NULL,'Face Pull',NULL,NULL,NULL,NULL),(25,'2023-03-19 22:33:40',NULL,'Levantamento Terra',NULL,NULL,NULL,NULL),(26,'2023-03-19 22:33:40',NULL,'Hiperextensão Lombar',NULL,NULL,NULL,NULL),(27,'2023-03-19 22:33:40',NULL,'Good Morning (Bom Dia)',NULL,NULL,NULL,NULL),(28,'2023-03-19 22:33:40',NULL,'Encolhimento com Barra',NULL,NULL,NULL,NULL),(29,'2023-03-19 22:33:40',NULL,'Remada Alta',NULL,NULL,NULL,NULL),(30,'2023-03-19 22:33:40',NULL,'Clean and Press',NULL,NULL,NULL,NULL),(31,'2023-03-19 22:33:40',NULL,'Encolhimento com Halteres',NULL,NULL,NULL,NULL),(33,'2023-03-19 22:33:40',NULL,'Remada Invertida',NULL,NULL,NULL,NULL),(34,'2023-03-19 22:33:40',NULL,'Elevação em Y Prono',NULL,NULL,NULL,NULL),(35,'2023-03-19 22:33:40',NULL,'Remada Curvada com Halteres',NULL,NULL,NULL,NULL),(36,'2023-03-19 22:33:40',NULL,'Scapular Wall Slide (Deslizamento Escapular na Parede)',NULL,NULL,NULL,NULL),(37,'2023-03-19 22:33:40',NULL,'Elevação em T Prona',NULL,NULL,NULL,NULL),(38,'2023-03-19 22:33:40',NULL,'Remada Sentada',NULL,NULL,NULL,NULL),(39,'2023-03-19 22:33:40',NULL,'Crucifixo Inverso com Halteres',NULL,NULL,NULL,NULL),(40,'2023-03-19 22:33:40',NULL,'Remada no Cabo em Pé',NULL,NULL,NULL,NULL),(42,'2023-03-19 22:33:40',NULL,'Encolhimento com Barra',NULL,NULL,NULL,NULL),(58,'2023-03-19 22:33:40',NULL,'Tríceps na Polia (Pushdown)',NULL,NULL,NULL,NULL),(59,'2023-03-19 22:33:40',NULL,'Supino Fechado',NULL,NULL,NULL,NULL),(60,'2023-03-19 22:33:40',NULL,'Tríceps Coice com Halter',NULL,NULL,NULL,NULL),(61,'2023-03-19 22:33:40',NULL,'Tríceps Testa com Halteres (Acima da Cabeça)',NULL,NULL,NULL,NULL),(62,'2023-03-19 22:33:40',NULL,'Tríceps Testa (Skullcrusher)',NULL,NULL,NULL,NULL),(63,'2023-03-19 22:33:40',NULL,'Paralelas (Dips)',NULL,NULL,NULL,NULL),(64,'2023-03-19 22:33:40',NULL,'Rosca de Punho Pronada com Barra',NULL,NULL,NULL,NULL),(65,'2023-03-19 22:33:40',NULL,'Rosca de Punho Pronada com Halteres',NULL,NULL,NULL,NULL),(66,'2023-03-19 22:33:40',NULL,'Rosca Inversa com Halteres',NULL,NULL,NULL,NULL),(67,'2023-03-19 22:33:40',NULL,'Rosca Martelo',NULL,NULL,NULL,NULL),(68,'2023-03-19 22:33:40',NULL,'Rosca de Punho',NULL,NULL,NULL,NULL),(69,'2023-03-19 22:33:40',NULL,'Rosca de Punho Inversa',NULL,NULL,NULL,NULL),(70,'2023-03-19 22:33:40',NULL,'Rosca Direta Inversa com Barra',NULL,NULL,NULL,NULL),(71,'2023-03-19 22:33:40',NULL,'Rosca de Punho',NULL,NULL,NULL,NULL),(72,'2023-03-19 22:33:40',NULL,'Rosca de Punho Inversa',NULL,NULL,NULL,NULL),(73,'2023-03-19 22:33:40',NULL,'Rosca de Punho',NULL,NULL,NULL,NULL),(74,'2023-03-19 22:33:40',NULL,'Rosca de Punho Inversa',NULL,NULL,NULL,NULL),(75,'2023-03-19 22:33:40',NULL,'Exercícios para Fortalecimento de Pegada',NULL,NULL,NULL,NULL),(76,'2023-03-19 22:33:40',NULL,'Elevação Lateral deitado de lado',NULL,NULL,NULL,NULL),(77,'2023-03-19 22:33:40',NULL,'Abdução de Quadril no Cabo em Pé',NULL,NULL,NULL,NULL),(78,'2023-03-19 22:33:40',NULL,'Agachamento Búlgaro',NULL,NULL,NULL,NULL),(79,'2023-03-19 22:33:40',NULL,'Clamshell (Concha)',NULL,NULL,NULL,NULL),(80,'2023-03-19 22:33:40',NULL,'Prancha Lateral com Elevação de Perna',NULL,NULL,NULL,NULL),(81,'2023-03-19 22:33:40',NULL,'Abdução de Quadril com Elástico',NULL,NULL,NULL,NULL),(82,'2023-03-19 22:33:40',NULL,'Elevação Pélvica com Barra (Hip Thrust)',NULL,NULL,NULL,NULL),(83,'2023-03-19 22:33:44',NULL,'Levantamento Terra',NULL,NULL,NULL,NULL),(84,'2023-03-19 22:33:49',NULL,'Afundo Caminhando (Walking Lunge)',NULL,NULL,NULL,NULL),(85,'2023-03-19 22:33:51',NULL,'Exercício de Kegel',NULL,NULL,NULL,NULL),(86,'2023-03-19 22:33:51',NULL,'Elevação Pélvica (Bridge)',NULL,NULL,NULL,NULL),(87,'2023-03-19 22:33:53',NULL,'Agachamento',NULL,NULL,NULL,NULL),(88,'2023-03-19 22:33:55',NULL,'Prancha',NULL,NULL,NULL,NULL),(89,'2023-03-19 22:33:56',NULL,'Dead Bug (Morto-Vivo)',NULL,NULL,NULL,NULL),(90,'2023-03-19 22:33:58',NULL,'Bird Dog (Cachorro-Pássaro)',NULL,NULL,NULL,NULL),(91,'2023-03-19 22:33:59',NULL,'Superman',NULL,NULL,NULL,NULL),(92,'2023-03-19 22:34:00',NULL,'Bird Dog',NULL,NULL,NULL,NULL),(93,'2023-03-19 22:34:02',NULL,'Elevação Pélvica (Bridging)',NULL,NULL,NULL,NULL),(94,'2023-03-19 22:34:03',NULL,'Rotação Russa (Russian Twist)',NULL,NULL,NULL,NULL),(95,'2023-03-19 22:34:04',NULL,'Prancha Lateral',NULL,NULL,NULL,NULL),(96,'2023-03-19 22:34:05',NULL,'Abdominal Bicicleta',NULL,NULL,NULL,NULL),(97,'2023-03-19 22:34:05',NULL,'Rotação Russa (Russian Twist)',NULL,NULL,NULL,NULL),(98,'2023-03-19 22:34:06',NULL,'Prancha Lateral',NULL,NULL,NULL,NULL),(99,'2023-03-19 22:34:07',NULL,'Wood Chop (Lenhador)',NULL,NULL,NULL,NULL),(100,'2023-03-19 22:34:07',NULL,'Abdominal (Sit-Up)',NULL,NULL,NULL,NULL),(101,'2023-03-19 22:34:09',NULL,'Elevação de Pernas',NULL,NULL,NULL,NULL),(102,'2023-03-19 22:34:12',NULL,'Abdominal Tradicional (Crunch)',NULL,NULL,NULL,NULL),(103,'2023-03-19 22:34:14',NULL,'Levantamento Terra',NULL,NULL,NULL,NULL),(104,'2023-03-19 22:34:16',NULL,'Extensão de Lombar (Back Extensions)',NULL,NULL,NULL,NULL),(105,'2023-03-19 22:34:17',NULL,'Good Morning (Bom Dia)',NULL,NULL,NULL,NULL),(106,'2023-03-19 22:34:18',NULL,'Levantamento Terra',NULL,NULL,NULL,NULL),(107,'2023-03-19 22:34:19',NULL,'Afundos (Lunges)',NULL,NULL,NULL,NULL),(108,'2023-03-19 22:34:22',NULL,'Good Morning (Bom Dia)',NULL,NULL,NULL,NULL),(109,'2023-03-19 22:34:23',NULL,'Stiff / RDL (Romanian Deadlift)',NULL,NULL,NULL,NULL),(110,'2023-03-19 22:34:26',NULL,'Flexão de Perna em Pé',NULL,NULL,NULL,NULL),(111,'2023-03-19 22:34:28',NULL,'Glute-Ham Raise (GHR)',NULL,NULL,NULL,NULL),(112,'2023-03-19 22:34:31',NULL,'Stiff-Leg Deadlift (Levantamento Terra com Pernas Rígidas)',NULL,NULL,NULL,NULL),(113,'2023-03-19 22:34:33',NULL,'Flexão de Perna Deitado (Lying Leg Curl)',NULL,NULL,NULL,NULL),(114,'2023-03-19 22:34:36',NULL,'Extensão de Lombar',NULL,NULL,NULL,NULL),(115,'2023-03-19 22:34:37',NULL,'Agachamento com Barra',NULL,NULL,NULL,NULL),(116,'2023-03-19 22:34:37',NULL,'Leg Press',NULL,NULL,NULL,NULL),(117,'2023-03-19 22:34:37',NULL,'Afundos (Lunges)',NULL,NULL,NULL,NULL),(118,'2023-03-19 22:34:37',NULL,'Cadeira Extensora',NULL,NULL,NULL,NULL),(119,'2023-03-19 22:34:37',NULL,'Agachamento Hack',NULL,NULL,NULL,NULL),(120,'2023-03-19 22:34:37',NULL,'Subida no Step (Step-up)',NULL,NULL,NULL,NULL),(121,'2023-03-19 22:34:37',NULL,'Agachamento',NULL,NULL,NULL,NULL),(122,'2023-03-19 22:34:37',NULL,'Leg Press',NULL,NULL,NULL,NULL),(123,'2023-03-19 22:34:37',NULL,'Afundos (Lunges)',NULL,NULL,NULL,NULL),(124,'2023-03-19 22:34:37',NULL,'Agachamento',NULL,NULL,NULL,NULL),(125,'2023-03-19 22:34:37',NULL,'Cadeira Extensora',NULL,NULL,NULL,NULL),(126,'2023-03-19 22:34:37',NULL,'Afundos (Lunges)',NULL,NULL,NULL,NULL),(127,'2023-03-19 22:34:37',NULL,'Elevação de Panturrilhas em Pé',NULL,NULL,NULL,NULL),(128,'2023-03-19 22:34:37',NULL,'Elevação de Panturrilhas Sentado',NULL,NULL,NULL,NULL),(129,'2023-03-19 22:34:37',NULL,'Elevação de Panturrilhas Donkey',NULL,NULL,NULL,NULL),(130,'2023-03-19 22:34:37',NULL,'Elevação de Panturrilhas em Pé',NULL,NULL,NULL,NULL),(131,'2023-03-19 22:34:37',NULL,'Elevação de Panturrilhas Sentado',NULL,NULL,NULL,NULL),(132,'2023-03-19 22:34:37',NULL,'Elevação de Panturrilhas Donkey',NULL,NULL,NULL,NULL),(133,'2023-03-19 22:34:37',NULL,'Desenvolvimento com Barra',NULL,NULL,NULL,NULL),(134,'2023-03-19 22:34:37',NULL,'Desenvolvimento Arnold',NULL,NULL,NULL,NULL),(135,'2023-03-19 22:34:37',NULL,'Elevação Lateral',NULL,NULL,NULL,NULL),(136,'2023-03-19 22:34:37',NULL,'Desenvolvimento Militar em Pé',NULL,NULL,NULL,NULL),(137,'2023-03-19 22:34:37',NULL,'Elevação Lateral com Halteres',NULL,NULL,NULL,NULL),(138,'2023-03-19 22:34:37',NULL,'Elevação Lateral no Cabo',NULL,NULL,NULL,NULL),(139,'2023-03-19 22:34:37',NULL,'Remada Alta com Barra',NULL,NULL,NULL,NULL),(140,'2023-03-19 22:34:37',NULL,'Desenvolvimento Sentado com Halteres',NULL,NULL,NULL,NULL),(141,'2023-03-19 22:34:37',NULL,'Face Pull no Cabo',NULL,NULL,NULL,NULL),(142,'2023-03-19 22:34:37',NULL,'Encolhimento com Barra',NULL,NULL,NULL,NULL),(143,'2023-03-19 22:34:37',NULL,'Encolhimento com Halteres',NULL,NULL,NULL,NULL),(144,'2023-03-19 22:34:37',NULL,'Remada Alta na Polia',NULL,NULL,NULL,NULL),(145,'2023-03-19 22:34:37',NULL,'Encolhimento com Barra',NULL,NULL,NULL,NULL),(146,'2023-03-19 22:34:37',NULL,'Remada Alta',NULL,NULL,NULL,NULL),(147,'2023-03-19 22:34:37',NULL,'Crucifixo Inverso (Reverse Flyes)',NULL,NULL,NULL,NULL),(148,'2023-03-19 22:34:37',NULL,'Afundo Lateral (Lateral Lunges)',NULL,NULL,NULL,NULL),(149,'2023-03-19 22:34:37',NULL,'Caminhada Lateral com Elástico',NULL,NULL,NULL,NULL),(150,'2023-03-19 22:34:37',NULL,'Subida no Step',NULL,NULL,NULL,NULL),(151,'2023-03-19 22:34:37',NULL,'Agachamento Búlgaro',NULL,NULL,NULL,NULL),(152,'2023-03-19 22:34:37',NULL,'Flexão de Pernas (Leg Curls)',NULL,NULL,NULL,NULL),(153,'2023-03-19 22:34:37',NULL,'Afundo Lateral',NULL,NULL,NULL,NULL),(154,'2023-03-19 22:34:37',NULL,'Desenvolvimento Sentado com Halteres',NULL,NULL,NULL,NULL),(155,'2023-03-19 22:34:37',NULL,'Elevação Lateral',NULL,NULL,NULL,NULL),(157,'2023-03-19 22:34:37',NULL,'Rosca Direta Inversa com Barra',NULL,NULL,NULL,NULL),(158,'2023-03-19 22:34:37',NULL,'Rosca de Punho',NULL,NULL,NULL,NULL),(159,'2023-03-19 22:34:37',NULL,'Rosca Martelo',NULL,NULL,NULL,NULL),(160,'2023-03-19 22:34:37',NULL,'Rosca Direta Inversa com Barra',NULL,NULL,NULL,NULL),(161,'2023-03-19 22:34:37',NULL,'Rosca de Punho',NULL,NULL,NULL,NULL),(162,'2023-03-19 22:34:37',NULL,'Rosca Martelo',NULL,NULL,NULL,NULL),(163,'2023-03-19 22:34:37',NULL,'Extensão de Punho com Halter',NULL,NULL,NULL,NULL),(164,'2023-03-19 22:34:37',NULL,'Rosca Inversa',NULL,NULL,NULL,NULL),(165,'2023-03-19 22:34:37',NULL,'Rosca de Punho com Barra',NULL,NULL,NULL,NULL),(166,'2023-03-19 22:34:37',NULL,'Agachamento com Barra',NULL,NULL,NULL,NULL),(167,'2023-03-19 22:34:37',NULL,'Afundos (Lunges)',NULL,NULL,NULL,NULL),(168,'2023-03-19 22:34:37',NULL,'Leg Press Sentado',NULL,NULL,NULL,NULL),(169,'2023-03-19 22:34:37',NULL,'Agachamento Búlgaro (Unilaterais)',NULL,NULL,NULL,NULL),(170,'2023-03-19 22:34:37',NULL,'Cadeira Abdutora Sentado',NULL,NULL,NULL,NULL),(171,'2023-03-19 22:34:37',NULL,'Afundo Lateral',NULL,NULL,NULL,NULL),(172,'2023-03-19 22:34:37',NULL,'Levantamento Terra Sumô',NULL,NULL,NULL,NULL),(173,'2023-03-19 22:34:37',NULL,'Agachamento Búlgaro',NULL,NULL,NULL,NULL),(174,'2023-03-19 22:34:37',NULL,'Cadeira Flexora Sentado',NULL,NULL,NULL,NULL),(175,'2023-03-19 22:34:37',NULL,'Agachamento com Barra com Base Estreita',NULL,NULL,NULL,NULL),(176,'2023-03-19 22:34:37',NULL,'Cadeira Flexora Sentado',NULL,NULL,NULL,NULL),(177,'2023-03-19 22:34:37',NULL,'Cadeira Adutora (Parte Interna da Coxa)',NULL,NULL,NULL,NULL),(178,'2023-03-19 22:34:37',NULL,'Rotação Externa de Ombro com Halter',NULL,NULL,NULL,NULL),(179,'2023-03-19 22:34:37',NULL,'Desenvolvimento com Barra',NULL,NULL,NULL,NULL),(180,'2023-03-19 22:34:37',NULL,'Remada Curvada com Barra',NULL,NULL,NULL,NULL),(182,'2023-03-19 22:34:37',NULL,'Rotação Externa de Ombro com Halter',NULL,NULL,NULL,NULL),(183,'2023-03-19 22:34:37',NULL,'Supino Inclinado com Halteres',NULL,NULL,NULL,NULL),(184,'2023-03-19 22:34:37',NULL,'Puxada Alta com Pegada Fechada',NULL,NULL,NULL,NULL),(185,'2023-03-19 22:34:37',NULL,'Remada Curvada',NULL,NULL,NULL,NULL),(186,'2023-03-19 22:34:37',NULL,'Remada Baixa no Cabo',NULL,NULL,NULL,NULL),(187,'2023-03-19 22:34:37',NULL,'Agachamento com Barra',NULL,NULL,NULL,NULL),(188,'2023-03-19 22:34:37',NULL,'Afundos (Lunges)',NULL,NULL,NULL,NULL),(189,'2023-03-19 22:34:37',NULL,'Leg Press',NULL,NULL,NULL,NULL),(190,'2023-03-19 22:34:37',NULL,'Elevação Pélvica (Glute Bridge)',NULL,NULL,NULL,NULL),(191,'2023-03-19 22:34:37',NULL,'Levantamento Terra Unilateral',NULL,NULL,NULL,NULL),(192,'2023-03-19 22:34:37',NULL,'Elevação Lateral deitado de lado',NULL,NULL,NULL,NULL),(193,'2023-03-19 22:34:37',NULL,'Elevação Pélvica (Hip Thrust)',NULL,NULL,NULL,NULL),(194,'2023-03-19 22:34:37',NULL,'Agachamento com Barra',NULL,NULL,NULL,NULL),(195,'2023-03-19 22:34:37',NULL,'Levantamento Terra',NULL,NULL,NULL,NULL),(196,'2023-03-19 22:34:37',NULL,'Levantamento Terra Unilateral',NULL,NULL,NULL,NULL),(197,'2023-03-19 22:34:37',NULL,'Leg Press',NULL,NULL,NULL,NULL),(198,'2023-03-19 22:34:37',NULL,'Caminhada Lateral com Elástico',NULL,NULL,NULL,NULL),(199,'2023-03-19 22:34:37',NULL,'Agachamento Búlgaro',NULL,NULL,NULL,NULL),(200,'2023-03-19 22:34:37',NULL,'Levantamento Terra',NULL,NULL,NULL,NULL),(201,'2023-03-19 22:34:37',NULL,'Subida no Step (Step Up)',NULL,NULL,NULL,NULL),(202,'2023-03-19 22:34:37',NULL,'Caminhada Lateral com Elástico',NULL,NULL,NULL,NULL),(203,'2023-03-19 22:34:37',NULL,'Levantamento Terra Unilateral',NULL,NULL,NULL,NULL),(204,'2023-03-19 22:34:37',NULL,'Elevação de Panturrilhas Sentado',NULL,NULL,NULL,NULL),(205,'2023-03-19 22:34:37',NULL,'Elevação Frontal com Halteres',NULL,NULL,NULL,NULL),(206,'2023-03-19 22:34:37',NULL,'Elevação de Ponta dos Pés',NULL,NULL,NULL,NULL),(207,'2023-03-19 22:34:37',NULL,'Rotação Externa de Ombro com Halter',NULL,NULL,NULL,NULL),(208,'2023-03-19 22:34:37',NULL,'Remada Baixa no Cabo',NULL,NULL,NULL,NULL),(209,'2023-03-19 22:34:37',NULL,'Elevação Lateral Inclinado (Bent-Over Lateral Raise)',NULL,NULL,NULL,NULL),(210,'2023-03-19 22:34:37',NULL,'Agachamento com Barra',NULL,NULL,NULL,NULL),(211,'2023-03-19 22:34:37',NULL,'Elevação de Pernas',NULL,NULL,NULL,NULL),(212,'2023-03-19 22:34:37',NULL,'Afundo (Lunge)',NULL,NULL,NULL,NULL),(213,'2023-03-19 22:34:37',NULL,'Elevação de Panturrilhas em Pé',NULL,NULL,NULL,NULL),(214,'2023-03-19 22:34:37',NULL,'Cadeira Flexora Sentado',NULL,NULL,NULL,NULL),(215,'2023-03-19 22:34:37',NULL,'Afundo Reverso (Reverse Lunge)',NULL,NULL,NULL,NULL);
/*!40000 ALTER TABLE `exercicios` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `exercise`
--

DROP TABLE IF EXISTS `exercise`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `exercise` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  `description` text NOT NULL,
  `info_link` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`)
) ENGINE=InnoDB AUTO_INCREMENT=200 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `exercise`
--

LOCK TABLES `exercise` WRITE;
/*!40000 ALTER TABLE `exercise` DISABLE KEYS */;
INSERT INTO `exercise` VALUES (84,'Rosca Martelo','Bicep curls are weight training exercises designed to strengthen the biceps brachii, brachialis, and brachioradialis muscles by bending the arm at the elbow. The effectiveness of the exercise depends on forearm rotation, with the biceps being strongest when the forearm is supinated (palms up). A proper bicep curl involves lifting a weight with a supinated grip, keeping the torso upright and elbows stationary, and using a full range of motion to maximize muscle activation. Variations like half-range curls can emphasize mid-range tension, while exercises like preacher curls and incline dumbbell curls target specific parts of the biceps. Common equipment includes dumbbells, which isolate muscles and allow a wider range of motion, and barbells, which stabilize the body and enable heavier lifting. Injuries, such as bicep tendon tears or ulnar neuropathy, can occur due to improper form, excessive weight, or overuse. Maintaining correct technique is crucial to avoid strain or damage.','https://en.wikipedia.org/wiki/hammer_curl'),(85,'Leg Press','The leg press is a compound weight training exercise where an individual pushes a weight or resistance away using their legs, typically performed on a leg press machine. It targets major lower body muscle groups, including the quadriceps, hamstrings, glutes, and calves, with variations in foot positioning and sled angle allowing emphasis on specific muscles. The exercise is effective for building lower body strength and can help improve squat performance. It can also be performed unilaterally (one-legged) to address strength imbalances, enhance stability, and mimic athletic movements. Additionally, a standing leg press variation involves pushing against a wall or resistance with one leg while the other remains on the floor, offering a functional alternative to traditional machine-based leg presses.','https://en.wikipedia.org/wiki/leg_press'),(86,'Elevação de Panturrilhas','Calf raises are an exercise targeting the triceps surae (gastrocnemius and soleus), tibialis posterior, and peroneal muscles through plantar flexion (ankle extension). They can be performed with a bent knee, often seated, to emphasize the soleus, as the gastrocnemius is less engaged in this position. Resistance can be added using weights or bodyweight, such as by draping one leg over the other. Standing calf raises, done with a straighter knee, engage the gastrocnemius more and are often performed on a raised surface to increase the stretch. These can be done with bodyweight or added resistance, such as barbells or specialized machines, and may require balance support. The leg press machine can also be used for straight-knee calf raises, focusing on the foot\'s movement while keeping the hip and knee immobile. The gastrocnemius, composed of fast-twitch fibers, benefits from heavy loads and low reps, while the soleus, a slow-twitch muscle, responds better to higher reps and lower loads.','https://en.wikipedia.org/wiki/calf_raise'),(87,'Desenvolvimento Arnold','The overhead press, also known as the shoulder press, strict press, or military press, is an upper-body weight training exercise where a weight is pressed overhead while standing or seated. It primarily targets the anterior deltoid muscles of the shoulder. The lift involves holding a barbell, dumbbells, or kettlebells at shoulder level and pressing it upward. Performing the exercise standing engages more muscles due to the need for balance, while seated variations are also common. Variations include the push press, which uses leg momentum, and unilateral or alternating presses with one arm at a time. The standing overhead press was once part of Olympic weightlifting but was removed in 1972 due to judging challenges.','https://en.wikipedia.org/wiki/arnold_press'),(88,'Supino Reto','The bench press is a weight training exercise where a person presses a weight upward while lying on a bench, primarily targeting the pectoralis major, anterior deltoids, and triceps brachii. Stabilizing muscles in the back, legs, and core also assist. Typically performed with a barbell or dumbbells, it is a key lift in powerlifting alongside the squat and deadlift and the sole lift in Paralympic powerlifting. Widely used in bodybuilding and strength training, the bench press enhances upper body strength, power, and muscle development.','https://en.wikipedia.org/wiki/Bench_press'),(89,'Remada Alta','The upright row is a weight training exercise that targets the trapezius, deltoids, and biceps. It involves lifting a weight with an overhand grip straight up to the collarbone. A narrower grip emphasizes the trapezius, while a wider grip engages the deltoids more. This exercise can be performed using barbells, dumbbells, kettlebells, or a cable machine. However, due to the internal rotation of the humerus during the movement, it may worsen shoulder impingement syndrome and is often considered contraindicated for individuals with shoulder issues. Many trainers recommend alternatives like shoulder shrugs for trapezius development and lateral raises for deltoids. If pain occurs during the exercise, it should be stopped immediately, as it may indicate a rotator cuff disorder.','https://en.wikipedia.org/wiki/upright_row'),(90,'Rosca Direta com Barra','Bicep curls are weight training exercises designed to strengthen the biceps brachii, brachialis, and brachioradialis muscles by bending the arm at the elbow. The effectiveness of the exercise depends on forearm rotation, with the biceps being most engaged when the palms face upward. A proper bicep curl involves lifting a weight with a supinated grip, keeping the torso upright and elbows stationary to maintain tension on the biceps. Full range of motion is generally recommended for maximum muscle activation, though advanced trainers may use partial ranges for specific emphasis. Variations like dumbbell curls, barbell curls, and preacher curls target the biceps differently, with dumbbells allowing greater isolation and barbells enabling heavier loads. Common injuries from bicep curls include biceps tendon tears, often caused by excessive weight or poor form, and ulnar neuropathy due to nerve compression. Proper technique and avoiding \"ego lifting\" are crucial to prevent injuries.','https://en.wikipedia.org/wiki/barbell_curl'),(91,'Desenvolvimento Militar','The overhead press, also known as the shoulder press, strict press, or military press, is an upper-body weight training exercise where a weight is pressed overhead while standing or seated. It primarily targets the anterior deltoid muscles of the shoulder. The lift involves holding a barbell, dumbbells, or kettlebells at shoulder level and pressing them upward. Performing the exercise standing requires more muscle engagement for balance compared to the seated version. Variations include the push press, which uses leg momentum, and unilateral or alternating presses, where weights are lifted one arm at a time. The standing overhead press was once part of Olympic weightlifting but was removed in 1972 due to judging difficulties.','https://en.wikipedia.org/wiki/military_press'),(92,'Rosca Direta com Halteres','Bicep curls are weight training exercises designed to strengthen the biceps brachii, brachialis, and brachioradialis muscles by bending the arm at the elbow. The effectiveness of the exercise depends on forearm rotation, with the biceps being most engaged when the forearm is supinated (palms facing up). A proper bicep curl involves lifting a weight with a supinated grip, keeping the torso upright and elbows stationary to maintain tension on the biceps. Full range of motion is generally recommended for maximum muscle activation, though advanced trainers may use partial ranges for specific emphasis. Variations like dumbbell curls allow for isolation and a wider range of motion, while barbell curls provide stability and enable heavier lifting. Common injuries from bicep curls include biceps tendon tears, often caused by excessive weight or poor form, and ulnar neuropathy due to nerve compression. Proper technique and avoiding overloading are crucial to prevent injuries.','https://en.wikipedia.org/wiki/dumbbell_curl'),(93,'Cadeira Flexora','The leg curl, or hamstring curl, is an isolation exercise that targets the hamstring muscles by flexing the lower leg against resistance toward the buttocks. It can be performed in three variations: seated, lying, or standing leg curls. This exercise specifically strengthens the hamstrings, and other exercises like the glute-ham raise and deadlift can also be used to target these muscles.','https://en.wikipedia.org/wiki/leg_curl_machine'),(94,'Rosca Concentrada','Bicep curls are weight training exercises designed to strengthen the biceps brachii, brachialis, and brachioradialis muscles by bending the arm at the elbow. The effectiveness of the exercise depends on forearm rotation, with the biceps being most engaged when the forearm is supinated (palms facing up). A full repetition involves curling the weight until the elbow is fully flexed and then slowly lowering it back to the starting position, maintaining an upright torso and stationary elbows to ensure proper tension on the biceps. Variations in range of motion, such as full or half-range, can be used to target specific muscle activation, with full range generally recommended for maximizing biceps engagement. Different equipment, like dumbbells or barbells, can be used, with dumbbells allowing for greater isolation and range of motion, while barbells enable heavier loading. Common injuries from bicep curls include biceps tendon tears, often caused by excessive weight or poor form, and ulnar neuropathy due to nerve compression during the exercise. Proper technique and avoiding \"ego lifting\" are crucial to prevent injuries.','https://en.wikipedia.org/wiki/concentration_curl'),(95,'Rosca de Punho','The wrist curl is a weight training exercise targeting the wrist flexor muscles in the front of the forearm. It is an isolation exercise often paired with the reverse wrist curl to ensure balanced development of forearm muscles. The exercise can be performed with a dumbbell or barbell. In the seated version, the lifter sits on a bench with forearms resting on the thighs or bench, palms facing up, and hands hanging off the edge. The wrist is bent back, and the hand is raised using forearm muscles, then slowly lowered. The standing version involves holding a barbell behind the body with palms facing backward, performing the same motion with a smaller range of motion but emphasizing peak contraction of the target muscles.','https://en.wikipedia.org/wiki/wrist_curl'),(96,'Extensão de Lombar','A back extension is an exercise targeting the lower, mid, and upper back, primarily engaging the erector spinae muscles, which run along the spine and consist of three smaller muscles: spinalis, longissimus, and iliocostalis. Often referred to as \"hyperextensions,\" this exercise is commonly performed using a hyperextension bench in gyms. However, the term is misleading, as hyperextension implies extending a joint beyond its normal range of motion, which is not the goal of this exercise. Instead, back extensions aim to extend the spine within its normal range, with the head and neck remaining neutral at the end range. Overextending the back beyond its normal range can be harmful, potentially leading to lumbar disc issues or muscle spasms. The exercise can be performed with or without equipment.','https://en.wikipedia.org/wiki/back_extension'),(97,'Remada na Máquina','Indoor rowing machines, or ergometers, simulate the action of rowing for exercise or training. They measure the work performed by the rower, often in metrics like distance, calories, or watts. Modern rowers typically use a flywheel for resistance, with variations like air, magnetic, or water resistance models. Rowing on these machines provides a full-body aerobic workout, engaging the legs, core, and arms. The rowing stroke consists of four phases: the catch (starting position), the drive (power generation), the finish (completing the stroke), and the recovery (returning to the starting position). Proper technique is essential to avoid injury, with emphasis on maintaining a neutral back posture and coordinating leg, core, and arm movements. Indoor rowing is widely used for fitness, competitive training, and performance testing, with standardized tests like the 2000-meter ergometer test commonly used to evaluate rowers. Competitions, such as the World Rowing Indoor Championships, feature events like the 2000-meter race and team challenges, often categorized by age, sex, and weight class.','https://en.wikipedia.org/wiki/rowing_machine'),(98,'Elevação Pélvica (Hip Thrust)','The pelvic thrust is a motion involving the pelvic region, used in various activities such as dance, exercise, and sexual activity. In the context of exercise, hip thrusts are a specific movement that targets the gluteus maximus muscle. To perform this exercise, the athlete assumes a reclined position and thrusts their hips upward to lift weights balanced on their lap. This exercise is effective for strengthening and developing the glute muscles.','https://en.wikipedia.org/wiki/hip_thrust'),(178,'Supino Reto (Teste)','The bench press is a weight training exercise where a person presses a weight upward while lying on a bench, primarily targeting the pectoralis major, anterior deltoids, and triceps brachii. Stabilizing muscles in the back, legs, and core also assist. Typically performed with a barbell or dumbbells, it is a key lift in powerlifting alongside the squat and deadlift and the sole lift in Paralympic powerlifting. Widely used in bodybuilding and strength training, the bench press enhances upper body strength, power, and muscle development.','teste'),(183,'Barra Fixa (Pull-Up)','                        A pull-up is an upper-body strength exercise. The pull-up is a closed-chain movement where the body is suspended by the hands, gripping a bar or other implement at a distance typically wider than shoulder-width, and pulled up. As this happens, the elbows flex and the shoulders adduct and extend to bring the elbows to the torso.\n\nPull-ups build up several muscles of the upper body, including the latissimus dorsi, trapezius, and biceps brachii. A pull-up may be performed with overhand (pronated), underhand (supinated)—sometimes referred to as a chin-up—neutral, or rotating hand position.\n\nPull-ups are used by some organizations as a component of fitness tests, and as a conditioning activity for some sports.\n        \n        ','https://en.wikipedia.org/wiki/Pull-up'),(186,'Tríceps Testa','The lying triceps extension, also known as skull crusher and French extension or French press, is a strength exercise used in many different forms of strength training. It is one of the most stimulating exercises to the entire triceps muscle group in the upper arm,[citation needed] and works the triceps from the elbow all the way to the latissimus dorsi. Due to its full use of the triceps muscle group, the lying triceps extensions are used by many as part of their training regimen.\n\nTriceps extensions are isolation exercises, meaning they use just one joint. They can also be beneficial for fixing imbalances in the triceps or rehabilitating from injury. In bodybuilding, this exercise is used to target the triceps for growth.[1]\n\nA study sponsored and done by the American Council on Exercise revealed that the extensions are around 70–90 percent effective compared to the triangle push up for the triceps.[2] However, extensions put no pressure on the wrists so they are an alternative for people with wrist strain or injury.[1] Despite the name, skull crushers are a safe, manageable way to effectively work the triceps.','https://en.wikipedia.org/wiki/Lying_triceps_extension'),(189,'Flexão de Braço','The push-up (press-up in British English) is a common calisthenics exercise beginning from the prone position. By raising and lowering the body using the arms, push-ups exercise the pectoral muscles, triceps, and anterior deltoids, with ancillary benefits to the rest of the deltoids, serratus anterior, coracobrachialis and the midsection as a whole.[1] Push-ups are a basic exercise used in civilian athletic training or physical education and commonly in military physical training. It is also a common form of punishment used in the military, school sport, and some martial arts disciplines for its humiliating factor (when one fails to do a specified amount) and for its lack of equipment. Variations of push-ups, such as wide-arm push-ups, diamond push-ups target specific muscle groups and provide further challenges.','https://en.wikipedia.org/wiki/Push-up');
/*!40000 ALTER TABLE `exercise` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `exercise_amm`
--

DROP TABLE IF EXISTS `exercise_amm`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `exercise_amm` (
  `id` int NOT NULL AUTO_INCREMENT,
  `exercise_id` int NOT NULL,
  `amm_id` int unsigned NOT NULL,
  `role` enum('stabilizer','eccentric','concentric') DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `unique_exercise_amm_role` (`exercise_id`,`amm_id`,`role`),
  KEY `exercise_amm_ibfk_2` (`amm_id`),
  CONSTRAINT `exercise_amm_ibfk_1` FOREIGN KEY (`exercise_id`) REFERENCES `exercise` (`id`) ON DELETE CASCADE,
  CONSTRAINT `exercise_amm_ibfk_2` FOREIGN KEY (`amm_id`) REFERENCES `articulation_movement_muscle` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=738 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `exercise_amm`
--

LOCK TABLES `exercise_amm` WRITE;
/*!40000 ALTER TABLE `exercise_amm` DISABLE KEYS */;
INSERT INTO `exercise_amm` VALUES (597,84,48,'concentric'),(605,84,49,'concentric'),(598,84,52,'concentric'),(606,84,53,'concentric'),(599,84,54,'concentric'),(600,84,73,'concentric'),(607,84,76,'concentric'),(611,84,92,'stabilizer'),(613,84,196,'stabilizer'),(614,84,198,'stabilizer'),(615,84,200,'stabilizer'),(616,84,206,'stabilizer'),(601,84,220,'eccentric'),(612,84,239,'stabilizer'),(602,84,240,'eccentric'),(603,84,241,'eccentric'),(604,84,242,'eccentric'),(610,84,243,'eccentric'),(608,84,247,'eccentric'),(609,84,248,'eccentric'),(637,85,81,'stabilizer'),(636,85,88,'stabilizer'),(635,85,92,'stabilizer'),(629,85,111,'eccentric'),(632,85,121,'concentric'),(630,85,162,'concentric'),(634,85,239,'stabilizer'),(633,85,244,'eccentric'),(628,85,252,'concentric'),(631,85,253,'eccentric'),(555,86,121,'concentric'),(557,86,124,'concentric'),(556,86,244,'eccentric'),(558,86,245,'eccentric'),(559,86,246,'stabilizer'),(487,87,1,'concentric'),(494,87,29,'stabilizer'),(495,87,31,'stabilizer'),(491,87,48,'concentric'),(496,87,92,'stabilizer'),(497,87,93,'stabilizer'),(488,87,176,'eccentric'),(493,87,180,'stabilizer'),(492,87,220,'eccentric'),(489,87,237,'concentric'),(490,87,238,'eccentric'),(523,88,1,'concentric'),(538,88,29,'stabilizer'),(539,88,31,'stabilizer'),(525,88,48,'concentric'),(527,88,51,'concentric'),(529,88,54,'concentric'),(532,88,56,'eccentric'),(534,88,57,'eccentric'),(536,88,58,'eccentric'),(544,88,81,'stabilizer'),(545,88,84,'stabilizer'),(543,88,88,'stabilizer'),(546,88,89,'stabilizer'),(547,88,92,'stabilizer'),(549,88,108,'stabilizer'),(550,88,109,'stabilizer'),(551,88,110,'stabilizer'),(552,88,111,'stabilizer'),(553,88,121,'stabilizer'),(554,88,124,'stabilizer'),(524,88,176,'eccentric'),(533,88,177,'concentric'),(541,88,178,'stabilizer'),(537,88,180,'stabilizer'),(542,88,208,'stabilizer'),(540,88,210,'stabilizer'),(535,88,211,'concentric'),(526,88,220,'eccentric'),(528,88,221,'eccentric'),(531,88,222,'concentric'),(548,88,239,'stabilizer'),(530,88,241,'eccentric'),(709,89,29,'stabilizer'),(710,89,31,'stabilizer'),(711,89,41,'stabilizer'),(713,89,43,'stabilizer'),(714,89,47,'stabilizer'),(705,89,47,'concentric'),(707,89,48,'concentric'),(703,89,50,'concentric'),(715,89,73,'stabilizer'),(701,89,180,'concentric'),(708,89,220,'eccentric'),(712,89,238,'stabilizer'),(702,89,254,'eccentric'),(704,89,258,'eccentric'),(706,89,261,'eccentric'),(716,89,262,'stabilizer'),(518,90,29,'stabilizer'),(519,90,31,'stabilizer'),(505,90,48,'concentric'),(506,90,52,'concentric'),(507,90,54,'concentric'),(508,90,73,'concentric'),(513,90,76,'concentric'),(515,90,92,'stabilizer'),(521,90,178,'stabilizer'),(517,90,180,'stabilizer'),(522,90,208,'stabilizer'),(520,90,210,'stabilizer'),(509,90,220,'eccentric'),(516,90,239,'stabilizer'),(510,90,240,'eccentric'),(511,90,241,'eccentric'),(512,90,242,'eccentric'),(514,90,243,'eccentric'),(650,91,1,'concentric'),(653,91,56,'eccentric'),(659,91,84,'stabilizer'),(658,91,88,'stabilizer'),(660,91,89,'stabilizer'),(656,91,92,'stabilizer'),(651,91,176,'eccentric'),(654,91,180,'concentric'),(652,91,222,'concentric'),(657,91,239,'stabilizer'),(655,91,254,'eccentric'),(592,92,29,'stabilizer'),(593,92,31,'stabilizer'),(579,92,48,'concentric'),(580,92,52,'concentric'),(581,92,54,'concentric'),(582,92,73,'concentric'),(587,92,76,'concentric'),(589,92,92,'stabilizer'),(595,92,178,'stabilizer'),(591,92,180,'stabilizer'),(596,92,208,'stabilizer'),(594,92,210,'stabilizer'),(583,92,220,'eccentric'),(590,92,239,'stabilizer'),(584,92,240,'eccentric'),(585,92,241,'eccentric'),(586,92,242,'eccentric'),(588,92,243,'eccentric'),(626,93,84,'stabilizer'),(625,93,88,'stabilizer'),(627,93,89,'stabilizer'),(624,93,92,'stabilizer'),(617,93,100,'concentric'),(619,93,103,'concentric'),(621,93,106,'concentric'),(623,93,239,'stabilizer'),(618,93,249,'eccentric'),(620,93,250,'eccentric'),(622,93,251,'eccentric'),(561,94,48,'concentric'),(562,94,52,'concentric'),(563,94,54,'concentric'),(564,94,73,'concentric'),(569,94,76,'concentric'),(577,94,92,'stabilizer'),(571,94,196,'stabilizer'),(572,94,198,'stabilizer'),(573,94,200,'stabilizer'),(574,94,202,'stabilizer'),(575,94,204,'stabilizer'),(576,94,206,'stabilizer'),(565,94,220,'eccentric'),(578,94,239,'stabilizer'),(566,94,240,'eccentric'),(567,94,241,'eccentric'),(568,94,242,'eccentric'),(570,94,243,'eccentric'),(724,95,54,'stabilizer'),(725,95,55,'stabilizer'),(717,95,64,'concentric'),(719,95,65,'concentric'),(721,95,66,'concentric'),(723,95,73,'stabilizer'),(727,95,225,'stabilizer'),(718,95,263,'eccentric'),(720,95,264,'eccentric'),(722,95,265,'eccentric'),(726,95,266,'stabilizer'),(504,96,81,'stabilizer'),(502,96,84,'stabilizer'),(501,96,88,'stabilizer'),(503,96,89,'stabilizer'),(500,96,92,'stabilizer'),(498,96,93,'concentric'),(499,96,239,'eccentric'),(696,97,29,'stabilizer'),(697,97,31,'stabilizer'),(686,97,48,'concentric'),(688,97,50,'concentric'),(693,97,81,'stabilizer'),(692,97,88,'stabilizer'),(690,97,92,'concentric'),(683,97,111,'eccentric'),(684,97,162,'concentric'),(699,97,178,'stabilizer'),(700,97,208,'stabilizer'),(698,97,210,'stabilizer'),(687,97,220,'eccentric'),(694,97,239,'stabilizer'),(682,97,252,'concentric'),(685,97,253,'eccentric'),(695,97,254,'stabilizer'),(689,97,258,'eccentric'),(691,97,260,'eccentric'),(728,98,79,'concentric'),(730,98,249,'stabilizer'),(731,98,250,'stabilizer'),(732,98,251,'stabilizer'),(729,98,267,'eccentric'),(733,98,268,'stabilizer'),(734,98,269,'stabilizer'),(735,98,270,'stabilizer'),(736,98,271,'stabilizer'),(737,98,272,'stabilizer'),(665,183,15,'concentric'),(663,183,18,'concentric'),(661,183,48,'concentric'),(667,183,92,'stabilizer'),(669,183,212,'stabilizer'),(662,183,220,'eccentric'),(668,183,239,'stabilizer'),(664,183,255,'eccentric'),(666,183,256,'eccentric'),(670,183,257,'stabilizer'),(645,186,29,'stabilizer'),(646,186,31,'stabilizer'),(638,186,48,'eccentric'),(641,186,56,'concentric'),(643,186,58,'concentric'),(648,186,178,'stabilizer'),(644,186,180,'stabilizer'),(649,186,208,'stabilizer'),(647,186,210,'stabilizer'),(642,186,211,'eccentric'),(639,186,220,'concentric'),(640,186,222,'eccentric'),(671,189,50,'concentric'),(674,189,56,'eccentric'),(680,189,84,'stabilizer'),(679,189,88,'stabilizer'),(681,189,89,'stabilizer'),(677,189,92,'stabilizer'),(675,189,181,'concentric'),(673,189,222,'concentric'),(678,189,239,'stabilizer'),(672,189,258,'eccentric'),(676,189,259,'eccentric');
/*!40000 ALTER TABLE `exercise_amm` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `exercise_search`
--

DROP TABLE IF EXISTS `exercise_search`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `exercise_search` (
  `id` int NOT NULL AUTO_INCREMENT,
  `exercise_id` int NOT NULL,
  `full_text` text NOT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_exercise_search_exercise` (`exercise_id`),
  FULLTEXT KEY `ft_exercise_search` (`full_text`),
  CONSTRAINT `fk_exercise_search_exercise` FOREIGN KEY (`exercise_id`) REFERENCES `exercise` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=15 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `exercise_search`
--

LOCK TABLES `exercise_search` WRITE;
/*!40000 ALTER TABLE `exercise_search` DISABLE KEYS */;
INSERT INTO `exercise_search` VALUES (1,88,'Supino Reto Peitoral Costas Bíceps Tríceps Core Quadríceps Panturrilha Porção Superior (Maior) Trapézio Médio Trapézio Inferior Cabeça Longa Cabeça Curta Braquial Cabeça Medial Cabeça Lateral Multífido Oblíquos Internos Transverso do Abdome Oblíquos Externos Reto do Abdome Vasto Lateral Vasto Medial Vasto Intermédio Reto Femoral Gastrocnêmio Sóleo Romboide Maior Trapézio Superior Levantador da Escápula Romboide Menor Eretores da Espinha Ombro (Glenoumeral) Escápula e Clavícula Cotovelo Coluna Torácica e Lombar Joelho Tornozelo Flexão Adução (Retração) Extensão Flexão Plantar Elevação'),(2,89,'Remada Alta Costas Bíceps Antebraço Trapézio Médio Trapézio Inferior Supraespinhal Redondo Menor Subescapular Cabeça Longa Braquiorradial Trapézio Superior Infraespinhal Escápula e Clavícula Ombro (Glenoumeral) Cotovelo Adução (Retração) Adução (Horizontal) Rotação Lateral (Rotação Externa) Rotação Medial (Rotação Interna) Flexão Elevação Extensão Depressão'),(3,90,'Rosca Direta com Barra Costas Bíceps Antebraço Core Trapézio Médio Trapézio Inferior Cabeça Longa Cabeça Curta Braquial Braquiorradial Supinador Reto do Abdome Romboide Maior Trapézio Superior Levantador da Escápula Romboide Menor Eretores da Espinha Escápula e Clavícula Cotovelo Antebraço (Rádio-ulnar) Coluna Torácica e Lombar Adução (Retração) Flexão Supinação Elevação Extensão Pronação'),(4,91,'Desenvolvimento Militar Peitoral Tríceps Core Costas Porção Superior (Maior) Cabeça Longa Oblíquos Internos Transverso do Abdome Oblíquos Externos Reto do Abdome Trapézio Superior Eretores da Espinha Ombro (Glenoumeral) Cotovelo Coluna Torácica e Lombar Escápula e Clavícula Flexão Extensão Elevação Depressão'),(5,92,'Rosca Direta com Halteres Costas Bíceps Antebraço Core Trapézio Médio Trapézio Inferior Cabeça Longa Cabeça Curta Braquial Braquiorradial Supinador Reto do Abdome Romboide Maior Trapézio Superior Levantador da Escápula Romboide Menor Eretores da Espinha Escápula e Clavícula Cotovelo Antebraço (Rádio-ulnar) Coluna Torácica e Lombar Adução (Retração) Flexão Supinação Elevação Extensão Pronação'),(6,93,'Cadeira Flexora Core Isquiotibiais Oblíquos Internos Transverso do Abdome Oblíquos Externos Reto do Abdome Bíceps Femoral Semitendíneo Semimembranáceo Eretores da Espinha Coluna Torácica e Lombar Joelho Flexão Extensão'),(7,94,'Rosca Concentrada Bíceps Antebraço Core Costas Cabeça Longa Cabeça Curta Braquial Braquiorradial Supinador Reto do Abdome Trapézio Superior Trapézio Médio Trapézio Inferior Romboide Menor Romboide Maior Levantador da Escápula Eretores da Espinha Cotovelo Antebraço (Rádio-ulnar) Coluna Torácica e Lombar Ombro (Glenoumeral) Flexão Supinação Extensão Pronação'),(8,95,'Rosca de Punho Bíceps Tríceps Antebraço Braquial Cabeça Longa Flexor Radial do Carpo Flexor Ulnar do Carpo Palmar Longo Braquiorradial Cabeça Lateral Cotovelo Ombro (Glenoumeral) Punho e Intercarpais Flexão Extensão'),(9,96,'Extensão de Lombar Core Multífido Oblíquos Internos Transverso do Abdome Oblíquos Externos Reto do Abdome Eretores da Espinha Coluna Torácica e Lombar Extensão Flexão'),(10,97,'Remada na Máquina Costas Bíceps Core Quadríceps Quadril Trapézio Médio Trapézio Inferior Cabeça Longa Multífido Transverso do Abdome Reto do Abdome Reto Femoral Iliopsoas Romboide Maior Levantador da Escápula Romboide Menor Eretores da Espinha Trapézio Superior Escápula e Clavícula Cotovelo Ombro (Glenoumeral) Coluna Torácica e Lombar Joelho Quadril Adução (Retração) Flexão Extensão Elevação Depressão'),(11,98,'Elevação Pélvica (Hip Thrust) Glúteos Isquiotibiais Core Máximo Bíceps Femoral Semitendíneo Semimembranáceo Transverso do Abdome Multífido Oblíquos Internos Oblíquos Externos Reto do Abdome Quadril Joelho Pelve Extensão Flexão Inclinação Posterior'),(12,183,'Barra Fixa (Pull-Up) Peitoral Costas Bíceps Core Serrátil Anterior Grande Dorsal Cabeça Longa Reto do Abdome Trapézio Superior Eretores da Espinha Escapulotorácica Ombro (Glenoumeral) Cotovelo Coluna Torácica e Lombar Rotação Superior Adução Flexão Elevação Extensão Abdução Rotação Inferior'),(13,186,'Tríceps Testa Costas Bíceps Tríceps Trapézio Médio Trapézio Inferior Cabeça Longa Cabeça Lateral Romboide Maior Trapézio Superior Levantador da Escápula Romboide Menor Escápula e Clavícula Cotovelo Adução (Retração) Flexão Extensão Elevação'),(14,189,'Flexão de Braço Bíceps Tríceps Core Peitoral Cabeça Longa Oblíquos Internos Transverso do Abdome Oblíquos Externos Reto do Abdome Serrátil Anterior Eretores da Espinha Ombro (Glenoumeral) Cotovelo Coluna Torácica e Lombar Escápula e Clavícula Flexão Extensão Rotação Superior Rotação Inferior');
/*!40000 ALTER TABLE `exercise_search` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `exercise_variation`
--

DROP TABLE IF EXISTS `exercise_variation`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `exercise_variation` (
  `id` int NOT NULL AUTO_INCREMENT,
  `base_exercise_id` int NOT NULL,
  `name` varchar(255) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`),
  KEY `base_exercise_id` (`base_exercise_id`),
  CONSTRAINT `exercise_variation_ibfk_1` FOREIGN KEY (`base_exercise_id`) REFERENCES `exercise` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=91 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `exercise_variation`
--

LOCK TABLES `exercise_variation` WRITE;
/*!40000 ALTER TABLE `exercise_variation` DISABLE KEYS */;
/*!40000 ALTER TABLE `exercise_variation` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `failed_jobs`
--

DROP TABLE IF EXISTS `failed_jobs`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `failed_jobs` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `uuid` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `connection` text CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `queue` text CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `payload` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `exception` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `failed_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `failed_jobs_uuid_unique` (`uuid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `failed_jobs`
--

LOCK TABLES `failed_jobs` WRITE;
/*!40000 ALTER TABLE `failed_jobs` DISABLE KEYS */;
/*!40000 ALTER TABLE `failed_jobs` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `job_batches`
--

DROP TABLE IF EXISTS `job_batches`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `job_batches` (
  `id` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `total_jobs` int NOT NULL,
  `pending_jobs` int NOT NULL,
  `failed_jobs` int NOT NULL,
  `failed_job_ids` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `options` mediumtext CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci,
  `cancelled_at` int DEFAULT NULL,
  `created_at` int NOT NULL,
  `finished_at` int DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `job_batches`
--

LOCK TABLES `job_batches` WRITE;
/*!40000 ALTER TABLE `job_batches` DISABLE KEYS */;
/*!40000 ALTER TABLE `job_batches` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `jobs`
--

DROP TABLE IF EXISTS `jobs`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `jobs` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `queue` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `payload` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `attempts` tinyint unsigned NOT NULL,
  `reserved_at` int unsigned DEFAULT NULL,
  `available_at` int unsigned NOT NULL,
  `created_at` int unsigned NOT NULL,
  PRIMARY KEY (`id`),
  KEY `jobs_queue_index` (`queue`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `jobs`
--

LOCK TABLES `jobs` WRITE;
/*!40000 ALTER TABLE `jobs` DISABLE KEYS */;
/*!40000 ALTER TABLE `jobs` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `migrations`
--

DROP TABLE IF EXISTS `migrations`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `migrations` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `migration` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `batch` int NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `migrations`
--

LOCK TABLES `migrations` WRITE;
/*!40000 ALTER TABLE `migrations` DISABLE KEYS */;
INSERT INTO `migrations` VALUES (1,'0001_01_01_000000_create_users_table',1),(2,'0001_01_01_000001_create_cache_table',1),(3,'0001_01_01_000002_create_jobs_table',1),(4,'2025_02_26_213001_create_oauth_auth_codes_table',1),(5,'2025_02_26_213002_create_oauth_access_tokens_table',1),(6,'2025_02_26_213003_create_oauth_refresh_tokens_table',1),(7,'2025_02_26_213004_create_oauth_clients_table',1),(8,'2025_02_26_213005_create_oauth_personal_access_clients_table',1);
/*!40000 ALTER TABLE `migrations` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `movements`
--

DROP TABLE IF EXISTS `movements`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `movements` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `createdAt` datetime DEFAULT CURRENT_TIMESTAMP,
  `updatedAt` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=44 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `movements`
--

LOCK TABLES `movements` WRITE;
/*!40000 ALTER TABLE `movements` DISABLE KEYS */;
INSERT INTO `movements` VALUES (1,'Abdução (Protração)','2022-11-20 22:12:06','2022-11-20 22:12:06'),(2,'Adução (Retração)','2022-11-20 22:12:07','2022-11-20 22:12:07'),(3,'Depressão','2022-11-20 22:12:07','2022-11-20 22:12:07'),(4,'Elevação','2022-11-20 22:12:07','2022-11-20 22:12:07'),(5,'Rotação Superior','2022-11-20 22:12:07','2022-11-20 22:12:07'),(6,'Rotação Inferior','2022-11-20 22:12:07','2022-11-20 22:12:07'),(7,'Flexão','2022-11-20 22:12:07','2022-11-20 22:12:07'),(8,'Extensão','2022-11-20 22:12:07','2022-11-20 22:12:07'),(9,'Adução','2022-11-20 22:12:07','2022-11-20 22:12:07'),(10,'Abdução','2022-11-20 22:12:07','2022-11-20 22:12:07'),(11,'Adução Transversa','2022-11-20 22:12:07','2022-11-20 22:12:07'),(12,'Flexão Transversa','2022-11-20 22:12:07','2022-11-20 22:12:07'),(13,'Abdução Transversa','2022-11-20 22:12:07','2022-11-20 22:12:07'),(14,'Extensão Transversa','2022-11-20 22:12:07','2022-11-20 22:12:07'),(15,'Rotação Medial (Rotação Interna)','2022-11-20 22:12:07','2022-11-20 22:12:07'),(16,'Rotação Lateral (Rotação Externa)','2022-11-20 22:12:07','2022-11-20 22:12:07'),(17,'Pronação','2022-11-20 22:12:07','2022-11-20 22:12:07'),(18,'Supinação','2022-11-20 22:12:08','2022-11-20 22:12:08'),(19,'Extensão / Hiperextensão','2022-11-20 22:12:08','2022-11-20 22:12:08'),(20,'Adução (Desvio Ulnar)','2022-11-20 22:12:08','2022-11-20 22:12:08'),(21,'Abdução (Desvio Radial)','2022-11-20 22:12:08','2022-11-20 22:12:08'),(22,'Oposição','2022-11-20 22:12:08','2022-11-20 22:12:08'),(23,'Flexão Lateral (Abdução)','2022-11-20 22:12:08','2022-11-20 22:12:08'),(24,'Redução (Adução)','2022-11-20 22:12:08','2022-11-20 22:12:08'),(25,'Rotação','2022-11-20 22:12:08','2022-11-20 22:12:08'),(26,'Flexão Plantar','2022-11-20 22:12:08','2022-11-20 22:12:08'),(27,'Dorsiflexão','2022-11-20 22:12:08','2022-11-20 22:12:08'),(28,'Inversão (Supinação)','2022-11-20 22:12:08','2022-11-20 22:12:08'),(29,'Eversão (Pronação)','2022-11-20 22:12:08','2022-11-20 22:12:08'),(30,'Circundução','2022-11-20 22:12:08','2022-11-20 22:12:08'),(31,'Adução (Horizontal)','2022-11-20 22:12:07','2022-11-20 22:12:07'),(32,'Abdução (Horizontal)','2022-11-20 22:12:07','2022-11-20 22:12:07'),(33,'Inclinação Anterior','2022-11-20 22:12:08','2022-11-20 22:12:08'),(34,'Inclinação Posterior','2022-11-20 22:12:08','2022-11-20 22:12:08'),(35,'Inclinação Lateral','2022-11-20 22:12:08','2022-11-20 22:12:08'),(37,'Nutação','2022-11-20 22:12:08','2022-11-20 22:12:08'),(38,'Contranutação','2022-11-20 22:12:08','2022-11-20 22:12:08'),(39,'Sínfise Púbica','2022-11-20 22:12:08','2022-11-20 22:12:08'),(40,'Desvio Lateral','2022-11-20 22:12:08','2022-11-20 22:12:08'),(41,'Dobradiça','2022-11-20 22:12:08','2022-11-20 22:12:08'),(42,'Rotação Medial da Tíbia','2022-11-20 22:12:08','2022-11-20 22:12:08'),(43,'Flexão Transversa','2022-11-20 22:12:07','2022-11-20 22:12:07');
/*!40000 ALTER TABLE `movements` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `muscle_group`
--

DROP TABLE IF EXISTS `muscle_group`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `muscle_group` (
  `id` int unsigned NOT NULL,
  `name` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `image` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `muscle_group`
--

LOCK TABLES `muscle_group` WRITE;
/*!40000 ALTER TABLE `muscle_group` DISABLE KEYS */;
INSERT INTO `muscle_group` VALUES (1,'Peitoral',NULL,'2023-02-18 19:56:28',NULL),(2,'Costas',NULL,'2023-02-18 19:56:28',NULL),(3,'Bíceps',NULL,'2023-02-18 19:56:28',NULL),(4,'Tríceps',NULL,'2023-02-18 19:56:28',NULL),(5,'Antebraço',NULL,'2023-02-18 19:56:28',NULL),(6,'Glúteos',NULL,'2023-02-18 19:56:28',NULL),(7,'Core',NULL,'2023-02-18 19:56:28',NULL),(8,'Isquiotibiais',NULL,'2023-02-18 19:56:28',NULL),(9,'Quadríceps',NULL,'2023-02-18 19:56:28',NULL),(10,'Panturrilha',NULL,'2023-02-18 19:56:28',NULL),(11,'Deltoide',NULL,'2023-02-18 20:08:02',NULL),(12,'Pescoço',NULL,'2023-02-19 21:25:38',NULL),(13,'Quadril',NULL,'2023-02-19 21:25:38',NULL);
/*!40000 ALTER TABLE `muscle_group` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `muscle_portion`
--

DROP TABLE IF EXISTS `muscle_portion`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `muscle_portion` (
  `id` int unsigned NOT NULL,
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `image` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `muscle_group_id` int unsigned NOT NULL,
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `muscle_portion_muscleGRoup` (`muscle_group_id`),
  CONSTRAINT `muscles_group_to_muscle_portion_foreign_key` FOREIGN KEY (`muscle_group_id`) REFERENCES `muscle_group` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `muscle_portion`
--

LOCK TABLES `muscle_portion` WRITE;
/*!40000 ALTER TABLE `muscle_portion` DISABLE KEYS */;
INSERT INTO `muscle_portion` VALUES (1,'Porção Superior (Maior)',NULL,1,'2023-02-18 20:13:01','2023-02-18 20:42:50'),(2,'Porção Medial (Maior)',NULL,1,'2023-02-18 20:13:01','2023-02-18 20:42:53'),(3,'Porção Inferior (Maior)',NULL,1,'2023-02-18 20:13:01','2023-02-18 20:42:55'),(4,'Menor',NULL,1,'2023-02-18 20:13:01','2023-02-18 20:24:08'),(5,'Serrátil Anterior',NULL,1,'2023-02-18 20:13:01','2023-02-18 20:24:14'),(6,'Subclávio',NULL,1,'2023-02-18 20:13:01',NULL),(7,'Grande Dorsal',NULL,2,'2023-02-18 20:17:23','2023-03-19 21:36:32'),(8,'Serrátil Posterior',NULL,2,'2023-02-18 20:17:23','2023-02-18 20:24:17'),(9,'Lombar',NULL,2,'2023-02-18 20:17:23',NULL),(10,'Trapézio Superior',NULL,2,'2023-02-18 20:21:20',NULL),(11,'Trapézio Médio',NULL,2,'2023-02-18 20:21:20',NULL),(12,'Trapézio Inferior',NULL,2,'2023-02-18 20:21:20',NULL),(13,'Romboide Menor',NULL,2,'2023-02-18 20:21:20',NULL),(14,'Romboide Maior',NULL,2,'2023-02-18 20:21:20',NULL),(15,'Levantador da Escápula',NULL,2,'2023-02-18 20:21:20',NULL),(16,'Cabeça Longa',NULL,3,'2023-02-18 20:27:11','2023-03-19 08:19:43'),(17,'Cabeça Curta',NULL,3,'2023-02-18 20:27:11','2023-03-19 08:19:24'),(18,'Braquial',NULL,3,'2023-02-18 20:27:11','2023-03-19 08:19:24'),(19,'Cabeça Longa',NULL,4,'2023-02-18 20:29:04','2023-03-19 08:20:17'),(20,'Cabeça Medial',NULL,4,'2023-02-18 20:29:04','2023-03-19 08:20:17'),(21,'Cabeça Lateral',NULL,4,'2023-02-18 20:29:04','2023-03-19 08:20:17'),(22,'Pronador Redondo',NULL,5,'2023-02-18 20:31:13','2023-03-19 08:21:06'),(23,'Flexor Radial do Carpo',NULL,5,'2023-02-18 20:31:13','2023-03-19 08:21:06'),(24,'Flexor Ulnar do Carpo',NULL,5,'2023-02-18 20:31:13','2023-03-19 08:21:06'),(25,'Palmar Longo',NULL,5,'2023-02-18 20:31:13','2023-03-19 08:21:06'),(26,'Médio',NULL,6,'2023-02-18 20:32:35',NULL),(27,'Mínimo',NULL,6,'2023-02-18 20:32:35',NULL),(28,'Máximo',NULL,6,'2023-02-18 20:32:35',NULL),(29,'Assoalho Pélvico',NULL,7,'2023-02-18 20:35:37',NULL),(30,'Transverso do Abdome',NULL,7,'2023-02-18 20:35:37','2023-02-18 20:36:12'),(31,'Multífido',NULL,7,'2023-02-18 20:35:37',NULL),(32,'Oblíquos Internos',NULL,7,'2023-02-18 20:35:37',NULL),(33,'Oblíquos Externos',NULL,7,'2023-02-18 20:35:37',NULL),(34,'Reto do Abdome',NULL,7,'2023-02-18 20:35:37',NULL),(35,'Eretores da Espinha',NULL,7,'2023-02-18 20:35:37',NULL),(36,'Bíceps Femoral',NULL,8,'2023-02-18 20:37:28',NULL),(37,'Semitendíneo',NULL,8,'2023-02-18 20:37:28',NULL),(38,'Semimembranáceo',NULL,8,'2023-02-18 20:37:28',NULL),(39,'Vasto Lateral',NULL,9,'2023-02-18 20:39:39',NULL),(40,'Vasto Medial',NULL,9,'2023-02-18 20:39:39',NULL),(41,'Vasto Intermédio',NULL,9,'2023-02-18 20:39:39',NULL),(42,'Reto Femoral',NULL,9,'2023-02-18 20:39:39',NULL),(43,'Gastrocnêmio',NULL,10,'2023-02-18 20:40:41',NULL),(44,'Sóleo',NULL,10,'2023-02-18 20:40:41',NULL),(45,'Anterior',NULL,11,'2023-02-18 20:41:36','2023-03-19 08:21:37'),(46,'Médio',NULL,11,'2023-02-18 20:41:36','2023-03-19 08:21:37'),(47,'Lateral',NULL,11,'2023-02-18 20:41:36','2023-03-19 08:21:37'),(48,'Esternocleidomastoideo',NULL,12,'2023-02-19 21:27:00',NULL),(49,'Esplênio',NULL,12,'2023-02-20 13:56:18',NULL),(50,'Tensor da Fáscia Lata',NULL,9,'2023-02-20 14:03:41',NULL),(51,'Sartório',NULL,9,'2023-02-20 14:03:45',NULL),(52,'Supraespinhal',NULL,2,'2023-02-20 14:03:48',NULL),(53,'Extensor Radial Longo do Carpo',NULL,5,'2023-02-20 14:03:51',NULL),(54,'Extensor Radial Curto do Carpo',NULL,5,'2023-02-20 14:03:55',NULL),(55,'Extensor Ulnar do Carpo',NULL,5,'2023-02-20 14:03:59',NULL),(56,'Adutor Longo',NULL,13,'2023-02-20 14:07:20','2023-02-20 14:11:08'),(57,'Adutor Curto',NULL,13,'2023-02-20 14:07:20','2023-02-20 14:11:08'),(58,'Adutor Magno',NULL,13,'2023-02-20 14:07:20','2023-02-20 14:11:08'),(59,'Grácil',NULL,13,'2023-02-20 14:07:20','2023-02-20 14:11:08'),(60,'Infraespinhal',NULL,2,'2023-02-20 14:09:07',NULL),(61,'Redondo Menor',NULL,2,'2023-02-20 14:09:41',NULL),(62,'Redondo Maior',NULL,2,'2023-02-20 14:13:41',NULL),(63,'Pectíneo',NULL,7,'2023-02-20 14:13:54','2023-03-19 21:06:16'),(64,'Piriforme',NULL,13,'2023-02-20 14:17:56',NULL),(65,'Quadrado Femoral',NULL,13,'2023-02-20 14:17:56',NULL),(66,'Obturador Externo',NULL,13,'2023-02-20 14:17:56',NULL),(67,'Obturador Interno',NULL,13,'2023-02-20 14:17:56',NULL),(68,'Gêmeo Superior',NULL,13,'2023-02-20 14:17:56',NULL),(69,'Tibial Anterior',NULL,10,'2023-02-20 14:19:58',NULL),(70,'Subescapular',NULL,2,'2023-02-20 14:21:20',NULL),(71,'Iliopsoas',NULL,13,'2023-02-20 14:21:20',NULL),(72,'Poplíteo',NULL,9,'2023-02-20 14:21:20',NULL),(73,'Posterior',NULL,11,'2023-03-19 20:37:48',NULL),(74,'Braquiorradial',NULL,5,'2023-03-19 20:43:12','2023-03-19 20:55:14'),(75,'Supinador',NULL,5,'2023-03-19 21:35:05',NULL),(76,'Adutor Magno',NULL,13,'2023-03-19 21:48:44',NULL),(77,'Grande Dorsal, fibras inferiores',NULL,2,'2025-02-27 17:49:03',NULL),(78,'Grande Dorsal, fibras superiores',NULL,2,'2025-02-27 17:49:21',NULL);
/*!40000 ALTER TABLE `muscle_portion` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `muscle_portion_articulation`
--

DROP TABLE IF EXISTS `muscle_portion_articulation`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `muscle_portion_articulation` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `muscle_portion_id` int unsigned NOT NULL,
  `articulation_id` int unsigned NOT NULL,
  PRIMARY KEY (`id`),
  KEY `muscle_portion_to_muscle_portion_articulation` (`muscle_portion_id`),
  KEY `articulation_to_muscle_portion_articulation` (`articulation_id`),
  CONSTRAINT `articulation_to_muscle_portion_articulation` FOREIGN KEY (`articulation_id`) REFERENCES `articulations` (`id`),
  CONSTRAINT `muscle_portion_to_muscle_portion_articulation` FOREIGN KEY (`muscle_portion_id`) REFERENCES `muscle_portion` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `muscle_portion_articulation`
--

LOCK TABLES `muscle_portion_articulation` WRITE;
/*!40000 ALTER TABLE `muscle_portion_articulation` DISABLE KEYS */;
/*!40000 ALTER TABLE `muscle_portion_articulation` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `oauth_access_tokens`
--

DROP TABLE IF EXISTS `oauth_access_tokens`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `oauth_access_tokens` (
  `id` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `user_id` bigint unsigned DEFAULT NULL,
  `client_id` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `scopes` text CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci,
  `revoked` tinyint(1) NOT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `expires_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `oauth_access_tokens_user_id_index` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `oauth_access_tokens`
--

LOCK TABLES `oauth_access_tokens` WRITE;
/*!40000 ALTER TABLE `oauth_access_tokens` DISABLE KEYS */;
INSERT INTO `oauth_access_tokens` VALUES ('1def2ff52ad7b8858029694bb2ef20cb0fd48c1aa6841ea8460d463c30d846058fcbfd005db3b945',3,'9e4eb6d3-68fe-40c2-8887-18b80305fa0d','carlosjr.ribas@mail.com','[]',0,'2025-02-26 22:37:39','2025-02-26 22:37:39','2026-02-26 22:37:39'),('45ed9915517743eb94afaf3248bf2fd658d97439be0e630b725b8f25e03f1f1b8e6608215e0e4122',2,'9e4eb6d3-68fe-40c2-8887-18b80305fa0d','carlosjr.ribas@teste.com','[]',0,'2025-02-26 22:32:05','2025-02-26 22:32:05','2026-02-26 22:32:05'),('f9839fc0d0287b4093b1a88f6e65d81720734d0370d4371b284e99649bb5c8c2f08bb3cce425b57e',3,'9e4eb6d3-68fe-40c2-8887-18b80305fa0d','carlosjr.ribas@mail.com','[]',0,'2025-02-26 22:37:54','2025-02-26 22:37:54','2026-02-26 22:37:54');
/*!40000 ALTER TABLE `oauth_access_tokens` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `oauth_auth_codes`
--

DROP TABLE IF EXISTS `oauth_auth_codes`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `oauth_auth_codes` (
  `id` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `user_id` bigint unsigned NOT NULL,
  `client_id` bigint unsigned NOT NULL,
  `scopes` text CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci,
  `revoked` tinyint(1) NOT NULL,
  `expires_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `oauth_auth_codes_user_id_index` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `oauth_auth_codes`
--

LOCK TABLES `oauth_auth_codes` WRITE;
/*!40000 ALTER TABLE `oauth_auth_codes` DISABLE KEYS */;
/*!40000 ALTER TABLE `oauth_auth_codes` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `oauth_clients`
--

DROP TABLE IF EXISTS `oauth_clients`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `oauth_clients` (
  `id` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `user_id` bigint unsigned DEFAULT NULL,
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `secret` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `provider` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `redirect` text CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `personal_access_client` tinyint(1) NOT NULL,
  `password_client` tinyint(1) NOT NULL,
  `revoked` tinyint(1) NOT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `oauth_clients_user_id_index` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `oauth_clients`
--

LOCK TABLES `oauth_clients` WRITE;
/*!40000 ALTER TABLE `oauth_clients` DISABLE KEYS */;
INSERT INTO `oauth_clients` VALUES ('9e4eb6d3-68fe-40c2-8887-18b80305fa0d',NULL,'carlos','loqrLbint26FE7VgGdcup93rvTeqLrzfxitFyWYe',NULL,'http://localhost',1,0,0,'2025-02-26 22:31:59','2025-02-26 22:31:59');
/*!40000 ALTER TABLE `oauth_clients` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `oauth_personal_access_clients`
--

DROP TABLE IF EXISTS `oauth_personal_access_clients`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `oauth_personal_access_clients` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `client_id` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `oauth_personal_access_clients`
--

LOCK TABLES `oauth_personal_access_clients` WRITE;
/*!40000 ALTER TABLE `oauth_personal_access_clients` DISABLE KEYS */;
INSERT INTO `oauth_personal_access_clients` VALUES (1,'9e4eb6d3-68fe-40c2-8887-18b80305fa0d','2025-02-26 22:31:59','2025-02-26 22:31:59');
/*!40000 ALTER TABLE `oauth_personal_access_clients` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `oauth_refresh_tokens`
--

DROP TABLE IF EXISTS `oauth_refresh_tokens`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `oauth_refresh_tokens` (
  `id` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `access_token_id` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `revoked` tinyint(1) NOT NULL,
  `expires_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `oauth_refresh_tokens_access_token_id_index` (`access_token_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `oauth_refresh_tokens`
--

LOCK TABLES `oauth_refresh_tokens` WRITE;
/*!40000 ALTER TABLE `oauth_refresh_tokens` DISABLE KEYS */;
/*!40000 ALTER TABLE `oauth_refresh_tokens` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `password_reset_tokens`
--

DROP TABLE IF EXISTS `password_reset_tokens`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `password_reset_tokens` (
  `email` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `token` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`email`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `password_reset_tokens`
--

LOCK TABLES `password_reset_tokens` WRITE;
/*!40000 ALTER TABLE `password_reset_tokens` DISABLE KEYS */;
/*!40000 ALTER TABLE `password_reset_tokens` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sessions`
--

DROP TABLE IF EXISTS `sessions`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `sessions` (
  `id` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `user_id` bigint unsigned DEFAULT NULL,
  `ip_address` varchar(45) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `user_agent` text CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci,
  `payload` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `last_activity` int NOT NULL,
  PRIMARY KEY (`id`),
  KEY `sessions_user_id_index` (`user_id`),
  KEY `sessions_last_activity_index` (`last_activity`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sessions`
--

LOCK TABLES `sessions` WRITE;
/*!40000 ALTER TABLE `sessions` DISABLE KEYS */;
/*!40000 ALTER TABLE `sessions` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user`
--

DROP TABLE IF EXISTS `user`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `user` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  `email` varchar(255) NOT NULL,
  `password` varchar(255) NOT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `id` (`id`),
  UNIQUE KEY `email` (`email`)
) ENGINE=InnoDB AUTO_INCREMENT=76 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user`
--

LOCK TABLES `user` WRITE;
/*!40000 ALTER TABLE `user` DISABLE KEYS */;
INSERT INTO `user` VALUES (6,'teste','teste@mail.com','$2a$10$BJKyeuHmhtiM4vbWMcpcbupkB1Gq/2W4wtCEEX5fEykke26swtZEu','2025-03-05 00:30:05'),(19,'carlos alberto ribas junior','carlosjr.ribas@gmail.com','$2a$10$x1YDfPkBwscVT8cEJ9eNVeCtBG.sdLPFSH8eTd/9rO96ClQtZJaN.','2025-03-07 00:39:02'),(35,'teste','carlosjr2.ribas@gmail.com','$2a$10$crIJSDax6Dtf5dgNbXtrruoGnMBBSTipu0FnAbwmWSa/QhqzOPRYe','2025-03-09 12:20:44'),(36,'teste','carlosjr3.ribas@gmail.com','$2a$10$Yr2a4Y0TPJog5xo.xThSeeIHkMHIjHL4dMHZ.u.FsqInzfkXdwG7q','2025-03-09 12:21:27'),(37,'Natasha','natashabarbozamocerino@hotmail.com','$2a$10$N.2rxalkPVLyGHMhtREA5uy2eWDY6kXBqIlQYbsC9XosMktTq6e/i','2025-03-09 12:42:49');
/*!40000 ALTER TABLE `user` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user_api_tokens`
--

DROP TABLE IF EXISTS `user_api_tokens`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `user_api_tokens` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `user_id` bigint NOT NULL,
  `token_name` varchar(255) NOT NULL,
  `token_hash` varchar(255) NOT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `unique_user_token_name` (`user_id`,`token_name`),
  KEY `idx_user_id` (`user_id`),
  KEY `idx_token_hash` (`token_hash`),
  CONSTRAINT `user_api_tokens_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `api_users` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=54 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user_api_tokens`
--

LOCK TABLES `user_api_tokens` WRITE;
/*!40000 ALTER TABLE `user_api_tokens` DISABLE KEYS */;
INSERT INTO `user_api_tokens` VALUES (28,1,'Teste','5+mJsC+d+UOKAHTo87y7h/X7c/fJe6GMFkoy9E5zCPI=','2025-04-10 02:43:13','2025-04-10 02:43:13',NULL),(29,1,'Nathy','KDEpt0p4LzRP9YEW0xF5aSblct3Nk3B7nRlmAH6+qbU=','2025-04-10 03:09:45','2025-04-10 03:09:45',NULL),(30,1,'Aplicativo de trienos','+afDZBQU79Wai8g1Z2MirvzXYR58M+l8F8sCQn6JjKA=','2025-04-10 16:27:52','2025-04-10 16:27:52',NULL),(36,1,'Dummy','YBuWcwEgOh+fUn/ZE+UuwYIK3BiCzlXFs1pLLZwGxKU=','2026-11-08 20:55:31','2026-11-08 20:55:31',NULL),(48,1,'Dummy2026-11-08T18:07:34-03:00','fq6x+6/4ZxP0SJ1i87dyGSgFrFXLEKg/I9nYLn8jdHs=','2026-11-08 21:07:34','2026-11-08 21:07:34',NULL),(49,1,'Dummy2026-11-08T18:10:18-03:00','e4OtU892TcsYedcCGh0GCiZSqtKobynf7JSWUcNzcLk=','2026-11-08 21:10:18','2026-11-08 21:10:18',NULL),(50,1,'Dummy2026-11-08T18:11:52-03:00','bS7SBWrdldPgWpvE68KLeg71m3CtI+WYxvEpSbGHNDI=','2026-11-08 21:11:52','2026-11-08 21:11:52',NULL),(51,1,'Dummy2026-11-08T18:12:18-03:00','jHyv1ciY8aTplsCf5hmwnnAfxi0fsERE3NzYRdMif6I=','2026-11-08 21:12:18','2026-11-08 21:12:18',NULL),(52,1,'Dummy2026-11-08T18:12:21-03:00','Akq6hjyEQlsqDHdfzVIpogTP5TqHHx/h5zdD4aixKGU=','2026-11-08 21:12:21','2026-11-08 21:12:21',NULL),(53,1,'Dummy2026-11-08T18:15:04-03:00','kl5GzRO/6q81qdwBy6GNMp/bv6Bnjmq3bBhISgFjC+Y=','2026-11-08 21:15:04','2026-11-08 21:15:04',NULL);
/*!40000 ALTER TABLE `user_api_tokens` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `users`
--

DROP TABLE IF EXISTS `users`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `users` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `email` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `email_verified_at` timestamp NULL DEFAULT NULL,
  `password` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `remember_token` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `users_email_unique` (`email`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `users`
--

LOCK TABLES `users` WRITE;
/*!40000 ALTER TABLE `users` DISABLE KEYS */;
INSERT INTO `users` VALUES (1,'Carlos Alberto Ribas Junior','carlosjr.ribas@gmail.com','2025-02-26 22:31:47','$2y$12$vclvUdGrlRbkqJpPyOxcDevNxAs2wpKxmTVca7g776IgUobKqtVte','QO3aT1PVza','2025-02-26 22:31:47','2025-02-26 22:31:47'),(2,'Carlos Alberto Ribas Junior','carlosjr.ribas@teste.com','2025-02-26 22:32:04','$2y$12$7oavP1rKOUo5pnUDDWQIXO.FfWfNEtxDGIP.4UayEk1cXEXQH1noC','qLShzc4z1k','2025-02-26 22:32:05','2025-02-26 22:32:05'),(3,'Carlos Alberto Ribas Junior','carlosjr.ribas@mail.com','2025-02-26 22:37:39','$2y$12$oFnR4XkSIlsbDHnpMa8XGOtqmV8GTeIwYPEuM./ob7FwQ.NuyxdDW','chdLm8FS6B','2025-02-26 22:37:39','2025-02-26 22:37:39');
/*!40000 ALTER TABLE `users` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Final view structure for view `auxiliar`
--

/*!50001 DROP VIEW IF EXISTS `auxiliar`*/;
/*!50001 SET @saved_cs_client          = @@character_set_client */;
/*!50001 SET @saved_cs_results         = @@character_set_results */;
/*!50001 SET @saved_col_connection     = @@collation_connection */;
/*!50001 SET character_set_client      = utf8mb4 */;
/*!50001 SET character_set_results     = utf8mb4 */;
/*!50001 SET collation_connection      = utf8mb4_0900_ai_ci */;
/*!50001 CREATE ALGORITHM=UNDEFINED */
/*!50013 DEFINER=`root`@`%` SQL SECURITY DEFINER */
/*!50001 VIEW `auxiliar` AS select `e`.`id` AS `id_exercicio`,`e`.`name` AS `nome_exercicio`,`mg`.`name` AS `grupo_muscular`,`mp`.`name` AS `porcao_muscular`,`a`.`name` AS `articulacao`,`m`.`name` AS `movimento` from ((((((`exercise` `e` join `exercise_amm` `ea` on((`ea`.`exercise_id` = `e`.`id`))) join `articulation_movement_muscle` `amm` on((`amm`.`id` = `ea`.`amm_id`))) join `muscle_portion` `mp` on((`mp`.`id` = `amm`.`muscle_portion_id`))) join `muscle_group` `mg` on((`mg`.`id` = `mp`.`muscle_group_id`))) join `movements` `m` on((`m`.`id` = `amm`.`movement_id`))) join `articulations` `a` on((`a`.`id` = `amm`.`articulation_id`))) order by `e`.`id` desc */;
/*!50001 SET character_set_client      = @saved_cs_client */;
/*!50001 SET character_set_results     = @saved_cs_results */;
/*!50001 SET collation_connection      = @saved_col_connection */;
SET @@SESSION.SQL_LOG_BIN = @MYSQLDUMP_TEMP_LOG_BIN;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2026-11-09  0:17:01
