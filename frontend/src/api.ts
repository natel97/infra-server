export type GetDeploymentsResponse = {
  id: string;
  name: string;
  type: string;
  visibility: string;
  environments?: string[];
};

export type CreateDeploymentRequest = {
  name: string;
  type: string;
};

export class APIHandler {
  constructor(
    private fetchHandler: typeof fetch = (...args) => fetch(...args)
  ) {
    // Allows for a mock fetch handler
  }

  getDeployments(): Promise<GetDeploymentsResponse[]> {
    return this.fetchHandler("/api/v1/service").then((val) => val.json());
  }

  createDeployment(
    body: CreateDeploymentRequest
  ): Promise<GetDeploymentsResponse> {
    return this.fetchHandler("/api/v1/service", {
      body: JSON.stringify(body),
      method: "POST",
    }).then((val) => val.json());
  }
}

export const API = new APIHandler();
