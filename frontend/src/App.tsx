import React from 'react';

import { BrowserRouter } from "react-router-dom";
import { Route, Routes } from "react-router-dom";

import './App.css';
import { Login } from './pages/login/Login';
import { Addmenu } from './pages/admin/Addmenu';





function App() {
  return (
    <BrowserRouter>
      <div className="router ">
        <Routes>
          <Route path="/login" element={<Login />} />
          <Route path="/addmenu" element={<Addmenu />} />
        </Routes>
      </div>
    </BrowserRouter>
  );
}

export default App;
