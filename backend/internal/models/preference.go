package models

type Preference struct {
    DeviceOrder   []string          `json:"deviceOrder"`
    HiddenDevices []string          `json:"hiddenDevices"`
    DevicePhotos  map[string][]byte `json:"devicePhotos"`
}