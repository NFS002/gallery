CREATE TABLE addresses
(
    id serial primary key,
    number     integer                     NOT NULL,
    address_1  character varying(255)      NOT NULL,
    address_2  character varying(255)      NOT NULL,
    city       character varying(50)       NOT NULL,
    postcode   character varying(10)       NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);