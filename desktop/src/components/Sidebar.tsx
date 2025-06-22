import React, { useEffect, useRef } from 'react';
import {
  Box,
  Drawer,
  List,
  ListItem,
  ListItemButton,
  ListItemIcon,
  ListItemText,
  IconButton,
  Typography,
  useTheme,
  alpha,
  Tooltip,
  Divider,
} from '@mui/material';
import {
  Dashboard,
  ShowChart,
  Notifications as NotificationsIcon,
  Settings,
  ChevronLeft as ChevronLeftIcon,
  AutoGraph as QuantIcon,
  Menu as MenuIcon,
  Code,
} from '@mui/icons-material';
import { useNavigate, useLocation } from 'react-router-dom';

interface SidebarProps {
  open: boolean;
  onClose: () => void;
}

const Sidebar: React.FC<SidebarProps> = ({ open, onClose }) => {
  const theme = useTheme();
  const navigate = useNavigate();
  const location = useLocation();
  const timeoutRef = useRef<NodeJS.Timeout>();

  useEffect(() => {
    if (open) {
      // Clear any existing timeout
      if (timeoutRef.current) {
        clearTimeout(timeoutRef.current);
      }
      // Set new timeout to close sidebar after 10 seconds
      timeoutRef.current = setTimeout(() => {
        onClose();
      }, 10000);
    }

    // Cleanup timeout on unmount or when open changes
    return () => {
      if (timeoutRef.current) {
        clearTimeout(timeoutRef.current);
      }
    };
  }, [open, onClose]);

  const menuItems = [
    { text: 'Dashboard', icon: <Dashboard />, path: '/' },
    { text: 'Services', icon: <Code />, path: '/services' },
    { text: 'Orchestra', icon: <Code />, path: '/orchestra' },
    { text: 'Chart', icon: <Code />, path: '/chart' },
    { text: 'Notifications', icon: <NotificationsIcon />, path: '/notifications' },
  ];

  const bottomMenuItems = [
    { text: 'Settings', icon: <Settings />, path: '/settings' },
  ];

  return (
    <Drawer
      variant="permanent"
      open={open}
      sx={{
        width: open ? 240 : 48,
        flexShrink: 0,
        '& .MuiDrawer-paper': {
          width: open ? 240 : 48,
          boxSizing: 'border-box',
          borderRight: '1px solid',
          borderColor: 'divider',
          transition: theme.transitions.create(['width', 'margin'], {
            easing: theme.transitions.easing.sharp,
            duration: theme.transitions.duration.enteringScreen,
          }),
          overflowX: 'hidden',
          bgcolor: 'background.paper',
          height: 'calc(100vh - 28px)',
          '&:hover': {
            boxShadow: '4px 0 8px rgba(0,0,0,0.1)',
          },
        },
      }}
    >
      <Box sx={{ 
        display: 'flex', 
        flexDirection: 'column', 
        height: '100%',
        position: 'relative'
      }}>
        <Box sx={{ 
          p: 2, 
          display: 'flex', 
          alignItems: 'center', 
          justifyContent: open ? 'space-between' : 'center',
          minHeight: 64
        }}>
          {open ? (
            <Box sx={{ display: 'flex', alignItems: 'center', gap: 1 }}>
              <ShowChart color="primary" sx={{ fontSize: 28 }} />
              <Tooltip title="God of trade, wealth, and merchants" placement="right">
                <Typography variant="h6" noWrap sx={{ fontWeight: 600 }}>
                  hxe
                </Typography>
              </Tooltip>
            </Box>
          ) : (
            <Tooltip title="Open menu" placement="right">
              <IconButton onClick={onClose} size="small">
                <MenuIcon />
              </IconButton>
            </Tooltip>
          )}
          {open && (
            <IconButton onClick={onClose} size="small">
              <ChevronLeftIcon />
            </IconButton>
          )}
        </Box>

        <Divider />

        <List sx={{ flexGrow: 1, p: 0 }}>
          {menuItems.map((item) => (
            <ListItem key={item.text} disablePadding sx={{ mb: 0 }}>
              <Tooltip title={!open ? item.text : ''} placement="right">
                <ListItemButton
                  selected={location.pathname === item.path}
                  onClick={() => navigate(item.path)}
                  sx={{
                    padding: 0,
                    margin: 0,
                    minHeight: 48,
                    height: 48,
                    width: open ? 'auto' : 48,
                    borderRadius: 0,
                    justifyContent: open ? 'initial' : 'center',
                    px: open ? 2.5 : 0,
                    '&.Mui-selected': {
                      backgroundColor: alpha(theme.palette.primary.main, 0.1),
                      '&:hover': {
                        backgroundColor: alpha(theme.palette.primary.main, 0.15),
                      },
                    },
                  }}
                >
                  <ListItemIcon
                    sx={{
                      minWidth: 0,
                      mr: open ? 3 : 0,
                      justifyContent: 'center',
                      width: 24,
                      height: 24,
                      display: 'flex',
                      alignItems: 'center',
                      color: location.pathname === item.path ? 'primary.main' : 'inherit',
                    }}
                  >
                    {item.icon}
                  </ListItemIcon>
                  {open && (
                    <ListItemText 
                      primary={item.text} 
                      primaryTypographyProps={{
                        fontSize: '0.9rem',
                        fontWeight: location.pathname === item.path ? 600 : 400,
                      }}
                    />
                  )}
                </ListItemButton>
              </Tooltip>
            </ListItem>
          ))}
        </List>

        <Divider />

        <List sx={{ p: 0 }}>
          {bottomMenuItems.map((item) => (
            <ListItem key={item.text} disablePadding sx={{ mb: 0 }}>
              <Tooltip title={!open ? item.text : ''} placement="right">
                <ListItemButton
                  selected={location.pathname === item.path}
                  onClick={() => navigate(item.path)}
                  sx={{
                    padding: 0,
                    margin: 0,
                    minHeight: 48,
                    height: 48,
                    width: open ? 'auto' : 48,
                    borderRadius: 0,
                    justifyContent: open ? 'initial' : 'center',
                    px: open ? 2.5 : 0,
                  }}
                >
                  <ListItemIcon
                    sx={{
                      minWidth: 0,
                      mr: open ? 3 : 0,
                      justifyContent: 'center',
                      width: 24,
                      height: 24,
                      display: 'flex',
                      alignItems: 'center',
                    }}
                  >
                    {item.icon}
                  </ListItemIcon>
                  {open && (
                    <ListItemText 
                      primary={item.text}
                      primaryTypographyProps={{
                        fontSize: '0.9rem',
                        fontWeight: location.pathname === item.path ? 600 : 400,
                      }}
                    />
                  )}
                </ListItemButton>
              </Tooltip>
            </ListItem>
          ))}
        </List>
      </Box>
    </Drawer>
  );
};

export default Sidebar; 