export const fitMapToMarkers = (map: google.maps.Map, markers: { lat: number; lng: number }[]) => {
    const bounds = new google.maps.LatLngBounds();
    markers.forEach((marker) => bounds.extend(marker));
    map.fitBounds(bounds);
  };