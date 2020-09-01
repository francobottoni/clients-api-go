CREATE TABLE client
(
    `id`             int(11) not null auto_increment,
    `dni`            int(11) not null ,
    `name`           varchar(150),
    `last_name`      varchar(150),
    `country_origin` varchar(150),
    `created_at`     DATETIME DEFAULT CURRENT_TIMESTAMP,
    `updated_at`     DATETIME ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`)
) engine = InnoDB
  DEFAULT charset = utf8;