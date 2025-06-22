import React from "react";
import { Routes, Route } from "react-router-dom";
import ServicesLayout from "./layout";
import ServicesList from "./list";
import ServiceForm from "./form";
import ServiceDetail from "./detail";

const Services: React.FC = () => {
  return (
    <ServicesLayout>
      <Routes>
        <Route path="/" element={<ServicesList />} />
        <Route path="/new" element={<ServiceForm />} />
        <Route path="/:id" element={<ServiceDetail />} />
      </Routes>
    </ServicesLayout>
  );
};

export default Services;


