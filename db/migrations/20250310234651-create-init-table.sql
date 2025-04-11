-- +migrate Up
CREATE TABLE categories (
    id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
    name VARCHAR(100) NOT NULL UNIQUE,
    slug VARCHAR(100) NOT NULL UNIQUE,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (id)
);
CREATE TABLE story_categories (
    story_id VARCHAR(24) NOT NULL,
    category_id BIGINT UNSIGNED NOT NULL,
    PRIMARY KEY (story_id, category_id),
    FOREIGN KEY (category_id) REFERENCES categories(id) ON DELETE CASCADE
);

-- +migrate Down
DROP TABLE IF EXISTS categories;
DROP TABLE IF EXISTS story_categories;