export const rssiInterpreter = (rssi: number) => {
    if (rssi >= -70) {
        return 'excellent';
    }
    if (rssi >= -85) {
        return 'normal';
    }
    if (rssi >= -100) {
        return 'poor';
    }
    return 'none';
}