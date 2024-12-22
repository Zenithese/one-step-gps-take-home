export interface DeviceResultListItem {
    device_id: string;
    display_name: string;
    active_status: string;
    latest_device_point: {
        device_state: {
            drive_status: string;
            lat: number;
            lng: number;
            adjusted_lat_lng: {
                lat: number;
                lng: number;
            };
        };
    };
  }