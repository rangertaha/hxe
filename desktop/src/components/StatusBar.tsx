import React from 'react';
import { Box, Typography, IconButton, Tooltip } from '@mui/material';
import { 
  Wifi, 
  Memory, 
  Speed, 
  Storage, 
  Security, 
  Notifications 
} from '@mui/icons-material';

interface StatusBarProps {
  onMenuClick: () => void;
}

const StatusBar: React.FC<StatusBarProps> = ({ onMenuClick }) => {
  return (
    <Box
      sx={{
        height: '28px',
        borderTop: '1px solid',
        borderColor: 'divider',
        bgcolor: 'background.paper',
        display: 'flex',
        alignItems: 'center',
        justifyContent: 'space-between',
        px: 2,
        position: 'fixed',
        bottom: 0,
        left: 0,
        right: 0,
        zIndex: 1000,
        fontSize: '0.75rem',
        color: 'text.secondary',
        backdropFilter: 'blur(8px)',
        boxShadow: '0 -1px 4px rgba(0,0,0,0.1)'
      }}
    >
      <Box sx={{ display: 'flex', alignItems: 'center', gap: 2 }}>
        <Box sx={{ display: 'flex', alignItems: 'center', gap: 0.5 }}>
          <Box
            sx={{
              width: 8,
              height: 8,
              borderRadius: '50%',
              bgcolor: 'success.main',
              mr: 0.5,
              boxShadow: '0 0 8px rgba(76, 175, 80, 0.5)'
            }}
          />
          System Online
        </Box>
        <Box>Last Update: {new Date().toLocaleTimeString()}</Box>
        <Box>API: Connected</Box>
        <Box>Version: 1.0.0</Box>
      </Box>

      <Box sx={{ display: 'flex', alignItems: 'center', gap: 1 }}>
        <Tooltip title="Memory Usage">
          <Box sx={{ display: 'flex', alignItems: 'center', gap: 0.5 }}>
            <Memory sx={{ fontSize: 16 }} />
            <Typography variant="caption">256MB / 1GB</Typography>
          </Box>
        </Tooltip>
        <Tooltip title="CPU Usage">
          <Box sx={{ display: 'flex', alignItems: 'center', gap: 0.5 }}>
            <Speed sx={{ fontSize: 16 }} />
            <Typography variant="caption">12%</Typography>
          </Box>
        </Tooltip>
        <Tooltip title="Network Status">
          <Box sx={{ display: 'flex', alignItems: 'center', gap: 0.5 }}>
            <Wifi sx={{ fontSize: 16 }} />
            <Typography variant="caption">1.2MB/s</Typography>
          </Box>
        </Tooltip>
        <Tooltip title="Storage">
          <Box sx={{ display: 'flex', alignItems: 'center', gap: 0.5 }}>
            <Storage sx={{ fontSize: 16 }} />
            <Typography variant="caption">45%</Typography>
          </Box>
        </Tooltip>
        <Tooltip title="Security Status">
          <IconButton size="small" sx={{ color: 'success.main' }}>
            <Security sx={{ fontSize: 16 }} />
          </IconButton>
        </Tooltip>
        <Tooltip title="Notifications">
          <IconButton size="small">
            <Notifications sx={{ fontSize: 16 }} />
          </IconButton>
        </Tooltip>
      </Box>
    </Box>
  );
};

export default StatusBar; 