import React from 'react'
import Output from './Output'

const SearchBar = ({data}) => {
  return (
    <div className="search">
        <div className="searchInput">
            <input type="text" />
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