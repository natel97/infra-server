import { useEffect, useState } from "react";
import "./App.css";
import { API, GetDeploymentsResponse } from "./api";

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

const CreateDeploymentsForm = ({ addDeployment }) => {
  const [name, setName] = useState("");
  const [type, setType] = useState("static-website");
  const [error, setError] = useState("");

  const submitForm = () => {
    if (!name || !type) {
      return;
    }

    API.createDeployment({ name, type })
      .then((response) => {
        setName("");
        setError("");
        addDeployment(response);
      })
      .catch((e: any) => {
        setError(JSON.stringify(e));
      });
  };
  return (
    <div>
      <input
        onChange={(e) => setName(e.target.value)}
        type="text"
        value={name}
        placeholder="Name"
      />
      <select value={type} onChange={(e) => setType(e.target.value)}>
        <option value="static-website" label="Static Website" />
        <option value="kubernetes-deployment" label="Kubernetes Deployment" />
      </select>
      <input type="button" value="Add Service" onClick={() => submitForm()} />
      {error && <p>{error}</p>}
    </div>
  );
};

function App() {
  const [deployments, setDeployments] = useState<GetDeploymentsResponse[]>([]);
  useEffect(() => {
    API.getDeployments().then((deployments) => setDeployments(deployments));
  }, []);

  return (
    <div>
      <ExistingSites deployments={deployments} />
      <CreateDeploymentsForm
        addDeployment={(deployment: GetDeploymentsResponse) =>
          setDeployments([...deployments, deployment])
        }
      />
    </div>
  );
}

export default App;
