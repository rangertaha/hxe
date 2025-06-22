import React, { useState, useEffect } from "react";
import { useParams } from "react-router-dom";
import {
  Box,
  Typography,
  IconButton,
  Divider,
  Stack,
  Chip,
  Grid,
  List,
  ListItem,
  ListItemText,
  ListItemIcon,
  Tabs,
  Tab,
  Paper,
  LinearProgress,
  Button,
} from "@mui/material";
import {
  Memory,
  Speed,
  Schedule,
  Update,
  Code,
  Description,
  PlayArrow,
  Stop,
  Terminal,
  MonitorHeart,
  ShowChart,
  Delete,
  Edit,
  ContentCopy,
  Download,
  Pause,
  Info,
  Settings,
  AutoGraph,
  Refresh,
  Storage,
  Timer,
  BugReport,
} from "@mui/icons-material";
import Client from "../../client";
import { Service } from "../../client/models";


interface TabPanelProps {
  children?: React.ReactNode;
  index: number;
  value: number;
}

const TabPanel = (props: TabPanelProps) => {
  const { children, value, index, ...other } = props;

  return (
    <div
      role="tabpanel"
      hidden={value !== index}
      id={`service-tabpanel-${index}`}
      aria-labelledby={`service-tab-${index}`}
      {...other}
    >
      {value === index && (
        <Box sx={{ p: 3 }}>
          {children}
        </Box>
      )}
    </div>
  );
};

