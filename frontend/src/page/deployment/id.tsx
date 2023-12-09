import Grid from "@mui/system/Unstable_Grid";
import { useEffect, useState } from "react";
import { useParams } from "react-router-dom";
import { GlowingListItem } from "../../components/specific";
import { Box } from "@mui/system";
import { API, GetDomainResponse } from "../../api";
import { Button } from "../../components/generic";
import styled from "styled-components";
import { useSelector } from "react-redux";

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
  position: fixed;
  bottom: 0;
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

const NewDeploymentSettings = ({ service, open, setOpen }) => {
  const [environment, setEnvironment] = useState(service?.environments[0]?.id);
  const [file, setFile] = useState<File | null>(null);

  if (!open) {
    return null;
  }

  const onSubmit = () => {
    if (file === null) return;
    API.deployZip({ environment, service: service.id, file }).then(() => {
      setOpen(false);
      setFile(null);
    });
  };

  return (
    <Modal title="New Deployment" open={open} setOpen={setOpen}>
      <div style={{ display: "flex", flexDirection: "column" }}>
        <input
          type="file"
          accept=".zip"
          onChange={(e) => setFile(e.target.files![0])}
        />
        <select
          onChange={(e) => setEnvironment(e.target.value)}
          value={environment}
        >
          {service.environments.map((environment) => (
            <option value={environment.id}>{environment.name}</option>
          ))}
        </select>
        <Button onClick={onSubmit} background="green">
          Create
        </Button>
      </div>
    </Modal>
  );
};

const StaticSettings = ({ service }) => {
  const [createDeployment, setCreateDeployment] = useState(false);
  return (
    <div>
      <NewDeploymentSettings
        service={service}
        open={createDeployment}
        setOpen={setCreateDeployment}
      />
      <Button onClick={() => setCreateDeployment(true)}>
        Create Deployment
      </Button>
    </div>
  );
};

const SpecificSettings = {
  "static-website": StaticSettings,
};

const SpecificDeploymentSettings = ({ service }) => {
  const CustomSettings = SpecificSettings[service.deploymentSettings.type];

  return (
    <GlowingListItem glow="gray">
      <h3>{service.deploymentSettings.type} Settings</h3>
      <DeploymentSettings settings={service.deploymentSettings} />
      <CustomSettings service={service} />
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

const NewDomainModal = ({ open, setOpen, deploymentID, environments }) => {
  const domains = useSelector((state) => state?.deployments?.domains);
  const [subdomain, setSubdomain] = useState("");
  const [environment, setEnvironment] = useState(environments[0]?.id);
  const [domain, setDomain] = useState(domains[0]?.id);
  useEffect(() => {
    setDomain(domains[0]?.id);
  }, [domains]);

  if (domains.length === 0) {
    return <div>No domains available :/</div>;
  }

  const createURL = () => {
    API.createURL(deploymentID, {
      domainId: domain,
      environmentId: environment,
      name: "todo: do",
      subdomain,
    }).then(() => {
      setSubdomain("");
    });
  };

  return (
    <Modal title="New Domain" open={open} setOpen={setOpen}>
      <div>
        <label htmlFor="subdomain">Subdomain</label>
        <input
          value={subdomain}
          onChange={(e) => setSubdomain(e.target.value)}
          id="subdomain"
          type="text"
        />
      </div>
      <div>
        Domain:{" "}
        <select onChange={(e) => setDomain(e.target.value)}>
          {domains.map((domain: GetDomainResponse) => (
            <option value={domain.id}>{domain.url}</option>
          ))}
        </select>
      </div>
      <div>
        New URL: {subdomain || "<web>"}.
        {domains.find((d) => d.id === domain).url}
      </div>
      <div>
        Environment:{" "}
        <select
          value={environment}
          onChange={(e) => setEnvironment(e.target.value)}
        >
          {environments.map((env) => (
            <option value={env.id}>{env.name}</option>
          ))}
        </select>
      </div>
      <Button onClick={() => createURL()} background="green">
        Create Domain
      </Button>
    </Modal>
  );
};

const Domains = ({ service }) => {
  const [newDomainOpen, setNewDomainOpen] = useState(false);
  return (
    <GlowingListItem glow="gray">
      <h3>Domains</h3>
      <hr />
      <Button onClick={() => setNewDomainOpen(true)}>New Domain</Button>
      <NewDomainModal
        deploymentID={service.id}
        open={newDomainOpen}
        setOpen={setNewDomainOpen}
        environments={service.environments}
      />
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
