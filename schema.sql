begin;

-- @author  Mikhail Kirillov <mikkirillov@yandex.ru>
-- @version 1.002
-- @date    2019-02-09

create table asso (
  id bigserial not null primary key,
  name character varying(255) not null,
  asso character varying(255) not null,
  created timestamp(0) without time zone not null default now(),
  counter bigint not null default 1,
  checked bool not null default false,
  unique(name,asso)
);

comment on table asso is 'таблица с ассоциациями';

comment on column asso.id is 'id записи';
comment on column asso.name is 'слово, к которому придумываем оссоциации';
comment on column asso.asso is 'ассоциация';
comment on column asso.created is 'дата создания ассоциации';
comment on column asso.counter is 'сколько раз ассоциация использовалась';
comment on column asso.checked is 'флаг того, что ассоциация проверена';

create table words (
  id bigserial not null primary key,
  name character varying(255) not null,
  created timestamp(0) without time zone not null default now(),
  unique(name)
);

comment on table words is 'таблица слов для начала игры';

comment on column words.id is 'id слова';
comment on column words.name is 'слово';
comment on column words.created is 'время создания';

insert into words(name) values ('гаджет'),('город'),('еда'),('животное'),('игра'),('имя'),('книга'),('компьютер'),
  ('лекарство'),('машина'),('музыка'),('наука'),('одежда'),('писатель'),('политик'),('предмет'),('птица'),('россия'),
  ('страна'),('ученый'),('фильм'),('школа');

end;
