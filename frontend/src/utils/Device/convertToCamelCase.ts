export function convertToCamelCase<T>(obj: T): T {
    if (Array.isArray(obj)) {
      return obj.map((v) => convertToCamelCase(v)) as unknown as T;
    } else if (obj !== null && obj?.constructor === Object) {
      return Object.keys(obj).reduce((result, key) => {
        const camelKey = key.replace(/_([a-z])/g, (g) => g[1].toUpperCase());
        (result as Record<string, unknown>)[camelKey] = convertToCamelCase(
          (obj as Record<string, unknown>)[key]
        );
        return result;
      }, {} as Record<string, unknown>) as T;
    }
    return obj;
  }