DROP TABLE IF EXISTS `posts`;

create table IF not exists `posts`
(
 `id`               INT(20) NOT NULL AUTO_INCREMENT,
 `name`             VARCHAR(50) NOT NULL,
 `text`             VARCHAR(50) NOT NULL,
    PRIMARY KEY (`id`)
);

INSERT INTO posts (name,text) VALUES
    ('田中',"こんにちは"),
    ('高橋',"ポテト食べたい"),
    ('渡辺',"ディズニー行きたい");