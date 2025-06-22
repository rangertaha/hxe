import React, { ReactNode } from "react";
import {
  AppBar,
  Box,
  CssBaseline,
  Drawer,
  IconButton,
  List,
  ListItem,
  ListItemButton,
  ListItemIcon,
  ListItemText,
  Toolbar,
  Typography,
  useTheme,
  InputBase,
  alpha,
  useMediaQuery,
  Menu,
  MenuItem,
  Popper,
  Paper,
  ClickAwayListener,
  Avatar,
  Divider,
} from "@mui/material";
import {
  Menu as MenuIcon,
  Dashboard,
  SmartToy,
  ListAlt,
  AccountBalance,
  ShowChart,
  Settings,
  ChevronLeft as ChevronLeftIcon,
  Search as SearchIcon,
  AutoGraph as QuantIcon,
  Clear as ClearIcon,
  Person as ProfileIcon,
  Logout as LogoutIcon,
  Notifications as NotificationsIcon,
  Security as SecurityIcon,
} from "@mui/icons-material";
import { useNavigate, useLocation, Routes, Route } from "react-router-dom";
import { useState } from "react";
import StatusBar from "./StatusBar";
import ToolMenu from "../components/ToolMenu";

const drawerWidth = 240;
const miniDrawerWidth = 65;

interface LayoutProps {
  children: ReactNode;
  title?: string;
  onMenuClick?: () => void;
  setAppTitle?: (title: string) => void;
}

interface NavItem {
  text: string;
  icon: ReactNode;
  path: string;
}

