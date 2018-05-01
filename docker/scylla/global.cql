CREATE KEYSPACE global WITH REPLICATION = {
	'class' : 'SimpleStrategy',
	'replication_factor' : 3
}
AND DURABLE_WRITES =  true;

USE global;

CREATE TABLE
	global.token_
(
	uuid        timeuuid,
	permissions frozen<map<timeuuid, tinyint>>,
	account     timeuuid,
	ip          text,
	PRIMARY KEY(uuid)
);

CREATE TABLE
	global.actor
(
	uuid timeuuid,
	hp   bigint,
	mp   bigint,
	x    double,
	y    double,
	z    double,
	PRIMARY KEY(uuid)
);

CREATE TABLE
	global.action
(
	uuid timeuuid,
	ts timestamp,
	PRIMARY KEY((uuid, ts), ts)
);
