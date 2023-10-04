CREATE TABLE cards (
    id bigserial not null primary key,
    word varchar not null unique,
    translate varchar not null
);