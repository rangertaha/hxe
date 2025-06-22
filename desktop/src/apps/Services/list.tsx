import React, { useState, useMemo, useEffect, useRef } from 'react';
import { useNavigate } from 'react-router-dom';

import {
  Box,
  Typography,
  IconButton,
  Menu,
  MenuItem,
  ListItemIcon,
  ListItemText,
  Divider,
  Chip,
  Tooltip,
  TextField,
  InputAdornment,
  Button,
  Paper,
  Grid,
  AppBar,
  Toolbar,
  Stack,
  CircularProgress,
} from '@mui/material';
import {
  MoreVert,
  Settings,
  FileDownload,
  FileUpload,
  Refresh,
  PlayArrow,
  Stop,
  Edit,
  Delete,
  AutoGraph,
  Speed,
  Security,
  BugReport,
  Code,
  Add,
  Menu as MenuIcon,
  Search,
  Download,
  Upload,
  FilterList,
} from '@mui/icons-material';
import { DataGrid, GridColDef, GridRenderCellParams, GridPaginationModel, GridMenuIcon } from '@mui/x-data-grid';
import ServiceDetail from './detail';
import ServiceForm from './form';
import Client from '../../client';
import { Service } from '../../client/models';

// Types
type ServiceState = 'RUNNING' | 'STOPPED' | 'FAILED' | 'COMPLETED' | 'TIMEOUT';

// Constants
const SCHEDULES = ['5m', '15m', '30m', '1h', '1d', '7d', '∞'];
const LAST_RUNS = ['1m', '5m', '10m', '30m', '1h', '2h', '1d', '2d'];
const ICONS = [<AutoGraph />, <Speed />, <Security />, <BugReport />, <Code />];

// Utility functions
const getStatusColor = (state: ServiceState) => {
  switch (state) {
    case 'RUNNING':
      return 'success';
    case 'STOPPED':
      return 'default';
    case 'FAILED':
      return 'error';
    case 'COMPLETED':
      return 'info';
    case 'TIMEOUT':
      return 'warning';
    default:
      return 'default';
  }
};

// Circular progress with label
const CircularProgressWithLabel = (props: {
  value: number;
  color: string;
  label: string;
}) => {
  return (
    <Box sx={{ position: 'relative', display: 'inline-flex', flexDirection: 'column', alignItems: 'center' }}>
      <CircularProgress
        variant="determinate"
        value={props.value}
        size={40}
        thickness={4}
        sx={{ color: props.color }}
      />
      <Box
        sx={{
          top: 0,
          left: 0,
          bottom: 0,
          right: 0,
          position: 'absolute',
          display: 'flex',
          alignItems: 'center',
          justifyContent: 'center',
        }}
      >
        <Typography variant="caption" component="div" color="text.secondary">
          {`${Math.round(props.value)}%`}
        </Typography>
      </Box>
      <Typography variant="caption" color="text.secondary" sx={{ mt: 0.5 }}>
        {props.label}
      </Typography>
    </Box>
  );
};

