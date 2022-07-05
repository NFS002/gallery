CREATE TABLE public.artists
(
    id serial primary key,
    first_name    character varying(25)       NOT NULL,
    last_name     character varying(25)       NOT NULL,
    email_address character varying(150)      NOT NULL,
    phone_number  character varying(11)       NOT NULL,
    address_id    integer                     NOT NULL,
    created_at    timestamp without time zone NOT NULL,
    updated_at    timestamp without time zone NOT NULL
)