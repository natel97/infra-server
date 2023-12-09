import { BrowserRouter, Link, Route, Routes } from "react-router-dom";
import { Provider, useDispatch } from "react-redux";
import store from "./redux/store";
import { Suspense, lazy, useEffect } from "react";
import { API } from "./api";
import { setDomains } from "./redux/deployment.reducer";

const Deployments = lazy(() => import("./page/home"));
const NewDeployment = lazy(() => import("./page/new"));
const DeploymentDetails = lazy(() => import("./page/deployment/id"));

const NotFoundComponent = () => {
  return (
    <h1>
      Path not found :/ <Link to="/">Go back home</Link>
    </h1>
  );
};

const App = () => {
  return (
    <div>
      <Provider store={store}>
        <GenerateInitialState />
        <BrowserRouter>
          <Routes>
            <Route
              path="/"
              element={
                <Suspense>
                  <Deployments />
                </Suspense>
              }
            />
            <Route
              path="/new/*"
              element={
                <Suspense>
                  <NewDeployment />
                </Suspense>
              }
            />
            <Route
              path="/deployment/:id"
              element={
                <Suspense>
                  <DeploymentDetails />
                </Suspense>
              }
            />
            <Route path="*" Component={NotFoundComponent} />
          </Routes>
        </BrowserRouter>
      </Provider>
    </div>
  );
};

const GenerateInitialState = () => {
  const dispatch = useDispatch();
  useEffect(() => {
    if (!dispatch) return;

    API.getDomains().then((domains) => dispatch(setDomains(domains)));
  }, [dispatch]);
  return null;
};

export default App;
