import { BrowserRouter, Link, Route, Routes } from "react-router-dom";
import { Provider } from "react-redux";
import store from "./redux/store";
import { Suspense, lazy } from "react";

const Deployments = lazy(() => import("./page/home"));
const NewDeployment = lazy(() => import("./page/new"));

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
        <BrowserRouter>
          <Routes>
            <Route path="/">
              <Suspense>
                <Deployments />
              </Suspense>
            </Route>

            <Route path="/new/*">
              <Suspense>
                <NewDeployment />
              </Suspense>
            </Route>
            <Route path="*" Component={NotFoundComponent} />
          </Routes>
        </BrowserRouter>
      </Provider>
    </div>
  );
};

export default App;
