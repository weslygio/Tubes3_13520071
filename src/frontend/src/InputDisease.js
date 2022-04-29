import React from 'react'
import DiseaseForm from './component/DiseaseForm'
import Header from './component/Header'
import Logo from './component/Logo'

const InputDisease = () => {
  return (
    <div>
        <Header/>
        <div className="otherpages">
            <Logo/>
            <DiseaseForm/>
        </div>
    </div>
  )
}

export default InputDisease