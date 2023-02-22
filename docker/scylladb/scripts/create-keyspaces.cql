-- dt_definition

CREATE KEYSPACE dt_definition WITH REPLICATION = { 'class' : 'NetworkTopologyStrategy','DC1' : 3};

USE dt_definition;

CREATE TABLE twin_component (
name text,
description text, 
active boolean, 
parent text,
data_definition text,
PRIMARY KEY((name)));

INSERT INTO twin_component ("name", "description", "active", "parent", "data_definition") VALUES ('Factory','Factory model', true, '', 'DATA_DEFINITION_FACTORY') ;
INSERT INTO twin_component ("name", "description", "active", "parent", "data_definition") VALUES ('Machine','Machine model', true, 'Factory', 'DATA_DEFINITION_MACHINE') ;
INSERT INTO twin_component ("name", "description", "active", "parent", "data_definition") VALUES ('Sensor','Sensor model', true, 'Machine', 'DATA_DEFINITION_SENSOR') ;

CREATE TABLE twin_instance (
name text,
component text,
active boolean,
parent text,
created_at timestamp,
PRIMARY KEY((name, component)));

INSERT INTO twin_instance ("name", "component", "active", "parent", "created_at") VALUES ('Factory_01', 'Factory', true, '', toTimestamp(now()));
INSERT INTO twin_instance ("name", "component", "active", "parent", "created_at") VALUES ('Machine_01', 'Machine', true, 'Factory_01', toTimestamp(now()));
INSERT INTO twin_instance ("name", "component", "active", "parent", "created_at") VALUES ('Machine_02', 'Machine', true, 'Factory_01', toTimestamp(now()));
INSERT INTO twin_instance ("name", "component", "active", "parent", "created_at") VALUES ('Sensor_01_01', 'Sensor', true, 'Machine_01', toTimestamp(now()));
INSERT INTO twin_instance ("name", "component", "active", "parent", "created_at") VALUES ('Sensor_01_02', 'Sensor', true, 'Machine_01', toTimestamp(now()));
INSERT INTO twin_instance ("name", "component", "active", "parent", "created_at") VALUES ('Sensor_02_01', 'Sensor', true, 'Machine_02', toTimestamp(now()));
INSERT INTO twin_instance ("name", "component", "active", "parent", "created_at") VALUES ('Sensor_02_02', 'Sensor', true, 'Machine_02', toTimestamp(now()));

-- dt_event_store

CREATE KEYSPACE dt_event_store WITH REPLICATION = { 'class' : 'NetworkTopologyStrategy','DC1' : 3};

USE dt_event_store;

CREATE TABLE event_store_data (
    intance text,
    component text,
    created_at timestamp,
    event_data text,
    primary key((intance, component), created_at)
)
WITH CLUSTERING ORDER BY (created_at DESC)
AND COMPACTION = {'class': 'DateTieredCompactionStrategy',
    'base_time_seconds': 3600,
    'max_sstable_age_days': 1};


INSERT INTO dt_event_store.event_store_data ("intance","component","created_at","event_data") VALUES ('Sensor_01_01','Sensor', toTimestamp(now()), 'Sensor_01_01_DATA') ;
INSERT INTO dt_event_store.event_store_data ("intance","component","created_at","event_data") VALUES ('Sensor_01_01','Sensor', toTimestamp(now()), 'Sensor_01_01_DATA') ;
INSERT INTO dt_event_store.event_store_data ("intance","component","created_at","event_data") VALUES ('Sensor_01_02','Sensor', toTimestamp(now()), 'Sensor_01_02_DATA') ;
INSERT INTO dt_event_store.event_store_data ("intance","component","created_at","event_data") VALUES ('Sensor_01_02','Sensor', toTimestamp(now()), 'Sensor_01_02_DATA') ;
INSERT INTO dt_event_store.event_store_data ("intance","component","created_at","event_data") VALUES ('Sensor_02_01','Sensor', toTimestamp(now()), 'Sensor_02_01_DATA') ;
INSERT INTO dt_event_store.event_store_data ("intance","component","created_at","event_data") VALUES ('Sensor_02_01','Sensor', toTimestamp(now()), 'Sensor_02_01_DATA') ;
INSERT INTO dt_event_store.event_store_data ("intance","component","created_at","event_data") VALUES ('Sensor_02_02','Sensor', toTimestamp(now()), 'Sensor_02_02_DATA') ;
INSERT INTO dt_event_store.event_store_data ("intance","component","created_at","event_data") VALUES ('Sensor_02_02','Sensor', toTimestamp(now()), 'Sensor_02_02_DATA') ;
 