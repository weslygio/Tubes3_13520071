import React from 'react';
import ReactDOM from 'react-dom/client';
import Home from './Home'
import DNATest from './DNATest';
import Searching from './Searching';
import InputDisease from './InputDisease';
import './index.css';
import {
  BrowserRouter,
  Routes,
  Route,
} from "react-router-dom";

const root = ReactDOM.createRoot(document.getElementById('root'));
root.render(
  <React.StrictMode>
    <BrowserRouter>
    <Routes>
      <Route path="/" element={<Home />}/>
      <Route path="inputdisease" element={<InputDisease />}/>
      <Route path="dnatest" element={<DNATest />}/>
      <Route path="searching" element={<Searching />}/>
    </Routes>
    </BrowserRouter>
  </React.StrictMode>
);

