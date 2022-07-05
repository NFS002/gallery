CREATE TABLE exhibition_dates
(
    id serial primary key,
    exhibition_id      integer                     NOT NULL,
    start_time         timestamp without time zone NOT NULL,
    end_time           timestamp without time zone NOT NULL,
    created_at         timestamp without time zone NOT NULL,
    updated_at         timestamp without time zone NOT NULL
);