const Layout: React.FC<LayoutProps> = ({
  children,
  title = "hxe",
  onMenuClick,
  setAppTitle,
}) => {
  const theme = useTheme();
  const navigate = useNavigate();
  const location = useLocation();
  const [mobileOpen, setMobileOpen] = useState(false);
  const [isCollapsed, setIsCollapsed] = useState(true);
  const isSmallScreen = useMediaQuery(theme.breakpoints.down("md"));
  const [searchQuery, setSearchQuery] = useState("");
  const [searchAnchorEl, setSearchAnchorEl] = useState<null | HTMLElement>(
    null
  );
  const [searchResults, setSearchResults] = useState<string[]>([]);
  const [isSearchVisible, setIsSearchVisible] = useState(false);
  const [isSearchFocused, setIsSearchFocused] = useState(false);
  const [profileAnchorEl, setProfileAnchorEl] = useState<null | HTMLElement>(
    null
  );
  const [notificationsAnchorEl, setNotificationsAnchorEl] =
    useState<null | HTMLElement>(null);

  // const menuItems = [
  //   { text: 'Dashboard', icon: <Dashboard />, path: '/' },
  //   { text: 'Tasks', icon: <SmartToy />, path: '/tasks' },
  //   { text: 'Notifications', icon: <NotificationsIcon />, path: '/notifications' },
  // ];

  // const bottomMenuItems = [
  //   { text: 'Settings', icon: <Settings />, path: '/settings' },
  // ];

  // const handleDrawerToggle = () => {
  //   setMobileOpen(!mobileOpen);
  // };

  // const handleSearchChange = (event: React.ChangeEvent<HTMLInputElement>) => {
  //   const query = event.target.value;
  //   setSearchQuery(query);

  //   // Simulate search results
  //   if (query.length > 0) {
  //     setSearchResults([
  //       `Search for "${query}" in Bots`,
  //       `Search for "${query}" in Orders`,
  //       `Search for "${query}" in Accounts`,
  //       `Search for "${query}" in Prices`,
  //     ]);
  //     setSearchAnchorEl(event.currentTarget);
  //   } else {
  //     setSearchResults([]);
  //     setSearchAnchorEl(null);
  //   }
  // };

  // const handleSearchClick = (result: string) => {
  //   // TODO: Implement search navigation
  //   console.log('Search clicked:', result);
  //   setSearchQuery('');
  //   setSearchResults([]);
  //   setSearchAnchorEl(null);
  // };

  // const handleSearchClose = () => {
  //   setSearchAnchorEl(null);
  // };

  // const handleProfileMenuOpen = (event: React.MouseEvent<HTMLElement>) => {
  //   setProfileAnchorEl(event.currentTarget);
  // };

  // const handleProfileMenuClose = () => {
  //   setProfileAnchorEl(null);
  // };

  // const handleNotificationsMenuOpen = (event: React.MouseEvent<HTMLElement>) => {
  //   setNotificationsAnchorEl(event.currentTarget);
  // };

  // const handleNotificationsMenuClose = () => {
  //   setNotificationsAnchorEl(null);
  // };

  // const handleSearchBlur = () => {
  //   setIsSearchFocused(false);
  //   if (!searchQuery) {
  //     setIsSearchVisible(false);
  //   }
  // };

  // const handleSearchFocus = () => {
  //   setIsSearchFocused(true);
  // };

  // const drawer = (
  //   <Box sx={{ display: 'flex', flexDirection: 'column', height: '100%' }}>
  //     <Toolbar>
  //       {isCollapsed ? (
  //         <IconButton sx={{ color: 'primary.main' }}>
  //           <QuantIcon />
  //         </IconButton>
  //       ) : (
  //         <Box sx={{ display: 'flex', alignItems: 'center', gap: 1 }}>
  //           <QuantIcon color="primary" />
  //           <Typography variant="h6" noWrap component="div">
  //             hxe
  //           </Typography>
  //         </Box>
  //       )}
  //     </Toolbar>
  //     <List sx={{ flexGrow: 1 }}>
  //       {menuItems.map((item) => (
  //         <ListItem key={item.text} disablePadding>
  //           <ListItemButton
  //             selected={location.pathname === item.path}
  //             onClick={() => {
  //               navigate(item.path);
  //               setMobileOpen(false);
  //             }}
  //             sx={{
  //               minHeight: 48,
  //               justifyContent: isCollapsed ? 'center' : 'initial',
  //               px: 2.5,
  //               '&.Mui-selected': {
  //                 backgroundColor: alpha(theme.palette.primary.main, 0.1),
  //                 '&:hover': {
  //                   backgroundColor: alpha(theme.palette.primary.main, 0.15),
  //                 },
  //               },
  //             }}
  //           >
  //             <ListItemIcon
  //               sx={{
  //                 minWidth: 0,
  //                 mr: isCollapsed ? 'auto' : 3,
  //                 justifyContent: 'center',
  //                 color: location.pathname === item.path ? 'primary.main' : 'inherit',
  //               }}
  //             >
  //               {item.icon}
  //             </ListItemIcon>
  //             {!isCollapsed && <ListItemText primary={item.text} />}
  //           </ListItemButton>
  //         </ListItem>
  //       ))}
  //     </List>
  //     <Divider />
  //     <List>
  //       {bottomMenuItems.map((item) => (
  //         <ListItem key={item.text} disablePadding>
  //           <ListItemButton
  //             selected={location.pathname === item.path}
  //             onClick={() => {
  //               navigate(item.path);
  //               setMobileOpen(false);
  //             }}
  //             sx={{
  //               minHeight: 48,
  //               justifyContent: isCollapsed ? 'center' : 'initial',
  //               px: 2.5,
  //             }}
  //           >
  //             <ListItemIcon
  //               sx={{
  //                 minWidth: 0,
  //                 mr: isCollapsed ? 'auto' : 3,
  //                 justifyContent: 'center',
  //               }}
  //             >
  //               {item.icon}
  //             </ListItemIcon>
  //             {!isCollapsed && <ListItemText primary={item.text} />}
  //           </ListItemButton>
  //         </ListItem>
  //       ))}
  //     </List>
  //   </Box>
  // );

  return (
    // <Box
    //   sx={{
    //     display: "flex",
    //     flexDirection: "column",
    //     height: "100vh",
    //     position: "relative",
    //   }}
    // >
      <Box
        sx={{
          flex: 1,
          // position: "relative",
          // overflow: "hidden",

          minHeight: 'calc(100vh - 28px)'
        }}
      >
        {children}
      </Box>
      // {/* <StatusBar onMenuClick={onMenuClick || (() => {})} /> */}
    // </Box>
  );
};

export default Layout;
