package models

type Preference struct {
    DeviceOrder   []string          `json:"device_order"`
    HiddenDevices []string          `json:"hidden_devices"`
    DevicePhotos  map[string][]byte `json:"device_photos"`
}