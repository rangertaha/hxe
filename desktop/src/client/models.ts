export interface APIError {
  error: string;
}

export interface Service {
  id: number;
  name: string;
  description: string;
  created_at: string;
  updated_at: string;
  // Fields from detail.tsx
  metrics?: { [key: string]: number };
  state?: string;
  uptime?: string;
  exec?: string;
  dir?: string;
  stopSignal?: string;
  retries?: number;
}
