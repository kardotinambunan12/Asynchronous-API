-- phpMyAdmin SQL Dump
-- version 5.2.1
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Waktu pembuatan: 03 Mar 2025 pada 07.50
-- Versi server: 10.4.32-MariaDB
-- Versi PHP: 8.0.30

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `spe_test`
--

-- --------------------------------------------------------

--
-- Struktur dari tabel `customer`
--

CREATE TABLE `customer` (
  `customer_id` varchar(100) NOT NULL,
  `customer_pan` varchar(100) NOT NULL,
  `customer_name` varchar(200) DEFAULT NULL,
  `tgl_rekam` varchar(100) DEFAULT NULL,
  `petugas_rekam` varchar(100) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data untuk tabel `customer`
--

INSERT INTO `customer` (`customer_id`, `customer_pan`, `customer_name`, `tgl_rekam`, `petugas_rekam`) VALUES
('7c44722f-d3a2-4b01-b45e-3be8b9fbd2af', '8327732737474787324', 'Sample', '2025-03-01 16:53:02', 'ADMIN'),
('aefa7b68-dc02-4436-81a2-123da02405b4', 'JhonChenamail@mail.com', 'Smack Down', '2025-03-01 13:17:06', 'Toba');

-- --------------------------------------------------------

--
-- Struktur dari tabel `merchant`
--

CREATE TABLE `merchant` (
  `merchant_id` varchar(36) NOT NULL,
  `merchant_name` varchar(200) DEFAULT NULL,
  `merchant_city` varchar(100) NOT NULL,
  `tgl_rekam` varchar(50) DEFAULT NULL,
  `petugas_rekam` varchar(100) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data untuk tabel `merchant`
--

INSERT INTO `merchant` (`merchant_id`, `merchant_name`, `merchant_city`, `tgl_rekam`, `petugas_rekam`) VALUES
('cc655cf4-f98b-44a6-be4c-bd258c087551', 'ABC', 'Solo', '2025-03-01 14:01:14', 'AD');

-- --------------------------------------------------------

--
-- Struktur dari tabel `transactions`
--

CREATE TABLE `transactions` (
  `request_id` varchar(32) NOT NULL,
  `customer_pan` varchar(200) DEFAULT NULL,
  `amount` varchar(100) NOT NULL,
  `transaction_datetime` varchar(50) DEFAULT NULL,
  `rrn` varchar(100) DEFAULT NULL,
  `bill_number` varchar(100) DEFAULT NULL,
  `customer_name` varchar(100) DEFAULT NULL,
  `merchant_id` varchar(100) DEFAULT NULL,
  `merchant_name` varchar(100) DEFAULT NULL,
  `merchant_city` varchar(100) DEFAULT NULL,
  `currency_code` varchar(100) DEFAULT NULL,
  `payment_status` varchar(100) DEFAULT NULL,
  `payment_description` varchar(100) DEFAULT NULL,
  `tgl_rekam` varchar(50) DEFAULT NULL,
  `petugas_rekam` varchar(100) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data untuk tabel `transactions`
--

INSERT INTO `transactions` (`request_id`, `customer_pan`, `amount`, `transaction_datetime`, `rrn`, `bill_number`, `customer_name`, `merchant_id`, `merchant_name`, `merchant_city`, `currency_code`, `payment_status`, `payment_description`, `tgl_rekam`, `petugas_rekam`) VALUES
('aefa7b68dc02443681a2123da02405b2', '8327732737474787324', '150000.12', '2024-03-01T14:30:00Z', '987654321012', 'INV00000002', 'John Doe', 'cc655cf4-f98b-44a6-be4c-bd258c087551', 'Toko Elektronik ABC', 'Jakarta', 'IDR', 'Completed', 'Pembayaran berhasil', '2025-03-03 13:27:52', 'ADMIN'),
('aefa7b68dc02443681a2123da02405b4', '8327732737474787324', '150000', '2024-03-01T14:30:00Z', '987654321012', 'INV00000001', 'John Doe', 'cc655cf4-f98b-44a6-be4c-bd258c087551', 'Toko Elektronik ABC', 'Jakarta', 'IDR', 'Completed', 'Pembayaran berhasil', '2025-03-03 11:52:30', 'ADMIN');

--
-- Indexes for dumped tables
--

--
-- Indeks untuk tabel `customer`
--
ALTER TABLE `customer`
  ADD PRIMARY KEY (`customer_pan`);

--
-- Indeks untuk tabel `merchant`
--
ALTER TABLE `merchant`
  ADD PRIMARY KEY (`merchant_id`);

--
-- Indeks untuk tabel `transactions`
--
ALTER TABLE `transactions`
  ADD PRIMARY KEY (`request_id`),
  ADD KEY `fk_merchant` (`merchant_id`),
  ADD KEY `fk_customer` (`customer_pan`);

--
-- Ketidakleluasaan untuk tabel pelimpahan (Dumped Tables)
--

--
-- Ketidakleluasaan untuk tabel `transactions`
--
ALTER TABLE `transactions`
  ADD CONSTRAINT `fk_customer` FOREIGN KEY (`customer_pan`) REFERENCES `customer` (`customer_pan`),
  ADD CONSTRAINT `fk_merchant` FOREIGN KEY (`merchant_id`) REFERENCES `merchant` (`merchant_id`);
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
