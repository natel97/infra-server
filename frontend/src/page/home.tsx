import { useEffect } from "react";
import { useDispatch, useSelector } from "react-redux";
import { API, GetDeploymentsResponse } from "../api";
import { Link } from "react-router-dom";
import { setDeployments } from "../redux/deployment.reducer";

type ExistingSitesProps = {
  deployments: GetDeploymentsResponse[];
};

const ExistingSites = ({
  deployments = [],
}: ExistingSitesProps): JSX.Element => {
  return (
    <div style={{ display: "flex", flexDirection: "column" }}>
      {deployments.map((value) => (
        <div
          style={{
            boxShadow: "0px 0px 6px 6px rgba(96, 96, 96, 0.25)",
            borderRadius: "4px",
            padding: "12px",
            margin: "8px",
          }}
        >
          <h2>{value.name}</h2>
          <div>{value.type}</div>
        </div>
      ))}
    </div>
  );
};

const Deployment = () => {
  const dispatch = useDispatch();
  const deployments: any = useSelector<any>(
    (state) => state?.deployments?.deployments
  );

  useEffect(() => {
    if (deployments.length) {
      return;
    }

    API.getDeployments().then((deployments) =>
      dispatch(setDeployments(deployments))
    );
  }, []);

  return (
    <div>
      <Link to="/new">Add Deployment</Link>
      <ExistingSites deployments={deployments} />
    </div>
  );
};

export default Deployment;
