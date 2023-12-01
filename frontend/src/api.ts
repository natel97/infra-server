export type GetDeploymentsResponse = {
  id: string;
  name: string;
  type: string;
  visibility: string;
  environments?: EnvironmentStub[];
};

export type EnvironmentStub = {
  id: string;
  name: string;
  status: string;
  lastDeploy: string;
};

export type GetSingleDeploymentResponse = {
  id: string;
  name: string;
  type: string;
  visibility: string;
  environments?: EnvironmentStub[];
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

  getDeploymentByID(id: string): Promise<GetSingleDeploymentResponse> {
    return this.fetchHandler(`/api/v1/service/${id}`).then((val) => val.json());
  }

  createEnvironment(id: string, name: string): Promise<void> {
    return this.fetchHandler(`/api/v1/service/${id}/environment`, {
      method: "POST",
      body: JSON.stringify({ name }),
    }).then((val) => val.json());
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
