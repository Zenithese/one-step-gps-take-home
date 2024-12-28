export function convertToSnakeCase<T>(obj: T): T {
    if (Array.isArray(obj)) {
      return obj.map((v) => convertToSnakeCase(v)) as unknown as T;
    } else if (obj !== null && typeof obj === 'object') {
      return Object.keys(obj).reduce((result, key) => {
        const snakeKey = key.replace(/[A-Z]/g, (letter) => `_${letter.toLowerCase()}`);
        (result as Record<string, unknown>)[snakeKey] = convertToSnakeCase(
          (obj as Record<string, unknown>)[key]
        );
        return result;
      }, {} as Record<string, unknown>) as T;
    }
    return obj;
  }