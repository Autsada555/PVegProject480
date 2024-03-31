import React from 'react';

import { BrowserRouter } from "react-router-dom";
import { Route, Routes } from "react-router-dom";

import './App.css';
import Login from './pages/login/login';


function App() {
  return (
    <BrowserRouter>
      <div className="router ">
        <Routes>
          <Route path="/login" element={<Login />} />d
        </Routes>
      </div>
    </BrowserRouter>
  );
}

export default App;
