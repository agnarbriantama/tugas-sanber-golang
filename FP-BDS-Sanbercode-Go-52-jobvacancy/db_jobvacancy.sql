-- phpMyAdmin SQL Dump
-- version 5.2.0
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Generation Time: Jan 05, 2024 at 12:42 PM
-- Server version: 10.4.25-MariaDB
-- PHP Version: 8.0.23

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `db_jobvacancy`
--

-- --------------------------------------------------------

--
-- Table structure for table `tb_apply`
--

CREATE TABLE `tb_apply` (
  `id_apply` int(11) NOT NULL,
  `id` int(11) NOT NULL,
  `id_user` int(11) NOT NULL,
  `status_lamaran` varchar(255) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `tb_apply`
--

INSERT INTO `tb_apply` (`id_apply`, `id`, `id_user`, `status_lamaran`) VALUES
(3, 4, 7, ''),
(4, 1, 5, ''),
(5, 5, 7, ''),
(7, 1, 7, 'pending'),
(11, 1, 7, 'pending');

-- --------------------------------------------------------

--
-- Table structure for table `tb_listjob`
--

CREATE TABLE `tb_listjob` (
  `id` int(11) NOT NULL,
  `title` varchar(255) NOT NULL,
  `company_name` varchar(255) NOT NULL,
  `company_desc` varchar(255) NOT NULL,
  `company_salary` int(11) NOT NULL,
  `company_status` int(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `tb_listjob`
--

INSERT INTO `tb_listjob` (`id`, `title`, `company_name`, `company_desc`, `company_salary`, `company_status`) VALUES
(1, 'Backend Dev', 'Sanbercode', 'kerja sebagai backend', 5000000, 1),
(2, 'Frontend Dev', 'Sanbercode', 'kerja sebagai frotend', 4500000, 2),
(3, 'Frontend Dev', 'Sanbercode', 'kerja sebagai frotend', 4500000, 2),
(4, 'Tampilan Admin', 'Sanbercode', 'kerja sebagai frotend', 4500000, 1),
(5, 'Coba', 'Sanbercode', '', 60000000, 1),
(6, 'WOW', 'codeee', 'ya kerja', 2147483647, 2);

-- --------------------------------------------------------

--
-- Table structure for table `users`
--

CREATE TABLE `users` (
  `id_user` int(11) NOT NULL,
  `email` varchar(255) NOT NULL,
  `password` varchar(255) NOT NULL,
  `username` varchar(255) NOT NULL,
  `role` varchar(255) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `users`
--

INSERT INTO `users` (`id_user`, `email`, `password`, `username`, `role`) VALUES
(4, '', '$2a$10$a7HQBtG2MUQCkO7NAegedu4BaY.z6TDfJfSDbEbpDFkqbe4aZW8ui', 'brian', ''),
(5, 'agnar@gmail.com', '$2a$10$VUMv/GBOOH2EWs1Ex1tNeeERAaElJhNhcThzIFJa4zllBRFx.1muS', 'agnar', 'admin'),
(7, 'brian@gmail.com', '$2a$10$LYRdjLEd2UKpzzY44ZBmRO6lbnA1k53NzuGTJUT201s1CW1/4ok.2', 'tama', 'guest'),
(9, 'briantama@gmail.com', '$2a$10$Z3uSpYxaVlWhqBQ4Furgce2W2zGRH4fUcJuAFRiWnUrhwSlK2SReu', 'agnar123', 'admin');

--
-- Indexes for dumped tables
--

--
-- Indexes for table `tb_apply`
--
ALTER TABLE `tb_apply`
  ADD PRIMARY KEY (`id_apply`);

--
-- Indexes for table `tb_listjob`
--
ALTER TABLE `tb_listjob`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `users`
--
ALTER TABLE `users`
  ADD PRIMARY KEY (`id_user`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `tb_apply`
--
ALTER TABLE `tb_apply`
  MODIFY `id_apply` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=12;

--
-- AUTO_INCREMENT for table `tb_listjob`
--
ALTER TABLE `tb_listjob`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=7;

--
-- AUTO_INCREMENT for table `users`
--
ALTER TABLE `users`
  MODIFY `id_user` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=10;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
