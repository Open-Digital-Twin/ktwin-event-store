-- ktwin

CREATE KEYSPACE ktwin WITH REPLICATION = { 'class' : 'SimpleStrategy', 'replication_factor' : 1 };

USE ktwin;

-- TODO: review primary key
CREATE TABLE twin_interface (
id text,
name text,
active boolean, 
parent text,
data_definition text,
created_at timestamp,
PRIMARY KEY((id)));

INSERT INTO twin_interface ("id", "name", "active", "parent", "data_definition", "created_at") VALUES ('FactoryId', 'Factory', true, '', 'DATA_DEFINITION_FACTORY', toTimestamp(now()));
INSERT INTO twin_interface ("id", "name", "active", "parent", "data_definition", "created_at") VALUES ('MachineId', 'Machine', true, 'Factory', 'DATA_DEFINITION_MACHINE', toTimestamp(now()));
INSERT INTO twin_interface ("id", "name", "active", "parent", "data_definition", "created_at") VALUES ('SensorId', 'Sensor', true, 'Machine', 'DATA_DEFINITION_SENSOR', toTimestamp(now()));

CREATE TABLE twin_instance (
id text,
name text,
interface_id text,
active boolean,
parent text,
last_event_data text,
created_at timestamp,
PRIMARY KEY((interface_id, id)));

INSERT INTO twin_instance ("id", "name", "interface_id", "active", "parent", "last_event_data", "created_at") VALUES ('Factory_01_id', 'Factory_01', 'Factory', true, '', '', toTimestamp(now()));
INSERT INTO twin_instance ("id", "name", "interface_id", "active", "parent", "last_event_data", "created_at") VALUES ('Machine_01_id', 'Machine_01', 'Machine', true, 'Factory_01', '', toTimestamp(now()));
INSERT INTO twin_instance ("id", "name", "interface_id", "active", "parent", "last_event_data", "created_at") VALUES ('Machine_02_id', 'Machine_02', 'Machine', true, 'Factory_01', '', toTimestamp(now()));
INSERT INTO twin_instance ("id", "name", "interface_id", "active", "parent", "last_event_data", "created_at") VALUES ('Sensor_01_01_id', 'Sensor_01_01', 'Sensor', true, 'Machine_01', '', toTimestamp(now()));
INSERT INTO twin_instance ("id", "name", "interface_id", "active", "parent", "last_event_data", "created_at") VALUES ('Sensor_01_02_id', 'Sensor_01_02', 'Sensor', true, 'Machine_01', '', toTimestamp(now()));
INSERT INTO twin_instance ("id", "name", "interface_id", "active", "parent", "last_event_data", "created_at") VALUES ('Sensor_02_01_id', 'Sensor_02_01', 'Sensor', true, 'Machine_02', '', toTimestamp(now()));
INSERT INTO twin_instance ("id", "name", "interface_id", "active", "parent", "last_event_data", "created_at") VALUES ('Sensor_02_02_id', 'Sensor_02_02', 'Sensor', true, 'Machine_02', '', toTimestamp(now()));

-- twin_event

CREATE TABLE twin_event (
instance_id text,
interface_id text,
event_data text,
created_at timestamp,
PRIMARY KEY((interface_id, instance_id), created_at)
) WITH CLUSTERING ORDER BY (created_at DESC);

INSERT INTO twin_event ("instance_id","interface_id","event_data","created_at") VALUES ('Sensor_01_01','Sensor', 'Sensor_01_01_DATA', toTimestamp(now()));
INSERT INTO twin_event ("instance_id","interface_id","event_data","created_at") VALUES ('Sensor_01_01','Sensor', 'Sensor_01_01_DATA', toTimestamp(now()));
INSERT INTO twin_event ("instance_id","interface_id","event_data","created_at") VALUES ('Sensor_01_02','Sensor', 'Sensor_01_02_DATA', toTimestamp(now()));
INSERT INTO twin_event ("instance_id","interface_id","event_data","created_at") VALUES ('Sensor_01_02','Sensor', 'Sensor_01_02_DATA', toTimestamp(now()));
INSERT INTO twin_event ("instance_id","interface_id","event_data","created_at") VALUES ('Sensor_02_01','Sensor', 'Sensor_02_01_DATA', toTimestamp(now()));
INSERT INTO twin_event ("instance_id","interface_id","event_data","created_at") VALUES ('Sensor_02_01','Sensor', 'Sensor_02_01_DATA', toTimestamp(now()));
INSERT INTO twin_event ("instance_id","interface_id","event_data","created_at") VALUES ('Sensor_02_02','Sensor', 'Sensor_02_02_DATA', toTimestamp(now()));
INSERT INTO twin_event ("instance_id","interface_id","event_data","created_at") VALUES ('Sensor_02_02','Sensor', 'Sensor_02_02_DATA', toTimestamp(now()));
