set search_path to public;

-- SELECT 'CREATE DATABASE auth' WHERE NOT EXISTS (SELECT FROM pg_database WHERE datname = 'auth');

CREATE DATABASE auth;
-- drop table if exists users;
create table users
(
    id                serial
        primary key,
    login             varchar(20)           not null,
    nickname          varchar(16)           not null,
    password          varchar(100)          not null,
    email             varchar(254)          not null,
    registration_date date                  not null,
    is_verified       boolean default false not null
);
--
-- drop table if exists verification_codes;
-- create table verification_codes
-- (
--     user_id           integer,
--     verification_code varchar(8),
--     expiration_time   timestamp
-- );
--
-- create function add_verification_code(id integer, code character, e_time timestamp without time zone) returns void
--     language plpgsql
-- as
-- $$
-- declare
--     count int;
-- begin
--     select count(*) from verification_codes where user_id = id into count;
--     if count > 0 then
--         update verification_codes
--         set verification_code = code and expiration_time = e_time
--         where user_id = id;
--     else
--         insert into verification_codes (user_id, verification_code, expiration_time)
--         values (id, code, e_time);
--     end if;
-- end;
-- $$;
--
-- alter function add_verification_code(integer, char, timestamp) owner to postgres;
--
-- drop table if exists sessions;
-- create table sessions(
--     access_token varchar(128) primary key ,
--     refresh_token varchar(128),
--     user_id int,
--     device_description varchar(128),
--     expire_at date
-- )