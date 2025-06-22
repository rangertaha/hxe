interface APIError {
  error: string;
}

export class BaseClient {
  protected baseURL: string;
  protected defaultHeaders: HeadersInit = {
    "Content-Type": "application/json",
    // You can add auth headers here
  };

  constructor(baseURL: string) {
    this.baseURL = baseURL;
  }

  protected async doRequest<T>(
    method: string,
    path: string,
    body?: unknown
  ): Promise<T> {
    const response = await fetch(`${this.baseURL}${path}`, {
      method,
      headers: this.defaultHeaders,
      body: body ? JSON.stringify(body) : undefined,
    });

    if (!response.ok) {
      const error = (await response.json()) as APIError;
      throw new Error(error.error || `HTTP error: ${response.status}`);
    }

    if (response.status === 204) {
      return {} as T;
    }

    return response.json() as Promise<T>;
  }

  protected async get<T>(path: string): Promise<T> {
    return this.doRequest<T>("GET", path);
  }

  protected async post<T>(path: string, body: unknown): Promise<T> {
    return this.doRequest<T>("POST", path, body);
  }

  protected async put<T>(path: string, body: unknown): Promise<T> {
    return this.doRequest<T>("PUT", path, body);
  }

  protected async delete(path: string): Promise<void> {
    await this.doRequest<void>("DELETE", path);
  }

  protected async options<T>(path: string): Promise<T> {
    return this.doRequest<T>("OPTIONS", path);
  }
} 