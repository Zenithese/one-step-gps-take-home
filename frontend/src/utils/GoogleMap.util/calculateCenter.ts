export const calculateCenter = (points: { lat: number; lng: number }[]): { lat: number; lng: number } => {
    const sum = points.reduce(
      (acc, point) => {
        acc.lat += point.lat;
        acc.lng += point.lng;
        return acc;
      },
      { lat: 0, lng: 0 }
    );
  
    return {
      lat: sum.lat / points.length,
      lng: sum.lng / points.length,
    };
  };