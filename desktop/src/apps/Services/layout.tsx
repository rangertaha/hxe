import React, { useState } from "react";
import { Outlet, useNavigate, useLocation } from "react-router-dom";
import { 
  Typography, 
  Box, 
  IconButton, 
  Stack,
  Divider,
  Menu,
  MenuItem,
  ListItemIcon,
  ListItemText,
} from "@mui/material";
import {
  MoreVert,
  Settings,
  Add,
  ArrowBack,
} from "@mui/icons-material";

interface ServicesLayoutProps {
  children?: React.ReactNode;
}

const ServicesLayout: React.FC<ServicesLayoutProps> = ({ children }) => {
  const [anchorEl, setAnchorEl] = useState<null | HTMLElement>(null);
  const navigate = useNavigate();
  const location = useLocation();
  const open = Boolean(anchorEl);

  const handleMenuClick = (event: React.MouseEvent<HTMLElement>) => {
    setAnchorEl(event.currentTarget);
  };

  const handleMenuClose = () => {
    setAnchorEl(null);
  };

  const handleCreateService = () => {
    navigate("new");
    handleMenuClose();
  };

  const handleBackToList = () => {
    navigate(".");
  };

  const isDetailView = location.pathname.includes("/services/");
  const isFormView = location.pathname.includes("/services/new");

  return (
    <Box sx={{ display: "flex", flexDirection: "column", height: "100%" }}>
      <Box 
        sx={{ 
          display: 'flex',
          alignItems: 'center',
          px: 2,
          height: 40,
          borderBottom: 1,
          borderColor: 'divider',
          bgcolor: 'background.paper',
        }}
      >
        <Typography variant="subtitle2" sx={{ fontWeight: 600 }}>
              Process Manager
            </Typography>
            <Box sx={{ flexGrow: 1 }} />
            <Stack direction="row" spacing={0.5} alignItems="center">
              <IconButton
                size="small"
                sx={{ p: 0.5 }}
                onClick={handleMenuClick}
                aria-label="more options"
                aria-controls={open ? "service-menu" : undefined}
                aria-haspopup="true"
                aria-expanded={open ? "true" : undefined}
              >
                <MoreVert fontSize="small" />
              </IconButton>
            </Stack>
            
        {/* {(isDetailView || isFormView) ? (
          <>
            <IconButton
              size="small"
              onClick={handleBackToList}
              sx={{ mr: 1 }}
            >
              <ArrowBack fontSize="small" />
            </IconButton>
            <Typography variant="subtitle2" sx={{ fontWeight: 600 }}>
              {isFormView ? "Create New Service" : "Service Details"}
            </Typography>
          </>
        ) : (
          <>
            <Typography variant="subtitle2" sx={{ fontWeight: 600 }}>
              Process Manager
            </Typography>
            <Box sx={{ flexGrow: 1 }} />
            <Stack direction="row" spacing={0.5} alignItems="center">
              <IconButton
                size="small"
                sx={{ p: 0.5 }}
                onClick={handleMenuClick}
                aria-label="more options"
                aria-controls={open ? "service-menu" : undefined}
                aria-haspopup="true"
                aria-expanded={open ? "true" : undefined}
              >
                <MoreVert fontSize="small" />
              </IconButton>
            </Stack>
          </>
        )} */}
      </Box>

      <Menu
        id="service-menu"
        anchorEl={anchorEl}
        open={open}
        onClose={handleMenuClose}
        MenuListProps={{
          "aria-labelledby": "menu-button",
        }}
      >
        <MenuItem onClick={handleCreateService} dense>
          <ListItemIcon>
            <Add fontSize="small" />
          </ListItemIcon>
          <ListItemText>Create Service</ListItemText>
        </MenuItem>
        <Divider />
        <MenuItem onClick={handleMenuClose} dense>
          <ListItemIcon>
            <Settings fontSize="small" />
          </ListItemIcon>
          <ListItemText>Settings</ListItemText>
        </MenuItem>
      </Menu>

      <Box sx={{ flex: "1 1 auto", overflow: 'auto' }}>
        {children || <Outlet />}
      </Box>
    </Box>
  );
};

export default ServicesLayout; 