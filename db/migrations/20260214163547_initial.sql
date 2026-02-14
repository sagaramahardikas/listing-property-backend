-- Create "listings" table
CREATE TABLE `listings` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `title` varchar(255) NOT NULL,
  `price` int NOT NULL,
  `description` text NULL,
  `facilities` text NULL,
  `banner` text NULL,
  `images` text NULL,
  `terms_and_conditions` text NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  PRIMARY KEY (`id`)
) CHARSET utf8mb4 COLLATE utf8mb4_unicode_ci;
