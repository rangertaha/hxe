import React from 'react';
import { AppBar, Toolbar, IconButton, Typography, Box } from '@mui/material';
import { Menu as MenuIcon, AutoGraph as QuantIcon } from '@mui/icons-material';

interface ToolMenuProps {
  onMenuClick: () => void;
  title?: string;
}

const ToolMenu: React.FC<ToolMenuProps> = ({ onMenuClick, title = 'hxe' }) => {
  return (
    <AppBar position="static" color="transparent" elevation={0}>
    <Toolbar>
      <IconButton
        edge="start"
        color="inherit"
        aria-label="menu"
        onClick={onMenuClick}
        sx={{ mr: 2 }}
      >
        <MenuIcon />
      </IconButton>
      <Box sx={{ display: 'flex', alignItems: 'center', gap: 1 }}>
        <QuantIcon color="primary" />
        <Typography variant="h6" component="div" sx={{ fontWeight: 500 }}>
          {title}
        </Typography>
      </Box>
    </Toolbar>
  </AppBar>
  );
};

export default ToolMenu; 