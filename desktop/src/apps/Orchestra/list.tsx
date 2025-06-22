import React, { useEffect, useRef } from 'react';
import * as go from 'gojs';
import {
  Box,
  AppBar,
  Toolbar,
  Typography,
  IconButton,
  Tooltip,
} from '@mui/material';
import {
  Add as AddIcon,
  Save as SaveIcon,
} from '@mui/icons-material';

const ListView: React.FC = () => {
  const diagramRef = useRef<HTMLDivElement>(null);
  const diagram = useRef<go.Diagram>();

  useEffect(() => {
    if (!diagramRef.current || diagram.current) return;

    const $ = go.GraphObject.make;
    diagram.current = $(go.Diagram, diagramRef.current, {
      initialContentAlignment: go.Spot.Center,
      "undoManager.isEnabled": true,
      layout: $(go.TreeLayout, { angle: 90, layerSpacing: 35 }),
    });

    // Node template
    diagram.current.nodeTemplate = $(
      go.Node,
      "Auto",
      $(
        go.Shape,
        "RoundedRectangle",
        { fill: "white" },
        new go.Binding("fill", "color")
      ),
      $(
        go.Panel,
        "Vertical",
        { margin: 8 },
        $(
          go.TextBlock,
          { margin: 8, font: "bold 14px sans-serif" },
          new go.Binding("text", "name")
        ),
        $(
          go.TextBlock,
          { margin: 8 },
          new go.Binding("text", "command")
        )
      )
    );

    // Link template
    diagram.current.linkTemplate = $(
      go.Link,
      { routing: go.Link.AvoidsNodes, corner: 5 },
      $(go.Shape),
      $(go.Shape, { toArrow: "Standard" })
    );

    // Initial data
    diagram.current.model = new go.GraphLinksModel(
      [
        { key: 1, name: "Start", command: "echo 'start'", color: "lightgreen" },
        { key: 2, name: "Process", command: "echo 'process'", color: "lightblue" },
        { key: 3, name: "End", command: "echo 'end'", color: "pink" },
      ],
      [
        { from: 1, to: 2 },
        { from: 2, to: 3 },
      ]
    );

    return () => {
      if (diagram.current) {
        diagram.current.div = null;
      }
    };
  }, []);

  const handleSave = () => {
    if (!diagram.current) return;
    const json = diagram.current.model.toJson();
    console.log('Saving diagram:', json);
  };

  return (
    <Box sx={{ 
      display: 'flex', 
      flexDirection: 'column', 
      height: '100%',
      bgcolor: 'background.default'
    }}>
      <AppBar position="static" color="default" elevation={1}>
        <Toolbar variant="dense">
          <Typography variant="h6" sx={{ flexGrow: 1 }}>
            Flow Editor
          </Typography>
          <Tooltip title="Save Flow">
            <IconButton onClick={handleSave} size="small">
              <SaveIcon />
            </IconButton>
          </Tooltip>
        </Toolbar>
      </AppBar>

      <Box 
        ref={diagramRef} 
        sx={{ 
          flex: 1,
          bgcolor: 'background.paper'
        }} 
      />
    </Box>
  );
};

export default ListView;
