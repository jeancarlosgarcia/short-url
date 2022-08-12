CREATE TABLE IF NOT EXISTS short_urls
(
    short      varchar(100) not null constraint urls_pk primary key,
    original   varchar(255),
    created_at timestamp default CURRENT_TIMESTAMP
);

create unique index urls_short_uindex on short_urls (short);