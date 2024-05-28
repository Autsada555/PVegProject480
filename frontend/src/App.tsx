import React from 'react';

import { BrowserRouter } from "react-router-dom";
import { Route, Routes } from "react-router-dom";

import './App.css';
import { Login } from './pages/login/Login';
import { Addmenu } from './pages/admin/Addmenu';
import { Homepage } from './pages/homepage/homepage';
import { Customer } from './pages/admin/Customer';




function App() {
  return (
    <BrowserRouter>
      <div className="router ">
        <Routes>
          <Route path="/login" element={<Login />} />
          <Route path="/addmenu" element={<Addmenu />} />
          <Route path="/homepage" element={<Homepage />} />
          <Route path="/customer" element={<Customer />} />

        </Routes>
      </div>
    </BrowserRouter>
  );
}

export default App;
