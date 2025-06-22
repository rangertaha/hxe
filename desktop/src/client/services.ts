import { BaseClient } from "./base";
import { Service } from "./models";

interface ServiceOperations {
  getService(id: number): Promise<Service>;
  listServices(): Promise<Service[]>;
  createService(
    service: Omit<Service, "id" | "created_at" | "updated_at">
  ): Promise<Service>;
  updateService(
    id: number,
    service: Partial<Omit<Service, "id" | "created_at" | "updated_at">>
  ): Promise<Service>;
  deleteService(id: number): Promise<void>;
  getServiceSchema(): Promise<any>;
  startService(id: number): Promise<Service>;
  stopService(id: number): Promise<Service>;
  restartService(id: number): Promise<Service>;
}

export class ServiceClient extends BaseClient implements ServiceOperations {
  async listServices(): Promise<Service[]> {
    return this.get<Service[]>("/api/services");
  }

  async getService(id: number): Promise<Service> {
    return this.get<Service>(`/api/services/${id}`);
  }

  async createService(
    service: Omit<Service, "id" | "created_at" | "updated_at">
  ): Promise<Service> {
    return this.post<Service>("/api/services", service);
  }

  async updateService(
    id: number,
    service: Partial<Omit<Service, "id" | "created_at" | "updated_at">>
  ): Promise<Service> {
    return this.put<Service>(`/api/services/${id}`, service);
  }

  async deleteService(id: number): Promise<void> {
    return this.delete(`/api/services/${id}`);
  }

  async getServiceSchema(): Promise<any> {
    return this.options<any>("/api/services/schema");
  }

  async startService(id: number): Promise<Service> {
    return this.post<Service>(`/api/services/${id}/start`, {});
  }

  async stopService(id: number): Promise<Service> {
    return this.post<Service>(`/api/services/${id}/stop`, {});
  }

  async restartService(id: number): Promise<Service> {
    return this.post<Service>(`/api/services/${id}/restart`, {});
  }
}

export type { Service, ServiceOperations };
export default ServiceClient;
