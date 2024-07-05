DROP DATABASE IF EXISTS smartri;
CREATE DATABASE smartri;

\c smartri;

create sequence accounts_seq;
-- drop table if exists users;
create table accounts
(
    id                integer default nextval('accounts_seq') not null
        primary key,
    login             varchar(20)                           not null,
    nickname          varchar(16)                           not null,
    password          varchar(100)                          not null,
    email             varchar(254)                          not null,
    registration_date date                                  not null,
    is_verified       boolean default false                 not null
);

-- drop table if exists verification_codes;
create table verification_codes
(
    id              serial primary key,
    user_id         integer,
    code            varchar(8),
    expiration_time timestamp
);

-- drop table if exists sessions;
create table sessions
(
    id              serial primary key,
    access_token    varchar(256) not null,
    refresh_token   varchar(256) not null,
    user_id         int          not null,
    device_identity varchar(128),
    expire_at       date         not null
);

insert into accounts(id, login, nickname, password, email, registration_date, is_verified)
values (1, 'admin', 'крутой_админ', '$2a$10$DNMPT2CvlYTS1/CRZnT7qO/jlWgA0v99EiIA5Fg.n2AM4H0Zm2Fnq', 'admin@ya.ru',
        current_date, true);
alter sequence accounts_seq restart with 2;