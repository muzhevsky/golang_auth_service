DROP DATABASE IF EXISTS smartri;
CREATE DATABASE smartri;

\c smartri;

create table user_data
(
    id      serial
        constraint user_data_pk
            primary key,
    account_id integer not null,
    age     integer not null,
    xp      integer not null,
    gender  char     not null
);

alter table user_data
    owner to postgres;

create table test_questions
(
    id   serial
        constraint questions_pk
            primary key,
    text varchar(256) not null
);

alter table test_questions
    owner to postgres;

create table test_answers
(
    id          serial
        constraint answers_pk
            primary key,
    question_id smallint     not null,
    text        varchar(128) not null
);

alter table test_answers
    owner to postgres;

create table user_answers
(
    id          serial
        constraint user_answers_pk
            primary key,
    account_id     integer not null,
    question_id integer not null
        constraint user_answers_questions_fk
            references test_questions,
    answer_id   integer not null
        constraint user_answers_answers_fk
            references test_answers
);

alter table user_answers
    owner to postgres;

create table skills
(
    id    serial
        constraint skills_pk
            primary key,
    title varchar(32)
);

alter table skills
    owner to postgres;

create table test_answers_values
(
    id        serial
        constraint test_answers_values_pk
            primary key,
    answer_id integer not null
        constraint test_answers_values_answers_fk
            references test_answers,
    skill_id  integer not null
        constraint test_answers_values_skills_fk
            references skills,
    points    integer not null
);

alter table test_answers_values
    owner to postgres;


CREATE TABLE "user_skills"
(
    account_id integer NOT NULL,
    skill_id smallint NOT NULL,
    xp integer NOT NULL DEFAULT 0,
    CONSTRAINT "user_skills_pk" PRIMARY KEY (account_id, skill_id)
);


CREATE TABLE "skill_changes"
(
    id serial NOT NULL,
    account_id integer NOT NULL,
    skill_id smallint NOT NULL,
    date date NOT NULL,
    action_id smallint NOT NULL,
    points integer NOT NULL,
    CONSTRAINT "user_skill_changes_pk" PRIMARY KEY (id)
);


create table "test_actions"
(
    id serial,
    title varchar(64) not null,
    constraint "test_actions_pk" primary key (id)
);


CREATE TABLE skill_normalizations
(
    skill_id smallint NOT NULL,
    minimum smallint NOT NULL,
    maximum smallint NOT NULL,
    CONSTRAINT "skill_normalization_pkey" PRIMARY KEY (skill_id)
);

ALTER TABLE "skill_changes"
    ADD CONSTRAINT "skill_changes_skill_fk" FOREIGN KEY (skill_id)
        REFERENCES "skills" (id) MATCH SIMPLE
        ON UPDATE NO ACTION
           ON DELETE NO ACTION
        NOT VALID,
    ADD CONSTRAINT "skill_changes_action_fk" FOREIGN KEY (action_id)
        REFERENCES "test_actions" (id) MATCH SIMPLE
        ON UPDATE NO ACTION
           ON DELETE NO ACTION
        NOT VALID;

