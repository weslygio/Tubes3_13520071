import React, { useState } from 'react';
import Output from './Output';
import axios from 'axios';

const SearchBar = () => {
  const [request, setRequest] = useState("");
  const [data, setData] = useState("{}");

  const submitForm = async event => {
    event.preventDefault();
      
    axios.get('http://localhost:8080/logs/' + encodeURI(request)).then((res) => setData(res.data)).catch((err) => {});

  }

  return (
    <div className="search">
        <div className="searchInput">
            <input type="text" placeHolder="<date><space><disease-name>" className="input-search" value = {request} onChange={(e) => setRequest(e.target.value)}/>
        </div>
        <div className='searchButton'>
          <input type = 'submit' value='Submit' className='btn-submit-searching' onClick = {submitForm}/>         
        </div>
          {data !== "{}" ? data.map((d) => <Output text={d.tanggal + " - " + d.namaPasien + " - " + d.namaPenyakit + " - " + (d.kemiripan * 100).toFixed(2) + "% - " + JSON.stringify(d.hasil)}/>) : null }
          {data == "" ? <Output text={"No data match the search"}/> : null}
    </div>
  )
}

export default SearchBar