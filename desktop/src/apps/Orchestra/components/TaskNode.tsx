import React, { memo } from 'react';
import { Handle, Position, NodeProps } from 'reactflow';
import { Paper, Typography, Box, IconButton } from '@mui/material';
import { Settings, PlayArrow, Stop } from '@mui/icons-material';

interface TaskData {
  label: string;
  type: string;
  command: string;
  status?: 'idle' | 'running' | 'completed' | 'failed';
}

const TaskNode: React.FC<NodeProps<TaskData>> = ({ data }) => {
  return (
    <Paper 
      elevation={2}
      sx={{
        p: 1,
        minWidth: 180,
        bgcolor: data.status === 'running' ? 'success.light' :
                data.status === 'failed' ? 'error.light' :
                data.status === 'completed' ? 'info.light' : 'background.paper'
      }}
    >
      <Handle type="target" position={Position.Top} />
      
      <Box sx={{ display: 'flex', alignItems: 'center', mb: 1 }}>
        <Typography variant="subtitle2" sx={{ flex: 1 }}>
          {data.label}
        </Typography>
        <IconButton size="small">
          <Settings fontSize="small" />
        </IconButton>
      </Box>

      <Typography variant="caption" display="block" color="text.secondary">
        Type: {data.type}
      </Typography>
      <Typography 
        variant="caption" 
        display="block" 
        color="text.secondary"
        sx={{ 
          whiteSpace: 'nowrap',
          overflow: 'hidden',
          textOverflow: 'ellipsis',
          maxWidth: '100%'
        }}
      >
        Command: {data.command}
      </Typography>

      <Box sx={{ display: 'flex', justifyContent: 'flex-end', mt: 1 }}>
        <IconButton size="small" color="success">
          <PlayArrow fontSize="small" />
        </IconButton>
        <IconButton size="small" color="error">
          <Stop fontSize="small" />
        </IconButton>
      </Box>

      <Handle type="source" position={Position.Bottom} />
    </Paper>
  );
};

export default memo(TaskNode); 