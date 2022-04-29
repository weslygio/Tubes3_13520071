import React from 'react'
import { Link } from "react-router-dom";

const Header = () => {
  return (
    <header>
        <ul className="menu">
        <li><Link to="/">Home</Link></li>
        <li><Link to="/inputdisease">Input Disease</Link></li>
        <li><Link to="/dnatest">DNA Test</Link></li>
        <li><Link to="/searching">Searching</Link></li>
        <li className="about-us">developed by sangci reborn :v</li>
        </ul>
    </header>
  )
}

export default Header