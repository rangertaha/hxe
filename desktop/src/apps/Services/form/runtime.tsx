import React from 'react';
import {
  Box,
  TextField,
  FormControl,
  InputLabel,
  Select,
  MenuItem,
  Grid,
  Typography,
  Switch,
  FormControlLabel,
} from '@mui/material';

const ServiceRuntime: React.FC = () => {
  return (
    <Box>
      <Typography variant="subtitle2" gutterBottom>
        Service Configuration
      </Typography>
      <Grid container spacing={2}>
        <Grid item xs={12}>
          <TextField
            fullWidth
            label="Service Name"
            name="name"
            required
            size="small"
          />
        </Grid>
        <Grid item xs={12}>
          <TextField
            fullWidth
            label="Description"
            name="description"
            multiline
            rows={2}
            size="small"
          />
        </Grid>
        <Grid item xs={12}>
          <FormControl fullWidth size="small">
            <InputLabel>Type</InputLabel>
            <Select
              name="type"
              label="Type"
              defaultValue="shell"
            >
              <MenuItem value="shell">Shell</MenuItem>
              <MenuItem value="python">Python</MenuItem>
              <MenuItem value="node">Node.js</MenuItem>
            </Select>
          </FormControl>
        </Grid>
        <Grid item xs={12}>
          <TextField
            fullWidth
            label="Working Directory"
            name="dir"
            placeholder="/path/to/service/directory"
            size="small"
          />
        </Grid>
        <Grid item xs={6}>
          <TextField
            fullWidth
            label="User"
            name="user"
            placeholder="user to run the service"
            size="small"
          />
        </Grid>
        <Grid item xs={6}>
          <TextField
            fullWidth
            label="Group"
            name="group"
            placeholder="group to run the service"
            size="small"
          />
        </Grid>
        <Grid item xs={12}>
          <Typography variant="subtitle2" gutterBottom sx={{ mt: 2 }}>
            Execution Commands
          </Typography>
        </Grid>
        <Grid item xs={12}>
          <TextField
            fullWidth
            label="Pre-execution Command"
            name="preExec"
            placeholder="Command to run before the main service"
            size="small"
          />
        </Grid>
        <Grid item xs={12}>
          <TextField
            fullWidth
            label="Main Execution Command"
            name="exec"
            placeholder="Main command to run the service"
            required
            size="small"
          />
        </Grid>
        <Grid item xs={12}>
          <TextField
            fullWidth
            label="Post-execution Command"
            name="postExec"
            placeholder="Command to run after the main service"
            size="small"
          />
        </Grid>
        <Grid item xs={12}>
          <Typography variant="subtitle2" gutterBottom sx={{ mt: 2 }}>
            Service Control
          </Typography>
        </Grid>
        <Grid item xs={12}>
          <FormControlLabel
            control={<Switch name="autostart" defaultChecked />}
            label="Auto Start"
          />
        </Grid>
        <Grid item xs={12}>
          <FormControl fullWidth size="small">
            <InputLabel>Stop Signal</InputLabel>
            <Select
              name="stopSignal"
              label="Stop Signal"
              defaultValue="SIGTERM"
            >
              <MenuItem value="SIGTERM">SIGTERM</MenuItem>
              <MenuItem value="SIGINT">SIGINT</MenuItem>
              <MenuItem value="SIGKILL">SIGKILL</MenuItem>
            </Select>
          </FormControl>
        </Grid>
        <Grid item xs={12}>
          <TextField
            fullWidth
            label="Exit Codes"
            name="exitCodes"
            placeholder="0,1,2 (comma-separated)"
            helperText="Comma-separated list of valid exit codes"
            size="small"
          />
        </Grid>
      </Grid>
    </Box>
  );
};

export default ServiceRuntime;