ALTER TABLE skill_normalizations ADD
    CONSTRAINT "skill_normalization_fk" FOREIGN KEY (skill_id)
        REFERENCES "skills" (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
    NOT VALID;

CREATE TABLE "avatars"
(
    account_id serial NOT NULL,
    hair_id smallint NOT NULL,
    hair_color int NOT NULL,
    eyes_id smallint NOT NULL,
    eyes_color int NOT NULL,
    clothes_id smallint NOT NULL,
    expression_id smallint NOT NULL,
    skin_color int NOT NULL,
    CONSTRAINT "avatar_pkey" PRIMARY KEY (account_id)
);


INSERT INTO public."skills"(id, title) VALUES(1, 'Confidences');
INSERT INTO public."skills"(id, title) VALUES(2, 'Socials');
INSERT INTO public."skills"(id, title) VALUES(3, 'Emotions');
INSERT INTO public."skills"(id, title) VALUES(4, 'Conflicts');
INSERT INTO public."skills"(id, title) VALUES(5, 'Creativity');
INSERT INTO public."skills"(id, title) VALUES(6, 'Thinking');
INSERT INTO public."skills"(id, title) VALUES(7, 'Purposes');
INSERT INTO public."skills"(id, title) VALUES(8, 'Responsibility');


INSERT INTO "test_questions"(id, text) VALUES(6,'Какой средний балл успеваемости у Вас в Вашем учебном заведении?');

INSERT INTO "test_answers"(id, question_id, text) VALUES(1, 6, '3');
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(1, 1, 1);
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(1, 2, 2);
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(1, 3, 1);
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(1, 4, 1);
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(1, 7, 1);
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(1, 8, 1);


INSERT INTO "test_answers"(id, question_id, text) VALUES(2, 6, '4');
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(2, 1, 2);
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(2, 2, 2);
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(2, 3, 2);
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(2, 4, 1);
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(2, 7, 2);
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(2, 8, 2);


INSERT INTO "test_answers"(id, question_id, text) VALUES(3, 6, '5');
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(3, 1, 3);
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(3, 2, 2);
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(3, 3, 3);
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(3, 4, 2);
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(3, 7, 3);
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(3, 8, 3);


INSERT INTO "test_questions"(id, text) VALUES(7,'Есть ли у Вас достижения в спорте  (награды, разряды)?');

INSERT INTO "test_answers"(id, question_id, text) VALUES(4, 7, 'Да');
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(4, 1, 3);
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(4, 2, 2);
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(4, 3, 2);
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(4, 7, 0);
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(4, 1, 3);
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(4, 8, 3);


INSERT INTO "test_answers"(id, question_id, text) VALUES(5, 7, 'Нет');


INSERT INTO "test_questions"(id, text) VALUES(8,'Есть ли у Вас достижения в творчестве (музыка, живопись, поэзия и так далее)?');

INSERT INTO "test_answers"(id, question_id, text) VALUES(6, 8, 'Да');
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(6, 2, 1);
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(6, 3, 2);
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(6, 5, 2);
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(6, 6, 2);
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(6, 8, 1);


INSERT INTO "test_answers"(id, question_id, text) VALUES(7, 8, 'Нет');


INSERT INTO "test_questions"(id, text) VALUES(9,'Есть ли у Вас достижения в учебных/научных отраслях (олимпиады, конференции и так далее)?');

INSERT INTO "test_answers"(id, question_id, text) VALUES(8, 9, 'Да');
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(8, 1, 3);
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(8, 2, 2);
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(8, 3, 2);
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(8, 7, 3);
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(8, 8, 3);


INSERT INTO "test_answers"(id, question_id, text) VALUES(9, 9, 'Нет');
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(9, 1, 1);


INSERT INTO "test_questions"(id, text) VALUES(10,'Посещали ли Вы детский сад?');

INSERT INTO "test_answers"(id, question_id, text) VALUES(10, 10, 'Да');
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(10, 2, 2);
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(10, 3, 2);
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(10, 4, 1);
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(10, 8, 2);


INSERT INTO "test_answers"(id, question_id, text) VALUES(11, 10, 'Нет');
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(11, 2, 1);
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(11, 3, 1);


INSERT INTO "test_questions"(id, text) VALUES(11,'Болеете ли Вы чаще, чем Ваши сверстники?
');

INSERT INTO "test_answers"(id, question_id, text) VALUES(12, 11, 'Да');
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(12, 2, 2);
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(12, 3, 2);
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(12, 4, 1);
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(12, 8, 2);


INSERT INTO "test_answers"(id, question_id, text) VALUES(13, 11, 'Нет');
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(13, 2, 1);
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(13, 3, 1);


INSERT INTO "test_questions"(id, text) VALUES(12,'Включите Вашу фантазию. Если бы Вы были геометрической фигурой, то какой? ');

INSERT INTO "test_answers"(id, question_id, text) VALUES(14, 12, 'Круг');
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(14, 1, 1);
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(14, 2, 3);
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(14, 3, 2);
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(14, 4, 2);
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(14, 7, 1);
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(14, 8, 1);


INSERT INTO "test_answers"(id, question_id, text) VALUES(15, 12, 'Квадрат');
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(15, 1, 1);
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(15, 2, 1);
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(15, 3, 2);
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(15, 7, 2);
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(15, 8, 1);


INSERT INTO "test_answers"(id, question_id, text) VALUES(16, 12, 'Прямоугольник');
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(16, 1, 1);
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(16, 2, 1);
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(16, 3, 1);
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(16, 7, 1);
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(16, 8, 1);


INSERT INTO "test_answers"(id, question_id, text) VALUES(17, 12, 'Треугольник');
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(17, 1, 1);
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(17, 2, 1);
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(17, 3, 1);
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(17, 6, 1);
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(17, 7, 3);
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(17, 8, 1);


INSERT INTO "test_answers"(id, question_id, text) VALUES(18, 12, 'Зигзаг');
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(18, 1, 1);
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(18, 2, 1);
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(18, 3, 1);
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(18, 5, 3);
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(18, 6, 3);
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(18, 7, 2);
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(18, 8, 1);


INSERT INTO "test_questions"(id, text) VALUES(13,'Нравится ли Вам быть в центре внимания окружающих?');

INSERT INTO "test_answers"(id, question_id, text) VALUES(19, 13, 'Да');
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(19, 1, 3);
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(19, 2, 2);
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(19, 3, 2);
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(19, 4, 1);
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(19, 5, 1);
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(19, 6, 2);
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(19, 8, 1);


INSERT INTO "test_answers"(id, question_id, text) VALUES(20, 13, 'Скорее да');
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(20, 1, 2);
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(20, 2, 2);
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(20, 3, 1);
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(20, 6, 1);


INSERT INTO "test_answers"(id, question_id, text) VALUES(21, 13, 'Трудно определиться');


INSERT INTO "test_answers"(id, question_id, text) VALUES(22, 13, 'Скорее нет');
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(22, 1, -1);
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(22, 2, -1);
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(22, 3, -1);


INSERT INTO "test_answers"(id, question_id, text) VALUES(23, 13, 'Нет');
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(23, 1, -2);
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(23, 2, -2);
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(23, 3, -2);


INSERT INTO "test_questions"(id, text) VALUES(14,'Согласны ли Вы со следующим высказыванием? Я считаю себя неуверенным человеком');

INSERT INTO "test_answers"(id, question_id, text) VALUES(24, 14, 'Да');
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(24, 1, -3);
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(24, 3, -2);
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(24, 4, -2);


INSERT INTO "test_answers"(id, question_id, text) VALUES(25, 14, 'Скорее да');
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(25, 1, -2);
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(25, 3, -1);
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(25, 4, -1);


INSERT INTO "test_answers"(id, question_id, text) VALUES(26, 14, 'Трудно определиться');
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(26, 1, -2);


INSERT INTO "test_answers"(id, question_id, text) VALUES(27, 14, 'Скорее нет');
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(27, 1, 1);
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(27, 3, 1);
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(27, 4, 1);


INSERT INTO "test_answers"(id, question_id, text) VALUES(28, 14, 'Нет');
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(28, 1, 3);
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(28, 3, 2);
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(28, 4, 2);


INSERT INTO "test_questions"(id, text) VALUES(15,'Согласны ли Вы со следующим высказыванием? У меня нет проблем в установлении контактов и  в общении');

INSERT INTO "test_answers"(id, question_id, text) VALUES(29, 15, 'Да');
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(29, 1, 2);
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(29, 2, 3);
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(29, 3, 2);
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(29, 4, 2);
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(29, 5, 1);


INSERT INTO "test_answers"(id, question_id, text) VALUES(30, 15, 'Скорее да');
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(30, 1, 1);
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(30, 2, 2);
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(30, 3, 1);
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(30, 4, 1);


INSERT INTO "test_answers"(id, question_id, text) VALUES(31, 15, 'Трудно определиться');


INSERT INTO "test_answers"(id, question_id, text) VALUES(32, 15, 'Скорее нет');
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(32, 1, -1);
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(32, 2, -1);
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(32, 3, -1);
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(32, 4, -1);


INSERT INTO "test_answers"(id, question_id, text) VALUES(33, 15, 'Нет');
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(33, 1, -2);
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(33, 2, -2);
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(33, 3, -2);
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(33, 4, -2);


INSERT INTO "test_questions"(id, text) VALUES(16,'Согласны ли Вы со следующим высказыванием? В конфликте я часто прибегаю к методам жесткого давления на оппонента');

INSERT INTO "test_answers"(id, question_id, text) VALUES(34, 16, 'Да');
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(34, 1, 2);
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(34, 2, -1);
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(34, 3, -2);
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(34, 4, -3);


INSERT INTO "test_answers"(id, question_id, text) VALUES(35, 16, 'Скорее да');
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(35, 1, 1);
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(35, 3, -1);
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(35, 4, -2);


INSERT INTO "test_answers"(id, question_id, text) VALUES(36, 16, 'Трудно определиться');
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(36, 2, 1);


INSERT INTO "test_answers"(id, question_id, text) VALUES(37, 16, 'Скорее нет');
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(37, 1, -1);
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(37, 2, 1);
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(37, 4, 1);


INSERT INTO "test_answers"(id, question_id, text) VALUES(38, 16, 'Нет');
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(38, 1, -2);
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(38, 3, 1);
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(38, 4, 2);


INSERT INTO "test_questions"(id, text) VALUES(17,'Согласны ли Вы со следующим высказыванием? Когда я оказываюсь неправым в споре, я болезненно это переживаю');

INSERT INTO "test_answers"(id, question_id, text) VALUES(39, 17, 'Да');
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(39, 1, -2);
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(39, 2, -2);
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(39, 3, -3);
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(39, 4, -3);


INSERT INTO "test_answers"(id, question_id, text) VALUES(40, 17, 'Скорее да');
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(40, 1, -1);
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(40, 2, -1);
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(40, 3, -2);
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(40, 4, -2);


INSERT INTO "test_answers"(id, question_id, text) VALUES(41, 17, 'Трудно определиться');
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(41, 2, -1);
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(41, 3, -1);
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(41, 4, -1);


INSERT INTO "test_answers"(id, question_id, text) VALUES(42, 17, 'Скорее нет');
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(42, 1, 1);
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(42, 2, 1);
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(42, 3, 1);
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(42, 4, 1);


INSERT INTO "test_answers"(id, question_id, text) VALUES(43, 17, 'Нет');
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(43, 1, 2);
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(43, 2, 2);
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(43, 3, 2);
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(43, 4, 2);


INSERT INTO "test_questions"(id, text) VALUES(18,'Согласны ли Вы со следующим высказыванием? Я всегда успеваю выполнить домашнее задание до конца крайнего срока');

INSERT INTO "test_answers"(id, question_id, text) VALUES(44, 18, 'Да');
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(44, 7, 3);
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(44, 8, 3);


INSERT INTO "test_answers"(id, question_id, text) VALUES(45, 18, 'Скорее да');
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(45, 7, 2);
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(45, 8, 2);


INSERT INTO "test_answers"(id, question_id, text) VALUES(46, 18, 'Трудно определиться');
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(46, 7, -1);
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(46, 8, -1);


INSERT INTO "test_answers"(id, question_id, text) VALUES(47, 18, 'Скорее нет');
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(47, 7, -2);
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(47, 8, -2);


INSERT INTO "test_answers"(id, question_id, text) VALUES(48, 18, 'Нет');
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(48, 7, -3);
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(48, 8, -3);


INSERT INTO "test_questions"(id, text) VALUES(19,'Согласны ли Вы со следующим высказыванием? Я всегда выполняю данное обещание.');

INSERT INTO "test_answers"(id, question_id, text) VALUES(49, 19, 'Да');
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(49, 1, 2);
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(49, 7, 3);
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(49, 8, 3);


INSERT INTO "test_answers"(id, question_id, text) VALUES(50, 19, 'Скорее да');
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(50, 1, 1);
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(50, 7, 2);
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(50, 8, 2);


INSERT INTO "test_answers"(id, question_id, text) VALUES(51, 19, 'Трудно определиться');


INSERT INTO "test_answers"(id, question_id, text) VALUES(52, 19, 'Скорее нет');
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(52, 1, -1);
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(52, 7, -1);
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(52, 8, -1);


INSERT INTO "test_answers"(id, question_id, text) VALUES(53, 19, 'Нет');


INSERT INTO "test_questions"(id, text) VALUES(20,'Согласны ли Вы со следующим высказыванием? Я часто отвлекаюсь при выполнении работы');

INSERT INTO "test_answers"(id, question_id, text) VALUES(54, 20, 'Да');
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(54, 7, -2);
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(54, 8, -3);


INSERT INTO "test_answers"(id, question_id, text) VALUES(55, 20, 'Скорее да');
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(55, 7, -1);
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(55, 8, -2);


INSERT INTO "test_answers"(id, question_id, text) VALUES(56, 20, 'Трудно определиться');


INSERT INTO "test_answers"(id, question_id, text) VALUES(57, 20, 'Скорее нет');
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(57, 7, 1);
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(57, 8, 1);


INSERT INTO "test_answers"(id, question_id, text) VALUES(58, 20, 'Нет');
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(58, 7, 2);
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(58, 8, 2);


INSERT INTO "test_questions"(id, text) VALUES(21,'Согласны ли Вы со следующим высказыванием? Я умею планировать свое время');

INSERT INTO "test_answers"(id, question_id, text) VALUES(59, 21, 'Да');
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(59, 1, 2);
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(59, 7, 3);
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(59, 8, 3);


INSERT INTO "test_answers"(id, question_id, text) VALUES(60, 21, 'Скорее да');
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(60, 1, 1);
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(60, 7, 2);
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(60, 8, 2);


INSERT INTO "test_answers"(id, question_id, text) VALUES(61, 21, 'Трудно определиться');


INSERT INTO "test_answers"(id, question_id, text) VALUES(62, 21, 'Скорее нет');
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(62, 7, -1);
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(62, 8, -1);


INSERT INTO "test_answers"(id, question_id, text) VALUES(63, 21, 'Нет');
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(63, 7, -2);
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(63, 8, -2);


INSERT INTO "test_questions"(id, text) VALUES(22,'Согласны ли Вы со следующим высказыванием? Мне нравится учиться.');

INSERT INTO "test_answers"(id, question_id, text) VALUES(64, 22, 'Да');
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(64, 1, 2);
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(64, 2, 3);
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(64, 7, 3);
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(64, 8, 3);


INSERT INTO "test_answers"(id, question_id, text) VALUES(65, 22, 'Скорее да');
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(65, 1, 1);
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(65, 2, 2);
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(65, 7, 2);
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(65, 8, 2);


INSERT INTO "test_answers"(id, question_id, text) VALUES(66, 22, 'Трудно определиться');


INSERT INTO "test_answers"(id, question_id, text) VALUES(67, 22, 'Скорее нет');
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(67, 7, -1);
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(67, 8, -1);


INSERT INTO "test_answers"(id, question_id, text) VALUES(68, 22, 'Нет');
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(68, 2, -1);
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(68, 7, -2);
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(68, 8, -2);


INSERT INTO "test_questions"(id, text) VALUES(23,'Согласны ли Вы со следующим высказыванием? Мне нравится организовывать игры для друзей');

INSERT INTO "test_answers"(id, question_id, text) VALUES(69, 23, 'Да');
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(69, 1, 2);
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(69, 2, 3);
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(69, 3, 2);
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(69, 4, 2);
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(69, 5, 2);
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(69, 6, 2);
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(69, 7, 3);
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(69, 8, 3);


INSERT INTO "test_answers"(id, question_id, text) VALUES(70, 23, 'Скорее да');
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(70, 1, 1);
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(70, 2, 2);
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(70, 3, 1);
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(70, 4, 1);
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(70, 5, 1);
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(70, 6, 1);
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(70, 7, 2);
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(70, 8, 2);


INSERT INTO "test_answers"(id, question_id, text) VALUES(71, 23, 'Трудно определиться');


INSERT INTO "test_answers"(id, question_id, text) VALUES(72, 23, 'Скорее нет');
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(72, 1, -1);
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(72, 2, 1);
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(72, 4, -1);
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(72, 5, -1);
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(72, 6, -1);
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(72, 7, -1);
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(72, 8, -1);


INSERT INTO "test_answers"(id, question_id, text) VALUES(73, 23, 'Нет');
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(73, 1, -2);
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(73, 3, -1);
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(73, 4, -2);
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(73, 5, -2);
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(73, 6, -2);
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(73, 7, -2);
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(73, 8, -2);


INSERT INTO "test_questions"(id, text) VALUES(24,'Согласны ли Вы со следующим высказыванием? Я постоянно увлекаюсь чем-то новым.');

INSERT INTO "test_answers"(id, question_id, text) VALUES(74, 24, 'Да');
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(74, 5, 2);
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(74, 6, 2);


INSERT INTO "test_answers"(id, question_id, text) VALUES(75, 24, 'Скорее да');
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(75, 5, 1);
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(75, 6, 1);


INSERT INTO "test_answers"(id, question_id, text) VALUES(76, 24, 'Трудно определиться');


INSERT INTO "test_answers"(id, question_id, text) VALUES(77, 24, 'Скорее нет');
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(77, 5, -1);
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(77, 6, -1);


INSERT INTO "test_answers"(id, question_id, text) VALUES(78, 24, 'Нет');
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(78, 5, -2);
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(78, 6, -2);


INSERT INTO "test_questions"(id, text) VALUES(25,'Согласны ли Вы со следующим высказыванием?  Мне сложно найти общий язык в общении с другими людьми.');

INSERT INTO "test_answers"(id, question_id, text) VALUES(79, 25, 'Да');
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(79, 1, -2);
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(79, 2, -3);
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(79, 3, -3);
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(79, 4, -3);


INSERT INTO "test_answers"(id, question_id, text) VALUES(80, 25, 'Скорее да');
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(80, 1, -1);
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(80, 2, -2);
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(80, 3, -2);
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(80, 4, -2);


INSERT INTO "test_answers"(id, question_id, text) VALUES(81, 25, 'Трудно определиться');


INSERT INTO "test_answers"(id, question_id, text) VALUES(82, 25, 'Скорее нет');
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(82, 1, 1);
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(82, 2, 1);
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(82, 3, 1);
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(82, 4, 1);


INSERT INTO "test_answers"(id, question_id, text) VALUES(83, 25, 'Нет');
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(83, 1, 2);
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(83, 2, 2);
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(83, 3, 2);
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(83, 4, 2);


INSERT INTO "test_questions"(id, text) VALUES(26,'Согласны ли Вы со следующим высказыванием? Для меня важно оставаться в хороших отношениях со всеми, занимаясь общим делом.');

INSERT INTO "test_answers"(id, question_id, text) VALUES(84, 26, 'Да');
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(84, 1, 2);
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(84, 2, 3);
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(84, 3, 2);
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(84, 4, 3);
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(84, 7, 2);
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(84, 8, 2);


INSERT INTO "test_answers"(id, question_id, text) VALUES(85, 26, 'Скорее да');
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(85, 1, 1);
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(85, 2, 2);
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(85, 3, 1);
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(85, 4, 2);
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(85, 7, 1);
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(85, 8, 1);


INSERT INTO "test_answers"(id, question_id, text) VALUES(86, 26, 'Трудно определиться');


INSERT INTO "test_answers"(id, question_id, text) VALUES(87, 26, 'Скорее нет');
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(87, 1, 0);
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(87, 2, -1);
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(87, 3, -1);
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(87, 4, -1);
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(87, 7, -1);
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(87, 8, -2);


INSERT INTO "test_answers"(id, question_id, text) VALUES(88, 26, 'Нет');
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(88, 1, 0);
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(88, 2, -2);
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(88, 3, -2);
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(88, 4, -2);
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(88, 7, -2);
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(88, 8, -3);


INSERT INTO "test_questions"(id, text) VALUES(27,'Согласны ли Вы со следующим высказыванием? Я часто обижаюсь на кого-либо.');

INSERT INTO "test_answers"(id, question_id, text) VALUES(89, 27, 'Да');
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(89, 1, -3);
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(89, 2, -3);
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(89, 3, -3);
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(89, 4, -3);


INSERT INTO "test_answers"(id, question_id, text) VALUES(90, 27, 'Скорее да');
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(90, 1, -2);
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(90, 2, -2);
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(90, 3, -2);
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(90, 4, -2);


INSERT INTO "test_answers"(id, question_id, text) VALUES(91, 27, 'Трудно определиться');
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(91, 1, -2);
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(91, 2, -2);
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(91, 3, -2);
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(91, 4, -1);


INSERT INTO "test_answers"(id, question_id, text) VALUES(92, 27, 'Скорее нет');
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(92, 1, 1);
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(92, 2, 1);
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(92, 3, 1);
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(92, 4, 1);


INSERT INTO "test_answers"(id, question_id, text) VALUES(93, 27, 'Нет');
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(93, 1, 2);
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(93, 2, 3);
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(93, 3, 2);
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(93, 4, 2);


INSERT INTO "test_questions"(id, text) VALUES(28,'Согласны ли Вы со следующим высказыванием? Для меня важно достижение цели.');

INSERT INTO "test_answers"(id, question_id, text) VALUES(94, 28, 'Да');
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(94, 1, 3);
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(94, 7, 3);
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(94, 8, 3);


INSERT INTO "test_answers"(id, question_id, text) VALUES(95, 28, 'Скорее да');
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(95, 1, 2);
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(95, 7, 2);
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(95, 8, 2);


INSERT INTO "test_answers"(id, question_id, text) VALUES(96, 28, 'Трудно определиться');


INSERT INTO "test_answers"(id, question_id, text) VALUES(97, 28, 'Скорее нет');
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(97, 1, -1);
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(97, 7, -2);
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(97, 8, -1);


INSERT INTO "test_answers"(id, question_id, text) VALUES(98, 28, 'Нет');
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(98, 1, -2);
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(98, 7, -3);
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(98, 8, -2);


INSERT INTO "test_questions"(id, text) VALUES(29,'Согласны ли Вы со следующим высказыванием? Я легко могу сдержать негативные эмоции.');

INSERT INTO "test_answers"(id, question_id, text) VALUES(99, 29, 'Да');
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(99, 1, 3);
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(99, 2, 3);
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(99, 3, 3);
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(99, 4, 2);


INSERT INTO "test_answers"(id, question_id, text) VALUES(100, 29, 'Скорее да');
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(100, 1, 2);
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(100, 2, 2);
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(100, 3, 2);
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(100, 4, 1);


INSERT INTO "test_answers"(id, question_id, text) VALUES(101, 29, 'Трудно определиться');


INSERT INTO "test_answers"(id, question_id, text) VALUES(102, 29, 'Скорее нет');
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(102, 1, -1);
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(102, 2, -1);
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(102, 3, -1);
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(102, 4, -1);


INSERT INTO "test_answers"(id, question_id, text) VALUES(103, 29, 'Нет');
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(103, 1, -2);
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(103, 2, -2);
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(103, 3, -3);
INSERT INTO "test_answers_values"(answer_id, skill_id, points) VALUES(103, 4, -2);

insert into skill_normalizations(skill_id, minimum, maximum)
values(1, -21, 42),
      (2, -13, 39),
      (3, -19, 34),
      (4, -21, 26),
      (5, -4, 11),
      (6, -4, 11),
      (7, -16, 34),
      (8, -17, 38);

insert into "test_actions"(title)
values('init_test');

insert into user_data(account_id, age, xp, gender)
values(1, 1337, 0, 'm');

insert into user_answers(account_id, question_id, answer_id)
values (1, 6, 1);

insert into user_skills(account_id, skill_id, xp)
values (1, 1, 0),
(1, 2, 0),
(1, 3, 0),
(1, 4, 0),
(1, 5, 0),
(1, 6, 0),
(1, 7, 0),
(1, 8, 0);