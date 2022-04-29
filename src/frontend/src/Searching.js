
import React from 'react'
import Header from './component/Header'
import Logo from './component/Logo'
import SearchBar from './component/SearchBar'
const Searching = () => {
  return (
    <div>
        <Header/>
        <div className="otherpages">
            <Logo/>
            <SearchBar/>
        </div>
    </div>
  )
}
export default Searching