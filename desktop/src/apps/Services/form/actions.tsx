import React, { useState } from 'react';
import {
  Box,
  TextField,
  Grid,
  Typography,
  IconButton,
  List,
  ListItem,
  ListItemText,
  ListItemSecondaryAction,
  Paper,
  Button,
  Divider,
  FormControl,
  InputLabel,
  Select,
  MenuItem,
  SelectChangeEvent,
} from '@mui/material';
import { Add as AddIcon, Delete as DeleteIcon } from '@mui/icons-material';

interface Action {
  name: string;
  command: string;
  type: 'start' | 'stop' | 'restart' | 'custom';
  description: string;
}

const ServiceActions: React.FC = () => {
  const [actions, setActions] = useState<Action[]>([]);
  const [newAction, setNewAction] = useState<Action>({
    name: '',
    command: '',
    type: 'custom',
    description: '',
  });

  const handleAddAction = () => {
    if (newAction.name && newAction.command) {
      setActions([...actions, newAction]);
      setNewAction({
        name: '',
        command: '',
        type: 'custom',
        description: '',
      });
    }
  };

  const handleRemoveAction = (index: number) => {
    setActions(actions.filter((_, i) => i !== index));
  };

  const handleTextChange = (field: keyof Action) => (
    event: React.ChangeEvent<HTMLInputElement>
  ) => {
    setNewAction({
      ...newAction,
      [field]: event.target.value,
    });
  };

  const handleSelectChange = (event: SelectChangeEvent) => {
    setNewAction({
      ...newAction,
      type: event.target.value as Action['type'],
    });
  };

  return (
    <Box>
      <Paper sx={{ p: 2, mb: 2 }}>
        <Typography variant="h6" gutterBottom>
          Add New Action
        </Typography>
        <Grid container spacing={2}>
          <Grid item xs={12} sm={6}>
            <TextField
              fullWidth
              label="Action Name"
              value={newAction.name}
              onChange={handleTextChange('name')}
              size="small"
            />
          </Grid>
          <Grid item xs={12} sm={6}>
            <FormControl fullWidth size="small">
              <InputLabel>Action Type</InputLabel>
              <Select
                value={newAction.type}
                label="Action Type"
                onChange={handleSelectChange}
              >
                <MenuItem value="start">Start</MenuItem>
                <MenuItem value="stop">Stop</MenuItem>
                <MenuItem value="restart">Restart</MenuItem>
                <MenuItem value="custom">Custom</MenuItem>
              </Select>
            </FormControl>
          </Grid>
          <Grid item xs={12}>
            <TextField
              fullWidth
              label="Command"
              value={newAction.command}
              onChange={handleTextChange('command')}
              size="small"
            />
          </Grid>
          <Grid item xs={12}>
            <TextField
              fullWidth
              label="Description"
              value={newAction.description}
              onChange={handleTextChange('description')}
              multiline
              rows={2}
              size="small"
            />
          </Grid>
          <Grid item xs={12}>
            <Button
              variant="contained"
              startIcon={<AddIcon />}
              onClick={handleAddAction}
              disabled={!newAction.name || !newAction.command}
            >
              Add Action
            </Button>
          </Grid>
        </Grid>
      </Paper>

      <Paper>
        <List>
          {actions.map((action, index) => (
            <React.Fragment key={index}>
              <ListItem>
                <ListItemText
                  primary={action.name}
                  secondary={
                    <>
                      <Typography component="span" variant="body2" color="text.primary">
                        Type: {action.type}
                      </Typography>
                      <br />
                      <Typography component="span" variant="body2">
                        Command: {action.command}
                      </Typography>
                      {action.description && (
                        <>
                          <br />
                          <Typography component="span" variant="body2">
                            Description: {action.description}
                          </Typography>
                        </>
                      )}
                    </>
                  }
                />
                <ListItemSecondaryAction>
                  <IconButton
                    edge="end"
                    aria-label="delete"
                    onClick={() => handleRemoveAction(index)}
                  >
                    <DeleteIcon />
                  </IconButton>
                </ListItemSecondaryAction>
              </ListItem>
              {index < actions.length - 1 && <Divider />}
            </React.Fragment>
          ))}
        </List>
      </Paper>
    </Box>
  );
};

export default ServiceActions;
