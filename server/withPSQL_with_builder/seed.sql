DROP TABLE IF EXISTS animals;

DROP TABLE IF EXISTS regions;

DROP TABLE IF EXISTS observation_devices;

CREATE TABLE
    IF NOT EXISTS regions (
        region_id integer NOT NULL,
        region_name character varying(255) NOT NULL,
        PRIMARY KEY (region_id)
    );

ALTER TABLE regions
ALTER COLUMN region_id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME region_id_seq START
    WITH
        1 INCREMENT BY 1 NO MINVALUE NO MAXVALUE CACHE 1
);

CREATE TABLE
    IF NOT EXISTS observation_devices (
        device_id integer NOT NULL,
        device_location character varying(255) NOT NULL,
        PRIMARY KEY (device_id)
    );

ALTER TABLE observation_devices
ALTER COLUMN device_id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME device_id_seq START
    WITH
        1 INCREMENT BY 1 NO MINVALUE NO MAXVALUE CACHE 1
);

CREATE TABLE
    IF NOT EXISTS animals (
        id integer NOT NULL,
        name character varying(255) NOT NULL,
        type character varying(255) NOT NULL,
        gender character varying(6) NOT NULL,
        age integer NOT NULL,
        created_at timestamp without time zone NOT NULL,
        updated_at timestamp without time zone NOT NULL,
        last_time_seen_at timestamp without time zone NOT NULL,
        seen_by_device_id integer NOT NULL,
        region_id integer NOT NULL,
        CONSTRAINT fk_region FOREIGN KEY (region_id) REFERENCES regions (region_id),
        CONSTRAINT fk_device FOREIGN KEY (seen_by_device_id) REFERENCES observation_devices (device_id),
        PRIMARY KEY (id)
    );

ALTER TABLE animals
ALTER COLUMN id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME animals_seq START
    WITH
        1 INCREMENT BY 1 NO MINVALUE NO MAXVALUE CACHE 1
);

INSERT INTO
    regions ("region_name")
VALUES
    ('Asia'),
    ('Africa'),
    ('Europe'),
    ('North America'),
    ('South America'),
    ('Australia');

INSERT INTO
    observation_devices ("device_location")
VALUES
    ('44.25:64.90'),
    ('124.50:332.09');

INSERT INTO
    animals (
        "name",
        "type",
        "gender",
        "age",
        "created_at",
        "updated_at",
        "last_time_seen_at",
        "seen_by_device_id",
        "region_id"
    )
VALUES
    (
        'Chloe',
        'lion',
        'female',
        5,
        '2024-03-16 00:00:00',
        '2024-03-16 00:00:00',
        '2024-03-16 00:00:00',
        1,
        2
    );

INSERT INTO
    animals (
        "name",
        "type",
        "gender",
        "age",
        "created_at",
        "updated_at",
        "last_time_seen_at",
        "seen_by_device_id",
        "region_id"
    )
VALUES
    (
        'Bill',
        'cougar',
        'male',
        7,
        '2024-03-16 00:00:00',
        '2024-03-16 00:00:00',
        '2024-03-16 00:00:00',
        2,
        4
    )