export async function changeSvgColor(svgPath: string, color: string): Promise<string | null> {
    try {
      const response = await fetch(svgPath);
      const svgString = await response.text();
  
      const parser = new DOMParser();
      const svgDoc = parser.parseFromString(svgString, "image/svg+xml");
  
      const svgElement = svgDoc.querySelector("svg");
      if (svgElement) {
        svgElement.setAttribute("fill", color);
  
        const serializer = new XMLSerializer();
        return serializer.serializeToString(svgElement);
      } else {
        console.error("SVG element not found in the fetched SVG document.");
        return null;
      }
    } catch (error) {
      console.error("Error fetching or applying color to SVG:", error);
      return null;
    }
  }