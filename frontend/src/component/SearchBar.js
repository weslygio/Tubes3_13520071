import React from 'react'
import Output from './Output'

const SearchBar = ({data, data2}) => {
  return (
    <div className="search">
        <div className="searchInput">
            <input type="text" className="input-search"/>
        </div>
        <div className='searchButton'>
          <input type = 'submit' value='Submit' className='btn-submit-searching'/>         
        </div>
        {/*
          <div className="dataResult">
              <Output text={data} text2={data2}/>
              <Output text={data} text2={data2}/>
              <Output text={data} text2={data2}/>
          </div>        
        */}
    </div>
  )
}

export default SearchBar