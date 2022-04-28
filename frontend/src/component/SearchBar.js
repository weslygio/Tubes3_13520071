import React from 'react'
import Output from './Output'

const SearchBar = ({data}) => {
  return (
    <div className="search">
        <div className="searchInput">
            <input type="text" className="input-search"/>
        </div>
        <div className='searchButton'>
          <input type = 'submit' value='Submit' className='btn-submit-searching'/>         
        </div>
        <div className="dataResult">
            <Output text={data}/>
            <Output text={data}/>
            <Output text={data}/>
        </div>
    </div>
  )
}

export default SearchBar