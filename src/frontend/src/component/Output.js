import React from 'react'

const Output = ({text}) => {
  return (
    <div className='outputrow'>
        <div className="output1"><p></p></div>
        <div className="output2"><p>{ text }</p></div>
        <div className="output1"><p></p></div>
    </div>
  )
}

export default Output