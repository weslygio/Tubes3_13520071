import React from 'react'
import Header from './component/Header'
import Logo from './component/Logo'
import Button from './component/Button'

const Home = () => {
  return (
    <div>
        <Header/>
        <div className="homeLogo">
            <Logo/>
        </div>
        <div className="homeBtn">
            <div className="btn1"><Button color = {'#FF7043'} text = {'DNA Test'} link = {'/dnatest'}/></div>
            <div><Button color = {'#FF7043'} text = {'Search Test Result'} link = {'/searching'}/></div>
        </div>
    </div>
  )
}

export default Home