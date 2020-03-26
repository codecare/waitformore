-- {{ ansible_managed }}
CREATE TABLE QueueDefinition (
  id    SERIAL NOT NULL,
  Title   varchar(255) NOT NULL,
  UserKey   varchar(255) NOT NULL,
  AdminKey  varchar(255) NOT NULL,
  LastPulled integer NOT NULL,
  LastCalled integer NOT NULL,
  created timestamp not null,
  updated timestamp not null,
  deleted timestamp null,
  version int not null default 0
);