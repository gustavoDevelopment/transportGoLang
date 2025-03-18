-- MySQL dump 10.13  Distrib 8.0.19, for Win64 (x86_64)
--
-- Host: localhost    Database: transport
-- ------------------------------------------------------
-- Server version	9.2.0

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

--
-- Table structure for table `cities`
--

DROP TABLE IF EXISTS `cities`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `cities` (
  `id` varchar(191) NOT NULL,
  `code` longtext,
  `name` longtext,
  `department_id` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_cities_departments` (`department_id`),
  CONSTRAINT `fk_cities_departments` FOREIGN KEY (`department_id`) REFERENCES `departments` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `cities`
--

LOCK TABLES `cities` WRITE;
/*!40000 ALTER TABLE `cities` DISABLE KEYS */;
INSERT INTO `cities` VALUES ('1','BGA','Bucaramanga','1'),('10','SM','Santa Marta','10'),('11','IQUI','Ibagué','4'),('2','MED','Medellín','1'),('3','CTG','Cartagena','3'),('4','BAR','Barranquilla','2'),('5','BOG','Bogotá','5'),('6','POP','Popayán','6'),('7','CALI','Cali','7'),('8','PST','Pasto','8'),('9','PERE','Pereira','9');
/*!40000 ALTER TABLE `cities` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `departments`
--

DROP TABLE IF EXISTS `departments`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `departments` (
  `id` varchar(191) NOT NULL,
  `code` longtext,
  `name` longtext,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `departments`
--

LOCK TABLES `departments` WRITE;
/*!40000 ALTER TABLE `departments` DISABLE KEYS */;
INSERT INTO `departments` VALUES ('1','ANT','Antioquia'),('10','MAG','Magdalena'),('2','ATL','Atlántico'),('3','BOL','Bolívar'),('4','BOY','Boyacá'),('5','CUN','Cundinamarca'),('6','HUI','Huila'),('7','VAL','Valle del Cauca'),('8','NAR','Nariño'),('9','RIS','Risaralda');
/*!40000 ALTER TABLE `departments` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `document_types`
--

DROP TABLE IF EXISTS `document_types`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `document_types` (
  `id` varchar(191) NOT NULL,
  `code` longtext,
  `name` longtext,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `document_types`
--

LOCK TABLES `document_types` WRITE;
/*!40000 ALTER TABLE `document_types` DISABLE KEYS */;
INSERT INTO `document_types` VALUES ('1','CC','Cédula de Ciudadanía'),('10','PEP','Permiso Especial de Permanencia'),('2','CE','Cédula de Extranjería'),('3','TI','Tarjeta de Identidad'),('4','RC','Registro Civil'),('5','PA','Pasaporte'),('6','NIT','Número de Identificación Tributaria'),('7','RUT','Registro Único Tributario'),('8','DNI','Documento Nacional de Identidad'),('9','PPT','Permiso por Protección Temporal');
/*!40000 ALTER TABLE `document_types` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `eps`
--

DROP TABLE IF EXISTS `eps`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `eps` (
  `id` varchar(191) NOT NULL,
  `code` longtext,
  `name` longtext,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `eps`
--

LOCK TABLES `eps` WRITE;
/*!40000 ALTER TABLE `eps` DISABLE KEYS */;
INSERT INTO `eps` VALUES ('1','EPS001','Salud Total'),('10','EPS010','Medimás'),('11','EPS011','SOS EPS'),('12','EPS012','Colsubsidio EPS'),('13','EPS013','Aliansalud'),('14','EPS014','Capital Salud'),('15','EPS015','Mutual Ser'),('16','EPS016','Ambuq'),('17','EPS017','Convida'),('18','EPS018','EPM Salud'),('19','EPS019','Emssanar'),('2','EPS002','Sanitas'),('20','EPS020','Comparta'),('21','EPS021','Saludvida'),('22','EPS022','Ecoopsos'),('23','EPS023','Asmet Salud'),('24','EPS024','Cajacopi EPS'),('25','EPS025','Dusakawi EPS'),('3','EPS003','Compensar'),('4','EPS004','Nueva EPS'),('5','EPS005','Sura'),('6','EPS006','Famisanar'),('7','EPS007','Coomeva'),('8','EPS008','Saludcoop'),('9','EPS009','Cafesalud');
/*!40000 ALTER TABLE `eps` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `insurance_companies`
--

DROP TABLE IF EXISTS `insurance_companies`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `insurance_companies` (
  `id` varchar(191) NOT NULL,
  `code` longtext,
  `name` longtext,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `insurance_companies`
--

LOCK TABLES `insurance_companies` WRITE;
/*!40000 ALTER TABLE `insurance_companies` DISABLE KEYS */;
INSERT INTO `insurance_companies` VALUES ('CO124','SEG002','Seguros Bolívar'),('CO125','SEG003','Sura Seguros'),('CO126','SEG004','Allianz Colombia'),('CO127','SEG005','Mapfre Colombia'),('CO128','SEG006','Liberty Seguros');
/*!40000 ALTER TABLE `insurance_companies` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `people`
--

DROP TABLE IF EXISTS `people`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `people` (
  `document_type` longtext,
  `document_number` longtext,
  `first_name` longtext,
  `second_name` longtext,
  `first_last_name` longtext,
  `second_last_name` longtext,
  `date_of_birth` longtext,
  `gender` longtext,
  `nationality` longtext,
  `phone_number` longtext,
  `email` longtext,
  `address` longtext
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `people`
--

LOCK TABLES `people` WRITE;
/*!40000 ALTER TABLE `people` DISABLE KEYS */;
/*!40000 ALTER TABLE `people` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `truck_brands`
--

DROP TABLE IF EXISTS `truck_brands`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `truck_brands` (
  `id` varchar(191) NOT NULL,
  `code` longtext,
  `name` longtext,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `truck_brands`
--

LOCK TABLES `truck_brands` WRITE;
/*!40000 ALTER TABLE `truck_brands` DISABLE KEYS */;
INSERT INTO `truck_brands` VALUES ('1','FREIGHT','Freightliner'),('10','DAF','DAF Trucks'),('2','KENW','Kenworth'),('3','MACK','Mack Trucks'),('4','INTL','International'),('5','HINO','Hino'),('6','ISUZU','Isuzu'),('7','MERC','Mercedes-Benz'),('8','SCANIA','Scania'),('9','VOLVO','Volvo Trucks');
/*!40000 ALTER TABLE `truck_brands` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `trucks`
--

DROP TABLE IF EXISTS `trucks`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `trucks` (
  `license_plate` longtext,
  `brand` longtext,
  `model` longtext,
  `year` bigint DEFAULT NULL,
  `color` longtext,
  `owner_name` longtext,
  `capacity` bigint DEFAULT NULL,
  `chassis_number` longtext,
  `engine_number` longtext
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `trucks`
--

LOCK TABLES `trucks` WRITE;
/*!40000 ALTER TABLE `trucks` DISABLE KEYS */;
/*!40000 ALTER TABLE `trucks` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Dumping routines for database 'transport'
--
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2025-03-18 12:05:01
