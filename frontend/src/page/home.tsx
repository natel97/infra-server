import { useEffect } from "react";
import { useDispatch, useSelector } from "react-redux";
import { API, GetDeploymentsResponse } from "../api";
import { Link, useNavigate } from "react-router-dom";
import { setDeployments } from "../redux/deployment.reducer";
import { GlowingListItem } from "../components/specific";
import { Button } from "../components/generic";

type ExistingSitesProps = {
  deployments: GetDeploymentsResponse[];
};

const ExistingSites = ({
  deployments = [],
}: ExistingSitesProps): JSX.Element => {
  const navigate = useNavigate();

  return (
    <div style={{ display: "flex", flexDirection: "column" }}>
      {deployments.map((value) => (
        <GlowingListItem
          glow="green"
          onClick={() => navigate(`/deployment/${value.id}`)}
        >
          <h2>{value.name}</h2>
          <div>{value.type}</div>
        </GlowingListItem>
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
      <Link to="/new">
        <Button>Add Deployment</Button>
      </Link>
      <ExistingSites deployments={deployments} />
    </div>
  );
};

export default Deployment;
