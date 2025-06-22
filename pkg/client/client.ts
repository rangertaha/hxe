// Models
interface Asset {
  id: number;
  name: string;
  symbol: string;
  description: string;
  created_at: string;
  updated_at: string;
}

interface Entity {
  id: number;
  name: string;
  description: string;
  created_at: string;
  updated_at: string;
}

interface APIError {
  error: string;
}

class Client {
  private baseURL: string;
  private defaultHeaders: HeadersInit = {
    'Content-Type': 'application/json',
  };

  constructor(baseURL: string) {
    this.baseURL = baseURL;
  }

  private async doRequest<T>(
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
      const error = await response.json() as APIError;
      throw new Error(error.error || `HTTP error: ${response.status}`);
    }

    return response.json() as Promise<T>;
  }

  // Asset operations
  async listAssets(): Promise<Asset[]> {
    return this.doRequest<Asset[]>('GET', '/api/assets');
  }

  async getAsset(id: number): Promise<Asset> {
    return this.doRequest<Asset>('GET', `/api/assets/${id}`);
  }

  async createAsset(asset: Omit<Asset, 'id' | 'created_at' | 'updated_at'>): Promise<Asset> {
    return this.doRequest<Asset>('POST', '/api/assets', asset);
  }

  async updateAsset(id: number, asset: Partial<Omit<Asset, 'id' | 'created_at' | 'updated_at'>>): Promise<Asset> {
    return this.doRequest<Asset>('PUT', `/api/assets/${id}`, asset);
  }

  async deleteAsset(id: number): Promise<void> {
    return this.doRequest<void>('DELETE', `/api/assets/${id}`);
  }

  // Entity operations
  async listEntities(): Promise<Entity[]> {
    return this.doRequest<Entity[]>('GET', '/api/entities');
  }

  async getEntity(id: number): Promise<Entity> {
    return this.doRequest<Entity>('GET', `/api/entities/${id}`);
  }

  async createEntity(entity: Omit<Entity, 'id' | 'created_at' | 'updated_at'>): Promise<Entity> {
    return this.doRequest<Entity>('POST', '/api/entities', entity);
  }

  async updateEntity(id: number, entity: Partial<Omit<Entity, 'id' | 'created_at' | 'updated_at'>>): Promise<Entity> {
    return this.doRequest<Entity>('PUT', `/api/entities/${id}`, entity);
  }

  async deleteEntity(id: number): Promise<void> {
    return this.doRequest<void>('DELETE', `/api/entities/${id}`);
  }
}

// Example usage:
/*
const client = new Client('http://localhost:8080');

// List assets
const assets = await client.listAssets();

// Create an asset
const newAsset = await client.createAsset({
  name: 'Bitcoin',
  symbol: 'BTC',
  description: 'Digital gold',
});

// Get an asset
const asset = await client.getAsset(newAsset.id);

// Update an asset
const updatedAsset = await client.updateAsset(asset.id, {
  description: 'Updated description',
});

// Delete an asset
await client.deleteAsset(asset.id);
*/

export default Client; 