CREATE KEYSPACE global WITH REPLICATION = {
	'class' : 'SimpleStrategy',
	'replication_factor' : 3
}
AND DURABLE_WRITES =  true;

USE global;

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
