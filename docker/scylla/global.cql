CREATE KEYSPACE global WITH REPLICATION = {
	'class' : 'SimpleStrategy',
	'replication_factor' : 3
}
AND DURABLE_WRITES =  true;

USE global;

CREATE TABLE
	global.tok
(
	uuid        uuid,
	permissions frozen<map<uuid, tinyint>>,
	account     uuid,
	ip          text,
	PRIMARY KEY(uuid)
);

CREATE TABLE
	global.actor
(
	uuid uuid,
	hp   bigint,
	mp   bigint,
	x    double,
	y    double,
	z    double,
	PRIMARY KEY(uuid)
);
