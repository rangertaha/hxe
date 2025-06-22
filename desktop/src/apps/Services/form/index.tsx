import React, { useState } from "react";
import { useNavigate } from "react-router-dom";
import { 
  Box, 
  Typography, 
  AppBar, 
  Toolbar, 
  IconButton, 
  Menu, 
  MenuItem,
  Tabs,
  Tab,
  Paper
} from "@mui/material";
import { 
  ArrowBack, 
  Code as CodeIcon,
  Info,
  PlayArrow,
  Memory,
  VpnKey,
  Add,
  Settings
} from "@mui/icons-material";
import ServiceActions from "./actions";
import ServiceRuntime from "./runtime";
import ServiceVariables from "./variables";


interface TabPanelProps {
  children?: React.ReactNode;
  index: number;
  value: number;
}

const TabPanel = (props: TabPanelProps) => {
  const { children, value, index, ...other } = props;
  return (
    <div role="tabpanel" hidden={value !== index} {...other}>
      {value === index && <Box sx={{ p: 2 }}>{children}</Box>}
    </div>
  );
};

const ServiceForm: React.FC = () => {
  const navigate = useNavigate();
  const [anchorEl, setAnchorEl] = useState<null | HTMLElement>(null);
  const [activeTab, setActiveTab] = useState(0);


  const handleTabChange = (event: React.SyntheticEvent, newValue: number) => {
    setActiveTab(newValue);
  };

  const handleSubmit = (event: React.FormEvent) => {
    event.preventDefault();
    // TODO: Implement service creation
    navigate("/services");
  };

  const handleCancel = () => {
    navigate("/services");
  };

  return (
    <Box
      sx={{
        display: "flex",
        flexDirection: "column",
        height: "100%",
      }}
    >
      <AppBar position="static" color="default" elevation={1} sx={{ mb: 0 }}>
        <Toolbar variant="dense" sx={{ minHeight: 40, p: 0 }}>
          <IconButton size="small" onClick={handleCancel} sx={{ mr: 1 }}>
            <ArrowBack fontSize="small" />
          </IconButton>
          <Typography variant="subtitle2" sx={{ mr: 2, fontWeight: 600 }}>
            New Service
          </Typography>
          <Box sx={{ flexGrow: 1 }} />

        </Toolbar>
      </AppBar>

      <Paper sx={{ flex: 1, display: 'flex', flexDirection: 'column', overflow: 'hidden' }}>
        <Tabs
          value={activeTab}
          onChange={handleTabChange}
          variant="scrollable"
          scrollButtons="auto"
          sx={{ borderBottom: 1, borderColor: 'divider' }}
        >
          <Tab icon={<Memory />} label="Runtime" />
          <Tab icon={<VpnKey />} label="Veriables" />
          <Tab icon={<Add />} label="Actions" />

        </Tabs>

        <Box sx={{ flex: 1, overflow: 'auto' }}>
          <TabPanel value={activeTab} index={0}>
            <ServiceRuntime />
          </TabPanel>
          <TabPanel value={activeTab} index={1}>
            <ServiceVariables />
          </TabPanel>
          <TabPanel value={activeTab} index={2}>
            <ServiceActions />
          </TabPanel>
        </Box>
      </Paper>
    </Box>
  );
};

export default ServiceForm;
