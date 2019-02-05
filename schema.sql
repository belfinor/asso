begin;


create table asso (
  id bigserial not null primary key,
  name character varying(255) not null,
  asso character varying(255) not null,
  counter bigint not null default 1,
  checked bool not null default false,
  unique(name,asso)
);

create table words (
  id bigserial not null primary key,
  name character varying(255) not null,
  unique(name)
);

end;

