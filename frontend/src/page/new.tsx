import { useState } from "react";
import { useDispatch } from "react-redux";
import { useNavigate } from "react-router-dom";
import { API } from "../api";
import { addDeployment } from "../redux/deployment.reducer";

const CreateDeploymentsForm = () => {
  const [name, setName] = useState("");
  const [type, setType] = useState("static-website");
  const [error, setError] = useState("");
  const dispatch = useDispatch();
  const navigate = useNavigate();

  const submitForm = () => {
    if (!name || !type) {
      return;
    }

    API.createDeployment({ name, type })
      .then((response) => {
        dispatch(addDeployment(response));
        navigate("/");
      })
      .catch((e: unknown) => {
        setError(JSON.stringify(e));
      });
  };

  return (
    <div
      style={{
        display: "flex",
        flexDirection: "column",
      }}
    >
      <h1>Add New Deployment</h1>
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

export default CreateDeploymentsForm;
