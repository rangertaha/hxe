import React, { useState } from "react";
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
  Settings
} from "@mui/icons-material";
import TradingView from './TradingView';

const ChartApp: React.FC = () => {
  const [anchorEl, setAnchorEl] = useState<null | HTMLElement>(null);
  const open = Boolean(anchorEl);
  
  const handleMenuClick = (event: React.MouseEvent<HTMLElement>) => {
    setAnchorEl(event.currentTarget);
  };

  const handleMenuClose = () => {
    setAnchorEl(null);
  };

  return (
    <Box sx={{ height: "100%", display: "flex", flexDirection: "column" }}>
      <Box sx={{ display: "flex", alignItems: "center", p: 2 }}>
        <Typography variant="subtitle2" sx={{ fontWeight: 600 }}>
          Chart
        </Typography>
        <Box sx={{ flexGrow: 1 }} />
        <Stack direction="row" spacing={0.5} alignItems="center">
          <IconButton
            size="small"
            sx={{ p: 0.5 }}
            onClick={handleMenuClick}
            aria-label="more options"
            aria-controls={open ? "task-menu" : undefined}
            aria-haspopup="true"
            aria-expanded={open ? "true" : undefined}
          >
            <MoreVert fontSize="small" />
          </IconButton>
        </Stack>
      </Box>

      <Menu
        id="task-menu"
        anchorEl={anchorEl}
        open={open}
        onClose={handleMenuClose}
        MenuListProps={{
          "aria-labelledby": "menu-button",
        }}
      >
        <Divider />
        <MenuItem onClick={handleMenuClose} dense>
          <ListItemIcon>
            <Settings fontSize="small" />
          </ListItemIcon>
          <ListItemText>Settings</ListItemText>
        </MenuItem>
      </Menu>

      <Box sx={{ flex: "1 1 auto" }}>
        <TradingView />
      </Box>
    </Box>
  );
};

export default ChartApp;
