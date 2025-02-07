import React from "react";
import ReactDOM from "react-dom/client"; // Updated import for React 18
import { BrowserRouter as Router, Route, Routes } from "react-router-dom";
import "./index.css";

const rootElement = document.getElementById("root");

// Create a root using ReactDOM.createRoot
const root = ReactDOM.createRoot(rootElement);

import Welcome from "./pages/welcome";


root.render(
    <React.StrictMode>
        <Router>
            <Routes>

             
             <Route path="/" element={<Welcome />} />


             
            </Routes>
        </Router>
    </React.StrictMode>
);
