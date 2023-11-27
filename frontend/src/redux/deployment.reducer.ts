import { createSlice } from "@reduxjs/toolkit";
import { GetDeploymentsResponse } from "../api";

type State = {
  deployments: GetDeploymentsResponse[];
};

const getDefaultState = (): State => ({
  deployments: [],
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
  },
});

export default deploymentSlice.reducer;

export const { addDeployment, setDeployments } = deploymentSlice.actions;
