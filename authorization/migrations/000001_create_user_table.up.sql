set search_path to public;

drop table users;
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

create table verification_codes
(
    user_id           integer,
    verification_code varchar(8),
    expiration_time   timestamp
);

create function add_verification_code(id integer, code character, e_time timestamp without time zone) returns void
    language plpgsql
as
$$
declare
    count int;
begin
    select count(*) from verification_codes where user_id = id into count;
    if count > 0 then
        update verification_codes
        set verification_code = code and expiration_time = e_time
        where user_id = id;
    else
        insert into verification_codes (user_id, verification_code, expiration_time)
        values (id, code, e_time);
    end if;
end;
$$;

alter function add_verification_code(integer, char, timestamp) owner to postgres;

