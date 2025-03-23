-- +migrate Up
CREATE TABLE stories (
    id BIGINT AUTO_INCREMENT PRIMARY KEY NOT NULL,
    user_id INT NOT NULL,
    title VARCHAR(255),
    subtitle VARCHAR(255),  -- Perbaiki penulisan
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

CREATE TABLE story_elements (
    id BIGINT AUTO_INCREMENT PRIMARY KEY NOT NULL,
    story_id BIGINT NOT NULL,
    type_data VARCHAR(10) NOT NULL CHECK (type_data IN ('text', 'image')), -- Perbaiki CHECK constraint
    content TEXT NOT NULL,
    position INT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (story_id) REFERENCES stories(id) ON DELETE CASCADE
);

-- +migrate Down
DROP TABLE IF EXISTS story_elements;
DROP TABLE IF EXISTS stories;
