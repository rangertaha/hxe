import React from "react";
import { Box, Typography } from "@mui/material";

const App: React.FC = () => {
  return (
    <Box
      sx={{
        display: "flex",
        flexDirection: "column",
        height: "100%",
      }}
    >
      {/* Header - fixed height */}
      <Box sx={{ flex: "0 0 48px", backgroundColor: "red" }}>
        <Typography
          variant="h5"
          component="h3"
          sx={{ p: 1, backgroundColor: "green" }}
        >
          Dashboard
        </Typography>
      </Box>

      {/* Content - takes remaining space */}
      <Box sx={{ flex: "1 1 auto", backgroundColor: "blue" }}>Content</Box>

      {/* Footer - fixed height */}
      <Box sx={{ p: 1, flex: "0 0 48px", backgroundColor: "green" }}>
        Footer
      </Box>
    </Box>
  );
};

export default App;
