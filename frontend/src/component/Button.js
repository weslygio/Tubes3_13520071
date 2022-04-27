import React from 'react'
import { Link } from "react-router-dom";

const Button = ({ text, color, link }) => {
  return (
    <button style ={{backgroundColor: color }} className="button"><Link to= { link }>{ text }</Link></button>
  )
}

export default Button