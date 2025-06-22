import { createTheme } from '@mui/material/styles';
import { DataGrid } from '@mui/x-data-grid';

declare module '@mui/material/styles' {
  interface Components {
    MuiDataGrid?: {
      styleOverrides?: {
        root?: {
          border?: string;
          backgroundColor?: string;
          backdropFilter?: string;
          borderRadius?: string;
          overflow?: string;
          '& .MuiDataGrid-cell'?: {
            borderColor?: string;
          };
          '& .MuiDataGrid-row'?: {
            backgroundColor?: string;
            '&:hover'?: {
              backgroundColor?: string;
            };
          };
          '& .MuiDataGrid-columnHeaders'?: {
            backgroundColor?: string;
            borderColor?: string;
            borderBottom?: string;
          };
          '& .MuiDataGrid-footerContainer'?: {
            backgroundColor?: string;
            borderColor?: string;
          };
          '& .MuiDataGrid-virtualScroller'?: {
            backgroundColor?: string;
          };
          '& .MuiDataGrid-virtualScrollerContent'?: {
            backgroundColor?: string;
          };
          '& .MuiDataGrid-virtualScrollerRenderZone'?: {
            backgroundColor?: string;
          };
          '& .MuiDataGrid-columnSeparator'?: {
            display?: string;
          };
          '& .MuiDataGrid-cell:focus'?: {
            outline?: string;
          };
          '& .MuiDataGrid-columnHeader:focus'?: {
            outline?: string;
          };
          '& .MuiDataGrid-columnHeader:focus-within'?: {
            outline?: string;
          };
        };
      };
    };
  }
}

export const cyberpunkTheme = createTheme({
  palette: {
    mode: 'dark',
    primary: {
      main: '#00f2ff',
      light: '#33f5ff',
      dark: '#00a9b3',
      contrastText: '#000',
    },
    secondary: {
      main: '#ff00ff',
      light: '#ff33ff',
      dark: '#b300b3',
      contrastText: '#fff',
    },
    background: {
      default: '#0a0a0f',
      paper: '#1a1a2e',
    },
    text: {
      primary: '#ffffff',
      secondary: '#b3b3b3',
    },
    error: {
      main: '#ff3d3d',
    },
    warning: {
      main: '#ffb74d',
    },
    info: {
      main: '#00f2ff',
    },
    success: {
      main: '#00ff9d',
    },
  },
  typography: {
    fontFamily: '"Roboto Mono", "Roboto", "Helvetica", "Arial", sans-serif',
    h1: {
      fontWeight: 700,
      letterSpacing: '0.02em',
    },
    h2: {
      fontWeight: 700,
      letterSpacing: '0.02em',
    },
    h3: {
      fontWeight: 600,
      letterSpacing: '0.01em',
    },
    h4: {
      fontWeight: 600,
      letterSpacing: '0.01em',
    },
    h5: {
      fontWeight: 500,
    },
    h6: {
      fontWeight: 500,
    },
    button: {
      textTransform: 'none',
      fontWeight: 500,
    },
  },
  shape: {
    borderRadius: 8,
  },
  components: {
    MuiCssBaseline: {
      styleOverrides: {
        body: {
          scrollbarColor: '#00f2ff #1a1a2e',
          '&::-webkit-scrollbar, & *::-webkit-scrollbar': {
            width: '8px',
            height: '8px',
          },
          '&::-webkit-scrollbar-thumb, & *::-webkit-scrollbar-thumb': {
            borderRadius: 8,
            backgroundColor: '#00f2ff',
            minHeight: 24,
          },
          '&::-webkit-scrollbar-track, & *::-webkit-scrollbar-track': {
            borderRadius: 8,
            backgroundColor: '#1a1a2e',
          },
        },
      },
    },
    MuiPaper: {
      styleOverrides: {
        root: {
          backgroundImage: 'none',
          boxShadow: '0 4px 20px rgba(0, 242, 255, 0.1)',
        },
      },
    },
    MuiButton: {
      styleOverrides: {
        root: {
          borderRadius: 8,
          padding: '8px 16px',
          transition: 'all 0.2s ease-in-out',
          '&:hover': {
            transform: 'translateY(-2px)',
            boxShadow: '0 4px 12px rgba(0, 242, 255, 0.2)',
          },
        },
        contained: {
          background: 'linear-gradient(45deg, #00f2ff 30%, #00a9b3 90%)',
          '&:hover': {
            background: 'linear-gradient(45deg, #33f5ff 30%, #00c4cf 90%)',
          },
        },
      },
    },
    MuiDataGrid: {
      styleOverrides: {
        root: {
          border: 'none',
          '& .MuiDataGrid-cell': {
            borderColor: 'rgba(255, 255, 255, 0.1)',
          },
          '& .MuiDataGrid-columnHeaders': {
            backgroundColor: 'rgba(0, 242, 255, 0.05)',
            borderBottom: '1px solid rgba(0, 242, 255, 0.1)',
          },
          '& .MuiDataGrid-row': {
            '&:hover': {
              backgroundColor: 'rgba(0, 242, 255, 0.05)',
            },
          },
        },
      },
    },
    MuiDrawer: {
      styleOverrides: {
        paper: {
          background: 'linear-gradient(180deg, #1a1a2e 0%, #0a0a0f 100%)',
          borderLeft: '1px solid rgba(0, 242, 255, 0.1)',
        },
      },
    },
  },
});

export default cyberpunkTheme; 