import React, { useState } from 'react';
import Output from './Output';
import axios from 'axios';

const SearchBar = () => {
  const [request, setRequest] = useState("");
  const [data, setData] = useState("{}"); // nerima data 
  const [status, setStatus] = useState(""); // nerima status error

  const submitForm = async event => {
    event.preventDefault();
      
    axios.get('http://localhost:8080/logs/' + encodeURI(request)).then((res) => {setData(res.data); setStatus(res.status);}).catch((err) => {setStatus(err.response.status);});
    console.log(data);
    console.log(status);
  }
  // kalo res status 406 = Query invalid, res status 200 = data ada/engga 
  return (
    <div className="search">
        <div className="searchInput">
            <input type="text" placeHolder="<date><space><disease-name>" className="input-search" value = {request} onChange={(e) => setRequest(e.target.value)}/>
        </div>
        <div className='searchButton'>
          <input type = 'submit' value='Submit' className='btn-submit-searching' onClick = {submitForm}/>         
        </div>
          {status === 406 ? <Output text={"Query invalid"} /> : 
          (data.length !== 0 ? data.map((d) => <Output text={d.tanggal + " - " + d.namaPasien + " - " + d.namaPenyakit + " - " + (d.kemiripan * 100).toFixed(2) + "% - " + JSON.stringify(d.hasil)}/>) : 
          (data == "" ? <Output text={"No data match the search"}/> : null))}
    </div>
  )
}

export default SearchBar