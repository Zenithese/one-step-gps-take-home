DROP TABLE IF EXISTS device_photo;
DROP TABLE IF EXISTS preference;
DROP TABLE IF EXISTS users;

CREATE TABLE users (
    user_id INT AUTO_INCREMENT PRIMARY KEY,
    username VARCHAR(255) NOT NULL UNIQUE,
    password_hash VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE preference (
    user_id INT NOT NULL,
    device_order JSON,
    hidden_devices JSON,
    PRIMARY KEY (user_id),
    FOREIGN KEY (user_id) REFERENCES users(user_id) ON DELETE CASCADE
);

CREATE TABLE device_photo (
    user_id INT NOT NULL,
    device_id VARCHAR(255) NOT NULL,
    photo_blob LONGBLOB NOT NULL,
    PRIMARY KEY (user_id, device_id),
    FOREIGN KEY (user_id) REFERENCES users(user_id) ON DELETE CASCADE
);

DELIMITER $$

CREATE TRIGGER set_default_preferences
BEFORE INSERT ON preference
FOR EACH ROW
BEGIN
    IF NEW.device_order IS NULL THEN
        SET NEW.device_order = '[]';
    END IF;
    IF NEW.hidden_devices IS NULL THEN
        SET NEW.hidden_devices = '[]';
    END IF;
END$$

DELIMITER ;