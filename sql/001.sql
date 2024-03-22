drop table if exists task;
create table task (
                      id SERIAL PRIMARY KEY ,
                      title TEXT,
                      description TEXT,
                      order_num INTEGER
);
insert into task (title, order_num) values ( 'task1', 555);
insert into task (title, order_num) values ( 'task2', 333);
insert into task (title, order_num) values ( 'task3', 111);