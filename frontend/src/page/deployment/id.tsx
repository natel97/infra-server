import { useEffect } from "react";
import { useSelector } from "react-redux";
import { useParams } from "react-router-dom";

const DeploymentID = () => {
  const deploymentID = useParams().id;
  const matchingItem = useSelector(
    (state: any) => state.deployments.deployments
  ).find((deployment: any) => deployment.id === deploymentID);
  useEffect(() => {
    // TODO: API call?
  }, [deploymentID]);

  return (
    <div>
      Single Deployment Details: {deploymentID} {JSON.stringify(matchingItem)}
    </div>
  );
};

export default DeploymentID;
