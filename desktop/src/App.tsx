import React, { useEffect, useState } from "react";
import { BrowserRouter as Router, Routes, Route } from "react-router-dom";
import { ThemeProvider } from "@mui/material/styles";
import CssBaseline from "@mui/material/CssBaseline";
import { Box } from "@mui/material";
import cyberpunkTheme from "./theme";
import Sidebar from "./components/Sidebar";
import Layout from "./components/Layout";
import StatusBar from "./components/StatusBar";
import { Dashboard, Settings, Services, Orchestra, Chart } from "./apps";
import { useKeyboardShortcuts } from "./hooks/useKeyboardShortcuts";
import { getFontSize } from "./utils/fontSize";
import "./App.css";

const App: React.FC = () => {
  const [sidebarOpen, setSidebarOpen] = useState(true);
  const [appTitle, setAppTitle] = useState("");

  useKeyboardShortcuts();

  useEffect(() => {
    // Initialize font size from localStorage
    const fontSize = getFontSize();
    document.documentElement.style.fontSize = `${fontSize}px`;
  }, []);

  const toggleSidebar = () => {
    setSidebarOpen(!sidebarOpen);
  };

  return (
    <ThemeProvider theme={cyberpunkTheme}>
      <CssBaseline />
      <Router>
        <Box sx={{ display: "flex", height: "100vh" }}>
          <Sidebar open={sidebarOpen} onClose={toggleSidebar} />
          <Box
            sx={{
              flex: 1,
              display: "flex",
              flexDirection: "column",
              height: "calc(100vh - 28px)",
              overflow: "hidden",
            }}
          >
            <Layout title={appTitle} setAppTitle={setAppTitle}>
              <Routes>
                <Route path="/" element={<Dashboard />} />
                <Route path="/services/*" element={<Services />} />
                <Route path="/settings/*" element={<Settings />} />
                <Route path="/orchestra" element={<Orchestra />} />
                <Route path="/chart" element={<Chart />} />
              </Routes>
            </Layout>
          </Box>
          <StatusBar onMenuClick={toggleSidebar} />
        </Box>
      </Router>
    </ThemeProvider>
  );
};

export default App;
