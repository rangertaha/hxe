import React, { useState } from 'react';
import {
  Box,
  Drawer,
  List,
  ListItem,
  ListItemIcon,
  ListItemText,
  ListItemButton,
  Typography,
  Divider,
  Paper,
} from '@mui/material';
import {
  Settings as SettingsIcon,
  Storage,
  Security,
  Notifications,
  Code,
  Cloud,
  Build,
} from '@mui/icons-material';
import GeneralSettings from './components/GeneralSettings';
import StorageSettings from './components/StorageSettings';

interface SettingsPageProps {
  children?: React.ReactNode;
}

const drawerWidth = 240;

const menuItems = [
  { id: 'general', label: 'General', icon: <SettingsIcon /> },
  { id: 'storage', label: 'Storage', icon: <Storage /> },
  { id: 'security', label: 'Security', icon: <Security /> },
  { id: 'notifications', label: 'Notifications', icon: <Notifications /> },
  { id: 'integrations', label: 'Integrations', icon: <Code /> },
  { id: 'cloud', label: 'Cloud Sync', icon: <Cloud /> },
  { id: 'advanced', label: 'Advanced', icon: <Build /> },
];

const SettingsPage: React.FC<SettingsPageProps> = () => {
  const [selectedItem, setSelectedItem] = useState('general');

  const handleMenuItemClick = (id: string) => {
    setSelectedItem(id);
  };

  const renderContent = () => {
    switch (selectedItem) {
      case 'general':
        return <GeneralSettings />;
      case 'storage':
        return <StorageSettings />;
      case 'security':
        return <Typography>Security Settings</Typography>;
      case 'notifications':
        return <Typography>Notification Settings</Typography>;
      case 'integrations':
        return <Typography>Integration Settings</Typography>;
      case 'cloud':
        return <Typography>Cloud Sync Settings</Typography>;
      case 'advanced':
        return <Typography>Advanced Settings</Typography>;
      default:
        return <Typography>Select a setting</Typography>;
    }
  };

  return (
    <Box sx={{ 
      display: 'flex', 
      height: '100%',
      overflow: 'hidden',
      position: 'relative',
    }}>
      <Drawer
        variant="permanent"
        sx={{
          width: drawerWidth,
          flexShrink: 0,
          '& .MuiDrawer-paper': {
            width: drawerWidth,
            boxSizing: 'border-box',
            borderRight: '1px solid',
            borderColor: 'divider',
            position: 'relative',
            height: '100%',
            overflow: 'hidden',
          },
        }}
      >
        <Box sx={{ p: 2 }}>
          <Typography variant="h6" sx={{ fontWeight: 600 }}>
            Settings
          </Typography>
        </Box>
        <Divider />
        <List sx={{ overflow: 'auto' }}>
          {menuItems.map((item) => (
            <ListItem key={item.id} disablePadding>
              <ListItemButton
                selected={selectedItem === item.id}
                onClick={() => handleMenuItemClick(item.id)}
              >
                <ListItemIcon>{item.icon}</ListItemIcon>
                <ListItemText primary={item.label} />
              </ListItemButton>
            </ListItem>
          ))}
        </List>
      </Drawer>
      <Box
        component="main"
        sx={{
          flexGrow: 1,
          p: 3,
          bgcolor: 'background.default',
          height: '100%',
          overflow: 'auto',
        }}
      >
        <Paper sx={{ p: 3, height: '100%', overflow: 'auto' }}>
          {renderContent()}
        </Paper>
      </Box>
    </Box>
  );
};

export default SettingsPage; 