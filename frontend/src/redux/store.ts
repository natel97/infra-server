import { configureStore } from "@reduxjs/toolkit";
import deploymentReducer from "./deployment.reducer";

const getStore = (preloadedState = {}) => {
  const store = configureStore({
    reducer: {
      deployments: deploymentReducer,
    },
    preloadedState,
  });

  return store;
};

const store = getStore({});

export default store;
