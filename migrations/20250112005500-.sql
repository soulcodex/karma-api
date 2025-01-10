-- +migrate Up
CREATE TABLE IF NOT EXISTS `karma_assignee` (
    `id`          VARCHAR(50)     NOT NULL,
    `username`    VARCHAR(255)    NOT NULL,
    `assigner`    VARCHAR(255)    NOT NULL,
    `count`       SMALLINT        NOT NULL,
    `created_at`  DATETIME(3)     NOT NULL,
    PRIMARY KEY (`id`),
    KEY `karma_assignee_username_idx` (`username`),
    KEY `karma_assignee_assigner_idx` (`assigner`),
    KEY `karma_assignee_count_idx` (`count`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
-- +migrate Down
DROP TABLE IF EXISTS `karma_assignee`;
