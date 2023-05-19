--@block
CREATE TABLE IF NOT EXISTS `user` (
    `national_id` VARCHAR(255) NOT NULL UNIQUE PRIMARY KEY,
    `first_name` VARCHAR(255) NOT NULL,
    `last_name` VARCHAR(255) NOT NULL,
    `language`  VARCHAR(255) NOT NULL,
    `birth_date` DATE NOT NULL,
    `email` VARCHAR(255) NOT NULL,
    `phone` VARCHAR(255) NOT NULL,
    `gender` VARCHAR(255) NOT NULL
    ) ENGINE=InnoDB DEFAULT CHARSET=utf8;

--@block

CREATE TABLE IF NOT EXISTS `patient` (
    `patient_id` INT NOT NULL UNIQUE PRIMARY KEY,
    `national_id` VARCHAR(255) NOT NULL UNIQUE REFERENCES `user`(`national_id`),
    `weight` FLOAT NOT NULL,
    `height` FLOAT NOT NULL,
    `blood_type` VARCHAR(255) NOT NULL,
    `allergies` TEXT NOT NULL,
    `conditions` TEXT NOT NULL
)ENGINE=InnoDB DEFAULT CHARSET=utf8;


--@block
CREATE TABLE IF NOT EXISTS `doctor` (
    `doctor_id` INT NOT NULL UNIQUE PRIMARY KEY,
    `national_id` VARCHAR(255) NOT NULL UNIQUE REFERENCES `user`(`national_id`),
    `specialization` VARCHAR(255) NOT NULL
)ENGINE=InnoDB DEFAULT CHARSET=utf8;

--@block
CREATE TABLE IF NOT EXISTS `connection` (
    `chat_id` BIGINT NOT NULL UNIQUE PRIMARY KEY,
    `national_id` VARCHAR(255) NOT NULL REFERENCES `user`(`national_id`),
    `question` VARCHAR(255) NULL    
)ENGINE=InnoDB DEFAULT CHARSET=utf8;

--@block
INSERT INTO `user` (`national_id`, `first_name`, `last_name`,  `language`, `birth_date`, `email`, `phone`, `gender`) VALUES
('123456789', 'John', 'Doe',  'en', '2000-01-01', 'XXXXXXXXXXXX', '0123456789', 'male');


--@block
INSERT INTO `user`(`national_id`, `first_name`, `last_name`,  `language`, `birth_date`, `email`, `phone`, `gender`) VALUES
('', '', '',  '', '2000-01-01', '', '', '');
--@block
DROP TABLE `patient`;

--@block 
DROP TABLE `connection`;

--@block 
DROP TABLE `doctor`;

--@block 
DROP TABLE `user`;