export interface Device {
    id: string;
    name: string;
    position: {
      lat: number;
      lng: number;
    };
    activeState: string;
    driveState: string;
    speed: number;
    address: string | null;
    rssi: string;
    hdop: string;
    gpsLev: string;
  }