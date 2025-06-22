import React from 'react';
import { Box } from '@mui/material';

interface DagFormProps {
  open: boolean;
  onClose: () => void;
  onSubmit?: () => void;
}

const DagForm: React.FC<DagFormProps> = () => {
  return <Box />;
};

export default DagForm; 