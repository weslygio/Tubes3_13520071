import React from 'react'
import { Link } from "react-router-dom";

const Button = ({ text, color, link }) => {
  return (
    <button style ={{backgroundColor: color, boxShadow:"2px 2px 5px 0px #888888"}} className="button"><Link to= { link }>{ text }</Link></button>
  )
}

export default Button