-- phpMyAdmin SQL Dump
-- version 5.2.1
-- https://www.phpmyadmin.net/
--
-- Host: db:3306
-- Generation Time: May 16, 2023 at 12:45 AM
-- Server version: 8.0.31
-- PHP Version: 8.1.17

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `ecommerce_db`
--

-- --------------------------------------------------------

--
-- Table structure for table `Address`
--

CREATE TABLE `Address` (
  `id` bigint UNSIGNED NOT NULL,
  `user_id` bigint UNSIGNED NOT NULL,
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `gender` tinyint(1) NOT NULL DEFAULT '0',
  `phone` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `province` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `district` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `ward` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `province_id` int DEFAULT NULL,
  `district_id` int DEFAULT NULL,
  `ward_code` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `street` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `deleted_at` datetime DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--
-- Dumping data for table `Address`
--

INSERT INTO `Address` (`id`, `user_id`, `name`, `gender`, `phone`, `province`, `district`, `ward`, `province_id`, `district_id`, `ward_code`, `street`, `created_at`, `updated_at`, `deleted_at`) VALUES
(10, 1, 'Nguyễn Văn Đạt', 0, '0987654321', 'Hồ Chí Minh', 'Thành Phố Thủ Đức', 'Phường Bình Trưng Tây', 202, 3695, '90767', '09 Nguyễn Chí Thanh', '2023-05-06 14:37:47', '2023-05-06 14:37:47', NULL),
(11, 1, 'Nguyễn Văn Đạt', 0, '0987654321', 'Hồ Chí Minh', 'Thành Phố Thủ Đức', 'Phường Bình Trưng Tây', 202, 3695, '90767', '09 Nguyễn Chí Thanh', '2023-05-10 13:15:40', '2023-05-10 13:15:40', NULL);

-- --------------------------------------------------------

--
-- Table structure for table `AdministrativeRegions`
--

CREATE TABLE `AdministrativeRegions` (
  `id` int NOT NULL,
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `name_en` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `code_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `code_name_en` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--
-- Dumping data for table `AdministrativeRegions`
--

INSERT INTO `AdministrativeRegions` (`id`, `name`, `name_en`, `code_name`, `code_name_en`, `deleted_at`) VALUES
(1, 'Đông Bắc Bộ', 'Northeast', 'dong_bac_bo', 'northest', NULL),
(2, 'Tây Bắc Bộ', 'Northwest', 'tay_bac_bo', 'northwest', NULL),
(3, 'Đồng bằng sông Hồng', 'Red River Delta', 'dong_bang_song_hong', 'red_river_delta', NULL),
(4, 'Bắc Trung Bộ', 'North Central Coast', 'bac_trung_bo', 'north_central_coast', NULL),
(5, 'Duyên hải Nam Trung Bộ', 'South Central Coast', 'duyen_hai_nam_trung_bo', 'south_central_coast', NULL),
(6, 'Tây Nguyên', 'Central Highlands', 'tay_nguyen', 'central_highlands', NULL),
(7, 'Đông Nam Bộ', 'Southeast', 'dong_nam_bo', 'southeast', NULL),
(8, 'Đồng bằng sông Cửu Long', 'Mekong River Delta', 'dong_bang_song_cuu_long', 'southwest', NULL);

-- --------------------------------------------------------

--
-- Table structure for table `AdministrativeUnits`
--

CREATE TABLE `AdministrativeUnits` (
  `id` int NOT NULL,
  `full_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `full_name_en` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `short_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `short_name_en` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `code_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `code_name_en` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--
-- Dumping data for table `AdministrativeUnits`
--

INSERT INTO `AdministrativeUnits` (`id`, `full_name`, `full_name_en`, `short_name`, `short_name_en`, `code_name`, `code_name_en`, `deleted_at`) VALUES
(1, 'Thành phố trực thuộc trung ương', 'Municipality', 'Thành phố', 'City', 'thanh_pho_truc_thuoc_trung_uong', 'municipality', NULL),
(2, 'Tỉnh', 'Province', 'Tỉnh', 'Province', 'tinh', 'province', NULL),
(3, 'Thành phố thuộc thành phố trực thuộc trung ương', 'Municipal city', 'Thành phố', 'City', 'thanh_pho_thuoc_thanh_pho_truc_thuoc_trung_uong', 'municipal_city', NULL),
(4, 'Thành phố thuộc tỉnh', 'Provincial city', 'Thành phố', 'City', 'thanh_pho_thuoc_tinh', 'provincial_city', NULL),
(5, 'Quận', 'Urban district', 'Quận', 'District', 'quan', 'urban_district', NULL),
(6, 'Thị xã', 'District-level town', 'Thị xã', 'Town', 'thi_xa', 'district_level_town', NULL),
(7, 'Huyện', 'District', 'Huyện', 'District', 'huyen', 'district', NULL),
(8, 'Phường', 'Ward', 'Phường', 'Ward', 'phuong', 'ward', NULL),
(9, 'Thị trấn', 'Commune-level town', 'Thị trấn', 'Township', 'thi_tran', 'commune_level_town', NULL),
(10, 'Xã', 'Commune', 'Xã', 'Commune', 'xa', 'commune', NULL);

-- --------------------------------------------------------

--
-- Table structure for table `Banner`
--

CREATE TABLE `Banner` (
  `id` bigint UNSIGNED NOT NULL,
  `user_id` bigint UNSIGNED NOT NULL,
  `title` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `collection` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `discount` int NOT NULL,
  `image` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `end_time` date NOT NULL,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `deleted_at` datetime DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--
-- Dumping data for table `Banner`
--

INSERT INTO `Banner` (`id`, `user_id`, `title`, `collection`, `discount`, `image`, `end_time`, `created_at`, `updated_at`, `deleted_at`) VALUES
(1, 1, 'New Hat TODAY', 'Highly Recommend', 10, 'https://www.promotionalcaps.com.au/media/codazon/slideshow/resize/1900x730/banner12.jpg', '2023-01-19', '2022-12-30 09:56:29', '2022-12-30 09:56:29', '2023-05-06 05:03:14'),
(2, 1, 'We suggest clothes to boys', 'Highly Recommend', 10, 'https://img.freepik.com/premium-vector/men-fashion-collection-social-media-banner-template-design_596383-109.jpg?w=2000', '2023-01-19', '2022-12-30 09:57:43', '2022-12-30 09:57:43', NULL),
(3, 1, 'We suggest clothes to boys', 'Highly Recommend', 10, 'https://img.freepik.com/free-vector/fashion-sale-banners-with-photo_52683-9828.jpg?w=2000', '2023-01-19', '2022-12-30 09:59:36', '2022-12-30 09:59:36', NULL),
(4, 1, 'Very Super Gud', 'Highly Recommend', 10, 'https://www.w3schools.com/howto/img_forest.jpg', '2022-11-09', '2023-05-06 07:32:12', '2023-05-06 07:32:12', NULL),
(5, 1, 'Very Super Gud', 'Highly Recommend', 10, 'https://www.w3schools.com/howto/img_forest.jpg', '2022-11-09', '2023-05-06 07:33:03', '2023-05-06 07:33:03', NULL),
(6, 1, 'Very Super Gud', 'Highly Recommend', 10, 'https://www.w3schools.com/howto/img_forest.jpg', '2022-11-09', '2023-05-06 14:40:08', '2023-05-06 14:40:08', NULL);

-- --------------------------------------------------------

--
-- Table structure for table `BannerDetail`
--

CREATE TABLE `BannerDetail` (
  `banner_id` bigint UNSIGNED NOT NULL,
  `product_id` bigint UNSIGNED NOT NULL,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `deleted_at` datetime DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--
-- Dumping data for table `BannerDetail`
--

INSERT INTO `BannerDetail` (`banner_id`, `product_id`, `created_at`, `updated_at`, `deleted_at`) VALUES
(1, 7, '2022-12-30 09:56:30', '2022-12-30 09:56:30', NULL),
(1, 8, '2022-12-30 09:56:30', '2022-12-30 09:56:30', NULL),
(2, 9, '2022-12-30 09:57:44', '2022-12-30 09:57:44', NULL),
(2, 10, '2022-12-30 09:57:44', '2022-12-30 09:57:44', NULL),
(3, 11, '2022-12-30 09:59:37', '2022-12-30 09:59:37', NULL),
(3, 12, '2022-12-30 09:59:37', '2022-12-30 09:59:37', NULL),
(1, 7, '2022-12-30 09:56:30', '2022-12-30 09:56:30', NULL),
(1, 8, '2022-12-30 09:56:30', '2022-12-30 09:56:30', NULL),
(2, 9, '2022-12-30 09:57:44', '2022-12-30 09:57:44', NULL),
(2, 10, '2022-12-30 09:57:44', '2022-12-30 09:57:44', NULL),
(3, 11, '2022-12-30 09:59:37', '2022-12-30 09:59:37', NULL),
(3, 12, '2022-12-30 09:59:37', '2022-12-30 09:59:37', NULL),
(1, 7, '2022-12-30 09:56:30', '2022-12-30 09:56:30', NULL),
(1, 8, '2022-12-30 09:56:30', '2022-12-30 09:56:30', NULL),
(2, 9, '2022-12-30 09:57:44', '2022-12-30 09:57:44', NULL),
(2, 10, '2022-12-30 09:57:44', '2022-12-30 09:57:44', NULL),
(3, 11, '2022-12-30 09:59:37', '2022-12-30 09:59:37', NULL),
(3, 12, '2022-12-30 09:59:37', '2022-12-30 09:59:37', NULL),
(5, 7, '2023-05-06 07:33:04', '2023-05-06 07:33:04', NULL),
(5, 8, '2023-05-06 07:33:04', '2023-05-06 07:33:04', NULL),
(5, 9, '2023-05-06 07:33:04', '2023-05-06 07:33:04', NULL),
(6, 7, '2023-05-06 14:40:08', '2023-05-06 14:40:08', NULL),
(6, 8, '2023-05-06 14:40:08', '2023-05-06 14:40:08', NULL),
(6, 9, '2023-05-06 14:40:08', '2023-05-06 14:40:08', NULL),
(1, 7, '2023-05-07 13:21:19', '2023-05-07 13:21:19', NULL),
(1, 8, '2023-05-07 13:21:19', '2023-05-07 13:21:19', NULL),
(1, 9, '2023-05-07 13:21:19', '2023-05-07 13:21:19', NULL),
(1, 10, '2023-05-07 13:21:19', '2023-05-07 13:21:19', NULL),
(1, 11, '2023-05-07 13:21:19', '2023-05-07 13:21:19', NULL),
(1, 12, '2023-05-07 13:21:19', '2023-05-07 13:21:19', NULL),
(1, 7, '2023-05-07 13:21:58', '2023-05-07 13:21:58', NULL),
(1, 8, '2023-05-07 13:21:58', '2023-05-07 13:21:58', NULL),
(1, 9, '2023-05-07 13:21:58', '2023-05-07 13:21:58', NULL),
(1, 10, '2023-05-07 13:21:58', '2023-05-07 13:21:58', NULL),
(1, 11, '2023-05-07 13:21:58', '2023-05-07 13:21:58', NULL),
(1, 12, '2023-05-07 13:21:58', '2023-05-07 13:21:58', NULL),
(1, 7, '2023-05-07 13:22:51', '2023-05-07 13:22:51', NULL),
(1, 8, '2023-05-07 13:22:51', '2023-05-07 13:22:51', NULL),
(1, 9, '2023-05-07 13:22:51', '2023-05-07 13:22:51', NULL),
(1, 10, '2023-05-07 13:22:51', '2023-05-07 13:22:51', NULL),
(1, 11, '2023-05-07 13:22:51', '2023-05-07 13:22:51', NULL),
(1, 12, '2023-05-07 13:22:51', '2023-05-07 13:22:51', NULL),
(1, 7, '2023-05-07 13:28:43', '2023-05-07 13:28:43', NULL),
(1, 8, '2023-05-07 13:28:43', '2023-05-07 13:28:43', NULL),
(1, 9, '2023-05-07 13:28:43', '2023-05-07 13:28:43', NULL),
(1, 10, '2023-05-07 13:28:43', '2023-05-07 13:28:43', NULL),
(1, 11, '2023-05-07 13:28:43', '2023-05-07 13:28:43', NULL),
(1, 12, '2023-05-07 13:28:43', '2023-05-07 13:28:43', NULL),
(1, 7, '2023-05-07 13:29:07', '2023-05-07 13:29:07', NULL),
(1, 8, '2023-05-07 13:29:07', '2023-05-07 13:29:07', NULL),
(1, 9, '2023-05-07 13:29:07', '2023-05-07 13:29:07', NULL),
(1, 10, '2023-05-07 13:29:07', '2023-05-07 13:29:07', NULL),
(1, 11, '2023-05-07 13:29:07', '2023-05-07 13:29:07', NULL),
(1, 12, '2023-05-07 13:29:07', '2023-05-07 13:29:07', NULL),
(1, 7, '2023-05-07 13:32:10', '2023-05-07 13:32:10', NULL),
(1, 8, '2023-05-07 13:32:10', '2023-05-07 13:32:10', NULL),
(1, 9, '2023-05-07 13:32:10', '2023-05-07 13:32:10', NULL),
(1, 10, '2023-05-07 13:32:10', '2023-05-07 13:32:10', NULL),
(1, 11, '2023-05-07 13:32:10', '2023-05-07 13:32:10', NULL),
(1, 12, '2023-05-07 13:32:10', '2023-05-07 13:32:10', NULL),
(1, 7, '2023-05-07 13:32:11', '2023-05-07 13:32:11', NULL),
(1, 8, '2023-05-07 13:32:11', '2023-05-07 13:32:11', NULL),
(1, 9, '2023-05-07 13:32:11', '2023-05-07 13:32:11', NULL),
(1, 10, '2023-05-07 13:32:11', '2023-05-07 13:32:11', NULL),
(1, 11, '2023-05-07 13:32:11', '2023-05-07 13:32:11', NULL),
(1, 12, '2023-05-07 13:32:11', '2023-05-07 13:32:11', NULL),
(1, 7, '2023-05-07 13:32:12', '2023-05-07 13:32:12', NULL),
(1, 8, '2023-05-07 13:32:12', '2023-05-07 13:32:12', NULL),
(1, 9, '2023-05-07 13:32:12', '2023-05-07 13:32:12', NULL),
(1, 10, '2023-05-07 13:32:12', '2023-05-07 13:32:12', NULL),
(1, 11, '2023-05-07 13:32:12', '2023-05-07 13:32:12', NULL),
(1, 12, '2023-05-07 13:32:12', '2023-05-07 13:32:12', NULL),
(1, 7, '2023-05-07 13:33:12', '2023-05-07 13:33:12', NULL),
(1, 8, '2023-05-07 13:33:12', '2023-05-07 13:33:12', NULL),
(1, 9, '2023-05-07 13:33:12', '2023-05-07 13:33:12', NULL),
(1, 10, '2023-05-07 13:33:12', '2023-05-07 13:33:12', NULL),
(1, 11, '2023-05-07 13:33:12', '2023-05-07 13:33:12', NULL),
(1, 12, '2023-05-07 13:33:12', '2023-05-07 13:33:12', NULL),
(1, 7, '2023-05-07 13:33:17', '2023-05-07 13:33:17', NULL),
(1, 8, '2023-05-07 13:33:17', '2023-05-07 13:33:17', NULL),
(1, 9, '2023-05-07 13:33:17', '2023-05-07 13:33:17', NULL),
(1, 10, '2023-05-07 13:33:17', '2023-05-07 13:33:17', NULL),
(1, 11, '2023-05-07 13:33:17', '2023-05-07 13:33:17', NULL),
(1, 12, '2023-05-07 13:33:17', '2023-05-07 13:33:17', NULL),
(1, 7, '2023-05-15 15:39:03', '2023-05-15 15:39:03', NULL),
(1, 8, '2023-05-15 15:39:03', '2023-05-15 15:39:03', NULL),
(1, 9, '2023-05-15 15:39:03', '2023-05-15 15:39:03', NULL),
(1, 10, '2023-05-15 15:39:03', '2023-05-15 15:39:03', NULL),
(1, 11, '2023-05-15 15:39:03', '2023-05-15 15:39:03', NULL),
(1, 12, '2023-05-15 15:39:03', '2023-05-15 15:39:03', NULL);

-- --------------------------------------------------------

--
-- Table structure for table `Cart`
--

CREATE TABLE `Cart` (
  `id` bigint UNSIGNED NOT NULL,
  `provider_id` bigint UNSIGNED NOT NULL,
  `user_id` bigint UNSIGNED NOT NULL,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `deleted_at` datetime DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--
-- Dumping data for table `Cart`
--

INSERT INTO `Cart` (`id`, `provider_id`, `user_id`, `created_at`, `updated_at`, `deleted_at`) VALUES
(5, 7, 1, '2023-01-02 08:01:37', '2023-01-02 08:01:37', '2023-01-02 08:03:06'),
(7, 7, 1, '2023-01-02 08:05:18', '2023-01-02 08:05:18', NULL),
(8, 8, 23, '2023-01-03 16:04:58', '2023-01-03 16:04:58', NULL);

-- --------------------------------------------------------

--
-- Table structure for table `CartItem`
--

CREATE TABLE `CartItem` (
  `id` bigint UNSIGNED NOT NULL,
  `cart_id` bigint UNSIGNED NOT NULL,
  `user_id` bigint UNSIGNED NOT NULL,
  `product_id` bigint UNSIGNED NOT NULL,
  `product_option_id` bigint UNSIGNED NOT NULL,
  `quantity` int UNSIGNED NOT NULL,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `deleted_at` datetime DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--
-- Dumping data for table `CartItem`
--

INSERT INTO `CartItem` (`id`, `cart_id`, `user_id`, `product_id`, `product_option_id`, `quantity`, `created_at`, `updated_at`, `deleted_at`) VALUES
(26, 7, 1, 12, 30, 10, '2023-04-02 13:02:16', '2023-04-02 13:02:16', NULL),
(27, 7, 1, 7, 30, 10, '2023-05-05 15:01:18', '2023-05-05 15:01:18', NULL);

-- --------------------------------------------------------

--
-- Table structure for table `Category`
--

CREATE TABLE `Category` (
  `id` bigint UNSIGNED NOT NULL,
  `category_parent_id` bigint UNSIGNED DEFAULT NULL,
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `image_path` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT 'https://www.englishclub.com/images/vocabulary/food/fish-seafood/fish-seafood.jpg',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `deleted_at` datetime DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--
-- Dumping data for table `Category`
--

INSERT INTO `Category` (`id`, `category_parent_id`, `name`, `image_path`, `created_at`, `updated_at`, `deleted_at`) VALUES
(1, NULL, 'Men\'s Fashion', 'https://freepngimg.com/thumb/model/28064-1-mens-fashion-transparent-background.png', '2022-12-06 16:47:03', '2022-12-06 16:47:03', NULL),
(2, NULL, 'Women\'s Fashion', 'https://freepngimg.com/thumb/anne_hathaway/131356-picture-anne-hathaway-hq-image-free.png', '2022-12-06 16:47:03', '2022-12-06 16:47:03', NULL),
(3, NULL, 'Accessories', 'https://images.squarespace-cdn.com/content/v1/5a49091d1f318dd63985ae87/1556138275926-IIQRRDOYKJSK7ZQ41ILQ/Capsule+Accessories+Collection+-+Must+Have+Jewelry+Pieces_01.JPG?format=1000w', '2022-12-06 16:47:03', '2022-12-06 16:47:03', NULL),
(15, 1, 'Hat', 'https://freepngimg.com/png/1143-baseball-cap-png-image', '2022-12-16 00:52:00', '2022-12-16 00:52:00', NULL),
(16, 1, 'Jean', 'https://freepngimg.com/thumb/jeans/6-jeans-png-image.png', '2022-12-16 00:52:00', '2022-12-16 00:52:00', NULL),
(17, 1, 'T-Shirt', 'https://freepngimg.com/thumb/tshirt/2-t-shirt-png-image.png', '2022-12-16 00:52:00', '2022-12-16 00:52:00', NULL),
(18, 2, 'Hat', 'https://freepngimg.com/png/1143-baseball-cap-png-image', '2022-12-16 00:52:20', '2022-12-16 00:52:20', NULL),
(19, 2, 'Jean', 'https://freepngimg.com/thumb/jeans/6-jeans-png-image.png', '2022-12-16 00:52:20', '2022-12-16 00:52:20', NULL),
(20, 2, 'T-Shirt', 'https://freepngimg.com/thumb/tshirt/2-t-shirt-png-image.png', '2022-12-16 00:52:20', '2022-12-16 00:52:20', NULL),
(22, 2, '{{$ran}}', 'https://d38b044pevnwc9.cloudfront.net/cutout-nuxt/enhancer/2.jpg', '2023-04-15 14:19:40', '2023-04-15 14:19:40', NULL),
(23, 22, '{{$ran}}', 'https://d38b044pevnwc9.cloudfront.net/cutout-nuxt/enhancer/2.jpg', '2023-04-15 14:20:57', '2023-04-15 14:20:57', NULL),
(24, 23, '{{$ran}}', 'https://d38b044pevnwc9.cloudfront.net/cutout-nuxt/enhancer/2.jpg', '2023-04-15 14:21:10', '2023-04-15 14:21:10', NULL),
(25, 24, '{{$ran}}', 'https://d38b044pevnwc9.cloudfront.net/cutout-nuxt/enhancer/2.jpg', '2023-04-15 14:21:17', '2023-04-15 14:21:17', NULL),
(26, 25, '{{$ran}}', 'https://d38b044pevnwc9.cloudfront.net/cutout-nuxt/enhancer/2.jpg', '2023-04-15 14:21:20', '2023-04-15 14:21:20', NULL),
(27, 26, '{{$ran}}', 'https://d38b044pevnwc9.cloudfront.net/cutout-nuxt/enhancer/2.jpg', '2023-04-15 14:21:25', '2023-04-15 14:21:25', NULL),
(28, 26, '{{$ran}}', 'https://d38b044pevnwc9.cloudfront.net/cutout-nuxt/enhancer/2.jpg', '2023-04-15 14:26:10', '2023-04-15 14:26:10', NULL),
(29, 26, '{{$ran}}', 'https://d38b044pevnwc9.cloudfront.net/cutout-nuxt/enhancer/2.jpg', '2023-04-15 14:26:12', '2023-04-15 14:26:12', NULL),
(30, 26, '{{$ran}}', 'https://d38b044pevnwc9.cloudfront.net/cutout-nuxt/enhancer/2.jpg', '2023-04-15 14:26:13', '2023-04-15 14:26:13', NULL),
(31, 26, '{{$ran}}', 'https://d38b044pevnwc9.cloudfront.net/cutout-nuxt/enhancer/2.jpg', '2023-04-15 14:26:13', '2023-04-15 14:26:13', NULL),
(32, 26, '{{$ran}}', 'https://d38b044pevnwc9.cloudfront.net/cutout-nuxt/enhancer/2.jpg', '2023-04-15 14:26:14', '2023-04-15 14:26:14', NULL),
(33, 26, '{{$ran}}', 'https://d38b044pevnwc9.cloudfront.net/cutout-nuxt/enhancer/2.jpg', '2023-04-15 14:26:15', '2023-04-15 14:26:15', NULL),
(34, 26, '{{$ran}}', 'https://d38b044pevnwc9.cloudfront.net/cutout-nuxt/enhancer/2.jpg', '2023-04-15 14:26:16', '2023-04-15 14:26:16', NULL),
(35, 26, '{{$ran}}', 'https://d38b044pevnwc9.cloudfront.net/cutout-nuxt/enhancer/2.jpg', '2023-04-15 14:26:18', '2023-04-15 14:26:18', NULL),
(36, 26, '{{$ran}}', 'https://d38b044pevnwc9.cloudfront.net/cutout-nuxt/enhancer/2.jpg', '2023-04-15 14:26:19', '2023-04-15 14:26:19', NULL);

-- --------------------------------------------------------

--
-- Table structure for table `ChatRoom`
--

CREATE TABLE `ChatRoom` (
  `id` bigint UNSIGNED NOT NULL,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `deleted_at` datetime DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- Dumping data for table `ChatRoom`
--

INSERT INTO `ChatRoom` (`id`, `created_at`, `updated_at`, `deleted_at`) VALUES
(1, '2023-05-13 19:17:08', '2023-05-13 19:17:08', NULL),
(2, '2023-05-13 19:17:08', '2023-05-13 19:17:08', NULL),
(3, '2023-05-13 19:17:08', '2023-05-13 19:17:08', NULL),
(4, '2023-05-13 19:17:08', '2023-05-13 19:17:08', NULL),
(5, '2023-05-13 19:17:08', '2023-05-13 19:17:08', NULL),
(6, '2023-05-13 20:19:04', '2023-05-13 20:19:04', NULL),
(7, '2023-05-13 20:19:23', '2023-05-13 20:19:23', NULL),
(8, '2023-05-13 20:21:47', '2023-05-13 20:21:47', NULL);

-- --------------------------------------------------------

--
-- Table structure for table `Comment`
--

CREATE TABLE `Comment` (
  `id` bigint UNSIGNED NOT NULL,
  `product_id` bigint UNSIGNED NOT NULL,
  `user_id` bigint UNSIGNED NOT NULL,
  `description` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `rating` int NOT NULL,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `deleted_at` datetime DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--
-- Dumping data for table `Comment`
--

INSERT INTO `Comment` (`id`, `product_id`, `user_id`, `description`, `rating`, `created_at`, `updated_at`, `deleted_at`) VALUES
(3, 1, 1, 'werewr', 3, '2022-12-17 16:48:05', '2022-12-17 16:48:05', NULL),
(4, 1, 1, 'fewrfr', 4, '2022-12-17 16:48:05', '2022-12-17 16:48:05', NULL),
(5, 1, 1, 'Something', 0, '2022-12-18 16:08:49', '2022-12-18 16:08:49', NULL),
(6, 1, 1, 'Something', 0, '2022-12-18 16:10:19', '2022-12-18 16:10:19', NULL),
(8, 1, 1, 'Something', 0, '2022-12-20 04:30:15', '2022-12-20 04:30:15', NULL),
(9, 1, 1, 'Something', 0, '2022-12-20 04:30:27', '2022-12-20 04:30:27', NULL),
(10, 1, 1, 'Something', 0, '2022-12-20 04:30:39', '2022-12-20 04:30:39', NULL),
(11, 1, 1, 'Something', 0, '2022-12-20 04:32:23', '2022-12-20 04:32:23', NULL),
(12, 1, 1, 'Something', 0, '2022-12-20 04:32:50', '2022-12-20 04:32:50', NULL),
(13, 1, 1, 'Gud, Gất Gud', 3, '2022-12-20 04:35:52', '2022-12-20 04:35:52', NULL),
(14, 1, 1, 'Gud, Gất Gud', 3, '2022-12-20 04:39:24', '2022-12-20 04:39:24', NULL),
(15, 1, 1, 'Gud, Gất Gud', 3, '2022-12-20 04:40:40', '2022-12-20 04:40:40', NULL),
(16, 1, 1, 'Gud, Gất Gud', 3, '2022-12-20 09:33:47', '2022-12-20 09:33:47', NULL),
(17, 12, 1, 'cũng khá thú vị', 4, '2023-01-01 04:39:09', '2023-01-01 04:39:09', NULL),
(18, 1, 1, '', 3, '2023-04-15 06:33:43', '2023-04-15 06:33:43', NULL),
(19, 1, 1, '', 3, '2023-04-15 06:35:06', '2023-04-15 06:35:06', NULL),
(20, 1, 1, '', 3, '2023-04-15 06:35:51', '2023-04-15 06:35:51', NULL),
(21, 1, 1, 'dssafdsf', 3, '2023-04-15 06:38:10', '2023-04-15 06:38:10', NULL),
(22, 1, 1, 'dssafdsf', 3, '2023-04-15 06:38:48', '2023-04-15 06:38:48', NULL),
(23, 1, 1, 'dssafdsf', 3, '2023-04-15 06:39:37', '2023-04-15 06:39:37', NULL),
(24, 1, 1, 'dssafdsf', 3, '2023-04-15 06:40:06', '2023-04-15 06:40:06', NULL),
(25, 1, 1, 'dssafdsf', 3, '2023-04-15 06:40:45', '2023-04-15 06:40:45', NULL);

-- --------------------------------------------------------

--
-- Table structure for table `CommentMedia`
--

CREATE TABLE `CommentMedia` (
  `id` bigint UNSIGNED NOT NULL,
  `comment_id` bigint UNSIGNED NOT NULL,
  `public_id` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `media_path` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `media_type` enum('IMAGE','VIDEO') CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `deleted_at` datetime DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--
-- Dumping data for table `CommentMedia`
--

INSERT INTO `CommentMedia` (`id`, `comment_id`, `public_id`, `media_path`, `media_type`, `created_at`, `updated_at`, `deleted_at`) VALUES
(10, 25, 'xyz', 'mnp', 'IMAGE', '2023-04-15 06:40:45', '2023-04-15 06:40:45', NULL);

-- --------------------------------------------------------

--
-- Table structure for table `Email`
--

CREATE TABLE `Email` (
  `id` bigint NOT NULL,
  `name` varchar(255) NOT NULL,
  `email` varchar(255) NOT NULL,
  `subject` varchar(255) NOT NULL,
  `descriptions` varchar(255) NOT NULL,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- Dumping data for table `Email`
--

INSERT INTO `Email` (`id`, `name`, `email`, `subject`, `descriptions`, `created_at`, `updated_at`) VALUES
(1, '', 'eNViDXT', '', 'I lớp du chịch chịch', '2023-03-22 01:28:21', '2023-03-22 01:28:21'),
(2, 'Nguyen Van A', 'eNViDXT', 'Title', 'I lớp du chịch chịch', '2023-04-22 15:46:30', '2023-04-22 15:46:30'),
(3, 'Nguyen Van A', 'eNViDXT', 'Title', 'I lớp du chịch chịch', '2023-04-28 13:57:47', '2023-04-28 13:57:47');

-- --------------------------------------------------------

--
-- Table structure for table `Favorite`
--

CREATE TABLE `Favorite` (
  `id` bigint UNSIGNED NOT NULL,
  `user_id` bigint UNSIGNED NOT NULL,
  `provider_id` bigint UNSIGNED NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- Dumping data for table `Favorite`
--

INSERT INTO `Favorite` (`id`, `user_id`, `provider_id`) VALUES
(1, 1, 2);

-- --------------------------------------------------------

--
-- Table structure for table `Message`
--

CREATE TABLE `Message` (
  `id` bigint UNSIGNED NOT NULL,
  `chat_room_id` bigint UNSIGNED NOT NULL,
  `from_user_id` bigint UNSIGNED NOT NULL,
  `to_user_id` bigint UNSIGNED NOT NULL,
  `content` varchar(255) NOT NULL,
  `seen` tinyint(1) NOT NULL DEFAULT '0',
  `type` enum('TEXT','MEDIA') NOT NULL,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `deleted_at` datetime DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- Dumping data for table `Message`
--

INSERT INTO `Message` (`id`, `chat_room_id`, `from_user_id`, `to_user_id`, `content`, `seen`, `type`, `created_at`, `updated_at`, `deleted_at`) VALUES
(82, 1, 1, 23, 'dfa', 0, 'TEXT', '2023-05-13 18:35:37', '2023-05-13 18:35:37', NULL),
(83, 2, 1, 24, 'dfa', 0, 'TEXT', '2023-05-13 18:35:39', '2023-05-13 18:35:39', NULL),
(84, 3, 1, 25, 'dfa', 0, 'TEXT', '2023-05-13 18:35:45', '2023-05-13 18:35:45', NULL),
(85, 4, 1, 26, 'dfa', 0, 'TEXT', '2023-05-13 18:35:50', '2023-05-13 18:35:50', NULL),
(86, 5, 1, 22, 'dfa', 1, 'TEXT', '2023-05-13 18:35:55', '2023-05-14 09:15:58', NULL),
(87, 5, 22, 1, 'dfa', 0, 'TEXT', '2023-05-13 18:36:00', '2023-05-13 18:36:00', NULL),
(90, 8, 30, 1, 'sadasd', 0, 'TEXT', '2023-05-13 20:21:48', '2023-05-13 20:21:48', NULL);

-- --------------------------------------------------------

--
-- Table structure for table `Notification`
--

CREATE TABLE `Notification` (
  `id` bigint UNSIGNED NOT NULL,
  `user_id` bigint UNSIGNED NOT NULL,
  `content` varchar(255) NOT NULL,
  `seen` tinyint(1) NOT NULL DEFAULT '0',
  `url` varchar(255) NOT NULL,
  `Image` varchar(255) NOT NULL,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `deleted_at` datetime DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- Dumping data for table `Notification`
--

INSERT INTO `Notification` (`id`, `user_id`, `content`, `seen`, `url`, `Image`, `created_at`, `updated_at`, `deleted_at`) VALUES
(1, 1, 'qưefqwef', 1, 'ádfasdf', '', '2023-05-03 12:28:14', '2023-05-03 12:33:02', NULL),
(2, 1, 'qưefqwef', 0, 'ádfasdf', '', '2023-05-03 12:28:20', '2023-05-03 12:28:20', NULL),
(3, 1, 'qưefqwef', 0, 'ádfasdf', '', '2023-05-03 12:28:20', '2023-05-03 12:28:20', NULL),
(4, 1, 'qưefqwef', 0, 'ádfasdf', '', '2023-05-03 12:28:20', '2023-05-03 12:28:20', NULL),
(5, 1, 'qưefqwef', 0, 'ádfasdf', '', '2023-05-03 12:28:21', '2023-05-03 12:28:21', NULL),
(6, 1, 'qưefqwef', 0, 'ádfasdf', '', '2023-05-03 12:28:21', '2023-05-03 12:28:21', NULL),
(7, 1, 'You have a new order from user: Nguyễn Văn Đạt', 0, 'api/v1/orders/91', '', '2023-05-03 14:37:33', '2023-05-03 14:37:33', NULL),
(8, 1, 'Create order successfully', 0, 'api/v1/orders/91', '', '2023-05-03 14:37:33', '2023-05-03 14:37:33', NULL),
(9, 1, 'You have a new order from user: Nguyễn Văn Đạt', 0, 'api/v1/orders/92', '', '2023-05-03 14:39:20', '2023-05-03 14:39:20', NULL),
(10, 1, 'Create order successfully', 0, 'api/v1/orders/92', '', '2023-05-03 14:39:20', '2023-05-03 14:39:20', NULL),
(11, 1, 'You have a new order from user: Nguyễn Văn Đạt', 0, 'api/v1/orders/93', '', '2023-05-03 14:44:10', '2023-05-03 14:44:10', NULL),
(12, 1, 'Create order successfully', 0, 'api/v1/orders/93', '', '2023-05-03 14:44:10', '2023-05-03 14:44:10', NULL),
(13, 1, 'You have a new order from user: Nguyễn Văn Đạt', 0, 'api/v1/orders/94', '', '2023-05-04 14:33:14', '2023-05-04 14:33:14', NULL),
(14, 1, 'Create order successfully', 0, 'api/v1/orders/94', '', '2023-05-04 14:33:14', '2023-05-04 14:33:14', NULL),
(15, 1, 'Create order successfully', 0, 'api/v1/orders/95', '', '2023-05-04 14:34:38', '2023-05-04 14:34:38', NULL),
(16, 1, 'Create order successfully', 0, 'api/v1/orders/95', '', '2023-05-04 14:34:38', '2023-05-04 14:34:38', NULL),
(17, 1, 'You have a new order from user: Nguyễn Văn Đạt', 0, 'api/v1/orders/96', '', '2023-05-04 14:37:22', '2023-05-04 14:37:22', NULL),
(18, 1, 'Create order successfully', 0, 'api/v1/orders/96', '', '2023-05-04 14:37:22', '2023-05-04 14:37:22', NULL),
(19, 1, 'You have a new order from user: Nguyễn Văn Đạt', 0, 'api/v1/orders/98', '', '2023-05-04 14:44:21', '2023-05-04 14:44:21', NULL),
(20, 1, 'Create order successfully', 0, 'api/v1/orders/98', '', '2023-05-04 14:44:21', '2023-05-04 14:44:21', NULL),
(21, 1, 'Create order successfully', 0, 'api/v1/orders/99', '', '2023-05-04 14:45:11', '2023-05-04 14:45:11', NULL),
(22, 1, 'You have a new order from user: Nguyễn Văn Đạt', 0, 'api/v1/orders/99', '', '2023-05-04 14:45:11', '2023-05-04 14:45:11', NULL),
(23, 1, 'Create order successfully', 0, 'api/v1/orders/99', '', '2023-05-04 14:45:12', '2023-05-04 14:45:12', NULL),
(24, 1, 'You have a new order from user: Nguyễn Văn Đạt', 0, 'api/v1/orders/99', '', '2023-05-04 14:45:12', '2023-05-04 14:45:12', NULL),
(25, 1, 'Create order successfully', 0, 'api/v1/orders/99', '', '2023-05-04 14:45:18', '2023-05-04 14:45:18', NULL),
(26, 1, 'You have a new order from user: Nguyễn Văn Đạt', 0, 'api/v1/orders/99', '', '2023-05-04 14:45:18', '2023-05-04 14:45:18', NULL),
(27, 1, 'Create order successfully', 0, 'api/v1/orders/99', '', '2023-05-04 14:45:28', '2023-05-04 14:45:28', NULL),
(28, 1, 'You have a new order from user: Nguyễn Văn Đạt', 0, 'api/v1/orders/99', '', '2023-05-04 14:45:28', '2023-05-04 14:45:28', NULL),
(29, 1, 'You have a new order from user: Nguyễn Văn Đạt', 0, 'api/v1/orders/100', '', '2023-05-05 14:31:59', '2023-05-05 14:31:59', NULL),
(30, 1, 'Create order successfully', 0, 'api/v1/orders/100', '', '2023-05-05 14:31:59', '2023-05-05 14:31:59', NULL),
(31, 1, 'You have a new order from user: Nguyễn Văn Đạt', 0, 'api/v1/orders/100', '', '2023-05-05 14:32:00', '2023-05-05 14:32:00', NULL),
(32, 1, 'Create order successfully', 0, 'api/v1/orders/100', '', '2023-05-05 14:32:00', '2023-05-05 14:32:00', NULL),
(33, 1, 'You have a new order from user: Nguyễn Văn Đạt', 0, 'api/v1/orders/100', '', '2023-05-05 14:32:05', '2023-05-05 14:32:05', NULL),
(34, 1, 'Create order successfully', 0, 'api/v1/orders/100', '', '2023-05-05 14:32:05', '2023-05-05 14:32:05', NULL),
(35, 1, 'Create order successfully', 0, 'api/v1/orders/100', '', '2023-05-05 14:32:15', '2023-05-05 14:32:15', NULL),
(36, 1, 'You have a new order from user: Nguyễn Văn Đạt', 0, 'api/v1/orders/100', '', '2023-05-05 14:32:15', '2023-05-05 14:32:15', NULL),
(37, 1, 'You have a new order from user: Nguyễn Văn Đạt', 0, 'api/v1/orders/101', '', '2023-05-05 14:33:35', '2023-05-05 14:33:35', NULL),
(38, 1, 'Create order successfully', 0, 'api/v1/orders/101', '', '2023-05-05 14:33:35', '2023-05-05 14:33:35', NULL),
(39, 1, 'You have a new order from user: Nguyễn Văn Đạt', 0, 'api/v1/orders/102', '', '2023-05-06 06:00:27', '2023-05-06 06:00:27', NULL),
(40, 1, 'Create order successfully', 0, 'api/v1/orders/102', '', '2023-05-06 06:00:27', '2023-05-06 06:00:27', NULL),
(41, 1, 'Create order successfully', 0, 'api/v1/orders/103', '', '2023-05-06 06:49:54', '2023-05-06 06:49:54', NULL),
(42, 1, 'You have a new order from user: Nguyễn Văn Đạt', 0, 'api/v1/orders/103', '', '2023-05-06 06:49:54', '2023-05-06 06:49:54', NULL),
(43, 1, 'Create order successfully', 0, '/user/order/104', '', '2023-05-06 14:53:39', '2023-05-06 14:53:39', NULL),
(44, 1, 'You have a new order from user: Nguyễn Văn Đạt', 0, '/brand-detail/order/104', '', '2023-05-06 14:53:39', '2023-05-06 14:53:39', NULL),
(45, 1, 'You have a new order from user: Nguyễn Văn Đạt', 0, '/brand-detail/order/105', '', '2023-05-06 15:13:43', '2023-05-06 15:13:43', NULL),
(46, 1, 'Create order successfully', 0, '/user/order/105', '', '2023-05-06 15:13:43', '2023-05-06 15:13:43', NULL),
(47, 1, 'Create order successfully', 0, '/user/order/106', '', '2023-05-07 07:48:15', '2023-05-07 07:48:15', NULL),
(48, 1, 'You have a new order from user: Nguyễn Văn Đạt', 0, '/brand-detail/order/106', '', '2023-05-07 07:48:15', '2023-05-07 07:48:15', NULL),
(49, 1, 'Your order has arrived', 0, 'localhost:3000/user/order/46', '', '2023-05-07 12:26:36', '2023-05-07 12:26:36', NULL),
(50, 1, 'Your order has arrived', 0, 'localhost:3000/user/order/46', '', '2023-05-07 12:30:57', '2023-05-07 12:30:57', NULL),
(51, 1, 'Create order successfully', 0, '/user/order/107', '', '2023-05-07 14:25:45', '2023-05-07 14:25:45', NULL),
(52, 1, 'You have a new order from user: Nguyễn Văn Đạt', 0, '/brand-detail/order/107', '', '2023-05-07 14:25:45', '2023-05-07 14:25:45', NULL),
(53, 1, 'You have a new order from user: Nguyễn Văn Đạt', 0, '/brand-detail/order/108', '', '2023-05-07 14:46:44', '2023-05-07 14:46:44', NULL),
(54, 1, 'Create order successfully', 0, '/user/order/108', '', '2023-05-07 14:46:44', '2023-05-07 14:46:44', NULL),
(55, 1, 'You have a new order from user: Nguyễn Văn Đạt', 0, '/brand-detail/order/109', '', '2023-05-07 14:49:10', '2023-05-07 14:49:10', NULL),
(56, 1, 'Create order successfully', 0, '/user/order/109', '', '2023-05-07 14:49:10', '2023-05-07 14:49:10', NULL),
(57, 1, 'You have a new order from user: Nguyễn Văn Đạt', 0, '/brand-detail/order/110', '', '2023-05-07 14:52:12', '2023-05-07 14:52:12', NULL),
(58, 1, 'Create order successfully', 0, '/user/order/110', '', '2023-05-07 14:52:12', '2023-05-07 14:52:12', NULL),
(59, 1, 'Create order successfully', 0, '/user/order/111', '', '2023-05-08 09:24:13', '2023-05-08 09:24:13', NULL),
(60, 1, 'You have a new order from user: Nguyễn Văn Đạt', 0, '/brand-detail/order/111', '', '2023-05-08 09:24:13', '2023-05-08 09:24:13', NULL),
(61, 1, 'You have a new order from user: Nguyễn Văn Đạt', 0, '/brand-detail/order/112', '', '2023-05-11 14:50:45', '2023-05-11 14:50:45', NULL),
(62, 1, 'Create order successfully', 0, '/user/order/112', '', '2023-05-11 14:50:45', '2023-05-11 14:50:45', NULL),
(63, 1, 'You have a new order from user: Nguyễn Văn Đạt', 0, '/brand-detail/order/113', '', '2023-05-11 14:51:40', '2023-05-11 14:51:40', NULL),
(64, 1, 'Create order successfully', 0, '/user/order/113', '', '2023-05-11 14:51:40', '2023-05-11 14:51:40', NULL),
(65, 1, 'You have a new order from user: Nguyễn Văn Đạt', 0, '/brand-detail/order/114', '', '2023-05-13 13:54:20', '2023-05-13 13:54:20', NULL),
(66, 1, 'Create order successfully', 0, '/user/order/114', '', '2023-05-13 13:54:20', '2023-05-13 13:54:20', NULL),
(67, 1, 'You have a new order from user: Nguyễn Văn Đạt', 0, '/brand-detail/order/115', '', '2023-05-13 14:12:19', '2023-05-13 14:12:19', NULL),
(68, 1, 'Create order successfully', 0, '/user/order/115', '', '2023-05-13 14:12:19', '2023-05-13 14:12:19', NULL),
(69, 1, 'Create order successfully', 0, '/user/order/116', '', '2023-05-13 14:12:53', '2023-05-13 14:12:53', NULL),
(70, 1, 'You have a new order from user: Nguyễn Văn Đạt', 0, '/brand-detail/order/116', '', '2023-05-13 14:12:53', '2023-05-13 14:12:53', NULL),
(71, 1, 'Create order successfully', 0, '/user/order/117', '', '2023-05-13 14:13:28', '2023-05-13 14:13:28', NULL),
(72, 1, 'You have a new order from user: Nguyễn Văn Đạt', 0, '/brand-detail/order/117', '', '2023-05-13 14:13:28', '2023-05-13 14:13:28', NULL),
(73, 1, 'You have a new order from user: Nguyễn Văn Đạt', 0, '/brand-detail/order/118', '', '2023-05-13 14:14:16', '2023-05-13 14:14:16', NULL),
(74, 1, 'Create order successfully', 0, '/user/order/118', '', '2023-05-13 14:14:16', '2023-05-13 14:14:16', NULL),
(75, 1, 'Create order successfully', 0, '/user/order/119', '', '2023-05-13 14:15:06', '2023-05-13 14:15:06', NULL),
(76, 1, 'You have a new order from user: Nguyễn Văn Đạt', 0, '/brand-detail/order/119', '', '2023-05-13 14:15:06', '2023-05-13 14:15:06', NULL),
(77, 1, 'You have a new order from user: Nguyễn Văn Đạt', 0, '/brand-detail/order/120', '', '2023-05-13 14:15:49', '2023-05-13 14:15:49', NULL),
(78, 1, 'Create order successfully', 0, '/user/order/120', '', '2023-05-13 14:15:49', '2023-05-13 14:15:49', NULL),
(79, 1, 'Create order successfully', 0, '/user/order/121', 'http://res.cloudinary.com/damzcas3k/image/upload/v1684051785/Product/itl4m7o3jsmtqb2mhtt1.png', '2023-05-14 08:11:32', '2023-05-14 08:11:32', NULL),
(80, 1, 'You have a new order from user: Nguyễn Văn Đạt', 0, '/brand-detail/order/121', 'http://res.cloudinary.com/damzcas3k/image/upload/v1684051785/Product/itl4m7o3jsmtqb2mhtt1.png', '2023-05-14 08:11:32', '2023-05-14 08:11:32', NULL),
(81, 1, 'You have a new order from user: Nguyễn Văn Đạt', 0, '/brand-detail/order/122', 'http://res.cloudinary.com/damzcas3k/image/upload/v1684051785/Product/itl4m7o3jsmtqb2mhtt1.png', '2023-05-14 08:11:40', '2023-05-14 08:11:40', NULL),
(82, 1, 'Create order successfully', 0, '/user/order/122', 'http://res.cloudinary.com/damzcas3k/image/upload/v1684051785/Product/itl4m7o3jsmtqb2mhtt1.png', '2023-05-14 08:11:40', '2023-05-14 08:11:40', NULL);

-- --------------------------------------------------------

--
-- Table structure for table `Order`
--

CREATE TABLE `Order` (
  `id` bigint UNSIGNED NOT NULL,
  `user_id` bigint UNSIGNED NOT NULL,
  `provider_id` bigint UNSIGNED NOT NULL,
  `payment_id` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `payment_url` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `cod` tinyint(1) NOT NULL DEFAULT '0',
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `gender` tinyint(1) NOT NULL DEFAULT '1',
  `phone` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `province` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `district` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `ward` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `street` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `quantity` int NOT NULL,
  `total` bigint NOT NULL,
  `discount` int NOT NULL DEFAULT '0',
  `status` enum('WAITING','CONFIRMED','DELIVERING','DELIVERED','CANCEL') CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT 'WAITING',
  `status_description` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `delivered_image` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `verify_delivered` tinyint(1) NOT NULL DEFAULT '0',
  `delivered_date` datetime DEFAULT NULL,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `deleted_at` datetime DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--
-- Dumping data for table `Order`
--

INSERT INTO `Order` (`id`, `user_id`, `provider_id`, `payment_id`, `payment_url`, `cod`, `name`, `gender`, `phone`, `province`, `district`, `ward`, `street`, `quantity`, `total`, `discount`, `status`, `status_description`, `delivered_image`, `verify_delivered`, `delivered_date`, `created_at`, `updated_at`, `deleted_at`) VALUES
(112, 1, 2, NULL, NULL, 0, 'Ngừi Địp', 1, '1234567890', 'Bình Chánh', 'Ninh Thuận', 'Quảng Bình', 'Trường chinh', 1, 1, 20, 'CANCEL', 'Provider Will call you soon', NULL, 0, NULL, '2023-05-11 14:50:45', '2023-05-11 15:01:05', NULL),
(113, 1, 2, NULL, NULL, 0, 'Ngừi Địp', 1, '1234567890', 'Bình Chánh', 'Ninh Thuận', 'Quảng Bình', 'Trường chinh', 1, 1, 20, 'CANCEL', 'Provider Will call you soon', NULL, 0, NULL, '2023-05-11 14:51:40', '2023-05-11 15:01:05', NULL),
(114, 1, 2, NULL, NULL, 0, 'Ngừi Địp', 1, '1234567890', 'Bình Chánh', 'Ninh Thuận', 'Quảng Bình', 'Trường chinh', 1, 1, 20, 'CANCEL', 'Provider Will call you soon', NULL, 0, NULL, '2023-05-13 13:54:20', '2023-05-15 13:36:25', NULL),
(115, 1, 2, NULL, NULL, 0, 'Ngừi Địp', 1, '1234567890', 'Bình Chánh', 'Ninh Thuận', 'Quảng Bình', 'Trường chinh', 1, 1, 20, 'CANCEL', 'Provider Will call you soon', NULL, 0, NULL, '2023-05-13 14:12:19', '2023-05-15 13:36:25', NULL),
(116, 1, 2, NULL, NULL, 0, 'Ngừi Địp', 1, '1234567890', 'Bình Chánh', 'Ninh Thuận', 'Quảng Bình', 'Trường chinh', 1, 1, 20, 'CANCEL', 'Provider Will call you soon', NULL, 0, NULL, '2023-05-13 14:12:53', '2023-05-15 13:36:25', NULL),
(117, 1, 2, NULL, NULL, 0, 'Ngừi Địp', 1, '1234567890', 'Bình Chánh', 'Ninh Thuận', 'Quảng Bình', 'Trường chinh', 1, 1, 20, 'CANCEL', 'Provider Will call you soon', NULL, 0, NULL, '2023-05-13 14:13:16', '2023-05-15 13:36:26', NULL),
(118, 1, 2, NULL, NULL, 0, 'Ngừi Địp', 1, '1234567890', 'Bình Chánh', 'Ninh Thuận', 'Quảng Bình', 'Trường chinh', 1, 1, 20, 'CANCEL', 'Provider Will call you soon', NULL, 0, NULL, '2023-05-13 14:13:47', '2023-05-15 13:36:26', NULL),
(119, 1, 2, NULL, NULL, 0, 'Ngừi Địp', 1, '1234567890', 'Bình Chánh', 'Ninh Thuận', 'Quảng Bình', 'Trường chinh', 1, 1, 20, 'CANCEL', 'Provider Will call you soon', NULL, 0, NULL, '2023-05-13 14:14:54', '2023-05-15 13:36:26', NULL),
(120, 1, 2, NULL, NULL, 0, 'Ngừi Địp', 1, '1234567890', 'Bình Chánh', 'Ninh Thuận', 'Quảng Bình', 'Trường chinh', 1, 1, 20, 'CANCEL', 'Provider Will call you soon', NULL, 0, NULL, '2023-05-13 14:15:46', '2023-05-15 13:36:26', NULL),
(121, 1, 2, NULL, NULL, 0, 'Ngừi Địp', 1, '1234567890', 'Bình Chánh', 'Ninh Thuận', 'Quảng Bình', 'Trường chinh', 1, 1, 20, 'CANCEL', 'Provider Will call you soon', NULL, 0, NULL, '2023-05-14 08:11:32', '2023-05-15 13:36:26', NULL),
(122, 1, 2, NULL, NULL, 0, 'Ngừi Địp', 1, '1234567890', 'Bình Chánh', 'Ninh Thuận', 'Quảng Bình', 'Trường chinh', 1, 1, 20, 'CANCEL', 'Provider Will call you soon', NULL, 0, NULL, '2023-05-14 08:11:40', '2023-05-15 13:36:27', NULL);

-- --------------------------------------------------------

--
-- Table structure for table `OrderItem`
--

CREATE TABLE `OrderItem` (
  `id` bigint UNSIGNED NOT NULL,
  `order_id` bigint UNSIGNED DEFAULT NULL,
  `product_id` bigint UNSIGNED DEFAULT NULL,
  `product_option_id` bigint UNSIGNED DEFAULT NULL,
  `provider_id` bigint UNSIGNED NOT NULL DEFAULT '1',
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `price` bigint UNSIGNED NOT NULL,
  `option` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `quantity` int UNSIGNED NOT NULL,
  `discount` int UNSIGNED NOT NULL DEFAULT '0',
  `image` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `deleted_at` datetime DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--
-- Dumping data for table `OrderItem`
--

INSERT INTO `OrderItem` (`id`, `order_id`, `product_id`, `product_option_id`, `provider_id`, `name`, `price`, `option`, `quantity`, `discount`, `image`, `created_at`, `updated_at`, `deleted_at`) VALUES
(1, 112, 1, 30, 2, 'Dummy Name', 1, 'Spec: Option', 1, 1, 'ádf', '2023-05-11 14:50:45', '2023-05-11 14:50:45', NULL),
(2, 113, 1, 30, 2, 'Dummy Name', 1, 'Spec: Option', 1, 1, 'ádf', '2023-05-11 14:51:39', '2023-05-11 14:51:39', NULL),
(3, 114, 1, 30, 2, 'Dummy Name', 1, 'Spec: Option', 1, 1, 'ádf', '2023-05-13 13:54:20', '2023-05-13 13:54:20', NULL),
(4, 115, 1, 30, 2, 'Dummy Name', 1, 'Spec: Option', 1, 1, 'ádf', '2023-05-13 14:12:19', '2023-05-13 14:12:19', NULL),
(5, 116, 1, 30, 2, 'Dummy Name', 1, 'Spec: Option', 1, 1, 'ádf', '2023-05-13 14:12:53', '2023-05-13 14:12:53', NULL),
(6, 117, 1, 30, 2, 'Dummy Name', 1, 'Spec: Option', 1, 1, 'ádf', '2023-05-13 14:13:15', '2023-05-13 14:13:15', NULL),
(7, 118, 1, 30, 2, 'Dummy Name', 1, 'Spec: Option', 1, 1, 'ádf', '2023-05-13 14:13:47', '2023-05-13 14:13:47', NULL),
(8, 119, 1, 30, 2, 'Dummy Name', 1, 'Spec: Option', 1, 1, 'ádf', '2023-05-13 14:14:53', '2023-05-13 14:14:53', NULL),
(9, 120, 1, 30, 2, 'Dummy Name', 1, 'Spec: Option', 1, 1, 'ádf', '2023-05-13 14:15:46', '2023-05-13 14:15:46', NULL),
(10, 121, 1, 30, 2, 'Dummy Name', 1, 'Spec: Option', 1, 1, 'ádf', '2023-05-14 08:11:32', '2023-05-14 08:11:32', NULL),
(11, 122, 1, 30, 2, 'Dummy Name', 1, 'Spec: Option', 1, 1, 'ádf', '2023-05-14 08:11:39', '2023-05-14 08:11:39', NULL);

-- --------------------------------------------------------

--
-- Table structure for table `Payment`
--

CREATE TABLE `Payment` (
  `id` varchar(255) NOT NULL,
  `status` tinyint(1) NOT NULL DEFAULT '0',
  `account_id` varchar(255) NOT NULL,
  `email` varchar(255) NOT NULL,
  `name` varchar(255) NOT NULL,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `deleted_at` datetime DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- Dumping data for table `Payment`
--

INSERT INTO `Payment` (`id`, `status`, `account_id`, `email`, `name`, `created_at`, `updated_at`, `deleted_at`) VALUES
('ád', 1, '2', 'abc@gmail.com', 'Nguyễn Văn A', '2023-05-10 23:54:35', '2023-05-10 23:54:35', NULL),
('abc', 1, '2', 'abc@gmail.com', 'Nguyễn Văn A', '2023-05-10 23:55:34', '2023-05-10 23:55:34', NULL);

-- --------------------------------------------------------

--
-- Table structure for table `Product`
--

CREATE TABLE `Product` (
  `id` bigint UNSIGNED NOT NULL,
  `provider_id` bigint UNSIGNED NOT NULL DEFAULT '1',
  `category_id` bigint UNSIGNED NOT NULL,
  `user_id` bigint UNSIGNED DEFAULT NULL,
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `price` bigint NOT NULL,
  `discount` int DEFAULT NULL,
  `length` int NOT NULL,
  `height` int NOT NULL,
  `weight` int NOT NULL,
  `width` int NOT NULL,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `deleted_at` datetime DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--
-- Dumping data for table `Product`
--

INSERT INTO `Product` (`id`, `provider_id`, `category_id`, `user_id`, `name`, `price`, `discount`, `length`, `height`, `weight`, `width`, `created_at`, `updated_at`, `deleted_at`) VALUES
(1, 4, 1, 1, 'HighLight', 100000, 5, 100, 100, 100, 100, '2022-12-17 16:46:42', '2022-12-17 16:46:42', NULL),
(7, 7, 1, 1, 'Nón mới ra', 700000, 30, 456, 3423, 1234, 534, '2022-12-30 09:48:29', '2022-12-30 09:48:29', NULL),
(8, 7, 1, 1, 'Nón mới ra phiên bản 2', 700000, 30, 0, 0, 0, 0, '2022-12-30 09:48:52', '2022-12-30 09:48:52', NULL),
(9, 7, 1, 1, 'Áo nam hàng hiệu', 900000, 30, 0, 0, 0, 0, '2022-12-30 09:49:53', '2022-12-30 09:49:53', NULL),
(10, 7, 1, 1, 'Áo nam hàng hiệu mới ra lò', 900000, 30, 0, 0, 0, 0, '2022-12-30 09:50:10', '2022-12-30 09:50:10', '2023-04-28 12:26:24'),
(11, 7, 1, 1, 'Áo nữ cute phô mai que', 95000, 30, 0, 0, 0, 0, '2022-12-30 09:50:53', '2022-12-30 09:50:53', NULL),
(12, 7, 1, 1, 'Áo nữ cute phô mai que ver 2', 95000, 30, 432, 431, 12, 43, '2022-12-30 09:51:06', '2022-12-30 09:51:06', NULL),
(15, 2, 1, 1, '1', 1, 1, 0, 0, 0, 0, '2023-01-06 12:06:17', '2023-01-06 12:06:17', NULL),
(16, 2, 1, 1, '1', 1, 1, 0, 0, 0, 0, '2023-01-06 12:06:55', '2023-01-06 12:06:55', NULL),
(20, 2, 1, 1, 'abc', 20000, 20, 0, 0, 0, 0, '2023-01-10 00:25:39', '2023-01-10 00:25:39', NULL),
(21, 2, 1, 1, 'abc', 20000, 20, 0, 0, 0, 0, '2023-01-10 01:27:32', '2023-01-10 01:27:32', NULL),
(22, 2, 1, 1, 'abc', 20000, 20, 0, 0, 0, 0, '2023-02-03 05:47:00', '2023-02-03 05:47:00', NULL),
(23, 2, 1, 1, 'abc', 20000, 20, 0, 0, 0, 0, '2023-03-26 16:24:12', '2023-03-26 16:24:12', NULL),
(24, 2, 1, 1, 'abc', 20000, 20, 0, 0, 0, 0, '2023-03-26 16:45:15', '2023-03-26 16:45:15', NULL),
(25, 2, 1, 1, 'abc', 20000, 20, 0, 0, 0, 0, '2023-03-26 16:47:30', '2023-03-26 16:47:30', NULL),
(26, 2, 1, 1, 'abc', 20000, 20, 0, 0, 0, 0, '2023-03-26 16:56:01', '2023-03-26 16:56:01', NULL),
(27, 2, 1, 1, 'abc', 20000, 20, 0, 0, 0, 0, '2023-03-26 16:56:11', '2023-03-26 16:56:11', NULL),
(28, 2, 1, 1, 'abc', 20000, 20, 0, 0, 0, 0, '2023-03-26 16:56:54', '2023-03-26 16:56:54', NULL),
(29, 2, 1, 1, 'abc', 20000, 20, 0, 0, 0, 0, '2023-03-26 17:01:42', '2023-03-26 17:01:42', NULL),
(30, 2, 1, 1, 'abc', 20000, 20, 0, 0, 0, 0, '2023-03-26 17:02:10', '2023-03-26 17:02:10', NULL),
(31, 2, 1, 1, 'abc', 20000, 20, 0, 0, 0, 0, '2023-03-26 17:07:05', '2023-03-26 17:07:05', NULL),
(32, 2, 1, 1, 'abc', 20000, 20, 0, 0, 0, 0, '2023-03-26 17:07:15', '2023-03-26 17:07:15', NULL),
(33, 2, 1, 1, 'abc', 20000, 20, 0, 0, 0, 0, '2023-03-26 17:09:04', '2023-03-26 17:09:04', NULL),
(34, 2, 1, 1, 'abc', 20000, 20, 0, 0, 0, 0, '2023-03-26 17:12:04', '2023-03-26 17:12:04', NULL),
(35, 2, 1, 1, 'abc', 20000, 20, 0, 0, 0, 0, '2023-03-26 17:12:42', '2023-03-26 17:12:42', NULL),
(36, 2, 1, 1, 'abc', 20000, 20, 0, 0, 0, 0, '2023-04-27 14:37:38', '2023-04-27 14:37:38', NULL),
(37, 2, 1, 1, 'abc', 20000, 20, 0, 0, 0, 0, '2023-04-27 14:41:09', '2023-04-27 14:41:09', NULL),
(38, 2, 1, 1, 'abc', 20000, 20, 0, 0, 0, 0, '2023-04-28 14:41:46', '2023-04-28 14:41:46', NULL),
(39, 2, 1, 1, 'abc', 20000, 20, 0, 0, 0, 0, '2023-05-15 11:05:49', '2023-05-15 11:05:49', NULL),
(40, 2, 1, 1, 'abc', 20000, 20, 100, 100, 100, 100, '2023-05-15 11:06:42', '2023-05-15 11:06:42', NULL),
(41, 2, 1, 1, 'abc', 20000, 20, 100, 100, 100, 100, '2023-05-16 00:26:45', '2023-05-16 00:26:45', NULL);

-- --------------------------------------------------------

--
-- Table structure for table `ProductDescriptions`
--

CREATE TABLE `ProductDescriptions` (
  `id` bigint NOT NULL,
  `product_id` bigint UNSIGNED NOT NULL,
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `public_id` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `descriptions_path` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `deleted_at` datetime DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--
-- Dumping data for table `ProductDescriptions`
--

INSERT INTO `ProductDescriptions` (`id`, `product_id`, `name`, `public_id`, `descriptions_path`, `created_at`, `updated_at`, `deleted_at`) VALUES
(1, 7, 'test', 'ProductDescriptions/mfsveip2ff46xnnsjeoo', 'http://res.cloudinary.com/damzcas3k/raw/upload/v1672393712/ProductDescriptions/mfsveip2ff46xnnsjeoo', '2022-12-30 09:48:31', '2022-12-30 09:48:31', NULL),
(2, 8, 'test', 'ProductDescriptions/qtx3iwlvdqokxw1aqgsk', 'http://res.cloudinary.com/damzcas3k/raw/upload/v1672393734/ProductDescriptions/qtx3iwlvdqokxw1aqgsk', '2022-12-30 09:48:54', '2022-12-30 09:48:54', NULL),
(3, 9, 'test', 'ProductDescriptions/cho3k9i2dc9yftdhlkya', 'http://res.cloudinary.com/damzcas3k/raw/upload/v1672393797/ProductDescriptions/cho3k9i2dc9yftdhlkya', '2022-12-30 09:49:56', '2022-12-30 09:49:56', NULL),
(4, 10, 'test', 'ProductDescriptions/sr9nqkhyg82avwfylbqo', 'http://res.cloudinary.com/damzcas3k/raw/upload/v1672393811/ProductDescriptions/sr9nqkhyg82avwfylbqo', '2022-12-30 09:50:11', '2022-12-30 09:50:11', NULL),
(5, 11, 'test', 'ProductDescriptions/czzbfti91suaoojhourh', 'http://res.cloudinary.com/damzcas3k/raw/upload/v1672393856/ProductDescriptions/czzbfti91suaoojhourh', '2022-12-30 09:50:55', '2022-12-30 09:50:55', NULL),
(6, 12, 'test', 'ProductDescriptions/nhor9db5slwlaanj4j7i', 'http://res.cloudinary.com/damzcas3k/raw/upload/v1672393868/ProductDescriptions/nhor9db5slwlaanj4j7i', '2022-12-30 09:51:07', '2022-12-30 09:51:07', NULL),
(8, 20, 'abc', 'xyz', '', '2023-01-10 00:25:40', '2023-01-10 00:25:40', NULL),
(9, 21, 'abc', 'xyz', '', '2023-01-10 01:27:32', '2023-01-10 01:27:32', NULL),
(10, 22, 'abc', 'xyz', '', '2023-02-03 05:47:00', '2023-02-03 05:47:00', NULL),
(11, 23, 'abc', 'xyz', '', '2023-03-26 16:24:12', '2023-03-26 16:24:12', NULL),
(12, 35, 'abc', 'xyz', '', '2023-03-26 17:12:42', '2023-03-26 17:12:42', NULL),
(13, 36, 'abc', 'xyz', '', '2023-04-27 14:38:35', '2023-04-27 14:38:35', NULL),
(14, 37, 'abc', 'xyz', 'mnp', '2023-04-27 14:41:18', '2023-04-27 14:41:18', NULL),
(15, 15, 'dasd', 'asdasd', 'asd', '2023-04-27 15:11:47', '2023-04-27 15:11:47', NULL),
(16, 1, 'gerg', 'werg', 'werg', '2023-04-27 15:12:13', '2023-04-27 15:12:13', NULL),
(17, 1, 'asd', 'ProductDescriptions/wne9jj48ufqp83fkxzwn', 'http://res.cloudinary.com/damzcas3k/image/upload/v1682608941/ProductDescriptions/wne9jj48ufqp83fkxzwn.jpg', '2023-04-27 15:12:19', '2023-04-27 15:12:19', NULL),
(18, 38, 'abc', 'xyz', 'mnp', '2023-04-28 14:41:47', '2023-04-28 14:41:47', NULL),
(19, 39, 'abc', 'xyz', 'mnp', '2023-05-15 11:05:49', '2023-05-15 11:05:49', NULL),
(20, 40, 'abc', 'xyz', 'mnp', '2023-05-15 11:06:42', '2023-05-15 11:06:42', NULL),
(21, 41, 'abc', 'xyz', 'mnp', '2023-05-16 00:26:46', '2023-05-16 00:26:46', NULL);

-- --------------------------------------------------------

--
-- Table structure for table `ProductMedia`
--

CREATE TABLE `ProductMedia` (
  `id` bigint UNSIGNED NOT NULL,
  `product_id` bigint UNSIGNED NOT NULL,
  `public_id` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `media_path` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `media_type` enum('IMAGE_FULL','IMAGE_ITEM','VIDEO','IMAGE') CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `deleted_at` datetime DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--
-- Dumping data for table `ProductMedia`
--

INSERT INTO `ProductMedia` (`id`, `product_id`, `public_id`, `media_path`, `media_type`, `created_at`, `updated_at`, `deleted_at`) VALUES
(1, 7, 'Product/ohkzwuw3zqdt3vdn778x', 'asdas', 'IMAGE', '2022-12-30 09:48:31', '2023-05-02 15:13:53', NULL),
(2, 7, 'Product/m2kxncmeuel2qu2inxqn', 'http://res.cloudinary.com/damzcas3k/image/upload/v1672393711/Product/m2kxncmeuel2qu2inxqn.jpg', 'IMAGE', '2022-12-30 09:48:31', '2022-12-30 09:48:31', NULL),
(3, 8, 'Product/dystsi1tr2hdet8tu4ap', 'http://res.cloudinary.com/damzcas3k/image/upload/v1672393733/Product/dystsi1tr2hdet8tu4ap.jpg', 'IMAGE', '2022-12-30 09:48:53', '2022-12-30 09:48:53', NULL),
(4, 8, 'Product/rbfpbjpmleadmhhcwpge', 'http://res.cloudinary.com/damzcas3k/image/upload/v1672393733/Product/rbfpbjpmleadmhhcwpge.jpg', 'IMAGE', '2022-12-30 09:48:53', '2022-12-30 09:48:53', NULL),
(5, 9, 'Product/sfnd6ueanwntehvwkqpe', 'http://res.cloudinary.com/damzcas3k/image/upload/v1672393794/Product/sfnd6ueanwntehvwkqpe.jpg', 'IMAGE', '2022-12-30 09:49:55', '2022-12-30 09:49:55', NULL),
(6, 9, 'Product/m8mse51ngrxuqzhdudif', 'http://res.cloudinary.com/damzcas3k/image/upload/v1672393796/Product/m8mse51ngrxuqzhdudif.jpg', 'IMAGE', '2022-12-30 09:49:55', '2022-12-30 09:49:55', NULL),
(7, 10, 'Product/k7htd5rmssz4d9xwejz1', 'http://res.cloudinary.com/damzcas3k/image/upload/v1672393811/Product/k7htd5rmssz4d9xwejz1.jpg', 'IMAGE', '2022-12-30 09:50:11', '2022-12-30 09:50:11', NULL),
(8, 11, 'Product/sgslhzidjx6ffjyhnxro', 'http://res.cloudinary.com/damzcas3k/image/upload/v1672393855/Product/sgslhzidjx6ffjyhnxro.jpg', 'IMAGE', '2022-12-30 09:50:55', '2022-12-30 09:50:55', NULL),
(9, 12, 'Product/xtfibx615kxet8mbl182', 'http://res.cloudinary.com/damzcas3k/image/upload/v1672393867/Product/xtfibx615kxet8mbl182.jpg', 'IMAGE', '2022-12-30 09:51:07', '2022-12-30 09:51:07', NULL),
(10, 15, 'Product/ks94koekpxlqnxvt1p6v', 'http://res.cloudinary.com/damzcas3k/image/upload/v1673006781/Product/ks94koekpxlqnxvt1p6v.png', 'IMAGE', '2023-01-06 12:06:21', '2023-01-06 12:06:21', NULL),
(11, 16, 'Product/geiekkzljbbtlptqqjmp', 'http://res.cloudinary.com/damzcas3k/image/upload/v1673006819/Product/geiekkzljbbtlptqqjmp.png', 'IMAGE', '2023-01-06 12:06:59', '2023-01-06 12:06:59', NULL),
(12, 20, 'xyz', 'mnp', 'IMAGE', '2023-01-10 00:25:39', '2023-01-10 00:25:39', NULL),
(13, 21, 'xyz', 'mnp', 'IMAGE', '2023-01-10 01:27:32', '2023-01-10 01:27:32', NULL),
(14, 22, 'xyz', 'mnp', 'IMAGE', '2023-02-03 05:47:00', '2023-02-03 05:47:00', NULL),
(15, 23, 'xyz', 'mnp', 'IMAGE', '2023-03-26 16:24:12', '2023-03-26 16:24:12', NULL),
(16, 24, 'xyz', 'mnp', 'IMAGE', '2023-03-26 16:45:15', '2023-03-26 16:45:15', NULL),
(17, 25, 'xyz', 'mnp', 'IMAGE', '2023-03-26 16:47:30', '2023-03-26 16:47:30', NULL),
(18, 26, 'xyz', 'mnp', 'IMAGE', '2023-03-26 16:56:01', '2023-03-26 16:56:01', NULL),
(19, 27, 'xyz', 'mnp', 'IMAGE', '2023-03-26 16:56:11', '2023-03-26 16:56:11', NULL),
(20, 28, 'xyz', 'mnp', 'IMAGE', '2023-03-26 16:56:54', '2023-03-26 16:56:54', NULL),
(21, 29, 'xyz', 'mnp', 'IMAGE', '2023-03-26 17:01:42', '2023-03-26 17:01:42', NULL),
(22, 30, 'xyz', 'mnp', 'IMAGE', '2023-03-26 17:02:11', '2023-03-26 17:02:11', NULL),
(23, 31, 'xyz', 'mnp', 'IMAGE', '2023-03-26 17:07:05', '2023-03-26 17:07:05', NULL),
(24, 32, 'xyz', 'mnp', 'IMAGE', '2023-03-26 17:07:15', '2023-03-26 17:07:15', NULL),
(25, 33, 'xyz', 'mnp', 'IMAGE', '2023-03-26 17:09:05', '2023-03-26 17:09:05', NULL),
(26, 34, 'xyz', 'mnp', 'IMAGE', '2023-03-26 17:12:04', '2023-03-26 17:12:04', NULL),
(27, 35, 'xyz', 'mnp', 'IMAGE', '2023-03-26 17:12:42', '2023-03-26 17:12:42', NULL),
(28, 36, 'xyz', 'mnp', 'IMAGE', '2023-04-27 14:37:38', '2023-04-27 14:37:38', NULL),
(29, 37, 'xyz', 'mnp', 'IMAGE', '2023-04-27 14:41:09', '2023-04-27 14:41:09', NULL),
(30, 38, 'xyz', 'mnp', 'IMAGE', '2023-04-28 14:41:46', '2023-04-28 14:41:46', NULL),
(31, 39, 'xyz', 'mnp', 'IMAGE', '2023-05-15 11:05:49', '2023-05-15 11:05:49', NULL),
(32, 40, 'xyz', 'mnp', 'IMAGE', '2023-05-15 11:06:42', '2023-05-15 11:06:42', NULL),
(33, 41, 'xyz', 'mnp', 'IMAGE', '2023-05-16 00:26:45', '2023-05-16 00:26:45', NULL);

-- --------------------------------------------------------

--
-- Table structure for table `ProductOption`
--

CREATE TABLE `ProductOption` (
  `id` bigint UNSIGNED NOT NULL,
  `product_id` bigint UNSIGNED NOT NULL,
  `specification_id` bigint UNSIGNED NOT NULL,
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `price` int NOT NULL,
  `quantity` int NOT NULL,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `deleted_at` datetime DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--
-- Dumping data for table `ProductOption`
--

INSERT INTO `ProductOption` (`id`, `product_id`, `specification_id`, `name`, `price`, `quantity`, `created_at`, `updated_at`, `deleted_at`) VALUES
(30, 35, 21, 'Him', 20000, 52, '2023-03-26 17:12:42', '2023-03-26 17:12:42', NULL),
(31, 35, 21, 'Who', 20000, 100, '2023-03-26 17:12:42', '2023-03-26 17:12:42', NULL),
(32, 35, 21, 'Whom', 20000, 100, '2023-03-26 17:12:42', '2023-03-26 17:12:42', NULL),
(90, 12, 41, 'Him', 20000, 100, '2023-03-26 17:43:05', '2023-03-26 17:43:05', NULL),
(91, 12, 41, 'Who', 20000, 100, '2023-03-26 17:43:05', '2023-03-26 17:43:05', NULL),
(92, 12, 41, 'Whom', 20000, 100, '2023-03-26 17:43:05', '2023-03-26 17:43:05', NULL),
(96, 36, 43, 'Him', 20000, 100, '2023-04-27 14:37:38', '2023-04-27 14:37:38', NULL),
(97, 36, 43, 'Who', 20000, 100, '2023-04-27 14:37:38', '2023-04-27 14:37:38', NULL),
(98, 36, 43, 'Whom', 20000, 100, '2023-04-27 14:37:38', '2023-04-27 14:37:38', NULL),
(99, 37, 44, 'Him', 20000, 100, '2023-04-27 14:41:09', '2023-04-27 14:41:09', NULL),
(100, 37, 44, 'Who', 20000, 100, '2023-04-27 14:41:09', '2023-04-27 14:41:09', NULL),
(101, 37, 44, 'Whom', 20000, 100, '2023-04-27 14:41:09', '2023-04-27 14:41:09', NULL),
(102, 38, 45, 'Him', 20000, 100, '2023-04-28 14:41:47', '2023-04-28 14:41:47', NULL),
(103, 38, 45, 'Who', 20000, 100, '2023-04-28 14:41:47', '2023-04-28 14:41:47', NULL),
(104, 38, 45, 'Whom', 20000, 100, '2023-04-28 14:41:47', '2023-04-28 14:41:47', NULL),
(105, 39, 46, 'Him', 20000, 100, '2023-05-15 11:05:49', '2023-05-15 11:05:49', NULL),
(106, 39, 46, 'Who', 20000, 100, '2023-05-15 11:05:49', '2023-05-15 11:05:49', NULL),
(107, 39, 46, 'Whom', 20000, 100, '2023-05-15 11:05:49', '2023-05-15 11:05:49', NULL),
(108, 40, 47, 'Him', 20000, 100, '2023-05-15 11:06:42', '2023-05-15 11:06:42', NULL),
(109, 40, 47, 'Who', 20000, 100, '2023-05-15 11:06:42', '2023-05-15 11:06:42', NULL),
(110, 40, 47, 'Whom', 20000, 100, '2023-05-15 11:06:42', '2023-05-15 11:06:42', NULL),
(111, 41, 48, 'Him', 20000, 100, '2023-05-16 00:26:46', '2023-05-16 00:26:46', NULL),
(112, 41, 48, 'Who', 20000, 100, '2023-05-16 00:26:46', '2023-05-16 00:26:46', NULL),
(113, 41, 48, 'Whom', 20000, 100, '2023-05-16 00:26:46', '2023-05-16 00:26:46', NULL);

-- --------------------------------------------------------

--
-- Table structure for table `ProductSpecification`
--

CREATE TABLE `ProductSpecification` (
  `id` bigint UNSIGNED NOT NULL,
  `product_id` bigint UNSIGNED NOT NULL,
  `properties` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `deleted_at` datetime DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--
-- Dumping data for table `ProductSpecification`
--

INSERT INTO `ProductSpecification` (`id`, `product_id`, `properties`, `created_at`, `updated_at`, `deleted_at`) VALUES
(18, 10, 'Hường xuân', '2023-03-26 17:07:15', '2023-03-26 17:07:15', NULL),
(20, 34, 'Hường xuân', '2023-03-26 17:12:04', '2023-03-26 17:12:04', NULL),
(21, 35, 'Hường xuân', '2023-03-26 17:12:42', '2023-03-26 17:12:42', NULL),
(41, 12, 'Hường xuân', '2023-03-26 17:43:05', '2023-03-26 17:43:05', NULL),
(43, 36, 'Hường xuân', '2023-04-27 14:37:38', '2023-04-27 14:37:38', NULL),
(44, 37, 'Hường xuân', '2023-04-27 14:41:09', '2023-04-27 14:41:09', NULL),
(45, 38, 'Hường xuân', '2023-04-28 14:41:46', '2023-04-28 14:41:46', NULL),
(46, 39, 'Hường xuân', '2023-05-15 11:05:49', '2023-05-15 11:05:49', NULL),
(47, 40, 'Hường xuân', '2023-05-15 11:06:42', '2023-05-15 11:06:42', NULL),
(48, 41, 'Hường xuân', '2023-05-16 00:26:45', '2023-05-16 00:26:45', NULL);

-- --------------------------------------------------------

--
-- Table structure for table `Provider`
--

CREATE TABLE `Provider` (
  `id` bigint UNSIGNED NOT NULL,
  `user_id` bigint UNSIGNED NOT NULL,
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `image_path` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT 'https://img.freepik.com/free-vector/shop-with-sign-we-are-open_52683-38687.jpg?w=2000',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `deleted_at` datetime DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--
-- Dumping data for table `Provider`
--

INSERT INTO `Provider` (`id`, `user_id`, `name`, `image_path`, `created_at`, `updated_at`, `deleted_at`) VALUES
(2, 1, 'Thần Bài', 'https://img.etimg.com/thumb/width-1200,height-900,imgsize-122620,resizemode-1,msid-75214721/industry/services/retail/future-group-negotiates-rents-for-its-1700-stores.jpg', '2022-12-16 00:54:44', '2022-12-16 00:54:44', '2023-01-30 13:35:10'),
(3, 1, 'Thiên Lợi', 'https://img.freepik.com/free-vector/shop-with-we-are-open-sign_23-2148557016.jpg?w=2000', '2022-12-16 00:55:05', '2022-12-16 00:55:05', NULL),
(4, 1, 'Ngẫu Trùng', 'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcTxwHLbVDRABAok4N222HgTdjHQfdZftcvO8w&usqp=CAU', '2022-12-16 00:55:05', '2022-12-16 00:55:05', NULL),
(7, 1, 'Chi nhánh 1', 'https://hanoispiritofplace.com/wp-content/uploads/2017/06/hinh-nen-conan-edogawa-58.jpg', '2022-12-30 09:47:57', '2022-12-30 09:47:57', NULL),
(8, 1, 'Chi nhánh 2', 'http://res.cloudinary.com/damzcas3k/image/upload/v1672393692/Product/i6stmcr9gnqvsvhbkjwt.jpg', '2022-12-30 09:48:15', '2022-12-30 09:48:15', NULL),
(9, 1, 'Chi nhánh thần bài 3', 'https://toigingiuvedep.vn/wp-content/uploads/2022/02/anh-than-bai.jpg', '2022-12-30 14:15:20', '2022-12-30 14:15:20', NULL),
(10, 1, 'admin', 'https://res.cloudinary.com/damzcas3k/image/upload/v1671291124/Product/cyydihkdpdhdv7ktlq5e.png', '2023-01-30 13:20:17', '2023-01-30 13:20:17', NULL),
(11, 1, 'admin', 'https://res.cloudinary.com/damzcas3k/image/upload/v1671291124/Product/cyydihkdpdhdv7ktlq5e.png', '2023-01-30 13:35:49', '2023-01-30 13:35:49', NULL),
(12, 1, 'admin', 'https://res.cloudinary.com/damzcas3k/image/upload/v1671291124/Product/cyydihkdpdhdv7ktlq5e.png', '2023-03-19 15:27:35', '2023-03-19 15:27:35', NULL),
(13, 1, 'Addd', 'https://res.cloudinary.com/damzcas3k/image/upload/v1671291124/Product/cyydihkdpdhdv7ktlq5e.png', '2023-03-21 15:49:15', '2023-03-21 15:49:15', NULL);

-- --------------------------------------------------------

--
-- Table structure for table `User`
--

CREATE TABLE `User` (
  `id` bigint UNSIGNED NOT NULL,
  `username` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `password` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `salt` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `birthday` date DEFAULT NULL,
  `gender` tinyint(1) DEFAULT '1',
  `email` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `phone` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `coin` int NOT NULL DEFAULT '0',
  `type` enum('BUYER','ADMIN') CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `avatar` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT 'https://4xucy2kyby51ggkud2tadg3d-wpengine.netdna-ssl.com/wp-content/uploads/sites/37/2017/02/IAFOR-Blank-Avatar-Image.jpg',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `deleted_at` timestamp NULL DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--
-- Dumping data for table `User`
--

INSERT INTO `User` (`id`, `username`, `password`, `salt`, `name`, `birthday`, `gender`, `email`, `phone`, `coin`, `type`, `avatar`, `created_at`, `updated_at`, `deleted_at`) VALUES
(1, 'eNViDAT', '$2a$04$YKNp5sLdAF.24i3533AB6u0t53gEh8Unl6kV9GAvPhhFgnEl0ZpgK', '2022-12-01 17:53:50.2323002 +0700 +07 m=+93.342681801', 'Nguyễn Văn Đạt', '2022-11-09', 1, 'ngvidat@gmail.com', '0775702201', 0, 'ADMIN', 'https://4xucy2kyby51ggkud2tadg3d-wpengine.netdna-ssl.com/wp-content/uploads/sites/37/2017/02/IAFOR-Blank-Avatar-Image.jpg', '2022-12-01 10:53:50', '2022-12-01 10:53:50', NULL),
(22, 'khanhsd0901', '$2a$04$2WU19Up7bFqLX7qZll.J7OHUFO5Q7LQ1Z.8rilPam/mRY0f4RUDi6', '2022-12-10 12:13:49.124386 +0700 +07 m=+189.306835601', 'Le Khan1h', '2001-01-01', 1, 'khanhsd091101@gmail.com', '09459581952', 0, 'BUYER', 'https://4xucy2kyby51ggkud2tadg3d-wpengine.netdna-ssl.com/wp-content/uploads/sites/37/2017/02/IAFOR-Blank-Avatar-Image.jpg', '2022-12-10 05:13:49', '2022-12-10 05:13:49', NULL),
(23, 'khanh', '$2a$04$GwL630qoYRHNeJAxKw6BuedX36jXY.kPj.z7P2aPvujpPzzuLavBy', '2022-12-10 12:16:37.0707576 +0700 +07 m=+357.253207201', 'Le Khan2h', '2001-01-01', 1, 'khanhsd022901@gmail.com', '09459589522', 0, 'BUYER', 'https://4xucy2kyby51ggkud2tadg3d-wpengine.netdna-ssl.com/wp-content/uploads/sites/37/2017/02/IAFOR-Blank-Avatar-Image.jpg', '2022-12-10 05:16:37', '2022-12-10 05:16:37', NULL),
(24, '11122222', '$2a$04$R4ksrJDAeuCjGJImmCdj..Gn0zdcT2KT36fjpUs2MNtW7wfY0yYvK', '2022-12-10 12:34:21.2439553 +0700 +07 m=+1421.426404901', 'Le Khan1h', '2001-01-01', 1, 'khanhsd0291101@gmail.com', '094592581952', 0, 'BUYER', 'https://4xucy2kyby51ggkud2tadg3d-wpengine.netdna-ssl.com/wp-content/uploads/sites/37/2017/02/IAFOR-Blank-Avatar-Image.jpg', '2022-12-10 05:34:21', '2022-12-10 05:34:21', NULL),
(25, 'test111', '$2a$04$B4Mf8LuI1oRnKhTrgz5AOuW3C1rIwORq8oj83VUMSBUOob8CxgcB.', '2022-12-10 13:32:18.6492071 +0700 +07 m=+4898.831656701', 'ssad', '2001-01-01', 1, 'sdsdawda', '1112', 0, 'BUYER', 'https://4xucy2kyby51ggkud2tadg3d-wpengine.netdna-ssl.com/wp-content/uploads/sites/37/2017/02/IAFOR-Blank-Avatar-Image.jpg', '2022-12-10 06:32:19', '2022-12-10 06:32:19', NULL),
(26, 'eNViDsT', '$2a$04$61bOtopdvZsVJInIkhYyUeny84.OuWOwFDzvYhz/QGZl.xYZXuzQi', '2022-12-21 09:33:58.6435133 +0700 +07 m=+42.628213001', 'Nguyễn Văn Đạt', '2022-11-09', 1, 'abcc@gaisl.com', '077502222301', 0, 'ADMIN', 'https://4xucy2kyby51ggkud2tadg3d-wpengine.netdna-ssl.com/wp-content/uploads/sites/37/2017/02/IAFOR-Blank-Avatar-Image.jpg', '2022-12-21 02:33:59', '2022-12-21 02:33:59', '2023-05-01 15:39:32'),
(27, 'eNViDXT', '$2a$04$94zB9qSUcr5.zZdE8M8jAueWtLAKJ3DOPxgmfs005ihVW/sfpK94O', '2022-12-21 09:34:32.6491548 +0700 +07 m=+76.633854501', 'Nguyễn Văn Đạt', '2022-11-09', 1, 'abcc@gsl.com', '0775001', 0, 'ADMIN', 'https://4xucy2kyby51ggkud2tadg3d-wpengine.netdna-ssl.com/wp-content/uploads/sites/37/2017/02/IAFOR-Blank-Avatar-Image.jpg', '2022-12-21 02:34:33', '2022-12-21 02:34:33', NULL),
(29, 'eNViDAXT', '$2a$04$YpqZwVNxenNxF0NMw5mq/.ww450ThAjL7E7XTz06MPGndzw0/JBc.', '2023-02-12 20:16:10.5689094 +0700 +07 m=+19.819286801', 'Nguyễn Văn Đạt', '2022-11-09', 1, 'abcc@gaiasl.com', '077502226301', 0, 'ADMIN', NULL, '2023-02-12 13:16:11', '2023-02-12 13:16:11', NULL),
(30, 'eNViDsT2', '$2a$04$msFlZO5jyQfAfx.2vLkzg.PIl6CH/jK9p8v9vfyq3WHdtvRlcIVUO', '2023-02-16 21:45:29.4101073 +0700 +07 m=+21.469918501', 'Nguyễn Văn Đạt', '2022-11-09', 1, 'abcc@gaisl2.com', '077502222302', 0, 'ADMIN', NULL, '2023-02-16 14:45:29', '2023-02-16 14:45:29', NULL);

--
-- Indexes for dumped tables
--

--
-- Indexes for table `Address`
--
ALTER TABLE `Address`
  ADD PRIMARY KEY (`id`),
  ADD KEY `user_id` (`user_id`),
  ADD KEY `province_code` (`province_id`),
  ADD KEY `ward_code` (`ward_code`),
  ADD KEY `district_code` (`district_id`);

--
-- Indexes for table `AdministrativeRegions`
--
ALTER TABLE `AdministrativeRegions`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `AdministrativeUnits`
--
ALTER TABLE `AdministrativeUnits`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `Banner`
--
ALTER TABLE `Banner`
  ADD PRIMARY KEY (`id`),
  ADD KEY `user_id` (`user_id`);

--
-- Indexes for table `BannerDetail`
--
ALTER TABLE `BannerDetail`
  ADD KEY `BannerDetail_ibfk_1` (`banner_id`),
  ADD KEY `BannerDetail_ibfk_2` (`product_id`);

--
-- Indexes for table `Cart`
--
ALTER TABLE `Cart`
  ADD PRIMARY KEY (`id`),
  ADD KEY `Cart_ibfk_1` (`provider_id`),
  ADD KEY `Cart_ibfk_2` (`user_id`);

--
-- Indexes for table `CartItem`
--
ALTER TABLE `CartItem`
  ADD PRIMARY KEY (`id`),
  ADD KEY `product_id` (`product_id`),
  ADD KEY `cart_id` (`user_id`),
  ADD KEY `CartItem_ibfk_3` (`cart_id`),
  ADD KEY `product_option_id` (`product_option_id`);

--
-- Indexes for table `Category`
--
ALTER TABLE `Category`
  ADD PRIMARY KEY (`id`),
  ADD KEY `category_parent_id` (`category_parent_id`);

--
-- Indexes for table `ChatRoom`
--
ALTER TABLE `ChatRoom`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `Comment`
--
ALTER TABLE `Comment`
  ADD PRIMARY KEY (`id`),
  ADD KEY `product_id` (`product_id`),
  ADD KEY `user_id` (`user_id`);

--
-- Indexes for table `CommentMedia`
--
ALTER TABLE `CommentMedia`
  ADD PRIMARY KEY (`id`),
  ADD KEY `comment_id` (`comment_id`);

--
-- Indexes for table `Email`
--
ALTER TABLE `Email`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `Favorite`
--
ALTER TABLE `Favorite`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `Message`
--
ALTER TABLE `Message`
  ADD PRIMARY KEY (`id`),
  ADD KEY `Message_ibfk_1` (`from_user_id`),
  ADD KEY `chat_room_id` (`chat_room_id`),
  ADD KEY `to_user_id` (`to_user_id`);

--
-- Indexes for table `Notification`
--
ALTER TABLE `Notification`
  ADD PRIMARY KEY (`id`),
  ADD KEY `user_id` (`user_id`);

--
-- Indexes for table `Order`
--
ALTER TABLE `Order`
  ADD PRIMARY KEY (`id`),
  ADD KEY `user_id` (`user_id`),
  ADD KEY `provider_id` (`provider_id`);

--
-- Indexes for table `OrderItem`
--
ALTER TABLE `OrderItem`
  ADD PRIMARY KEY (`id`),
  ADD KEY `order_id` (`order_id`),
  ADD KEY `product_id` (`product_id`),
  ADD KEY `product_option_id` (`product_option_id`),
  ADD KEY `provider_id` (`provider_id`);

--
-- Indexes for table `Product`
--
ALTER TABLE `Product`
  ADD PRIMARY KEY (`id`),
  ADD KEY `category_id` (`category_id`),
  ADD KEY `discount_id` (`discount`),
  ADD KEY `user_id` (`user_id`),
  ADD KEY `provider_id` (`provider_id`);
ALTER TABLE `Product` ADD FULLTEXT KEY `name` (`name`);

--
-- Indexes for table `ProductDescriptions`
--
ALTER TABLE `ProductDescriptions`
  ADD PRIMARY KEY (`id`),
  ADD KEY `product_id` (`product_id`);

--
-- Indexes for table `ProductMedia`
--
ALTER TABLE `ProductMedia`
  ADD PRIMARY KEY (`id`),
  ADD KEY `product_id` (`product_id`);

--
-- Indexes for table `ProductOption`
--
ALTER TABLE `ProductOption`
  ADD PRIMARY KEY (`id`),
  ADD KEY `ProductOption_ibfk_1` (`product_id`),
  ADD KEY `specification_id` (`specification_id`);

--
-- Indexes for table `ProductSpecification`
--
ALTER TABLE `ProductSpecification`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `product_id_2` (`product_id`),
  ADD KEY `product_id` (`product_id`);

--
-- Indexes for table `Provider`
--
ALTER TABLE `Provider`
  ADD PRIMARY KEY (`id`),
  ADD KEY `user_id` (`user_id`);

--
-- Indexes for table `User`
--
ALTER TABLE `User`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `email` (`email`),
  ADD UNIQUE KEY `username` (`username`),
  ADD UNIQUE KEY `phone` (`phone`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `Address`
--
ALTER TABLE `Address`
  MODIFY `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=12;

--
-- AUTO_INCREMENT for table `Banner`
--
ALTER TABLE `Banner`
  MODIFY `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=7;

--
-- AUTO_INCREMENT for table `Cart`
--
ALTER TABLE `Cart`
  MODIFY `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=9;

--
-- AUTO_INCREMENT for table `CartItem`
--
ALTER TABLE `CartItem`
  MODIFY `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=28;

--
-- AUTO_INCREMENT for table `Category`
--
ALTER TABLE `Category`
  MODIFY `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=37;

--
-- AUTO_INCREMENT for table `ChatRoom`
--
ALTER TABLE `ChatRoom`
  MODIFY `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=9;

--
-- AUTO_INCREMENT for table `Comment`
--
ALTER TABLE `Comment`
  MODIFY `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=26;

--
-- AUTO_INCREMENT for table `CommentMedia`
--
ALTER TABLE `CommentMedia`
  MODIFY `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=11;

--
-- AUTO_INCREMENT for table `Email`
--
ALTER TABLE `Email`
  MODIFY `id` bigint NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=4;

--
-- AUTO_INCREMENT for table `Favorite`
--
ALTER TABLE `Favorite`
  MODIFY `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=2;

--
-- AUTO_INCREMENT for table `Message`
--
ALTER TABLE `Message`
  MODIFY `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=91;

--
-- AUTO_INCREMENT for table `Notification`
--
ALTER TABLE `Notification`
  MODIFY `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=83;

--
-- AUTO_INCREMENT for table `Order`
--
ALTER TABLE `Order`
  MODIFY `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=123;

--
-- AUTO_INCREMENT for table `OrderItem`
--
ALTER TABLE `OrderItem`
  MODIFY `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=12;

--
-- AUTO_INCREMENT for table `Product`
--
ALTER TABLE `Product`
  MODIFY `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=42;

--
-- AUTO_INCREMENT for table `ProductDescriptions`
--
ALTER TABLE `ProductDescriptions`
  MODIFY `id` bigint NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=22;

--
-- AUTO_INCREMENT for table `ProductMedia`
--
ALTER TABLE `ProductMedia`
  MODIFY `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=34;

--
-- AUTO_INCREMENT for table `ProductOption`
--
ALTER TABLE `ProductOption`
  MODIFY `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=114;

--
-- AUTO_INCREMENT for table `ProductSpecification`
--
ALTER TABLE `ProductSpecification`
  MODIFY `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=49;

--
-- AUTO_INCREMENT for table `Provider`
--
ALTER TABLE `Provider`
  MODIFY `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=14;

--
-- AUTO_INCREMENT for table `User`
--
ALTER TABLE `User`
  MODIFY `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=32;

--
-- Constraints for dumped tables
--

--
-- Constraints for table `Address`
--
ALTER TABLE `Address`
  ADD CONSTRAINT `Address_ibfk_3` FOREIGN KEY (`user_id`) REFERENCES `User` (`id`) ON UPDATE CASCADE;

--
-- Constraints for table `Banner`
--
ALTER TABLE `Banner`
  ADD CONSTRAINT `Banner_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `User` (`id`) ON UPDATE CASCADE;

--
-- Constraints for table `BannerDetail`
--
ALTER TABLE `BannerDetail`
  ADD CONSTRAINT `BannerDetail_ibfk_1` FOREIGN KEY (`banner_id`) REFERENCES `Banner` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
  ADD CONSTRAINT `BannerDetail_ibfk_2` FOREIGN KEY (`product_id`) REFERENCES `Product` (`id`) ON DELETE CASCADE ON UPDATE CASCADE;

--
-- Constraints for table `Cart`
--
ALTER TABLE `Cart`
  ADD CONSTRAINT `Cart_ibfk_1` FOREIGN KEY (`provider_id`) REFERENCES `Provider` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
  ADD CONSTRAINT `Cart_ibfk_2` FOREIGN KEY (`user_id`) REFERENCES `User` (`id`) ON DELETE CASCADE ON UPDATE CASCADE;

--
-- Constraints for table `CartItem`
--
ALTER TABLE `CartItem`
  ADD CONSTRAINT `CartItem_ibfk_1` FOREIGN KEY (`product_id`) REFERENCES `Product` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
  ADD CONSTRAINT `CartItem_ibfk_2` FOREIGN KEY (`user_id`) REFERENCES `User` (`id`),
  ADD CONSTRAINT `CartItem_ibfk_3` FOREIGN KEY (`cart_id`) REFERENCES `Cart` (`id`) ON DELETE CASCADE,
  ADD CONSTRAINT `CartItem_ibfk_4` FOREIGN KEY (`product_option_id`) REFERENCES `ProductOption` (`id`) ON DELETE CASCADE ON UPDATE CASCADE;

--
-- Constraints for table `Category`
--
ALTER TABLE `Category`
  ADD CONSTRAINT `Category_ibfk_1` FOREIGN KEY (`category_parent_id`) REFERENCES `Category` (`id`);

--
-- Constraints for table `Comment`
--
ALTER TABLE `Comment`
  ADD CONSTRAINT `Comment_ibfk_1` FOREIGN KEY (`product_id`) REFERENCES `Product` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
  ADD CONSTRAINT `Comment_ibfk_2` FOREIGN KEY (`user_id`) REFERENCES `User` (`id`) ON UPDATE CASCADE;

--
-- Constraints for table `CommentMedia`
--
ALTER TABLE `CommentMedia`
  ADD CONSTRAINT `CommentMedia_ibfk_1` FOREIGN KEY (`comment_id`) REFERENCES `Comment` (`id`) ON DELETE CASCADE ON UPDATE CASCADE;

--
-- Constraints for table `Message`
--
ALTER TABLE `Message`
  ADD CONSTRAINT `Message_ibfk_1` FOREIGN KEY (`from_user_id`) REFERENCES `User` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  ADD CONSTRAINT `Message_ibfk_3` FOREIGN KEY (`chat_room_id`) REFERENCES `ChatRoom` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  ADD CONSTRAINT `Message_ibfk_4` FOREIGN KEY (`to_user_id`) REFERENCES `User` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT;

--
-- Constraints for table `Notification`
--
ALTER TABLE `Notification`
  ADD CONSTRAINT `Notification_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `User` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT;

--
-- Constraints for table `Order`
--
ALTER TABLE `Order`
  ADD CONSTRAINT `Order_ibfk_3` FOREIGN KEY (`user_id`) REFERENCES `User` (`id`),
  ADD CONSTRAINT `Order_ibfk_4` FOREIGN KEY (`provider_id`) REFERENCES `Provider` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT;

--
-- Constraints for table `OrderItem`
--
ALTER TABLE `OrderItem`
  ADD CONSTRAINT `OrderItem_ibfk_2` FOREIGN KEY (`order_id`) REFERENCES `Order` (`id`) ON UPDATE CASCADE,
  ADD CONSTRAINT `OrderItem_ibfk_4` FOREIGN KEY (`product_id`) REFERENCES `Product` (`id`) ON UPDATE CASCADE,
  ADD CONSTRAINT `OrderItem_ibfk_5` FOREIGN KEY (`product_option_id`) REFERENCES `ProductOption` (`id`),
  ADD CONSTRAINT `OrderItem_ibfk_6` FOREIGN KEY (`provider_id`) REFERENCES `Provider` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT;

--
-- Constraints for table `Product`
--
ALTER TABLE `Product`
  ADD CONSTRAINT `Product_ibfk_1` FOREIGN KEY (`category_id`) REFERENCES `Category` (`id`) ON UPDATE CASCADE,
  ADD CONSTRAINT `Product_ibfk_4` FOREIGN KEY (`user_id`) REFERENCES `User` (`id`) ON UPDATE CASCADE,
  ADD CONSTRAINT `Product_ibfk_5` FOREIGN KEY (`provider_id`) REFERENCES `Provider` (`id`);

--
-- Constraints for table `ProductDescriptions`
--
ALTER TABLE `ProductDescriptions`
  ADD CONSTRAINT `ProductDescriptions_ibfk_1` FOREIGN KEY (`product_id`) REFERENCES `Product` (`id`) ON DELETE CASCADE ON UPDATE RESTRICT;

--
-- Constraints for table `ProductMedia`
--
ALTER TABLE `ProductMedia`
  ADD CONSTRAINT `ProductMedia_ibfk_1` FOREIGN KEY (`product_id`) REFERENCES `Product` (`id`) ON DELETE CASCADE ON UPDATE CASCADE;

--
-- Constraints for table `ProductOption`
--
ALTER TABLE `ProductOption`
  ADD CONSTRAINT `ProductOption_ibfk_1` FOREIGN KEY (`product_id`) REFERENCES `Product` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
  ADD CONSTRAINT `ProductOption_ibfk_2` FOREIGN KEY (`specification_id`) REFERENCES `ProductSpecification` (`id`) ON DELETE CASCADE ON UPDATE CASCADE;

--
-- Constraints for table `ProductSpecification`
--
ALTER TABLE `ProductSpecification`
  ADD CONSTRAINT `ProductSpecification_ibfk_1` FOREIGN KEY (`product_id`) REFERENCES `Product` (`id`) ON DELETE CASCADE ON UPDATE CASCADE;

--
-- Constraints for table `Provider`
--
ALTER TABLE `Provider`
  ADD CONSTRAINT `Provider_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `User` (`id`);
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
