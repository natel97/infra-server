export type GetDeploymentsResponse = {
  id: string;
  name: string;
  type: string;
  visibility: string;
  environments?: EnvironmentStub[];
};

export type GetDomainResponse = {
  url: string;
  id: string;
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

export type DeployZip = {
  file: File;
  environment: string;
  service: string;
};

export type CreateDeploymentRequest = {
  name: string;
  type: string;
};

export type CreateURLRequest = {
  environmentId: string;
  domainId: string;
  subdomain: string;
  name: string;
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

  deployZip(deploy: DeployZip): Promise<void> {
    const data = new FormData();
    data.append("name", deploy.file.name);
    data.append("file", deploy.file);
    data.append("environment", deploy.environment);
    data.append("service", deploy.service);
    return this.fetchHandler("/api/v1/upload", {
      method: "POST",
      body: data,
    }).then((response) => response.json());
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

  getDomains(): Promise<GetDomainResponse[]> {
    return this.fetchHandler("/api/v1/domain").then((val) => val.json());
  }

  createURL(deploymentID: string, body: CreateURLRequest): Promise<void> {
    return this.fetchHandler(`/api/v1/service/${deploymentID}/url`, {
      body: JSON.stringify(body),
      method: "POST",
    }).then((res) => res.json());
  }
}

export const API = new APIHandler();
