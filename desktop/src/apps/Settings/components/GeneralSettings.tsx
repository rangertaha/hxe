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
} from '@mui/material';
import {
  Language,
  Palette,
  Speed,
} from '@mui/icons-material';

const GeneralSettings: React.FC = () => {
  return (
    <Box>
      <Typography variant="h6" sx={{ mb: 3 }}>
        General Settings
      </Typography>

      <List>
        <ListItem>
          <ListItemIcon>
            <Language />
          </ListItemIcon>
          <ListItemText 
            primary="Language" 
            secondary="Select your preferred language"
          />
          <FormControl size="small" sx={{ minWidth: 120 }}>
            <Select defaultValue="en">
              <MenuItem value="en">English</MenuItem>
              <MenuItem value="es">Español</MenuItem>
              <MenuItem value="fr">Français</MenuItem>
            </Select>
          </FormControl>
        </ListItem>

        <Divider />

        <ListItem>
          <ListItemIcon>
            <Palette />
          </ListItemIcon>
          <ListItemText 
            primary="Theme" 
            secondary="Choose your preferred theme"
          />
          <FormControl size="small" sx={{ minWidth: 120 }}>
            <Select defaultValue="dark">
              <MenuItem value="dark">Dark</MenuItem>
              <MenuItem value="light">Light</MenuItem>
              <MenuItem value="system">System</MenuItem>
            </Select>
          </FormControl>
        </ListItem>

        <Divider />

        <ListItem>
          <ListItemIcon>
            <Speed />
          </ListItemIcon>
          <ListItemText 
            primary="Performance Mode" 
            secondary="Optimize for performance or battery life"
          />
          <FormControl size="small" sx={{ minWidth: 120 }}>
            <Select defaultValue="balanced">
              <MenuItem value="performance">Performance</MenuItem>
              <MenuItem value="balanced">Balanced</MenuItem>
              <MenuItem value="battery">Battery Saver</MenuItem>
            </Select>
          </FormControl>
        </ListItem>
      </List>
    </Box>
  );
};

export default GeneralSettings; 