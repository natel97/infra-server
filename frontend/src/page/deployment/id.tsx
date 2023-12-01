import Grid from "@mui/system/Unstable_Grid";
import { useEffect, useState } from "react";
import { useParams } from "react-router-dom";
import { GlowingListItem } from "../../components/specific";
import { Box } from "@mui/system";
import { API } from "../../api";
import { Button } from "../../components/generic";
import styled from "styled-components";

const ListOfEnvironments = ({ environments }) => {
  if (environments === null || environments.length === 0) {
    return <h4>No Environments</h4>;
  }

  return (
    <table style={{ width: "100%" }}>
      <tbody>
        {environments.map((env: any) => (
          <tr>
            <td>{env.name}</td>
            <td>{env.status}</td>
            <td>{env.lastUpdated}</td>
          </tr>
        ))}
      </tbody>
    </table>
  );
};

const Backdrop = styled.div`
  width: 100vw;
  height: 100vh;
  position: absolute;
  top: 0;
  left: 0;
  background: #000000cc;
  display: flex;
  align-items: center;
  justify-content: center;
`;

const Modal = ({ children, title, open, setOpen }) => {
  if (!open) return null;

  return (
    <Backdrop onClick={() => setOpen(false)}>
      <GlowingListItem glow="green" onClick={(e) => e.stopPropagation()}>
        <h2>{title}</h2>
        <hr />
        <div>{children}</div>
      </GlowingListItem>
    </Backdrop>
  );
};

const NewEnvironmentModal = ({ open, setOpen, id, onAdd }) => {
  const [name, setName] = useState("");

  const createEnvironment = () => {
    API.createEnvironment(id, name).then((val) => {
      setName("");
      setOpen(false);
      onAdd(val);
    });
  };

  return (
    <Modal open={open} setOpen={setOpen} title="Add Environment">
      <div>Environment Name</div>
      <input
        type="text"
        onChange={(e) => setName(e.target.value)}
        value={name}
      />
      <Button background="green" onClick={createEnvironment}>
        Create
      </Button>
    </Modal>
  );
};

const Environments = ({ service }) => {
  console.log({ service });
  const [open, setOpen] = useState(false);

  return (
    <GlowingListItem glow="gray">
      <h3>Environments</h3>
      <hr />
      <ListOfEnvironments environments={service.environments} />
      <Button onClick={() => setOpen(true)}>New Environment</Button>
      <NewEnvironmentModal
        id={service.id}
        open={open}
        setOpen={setOpen}
        onAdd={(e) =>
          (service.environments = [...(service.environments || []), e])
        }
      />
    </GlowingListItem>
  );
};

const DeploymentSettings = ({ settings }) => {
  if (settings === null) {
    return <h4>No Settings</h4>;
  }

  return (
    <table style={{ width: "100%" }}>
      <tbody>
        {settings.settings.map((setting) => (
          <tr>
            <td>{setting.name}</td>
            <td>{setting.value}</td>
            <td style={{ cursor: "pointer" }}>Edit</td>
          </tr>
        ))}
      </tbody>
    </table>
  );
};

const SpecificDeploymentSettings = ({ service }) => {
  return (
    <GlowingListItem glow="gray">
      <h3>{service.deploymentSettings.type} Settings</h3>
      <DeploymentSettings settings={service.deploymentSettings} />
    </GlowingListItem>
  );
};

const HistoryLog = ({ history }) => {
  if (history === null || history.length === 0) {
    return <h4>No History</h4>;
  }
};

const History = ({ service }) => {
  return (
    <GlowingListItem glow="gray" style={{ height: "100%" }}>
      <h3>History</h3>
      <hr />
      <HistoryLog history={service.deployHistory} />
      <div />
    </GlowingListItem>
  );
};

const EnvironmentVariables = ({ service }) => {
  return (
    <GlowingListItem glow="gray">
      <h3>EnvironmentVariables</h3>
    </GlowingListItem>
  );
};

const Vault = ({ service }) => {
  return (
    <GlowingListItem glow="gray">
      <h3>Vault</h3>
    </GlowingListItem>
  );
};

const Provisions = ({ service }) => {
  return (
    <GlowingListItem glow="gray">
      <h3>Provisions</h3>
    </GlowingListItem>
  );
};

const Domains = ({ service }) => {
  return (
    <GlowingListItem glow="gray">
      <h3>Domains</h3>
    </GlowingListItem>
  );
};

const DangerZone = ({ service }) => {
  return (
    <GlowingListItem glow="gray">
      <h3>DangerZone</h3>
    </GlowingListItem>
  );
};

const DeploymentID = () => {
  const deploymentID = useParams().id;
  const [deployment, setDeployment] = useState<any>(null);
  useEffect(() => {
    if (!deploymentID) return;
    API.getDeploymentByID(deploymentID).then((deployment) => {
      setDeployment(deployment);
    });
  }, [deploymentID]);

  if (!deployment) {
    return <h2>Loading</h2>;
  }

  return (
    <div>
      <h1>Deployment Details</h1>
      <h2>{deployment.name}</h2>

      <Box>
        <Grid container spacing={2} padding="1rem">
          <Grid xs={6}>
            <Environments service={deployment} />
          </Grid>
          <Grid xs={6}>
            <History service={deployment} />
          </Grid>
          <Grid xs={6}>
            <SpecificDeploymentSettings service={deployment} />
          </Grid>
          <Grid xs={12}>
            <EnvironmentVariables service={deployment} />
          </Grid>
          <Grid xs={12}>
            <Vault service={deployment} />
          </Grid>
          <Grid xs={12}>
            <Provisions service={deployment} />
          </Grid>
          <Grid xs={12}>
            <Domains service={deployment} />
          </Grid>
          <Grid xs={12}>
            <DangerZone service={deployment} />
          </Grid>
        </Grid>
      </Box>
    </div>
  );
};

export default DeploymentID;