const ServiceDetail: React.FC = () => {
  const { id } = useParams<{ id: string }>();
  const [activeTab, setActiveTab] = useState(0);
  const [service, setService] = useState<Service | null>(null);
  const [apiClient] = useState(() => new Client("http://localhost:8080"));

  useEffect(() => {
    const fetchService = async () => {
      if (id) {
        try {
          const result = await apiClient.services.getService(parseInt(id, 10));
          setService(result);
        } catch (error) {
          console.error('Failed to fetch service:', error);
        }
      }
    };

    fetchService();
  }, [id, apiClient]);

  const handleTabChange = (event: React.SyntheticEvent, newValue: number) => {
    setActiveTab(newValue);
  };

  if (!service) {
    return (
      <Box sx={{ p: 3, textAlign: 'center' }}>
        <Typography>Loading service details...</Typography>
      </Box>
    );
  }

  const getMetric = (key: string): number => {
    return service.metrics?.[key] || 0;
  };

  const RuntimeTab = () => (
    <Grid container spacing={3}>
      <Grid item xs={12} md={6}>
        <Paper sx={{ p: 2 }}>
          <Typography variant="h6" gutterBottom>
            Service Status
          </Typography>
          <List>
            <ListItem>
              <ListItemIcon>
                <Speed />
              </ListItemIcon>
              <ListItemText 
                primary="Status" 
                secondary={
                  <Chip 
                    label={service.state} 
                    color={
                      service.state === 'RUNNING' ? 'success' :
                      service.state === 'FAILED' ? 'error' :
                      service.state === 'TIMEOUT' ? 'warning' :
                      'default'
                    }
                    size="small"
                  />
                }
              />
            </ListItem>
            <ListItem>
              <ListItemIcon>
                <Timer />
              </ListItemIcon>
              <ListItemText 
                primary="Uptime" 
                secondary={service.uptime || 'N/A'}
              />
            </ListItem>
            <ListItem>
              <ListItemIcon>
                <Memory />
              </ListItemIcon>
              <ListItemText 
                primary="Memory Usage" 
                secondary={`${getMetric('memory')} MB`}
              />
            </ListItem>
            <ListItem>
              <ListItemIcon>
                <Storage />
              </ListItemIcon>
              <ListItemText 
                primary="CPU Usage" 
                secondary={`${getMetric('cpu')}%`}
              />
            </ListItem>
          </List>
        </Paper>
      </Grid>
      <Grid item xs={12} md={6}>
        <Paper sx={{ p: 2 }}>
          <Typography variant="h6" gutterBottom>
            Service Actions
          </Typography>
          <Stack direction="row" spacing={2} sx={{ mb: 2 }}>
            <Button
              variant="contained"
              color="primary"
              startIcon={<PlayArrow />}
              disabled={service.state === 'RUNNING'}
            >
              Start
            </Button>
            <Button
              variant="contained"
              color="error"
              startIcon={<Stop />}
              disabled={service.state !== 'RUNNING'}
            >
              Stop
            </Button>
            <Button
              variant="outlined"
              startIcon={<Refresh />}
            >
              Restart
            </Button>
          </Stack>
          <Stack direction="row" spacing={2}>
            <Button
              variant="outlined"
              startIcon={<Edit />}
            >
              Edit
            </Button>
            <Button
              variant="outlined"
              color="error"
              startIcon={<Delete />}
            >
              Delete
            </Button>
          </Stack>
        </Paper>
      </Grid>
    </Grid>
  );

  const ConfigurationTab = () => (
    <Grid container spacing={3}>
      <Grid item xs={12} md={6}>
        <Paper sx={{ p: 2 }}>
          <Typography variant="h6" gutterBottom>
            Basic Information
          </Typography>
          <List>
            <ListItem>
              <ListItemIcon>
                <Info />
              </ListItemIcon>
              <ListItemText 
                primary="Name" 
                secondary={service.name}
              />
            </ListItem>
            <ListItem>
              <ListItemIcon>
                <Code />
              </ListItemIcon>
              <ListItemText 
                primary="Type" 
                secondary={service.exec ? (service.exec.startsWith('python') ? 'Python' :
                          service.exec.startsWith('node') ? 'Node.js' : 'Shell') : 'N/A'}
              />
            </ListItem>
            <ListItem>
              <ListItemIcon>
                <Settings />
              </ListItemIcon>
              <ListItemText 
                primary="Working Directory" 
                secondary={service.dir}
              />
            </ListItem>
          </List>
        </Paper>
      </Grid>
      <Grid item xs={12} md={6}>
        <Paper sx={{ p: 2 }}>
          <Typography variant="h6" gutterBottom>
            Execution Details
          </Typography>
          <List>
            <ListItem>
              <ListItemIcon>
                <PlayArrow />
              </ListItemIcon>
              <ListItemText 
                primary="Command" 
                secondary={service.exec}
              />
            </ListItem>
            <ListItem>
              <ListItemIcon>
                <BugReport />
              </ListItemIcon>
              <ListItemText 
                primary="Stop Signal" 
                secondary={service.stopSignal}
              />
            </ListItem>
            <ListItem>
              <ListItemIcon>
                <Timer />
              </ListItemIcon>
              <ListItemText 
                primary="Timeout" 
                secondary={`${service.retries || 0} seconds`}
              />
            </ListItem>
          </List>
        </Paper>
      </Grid>
    </Grid>
  );

  const MetricsTab = () => (
    <Grid container spacing={3}>
      <Grid item xs={12}>
        <Paper sx={{ p: 2 }}>
          <Typography variant="h6" gutterBottom>
            Performance Metrics
          </Typography>
          <Grid container spacing={2}>
            <Grid item xs={12} md={4}>
              <Paper sx={{ p: 2, textAlign: 'center' }}>
                <Typography variant="subtitle2" color="text.secondary">
                  Memory Usage
                </Typography>
                <Typography variant="h4">
                  {getMetric('memory')} MB
                </Typography>
              </Paper>
            </Grid>
            <Grid item xs={12} md={4}>
              <Paper sx={{ p: 2, textAlign: 'center' }}>
                <Typography variant="subtitle2" color="text.secondary">
                  CPU Usage
                </Typography>
                <Typography variant="h4">
                  {getMetric('cpu')}%
                </Typography>
              </Paper>
            </Grid>
            <Grid item xs={12} md={4}>
              <Paper sx={{ p: 2, textAlign: 'center' }}>
                <Typography variant="subtitle2" color="text.secondary">
                  Uptime
                </Typography>
                <Typography variant="h4">
                  {service.uptime || 'N/A'}
                </Typography>
              </Paper>
            </Grid>
          </Grid>
        </Paper>
      </Grid>
    </Grid>
  );

  return (
    <Box sx={{ height: '100%', display: 'flex', flexDirection: 'column' }}>
      <Paper sx={{ flex: 1, display: 'flex', flexDirection: 'column', overflow: 'hidden' }}>
        <Tabs
          value={activeTab}
          onChange={handleTabChange}
          variant="scrollable"
          scrollButtons="auto"
          sx={{ borderBottom: 1, borderColor: 'divider' }}
        >
          <Tab icon={<AutoGraph />} label="Runtime" />
          <Tab icon={<Settings />} label="Configuration" />
          <Tab icon={<Speed />} label="Metrics" />
        </Tabs>

        <Box sx={{ flex: 1, overflow: 'auto' }}>
          <TabPanel value={activeTab} index={0}>
            <RuntimeTab />
          </TabPanel>
          <TabPanel value={activeTab} index={1}>
            <ConfigurationTab />
          </TabPanel>
          <TabPanel value={activeTab} index={2}>
            <MetricsTab />
          </TabPanel>
        </Box>
      </Paper>
    </Box>
  );
};

export default ServiceDetail;