// Main component
const ServicesList: React.FC = () => {
  const navigate = useNavigate();
  const [searchQuery, setSearchQuery] = useState('');
  const [services, setServices] = useState<Service[]>([]);
  const [filteredServices, setFilteredServices] = useState<Service[]>([]);
  const [page, setPage] = useState(0);
  const [loading, setLoading] = useState(false);
  const [hasMore, setHasMore] = useState(true);
  const pageSize = 20;
  const [anchorEl, setAnchorEl] = useState<null | HTMLElement>(null);
  const open = Boolean(anchorEl);
  const [selectedService, setSelectedService] = useState<Service | null>(null);
  const [detailDrawerOpen, setDetailDrawerOpen] = useState(false);
  const [formDrawerOpen, setFormDrawerOpen] = useState(false);
  const [apiClient] = useState(() => new Client('http://localhost:8080'));

  const fetchServices = async (pageNum: number) => {
    try {
      setLoading(true);
      const result = await apiClient.services.listServices();
      console.log(result);
      // Simulate pagination since the backend doesn't support it yet
      const start = pageNum * pageSize;
      const end = start + pageSize;
      const newServices = result.slice(start, end);
      
      if (newServices.length < pageSize) {
        setHasMore(false);  
      }
      
      setServices(prev => [...prev, ...newServices]);
      setLoading(false);
    } catch (error) {
      console.error('Failed to fetch services:', error);
      setLoading(false);
    }
  };

  useEffect(() => {
    fetchServices(0);
  }, [apiClient]);

  // Setup scroll observer
  const scrollRef = useRef<HTMLDivElement>(null);
  
  useEffect(() => {
    const options = {
      root: scrollRef.current,
      rootMargin: '20px',
      threshold: 0.1,
    };

    const observer = new IntersectionObserver((entries) => {
      const target = entries[0];
      if (target.isIntersecting && hasMore && !loading) {
        setPage(prev => {
          const nextPage = prev + 1;
          fetchServices(nextPage);
          return nextPage;
        });
      }
    }, options);

    const scrollArea = document.querySelector('.MuiDataGrid-virtualScroller');
    if (scrollArea) {
      observer.observe(scrollArea);
    }

    return () => {
      if (scrollArea) {
        observer.unobserve(scrollArea);
      }
    };
  }, [hasMore, loading]);

  useEffect(() => {
    if (!searchQuery) {
      setFilteredServices(services);
      return;
    }
    setFilteredServices(services.filter(service =>
      service.name.toLowerCase().includes(searchQuery.toLowerCase())
    ));
  }, [services, searchQuery]);

  const columns: GridColDef[] = [
    { 
      field: 'name', 
      headerName: 'Name',
      flex: 2,
      minWidth: 100
    },
    { 
      field: 'description', 
      headerName: 'Description',
      flex: 3,
      minWidth: 200
    },
    {
      field: 'state',
      headerName: 'Status',
      flex: 1,
      minWidth: 120,
      renderCell: (params) => (
        <Chip label={params.value} color={getStatusColor(params.value as ServiceState)} size="small" />
      ),
    },
    {
      field: 'memory',
      headerName: 'Memory',
      flex: 1,
      minWidth: 100,
      valueFormatter: (params) => {
        const mem = params.value || 0;
        if (mem < 1024) return `${mem} KB`;
        return `${(mem/1024).toFixed(1)} MB`;
      },
    },
    {
      field: 'cpu',
      headerName: 'CPU',
      flex: 1,
      minWidth: 80,
      valueFormatter: (params) => `${params.value || 0}%`,
    },
    {
      field: 'publishers',
      headerName: 'Pub',
      flex: 0.5,
      minWidth: 60,
      align: 'center',
      headerAlign: 'center',
    },
    {
      field: 'subscribers',
      headerName: 'Sub',
      flex: 0.5,
      minWidth: 60,
      align: 'center',
      headerAlign: 'center',
    },
    {
      field: 'uptime',
      headerName: 'Uptime',
      flex: 1.5,
      minWidth: 160,
    },
    {
      field: 'updated',
      headerName: 'Updated',
      flex: 1.5,
      minWidth: 160,
      valueFormatter: (params) => {
        return new Date(params.value).toLocaleTimeString();
      },
    },
    {
      field: 'actions',
      headerName: 'Actions',
      flex: 0.5,
      minWidth: 100,
      renderCell: (params) => (
        <IconButton size="small" onClick={(e) => e.stopPropagation()}>
          <MoreVert />
        </IconButton>
      ),
    },
  ];

  const rows = useMemo(() => services.map(service => ({ ...service, actions: 'actions' })), [services]);

  const total = services.length;
  const running = services.filter(t => t.state === 'RUNNING').length;
  const failed = services.filter(t => t.state === 'FAILED' || t.state === 'TIMEOUT').length;
  const completed = services.filter(t => t.state === 'COMPLETED').length;

  const typeStats = useMemo(() => {
    const acc: Record<string, number> = {};
    services.forEach(service => {
      const type = service.exec ? (service.exec.startsWith('python') ? 'Python' :
                  service.exec.startsWith('node') ? 'Node.js' : 'Shell') : 'N/A';
      acc[type] = (acc[type] || 0) + 1;
    });
    return acc;
  }, [services]);

  const handleMenuClick = (event: React.MouseEvent<HTMLElement>) => {
    setAnchorEl(event.currentTarget);
  };

  const handleMenuClose = () => {
    setAnchorEl(null);
  };

  const handleRowClick = (params: any) => {
    navigate(`/services/${params.row.id}`);
  };

  const handleCloseDrawer = () => {
    setDetailDrawerOpen(false);
    setSelectedService(null);
  };

  const handleAddClick = () => {
    navigate('/services/new');
  };

  const handleFormClose = () => {
    setFormDrawerOpen(false);
  };

  const handleFormSubmit = () => {
    // Refresh the service list after creating a new service
    fetchServices(0);
  };

  return (
    <Box sx={{ 
      display: 'flex', 
      flexDirection: 'column', 
      height: '100%',
    }}>
      <AppBar position="static" color="default" elevation={1} sx={{ mb: 0 }}>
        <Toolbar variant="dense" sx={{ minHeight: 40, p: 0 }}>
          <Typography variant="subtitle2" sx={{ mr: 2, fontWeight: 600 }}>
            Services
          </Typography>
          <TextField
            size="small"
            placeholder="Search services..."
            value={searchQuery}
            onChange={(e) => setSearchQuery(e.target.value)}
            sx={{
              width: '25%',
              mr: 2,
              '& .MuiOutlinedInput-root': {
                bgcolor: 'transparent',
                height: 40,
                borderRadius: 0,
                transition: 'border-bottom 0.2s',
                '& fieldset': {
                  border: 'none',
                },
                '&:hover fieldset': {
                  border: 'none',
                },
                '&.Mui-focused': {
                  borderBottom: '2px solid',
                  borderColor: 'primary.main',
                },
                '&.Mui-focused fieldset': {
                  border: 'none',
                },
              },
              '& .MuiInputBase-input': {
                padding: '8px 12px',
                fontSize: '0.875rem',
              },
              '& .MuiInputAdornment-root': {
                ml: 1,
                color: 'text.secondary',
                '& .MuiSvgIcon-root': {
                  fontSize: '1.25rem',
                },
              },
            }}
            InputProps={{
              startAdornment: (
                <InputAdornment position="start">
                  <Search fontSize="small" />
                </InputAdornment>
              ),
            }}
          />
          <Stack direction="row" spacing={3}>
            <Tooltip title="Running Services" placement="bottom">
              <Box>
                <CircularProgressWithLabel
                  value={Math.round((running / total) * 100) || 0}
                  color="success.main"
                  label="Running"
                />
              </Box>
            </Tooltip>
            <Tooltip title="Failed Services" placement="bottom">
              <Box>
                <CircularProgressWithLabel
                  value={Math.round((failed / total) * 100) || 0}
                  color="error.main"
                  label="Failed"
                />
              </Box>
            </Tooltip>
            <Tooltip title="Completed Services" placement="bottom">
              <Box>
                <CircularProgressWithLabel
                  value={Math.round((completed / total) * 100) || 0}
                  color="info.main"
                  label="Completed"
                />
              </Box>
            </Tooltip>
          </Stack>
          <Box sx={{ flexGrow: 1 }} />
          {/* <Tooltip title="Add Service">
            <IconButton 
              color="primary" 
              size="small" 
              sx={{ p: 0.5 }}
              onClick={handleAddClick}
            >
              <Add fontSize="small" />
            </IconButton>
          </Tooltip> */}
        </Toolbar>
      </AppBar>

      <Box sx={{ 
        flex: 1,
        minHeight: 0,
        '& .MuiDataGrid-root': {
          border: 'none',
        },
        '& .MuiDataGrid-cell': {
          borderColor: 'divider',
        },
        '& .MuiDataGrid-columnHeaders': {
          borderColor: 'divider',
          bgcolor: 'background.paper',
        },
      }}>
        <DataGrid
          rows={rows}
          columns={columns}
          onRowClick={handleRowClick}
          // disableColumnMenu
          // disableRowSelectionOnClick
          hideFooter
          sx={{
            '& .MuiDataGrid-cell:focus': {
              outline: 'none',
            },
          }}
        />
      </Box>
    </Box>
  );
};

export default ServicesList;
