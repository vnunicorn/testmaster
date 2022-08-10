CREATE TABLE `testmaster`.`users` (
  `id` bigint PRIMARY KEY AUTO_INCREMENT,
  `username` varchar(255),
  `mobile` varchar(15),
  `email` varchar(255),
  `password` varchar(255) NOT NULL,
  `first_name` varchar(255),
  `middle_name` varchar(255),
  `last_name` varchar(255),
  `last_login` timestamp,
  `created_at` timestamp DEFAULT (now()),
  `updated_at` timestamp DEFAULT (now())
);