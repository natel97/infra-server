import { createSlice } from "@reduxjs/toolkit";
import { GetDeploymentsResponse, GetDomainResponse } from "../api";

type State = {
  deployments: GetDeploymentsResponse[];
  domains: GetDomainResponse[];
};

const getDefaultState = (): State => ({
  deployments: [],
  domains: [],
});

const deploymentSlice = createSlice({
  name: "deployment",
  initialState: getDefaultState(),
  reducers: {
    addDeployment: (state, action) => {
      console.log({ state });
      state.deployments.push(action.payload);
    },
    setDeployments: (state, action) => {
      state.deployments = action.payload;
    },
    setDomains: (state, action) => {
      state.domains = action.payload;
    },
  },
});

export default deploymentSlice.reducer;

export const { addDeployment, setDeployments, setDomains } =
  deploymentSlice.actions;
