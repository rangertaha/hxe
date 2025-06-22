import React from 'react';
import {
  Box,
  Typography,
  List,
  ListItem,
  ListItemIcon,
  ListItemText,
  FormControl,
  InputLabel,
  Select,
  MenuItem,
  Divider,
  TextField,
  Button,
  Stack,
  Switch,
  FormControlLabel,
} from '@mui/material';
import {
  Storage,
  Delete,
  Backup,
  Restore,
  Dataset,
} from '@mui/icons-material';

const StorageSettings: React.FC = () => {
  return (
    <Box>
      <Typography variant="h6" sx={{ mb: 3 }}>
        Storage Settings
      </Typography>

      <List>
        <ListItem>
          <ListItemIcon>
            <Storage />
          </ListItemIcon>
          <ListItemText 
            primary="Data Retention" 
            secondary="How long to keep historical data"
          />
          <FormControl size="small" sx={{ minWidth: 120 }}>
            <Select defaultValue="30">
              <MenuItem value="7">7 days</MenuItem>
              <MenuItem value="30">30 days</MenuItem>
              <MenuItem value="90">90 days</MenuItem>
              <MenuItem value="365">1 year</MenuItem>
            </Select>
          </FormControl>
        </ListItem>

        <Divider />

        <ListItem>
          <ListItemIcon>
            <Delete />
          </ListItemIcon>
          <ListItemText 
            primary="Cleanup Schedule" 
            secondary="When to perform data cleanup"
          />
          <FormControl size="small" sx={{ minWidth: 120 }}>
            <Select defaultValue="daily">
              <MenuItem value="hourly">Hourly</MenuItem>
              <MenuItem value="daily">Daily</MenuItem>
              <MenuItem value="weekly">Weekly</MenuItem>
            </Select>
          </FormControl>
        </ListItem>

        <Divider />

        <ListItem>
          <ListItemIcon>
            <Storage />
          </ListItemIcon>
          <ListItemText 
            primary="Storage Location" 
            secondary="Where to store application data"
          />
          <TextField 
            size="small" 
            defaultValue="/var/lib/hxe" 
            sx={{ minWidth: 200 }}
          />
        </ListItem>
      </List>

      <Box sx={{ mt: 4 }}>
        <Typography variant="subtitle1" sx={{ mb: 2 }}>
          Backup & Restore
        </Typography>
        <Stack direction="row" spacing={2}>
          <Button
            variant="outlined"
            startIcon={<Backup />}
          >
            Backup Data
          </Button>
          <Button
            variant="outlined"
            startIcon={<Restore />}
          >
            Restore Data
          </Button>
        </Stack>
      </Box>

      <Box sx={{ mt: 4 }}>
        <Typography variant="subtitle1" sx={{ mb: 2, display: 'flex', alignItems: 'center', gap: 1 }}>
          <Dataset /> InfluxDB Configuration
        </Typography>
        <List>
          <ListItem>
            <ListItemIcon>
              <Dataset />
            </ListItemIcon>
            <ListItemText 
              primary="Enable InfluxDB" 
              secondary="Store metrics and time-series data in InfluxDB"
            />
            <FormControlLabel
              control={<Switch defaultChecked />}
              label=""
            />
          </ListItem>

          <Divider />

          <ListItem>
            <ListItemIcon>
              <Dataset />
            </ListItemIcon>
            <ListItemText 
              primary="InfluxDB URL" 
              secondary="Connection URL for InfluxDB server"
            />
            <TextField 
              size="small" 
              defaultValue="http://localhost:8086" 
              sx={{ minWidth: 200 }}
            />
          </ListItem>

          <Divider />

          <ListItem>
            <ListItemIcon>
              <Dataset />
            </ListItemIcon>
            <ListItemText 
              primary="Database Name" 
              secondary="Name of the InfluxDB database"
            />
            <TextField 
              size="small" 
              defaultValue="hxe" 
              sx={{ minWidth: 200 }}
            />
          </ListItem>

          <Divider />

          <ListItem>
            <ListItemIcon>
              <Dataset />
            </ListItemIcon>
            <ListItemText 
              primary="Retention Policy" 
              secondary="How long to keep time-series data"
            />
            <FormControl size="small" sx={{ minWidth: 120 }}>
              <Select defaultValue="30d">
                <MenuItem value="7d">7 days</MenuItem>
                <MenuItem value="30d">30 days</MenuItem>
                <MenuItem value="90d">90 days</MenuItem>
                <MenuItem value="1y">1 year</MenuItem>
              </Select>
            </FormControl>
          </ListItem>

          <Divider />

          <ListItem>
            <ListItemIcon>
              <Dataset />
            </ListItemIcon>
            <ListItemText 
              primary="Batch Size" 
              secondary="Number of points to write in a single batch"
            />
            <TextField 
              size="small" 
              type="number"
              defaultValue="5000" 
              sx={{ minWidth: 120 }}
            />
          </ListItem>

          <Divider />

          <ListItem>
            <ListItemIcon>
              <Dataset />
            </ListItemIcon>
            <ListItemText 
              primary="Write Interval" 
              secondary="How often to write data to InfluxDB"
            />
            <FormControl size="small" sx={{ minWidth: 120 }}>
              <Select defaultValue="10s">
                <MenuItem value="5s">5 seconds</MenuItem>
                <MenuItem value="10s">10 seconds</MenuItem>
                <MenuItem value="30s">30 seconds</MenuItem>
                <MenuItem value="1m">1 minute</MenuItem>
              </Select>
            </FormControl>
          </ListItem>
        </List>
      </Box>
    </Box>
  );
};

export default StorageSettings; 