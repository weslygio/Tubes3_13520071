import React from 'react'
import Header from './component/Header'
import Logo from './component/Logo'
import Form from './component/Form'
import Output from './component/Output'

const DNATest = () => {
  return (
    <div>
        <Header/>
        <div className="otherpages">
            <Logo/>
            <Form/>
        </div>
    </div>
  )
}

export default DNATest