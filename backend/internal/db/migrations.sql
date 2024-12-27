DROP TABLE IF EXISTS preference;
CREATE TABLE preference (
    id INT PRIMARY KEY,
    device_order JSON DEFAULT NULL,
    hidden_devices JSON DEFAULT NULL
);

DROP TABLE IF EXISTS device_photo;
CREATE TABLE device_photo (
    device_id VARCHAR(255) PRIMARY KEY,
    photo_blob LONGBLOB NOT NULL
);

INSERT INTO preference
  (id, device_order, hidden_devices)
VALUES
  (1, '[]', '[]');