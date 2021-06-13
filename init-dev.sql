DROP TABLE IF EXISTS INSTANCE_GPU;
DROP TABLE IF EXISTS INSTANCE;
DROP TABLE IF EXISTS GPU;
DROP TABLE IF EXISTS NODE;
DROP TABLE IF EXISTS USER;

CREATE TABLE IF NOT EXISTS NODE (
    node_id INT(10) NOT NULL AUTO_INCREMENT PRIMARY KEY,
    storage_group_id INT(10) NOT NULL,
    host_name varchar(32) NOT NULL,
    node_name varchar(32) NOT NULL,
    host_ip varchar(15) NOT NULL,
    port varchar(5) NOT NULL,
    excluded TINYINT(1) NOT NULL DEFAULT 0,
    exclusive TINYINT(1) NOT NULL DEFAULT 0,
    registered_at DATETIME,
    excluded_at DATETIME,
    last_allocated_at DATETIME,
    last_heartbeat_at DATETIME
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE IF NOT EXISTS GPU (
    id INT(10) NOT NULL AUTO_INCREMENT PRIMARY KEY,
    node_id INT(10) NOT NULL,
    model_name varchar(32) NOT NULL,
    bus_id varchar(32) NOT NULL,
    slot_index INT(10) NOT NULL,
    excluded TINYINT(1) NOT NULL DEFAULT 0,
    exclusive TINYINT(1) NOT NULL DEFAULT 0,
    excluded_at DATETIME,
    last_allocated_at DATETIME,
    FOREIGN KEY (node_id) REFERENCES NODE (node_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE IF NOT EXISTS USER (
    id INT(10) NOT NULL AUTO_INCREMENT PRIMARY KEY,
    login_id varchar(32) NOT NULL,
    password varchar(255) NOT NULL,
    user_name varchar(32) NOT NULL,
    phone VARCHAR(32) NOT NULL,
    email VARCHAR(32) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE IF NOT EXISTS INSTANCE (
    id INT(10) NOT NULL AUTO_INCREMENT PRIMARY KEY,
    node_id INT(10) NOT NULL,
    user_id INT(10) NOT NULL,
    name VARCHAR(48) NOT NULL,
    instance_hash VARCHAR(255) NOT NULL,
    image VARCHAR(128) NOT NULL,
    created_at DATETIME,
    expired_at DATETIME,
    deleted_at DATETIME,
    FOREIGN KEY (node_id) REFERENCES NODE (node_id),
    FOREIGN KEY (user_id) REFERENCES USER (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE IF NOT EXISTS INSTANCE_GPU (
    id INT(10) NOT NULL AUTO_INCREMENT PRIMARY KEY,
    instance_id INT(10) NOT NULL,
    gpu_id INT(10) NOT NULL,
    created_at DATETIME NOT NULL,
    FOREIGN KEY (instance_id) REFERENCES INSTANCE (id),
    FOREIGN KEY (gpu_id) REFERENCES GPU (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

INSERT INTO NODE VALUES (default, 1, 'ubuntu', 'server-1', '192.168.1.15', '8081', 0, 0,
                         null, null, null, null);

INSERT INTO NODE VALUES (default, 1, 'ubuntu', 'server-2', '192.168.1.16', '8081', 0, 0,
                         null, null, null, null);

INSERT INTO NODE VALUES (default, 1, 'ubuntu', 'server-3', '192.168.1.17', '8081', 0, 0,
                         null, null, null, null);