package repositories

import (
    "database/sql"
    "encoding/json"
    "log"
    "backend/internal/models"
)

type PreferenceRepository struct {
    DB *sql.DB
}

func NewPreferenceRepository(db *sql.DB) *PreferenceRepository {
    return &PreferenceRepository{DB: db}
}

func (r *PreferenceRepository) EnsurePreferencesRecord(userID int) error {
    _, err := r.DB.Exec("INSERT IGNORE INTO preference (user_id) VALUES (?)", userID)
    if err != nil {
        log.Printf("Error in EnsurePreferencesRecord: %v", err)
        return err
    }
    return nil
}

func (r *PreferenceRepository) SavePreferences(userID int, prefs *models.Preference) error {
    tx, err := r.DB.Begin()
    if err != nil {
        return err
    }

    if err := r.EnsurePreferencesRecord(userID); err != nil {
        tx.Rollback()
        return err
    }

    query := "UPDATE preference SET "
    args := []interface{}{}

    if prefs.DeviceOrder != nil {
        deviceOrderJSON, err := json.Marshal(prefs.DeviceOrder)
        if err != nil {
            tx.Rollback()
            return err
        }
        query += "device_order = ?, "
        args = append(args, deviceOrderJSON)
    }

    if prefs.HiddenDevices != nil {
        hiddenDevicesJSON, err := json.Marshal(prefs.HiddenDevices)
        if err != nil {
            tx.Rollback()
            return err
        }
        query += "hidden_devices = ?, "
        args = append(args, hiddenDevicesJSON)
    }

    if len(args) > 0 {
        query = query[:len(query)-2]
        query += " WHERE user_id = ?"
        args = append(args, userID)

        _, err := tx.Exec(query, args...)
        if err != nil {
            tx.Rollback()
            return err
        }
    }

    if prefs.DevicePhotos != nil {
        for deviceID, photo := range prefs.DevicePhotos {
            _, err := tx.Exec(`
                INSERT INTO device_photo (user_id, device_id, photo_blob)
                VALUES (?, ?, ?)
                ON DUPLICATE KEY UPDATE
                photo_blob = VALUES(photo_blob)
            `, userID, deviceID, photo)
            if err != nil {
                tx.Rollback()
                return err
            }
        }
    }

    return tx.Commit()
}

func (r *PreferenceRepository) GetPreferences(userID int) (*models.Preference, error) {
    prefs := &models.Preference{}

    if err := r.EnsurePreferencesRecord(userID); err != nil {
        return nil, err
    }

    row := r.DB.QueryRow("SELECT device_order, hidden_devices FROM preference WHERE user_id = ?", userID)
    var deviceOrderJSON, hiddenDevicesJSON []byte
    err := row.Scan(&deviceOrderJSON, &hiddenDevicesJSON)
    if err != nil {
        return nil, err
    }

    // log.Printf("deviceOrderJSON: %s", deviceOrderJSON)
    // log.Printf("hiddenDevicesJSON: %s", hiddenDevicesJSON)

    if err := json.Unmarshal(deviceOrderJSON, &prefs.DeviceOrder); err != nil {
        return nil, err
    }
    if err := json.Unmarshal(hiddenDevicesJSON, &prefs.HiddenDevices); err != nil {
        return nil, err
    }

    rows, err := r.DB.Query("SELECT device_id, photo_blob FROM device_photo WHERE user_id = ?", userID)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    prefs.DevicePhotos = make(map[string][]byte)

    for rows.Next() {
        var deviceID string
        var photoBlob []byte
        if err := rows.Scan(&deviceID, &photoBlob); err != nil {
            log.Printf("Error scanning device_photo row: %v", err)
            continue
        }
        prefs.DevicePhotos[deviceID] = photoBlob
    }

    return prefs, nil
}