package repositories

import (
    "database/sql"
    "encoding/json"
    "log"
    "errors"
    "backend/internal/models"
)

type PreferenceRepository struct {
    DB *sql.DB
}

func NewPreferenceRepository(db *sql.DB) *PreferenceRepository {
    return &PreferenceRepository{DB: db}
}

func (r *PreferenceRepository) SavePreferences(prefs *models.Preference) error {
    tx, err := r.DB.Begin()
    if err != nil {
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
        query += " WHERE id = 1"

        _, err := tx.Exec(query, args...)
        if err != nil {
            tx.Rollback()
            return err
        }
    }

    if prefs.DevicePhotos != nil {
        for deviceID, photo := range prefs.DevicePhotos {
            _, err := tx.Exec(`
                INSERT INTO device_photo (device_id, photo_blob)
                VALUES (?, ?)
                ON DUPLICATE KEY UPDATE
                photo_blob = VALUES(photo_blob)
            `, deviceID, photo)
            if err != nil {
                tx.Rollback()
                return err
            }
        }
    }

    return tx.Commit()
}

func (r *PreferenceRepository) GetPreferences() (*models.Preference, error) {
    prefs := &models.Preference{}

    row := r.DB.QueryRow("SELECT device_order, hidden_devices FROM preference WHERE id = 1")
    var deviceOrderJSON, hiddenDevicesJSON []byte
    err := row.Scan(&deviceOrderJSON, &hiddenDevicesJSON)
    if err != nil {
        if errors.Is(err, sql.ErrNoRows) {
            return prefs, nil
        }
        return nil, err
    }

    if err := json.Unmarshal(deviceOrderJSON, &prefs.DeviceOrder); err != nil {
        return nil, err
    }
    if err := json.Unmarshal(hiddenDevicesJSON, &prefs.HiddenDevices); err != nil {
        return nil, err
    }

    rows, err := r.DB.Query("SELECT device_id, photo_blob FROM device_photo")
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    prefs.DevicePhotos = make(map[string][]byte)
    for rows.Next() {
        var deviceID string
        var photoBlob []byte
        if err := rows.Scan(&deviceID, &photoBlob); err != nil {
            log.Println(err)
            continue
        }
        prefs.DevicePhotos[deviceID] = photoBlob
    }

    return prefs, nil
}