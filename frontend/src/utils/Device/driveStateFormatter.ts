export const driveStateFormatter = (driveStatus: string) => {
    if (driveStatus === 'off') {
        return 'Ignition off';
    }
    return driveStatus.charAt(0).toUpperCase() + driveStatus.slice(1);
}