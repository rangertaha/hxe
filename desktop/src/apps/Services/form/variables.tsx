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
  FormControlLabel,
  Switch,
  Tooltip,
  Alert,
  InputAdornment,
  Collapse,
} from '@mui/material';
import {
  Add as AddIcon,
  Delete as DeleteIcon,
  Visibility as VisibilityIcon,
  VisibilityOff as VisibilityOffIcon,
  Lock as LockIcon,
  LockOpen as LockOpenIcon,
  Edit as EditIcon,
  Save as SaveIcon,
  Cancel as CancelIcon,
} from '@mui/icons-material';

interface Variable {
  key: string;
  value: string;
  isSecret: boolean;
}

const ServiceVariables: React.FC = () => {
  const [variables, setVariables] = useState<Variable[]>([]);
  const [newVariable, setNewVariable] = useState<Variable>({
    key: '',
    value: '',
    isSecret: false,
  });
  const [showSecret, setShowSecret] = useState(false);
  const [error, setError] = useState<string | null>(null);
  const [editingIndex, setEditingIndex] = useState<number | null>(null);
  const [editVariable, setEditVariable] = useState<Variable | null>(null);

  const validateKey = (key: string, excludeIndex?: number): boolean => {
    // Check if key is empty
    if (!key.trim()) return false;
    
    // Check if key contains only valid characters (letters, numbers, underscores)
    if (!/^[a-zA-Z0-9_]+$/.test(key)) return false;
    
    // Check if key already exists (excluding the current variable when editing)
    if (variables.some((v, i) => v.key === key && i !== excludeIndex)) return false;
    
    return true;
  };

  const handleAddVariable = () => {
    setError(null);
    
    if (!validateKey(newVariable.key)) {
      setError('Invalid key. Use only letters, numbers, and underscores. Key must be unique.');
      return;
    }

    setVariables([...variables, newVariable]);
    setNewVariable({
      key: '',
      value: '',
      isSecret: false,
    });
  };

  const handleRemoveVariable = (index: number) => {
    setVariables(variables.filter((_, i) => i !== index));
    if (editingIndex === index) {
      setEditingIndex(null);
      setEditVariable(null);
    }
  };

  const handleEditVariable = (index: number) => {
    setEditingIndex(index);
    setEditVariable({ ...variables[index] });
  };

  const handleSaveEdit = (index: number) => {
    if (!editVariable) return;

    setError(null);
    if (!validateKey(editVariable.key, index)) {
      setError('Invalid key. Use only letters, numbers, and underscores. Key must be unique.');
      return;
    }

    const newVariables = [...variables];
    newVariables[index] = editVariable;
    setVariables(newVariables);
    setEditingIndex(null);
    setEditVariable(null);
  };

  const handleCancelEdit = () => {
    setEditingIndex(null);
    setEditVariable(null);
    setError(null);
  };

  const handleKeyPress = (event: React.KeyboardEvent) => {
    if (event.key === 'Enter') {
      event.preventDefault();
      if (editingIndex !== null) {
        handleSaveEdit(editingIndex);
      } else {
        handleAddVariable();
      }
    }
  };

  return (
    <Box>
      <Paper sx={{ p: 2, mb: 2 }}>
        <Typography variant="h6" gutterBottom>
          Add Environment Variable
        </Typography>
        {error && (
          <Alert severity="error" sx={{ mb: 2 }}>
            {error}
          </Alert>
        )}
        <Grid container spacing={2} alignItems="center">
          <Grid item xs={12} sm={4}>
            <TextField
              fullWidth
              label="Variable Name"
              value={newVariable.key}
              onChange={(e) => {
                setNewVariable({ ...newVariable, key: e.target.value });
                setError(null);
              }}
              onKeyPress={handleKeyPress}
              size="small"
              placeholder="e.g., API_KEY"
              InputProps={{
                startAdornment: (
                  <InputAdornment position="start">
                    {newVariable.isSecret ? <LockIcon fontSize="small" /> : <LockOpenIcon fontSize="small" />}
                  </InputAdornment>
                ),
              }}
            />
          </Grid>
          <Grid item xs={12} sm={6}>
            <TextField
              fullWidth
              label="Value"
              value={newVariable.value}
              onChange={(e) => setNewVariable({ ...newVariable, value: e.target.value })}
              onKeyPress={handleKeyPress}
              type={newVariable.isSecret && !showSecret ? 'password' : 'text'}
              size="small"
              placeholder={newVariable.isSecret ? '••••••••' : 'Enter value'}
              InputProps={{
                endAdornment: newVariable.isSecret && (
                  <InputAdornment position="end">
                    <IconButton
                      onClick={() => setShowSecret(!showSecret)}
                      edge="end"
                      size="small"
                    >
                      {showSecret ? <VisibilityOffIcon /> : <VisibilityIcon />}
                    </IconButton>
                  </InputAdornment>
                ),
              }}
            />
          </Grid>
          <Grid item xs={12} sm={2}>
            <FormControlLabel
              control={
                <Switch
                  checked={newVariable.isSecret}
                  onChange={(e) => setNewVariable({ ...newVariable, isSecret: e.target.checked })}
                  size="small"
                />
              }
              label="Secret"
            />
          </Grid>
          <Grid item xs={12}>
            <Button
              variant="contained"
              startIcon={<AddIcon />}
              onClick={handleAddVariable}
              disabled={!newVariable.key || !newVariable.value}
            >
              Add Variable
            </Button>
          </Grid>
        </Grid>
      </Paper>

      <Paper>
        <List>
          {variables.map((variable, index) => (
            <React.Fragment key={index}>
              <ListItem>
                {editingIndex === index && editVariable ? (
                  <Grid container spacing={2} alignItems="center">
                    <Grid item xs={12} sm={4}>
                      <TextField
                        fullWidth
                        label="Variable Name"
                        value={editVariable.key}
                        onChange={(e) => setEditVariable({ ...editVariable, key: e.target.value })}
                        onKeyPress={handleKeyPress}
                        size="small"
                        InputProps={{
                          startAdornment: (
                            <InputAdornment position="start">
                              {editVariable.isSecret ? <LockIcon fontSize="small" /> : <LockOpenIcon fontSize="small" />}
                            </InputAdornment>
                          ),
                        }}
                      />
                    </Grid>
                    <Grid item xs={12} sm={6}>
                      <TextField
                        fullWidth
                        label="Value"
                        value={editVariable.value}
                        onChange={(e) => setEditVariable({ ...editVariable, value: e.target.value })}
                        onKeyPress={handleKeyPress}
                        type={editVariable.isSecret && !showSecret ? 'password' : 'text'}
                        size="small"
                        InputProps={{
                          endAdornment: editVariable.isSecret && (
                            <InputAdornment position="end">
                              <IconButton
                                onClick={() => setShowSecret(!showSecret)}
                                edge="end"
                                size="small"
                              >
                                {showSecret ? <VisibilityOffIcon /> : <VisibilityIcon />}
                              </IconButton>
                            </InputAdornment>
                          ),
                        }}
                      />
                    </Grid>
                    <Grid item xs={12} sm={2}>
                      <FormControlLabel
                        control={
                          <Switch
                            checked={editVariable.isSecret}
                            onChange={(e) => setEditVariable({ ...editVariable, isSecret: e.target.checked })}
                            size="small"
                          />
                        }
                        label="Secret"
                      />
                    </Grid>
                  </Grid>
                ) : (
                  <ListItemText
                    primary={
                      <Box sx={{ display: 'flex', alignItems: 'center', gap: 1 }}>
                        <Typography variant="body2">
                          {variable.key}
                        </Typography>
                        {variable.isSecret && (
                          <Tooltip title="Secret Variable">
                            <LockIcon fontSize="small" color="action" />
                          </Tooltip>
                        )}
                      </Box>
                    }
                    secondary={
                      <Typography
                        variant="body2"
                        color="text.secondary"
                        sx={{
                          fontFamily: 'monospace',
                          wordBreak: 'break-all',
                        }}
                      >
                        {variable.isSecret ? '••••••••' : variable.value}
                      </Typography>
                    }
                  />
                )}
                <ListItemSecondaryAction>
                  {editingIndex === index ? (
                    <Box sx={{ display: 'flex', gap: 1 }}>
                      <Tooltip title="Save">
                        <IconButton
                          edge="end"
                          aria-label="save"
                          onClick={() => handleSaveEdit(index)}
                        >
                          <SaveIcon />
                        </IconButton>
                      </Tooltip>
                      <Tooltip title="Cancel">
                        <IconButton
                          edge="end"
                          aria-label="cancel"
                          onClick={handleCancelEdit}
                        >
                          <CancelIcon />
                        </IconButton>
                      </Tooltip>
                    </Box>
                  ) : (
                    <Box sx={{ display: 'flex', gap: 1 }}>
                      <Tooltip title="Edit">
                        <IconButton
                          edge="end"
                          aria-label="edit"
                          onClick={() => handleEditVariable(index)}
                        >
                          <EditIcon />
                        </IconButton>
                      </Tooltip>
                      <Tooltip title="Delete">
                        <IconButton
                          edge="end"
                          aria-label="delete"
                          onClick={() => handleRemoveVariable(index)}
                        >
                          <DeleteIcon />
                        </IconButton>
                      </Tooltip>
                    </Box>
                  )}
                </ListItemSecondaryAction>
              </ListItem>
              {index < variables.length - 1 && <Divider />}
            </React.Fragment>
          ))}
          {variables.length === 0 && (
            <ListItem>
              <ListItemText
                primary={
                  <Typography color="text.secondary" align="center">
                    No variables added yet
                  </Typography>
                }
              />
            </ListItem>
          )}
        </List>
      </Paper>
    </Box>
  );
};

export default ServiceVariables;
