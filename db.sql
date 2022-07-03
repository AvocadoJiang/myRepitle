CREATE DATABASE `myReptile` /*!40100 DEFAULT CHARACTER SET utf8 */;

-- myReptile.job_list definition

CREATE TABLE `job_list` (
    `jobid` varchar(100) NOT NULL,
    `coid` varchar(100) DEFAULT NULL,
    `job_href` varchar(300) DEFAULT NULL,
    `job_name` varchar(100) DEFAULT NULL,
    `job_title` varchar(100) DEFAULT NULL,
    `company_href` varchar(100) DEFAULT NULL,
    `company_name` varchar(100) DEFAULT NULL,
    `providesalary_text` varchar(100) DEFAULT NULL,
    `workarea` varchar(100) DEFAULT NULL,
    `workarea_text` varchar(100) DEFAULT NULL,
    `workyear` varchar(100) DEFAULT NULL,
    `jobwelf` varchar(100) DEFAULT NULL,
    UNIQUE KEY `job_list_UN` (`jobid`,`coid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;