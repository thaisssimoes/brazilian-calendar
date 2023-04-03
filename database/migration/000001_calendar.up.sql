CREATE TABLE CALENDARIO(
   data varchar(10) not null,
   dia int not null,
   mes int not null,
   ano int not null,
   feriado bit not null,
   dia_util bit not null
);

create index idx_data_calendario on CALENDARIO(data);

create index idx_ano_mes_calendario on CALENDARIO(ano, mes);

create index idx_ano_mes_dia_calendario on CALENDARIO(ano, mes,dia);