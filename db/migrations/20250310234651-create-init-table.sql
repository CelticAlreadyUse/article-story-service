
-- +migrate Up
CREATE TABLE stories(
    id serial primary key,
    user_id INT NOT NULL,
    created_at timestamp not null default now(),
    expires_at timestamp not null
)

CREATE TABLE story_elements(
    id serial primary key,
    story_id INT references stories(id) on delete cascade,
    type varchar(10) not null check (type in ("text","image")),
    content text not null,
    position int not null,
    created_at timestamp not null default now()
)
-- +migrate Down
