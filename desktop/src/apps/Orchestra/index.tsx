import React, { useState } from "react";
import ListView from "./list";
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
  FilterList,
  MoreVert,
  Settings
} from "@mui/icons-material";

const Dags: React.FC = () => {
  const [anchorEl, setAnchorEl] = useState<null | HTMLElement>(null);
  const open = Boolean(anchorEl);

  const handleMenuClick = (event: React.MouseEvent<HTMLElement>) => {
    setAnchorEl(event.currentTarget);
  };

  const handleMenuClose = () => {
    setAnchorEl(null);
  };

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
        <ListView />
      </Box>
    </Box>
  );
};

export default Dags;
