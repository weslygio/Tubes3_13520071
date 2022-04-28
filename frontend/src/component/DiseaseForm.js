/*

import React from 'react'

const DiseaseForm = () => {
  return (
    <div className='formrow'>
        <div className="output1"><p></p></div>
        <div className="output2-form">
            <form className='add-form' >
                <div className='form-control'>
                    <p><label>Disease Name</label></p>
                    <p><input type = 'text' placeholder='Name' className='input-form'/></p>
                </div>
                <div className='form-control form-control check'>
                    <p><label>DNA Sequence</label></p>
                    <p><input type = 'file' /></p>
                </div>
                <input type = 'submit' value='Submit' className='btn-submit'/>
            </form>
        </div>
        <div className="output1"><p></p></div>
    </div>
  )
}

export default DiseaseForm

*/
import React, { useState } from 'react';
import axios from 'axios';
import Output from './Output';
import { render } from '@testing-library/react';

const DiseaseForm = () => {
    const [name, setName] = useState("");
    const [selectedFile, setSelectedFile] = useState("");
    const [status, setStatus] = useState("");

    const printState = (s) => {
        console.log(s);
    }

    const submitForm = async event => {
        event.preventDefault();

        console.log(name);
        console.log(selectedFile);

        const json = JSON.stringify({ "namaPenyakit": name, "dnaSequence": selectedFile });
        //console.log(json);
      
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
        axios.post('http://localhost:8080/diseases', json).then((res) => {printState(res.status);
        console.log("hi"); setStatus(res.status);}).catch((err) => {console.log(err.response.statusText); console.log(err.response.data); setStatus(err.response.status);});

    }

    function processFile(e) {
        console.log('helloFile');
        var file = e.target.files[0];
        var reader = new FileReader();
        reader.onload = function(e) {
            // The file's text will be printed here
            console.log(e.target.result)
            setSelectedFile(e.target.result)
        };
        
        reader.readAsText(file);
    }

    function printMessage(e) {
        if (e == "406"){
            return "The DNA sequence is not valid"          
        } else if (e == "409"){
            return "The disease has been inputed before"
        } else {
            return "Disease successfully inputed"        
        }

    }

    return (
    <div className='formrow'>
        <div className="output1"><p></p></div>
        <div className="output2-form">
            <form className='add-form' >

                <div className='form-control'>
                    <p><label>Disease Name</label></p>
                    <p><input type = 'text' placeholder='Name' className='input-form' value = {name} onChange={(e) => setName(e.target.value)}/></p>
                </div>

                <div className='form-control form-control check'>
                    <p><label>DNA Sequence</label></p>
                    <p><input type = 'file' onChange={(e) => processFile(e)}/></p>
                </div>
                <input type = 'submit' className='btn-submit' onClick = {submitForm}/>
                <div>{status != "" ? <Output text={printMessage(status)} /> : null}</div>
            </form>
        </div>
        <div className="output1"><p></p></div>
    </div>
    )
}

export default DiseaseForm