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
            <Output text= {'1 April 2022 - Mhs IF - HIV - 75% - False'}/>
            <Output text= {'1 April 2022 - Mhs IF - HIV - 75% - False'}/>
        </div>
    </div>
  )
}

export default DNATest