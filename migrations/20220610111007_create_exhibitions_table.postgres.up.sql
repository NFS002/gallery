CREATE TABLE exhibitions
(
    id serial primary key,
    title         varchar(100) NOT NULL,
    host_id       integer                     NOT NULL,
    artist_id     integer                     NOT NULL,
    notes         text                        NOT NULL,
    created_at    timestamp without time zone NOT NULL,
    updated_at    timestamp without time zone NOT NULL
);