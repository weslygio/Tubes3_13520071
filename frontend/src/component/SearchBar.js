import React, { useState } from 'react';
import Output from './Output';
import axios from 'axios';

const SearchBar = () => {
  const [request, setRequest] = useState("");
  const [data, setData] = useState("{}");

  const submitForm = async event => {
    event.preventDefault();
    console.log(request);
    console.log(encodeURI(request));
    console.log(request);
    // console.log(selectedFile);

    // const json = JSON.stringify({ "namaPenyakit": name, "dnaSequence": selectedFile });
    // console.log(json);
  
    // axios
    //     .post('http://localhost:8080/diseases', json)
    //     .then((res) => {
    //         console.log(res.data);
    //         console.log(res.status);
    //         console.log(res.statusText);
    //         console.log(res.headers);
    //         console.log(res.config);
    //         console.log("hi");
    //     })
    //     .catch((err) => console.log(err));       
    axios.get('http://localhost:8080/logs/' + encodeURI(request)).then((res) => {console.log(res.data);
    console.log("hi"); setData(res.data);}).catch((err) => {console.log(err.response.statusText); console.log(err.response.data);});

  }

  return (
    <div className="search">
        <div className="searchInput">
            <input type="text" placeHolder="<date><space><disease-name>" className="input-search" value = {request} onChange={(e) => setRequest(e.target.value)}/>
        </div>
        <div className='searchButton'>
          <input type = 'submit' value='Submit' className='btn-submit-searching' onClick = {submitForm}/>         
        </div>
          {data != "{}" ? data.map((d) => <Output text={d.tanggal + " - " + d.namaPasien + " - " + d.namaPenyakit + " - " + (d.kemiripan * 100).toFixed(2) + "% - " + JSON.stringify(d.hasil)}/>) : null}
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