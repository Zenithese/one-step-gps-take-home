export const convertKmhToMph = (kmh: number): number => {
  const conversionFactor = 0.621371;
  return kmh * conversionFactor;
};
