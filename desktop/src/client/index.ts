import { ServiceClient } from "./services";


export class Client {
  private baseURL: string;

  constructor(baseURL: string) {
    this.baseURL = baseURL;
  }

  // Use a getter for lazy initialization
  private _services: ServiceClient | undefined;
  public get services(): ServiceClient {
    if (!this._services) {
      this._services = new ServiceClient(this.baseURL);
    }
    return this._services;
  }
}

export default Client;


// Usage exampleimport Client from './client';
// import type { Service } from './client'; // Import types easily

// const apiClient = new Client('http://localhost:8080');

// // Access the services resource client
// const services: Service[] = await apiClient.services.listServices();

// Access another resource client (if it existed)
// const users = await apiClient.users.listUsers();
