import React from 'react';
import ReactDOM from 'react-dom/client';
import Home from './Home'
import reportWebVitals from './reportWebVitals';
import {
  BrowserRouter,
  Routes,
  Route,
} from "react-router-dom";
import DNATest from './DNATest';
import Searching from './Searching';
import InputDisease from './InputDisease';
import './index.css';

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


// If you want to start measuring performance in your app, pass a function
// to log results (for example: reportWebVitals(console.log))
// or send to an analytics endpoint. Learn more: https://bit.ly/CRA-vitals
reportWebVitals();
