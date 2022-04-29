import React, { useState } from 'react';
import Output from './Output'
import axios from 'axios';

const Form = () => {
    const [name, setName] = useState(""); // namaPasien
    const [disease, setDisease] = useState(""); // namaPenyakit
    const [selectedFile, setSelectedFile] = useState(""); // dnaSequence
    const [result, setResult] = useState("{}");
    const [error, setError] = useState("");

    function processFile(e) {
        var file = e.target.files[0];
        var reader = new FileReader();
        reader.onload = function(e) {
            console.log(e.target.result)
            setSelectedFile(e.target.result)
        };
        
        reader.readAsText(file);
    }

    const submitForm = async event => {
        event.preventDefault();

        const json = JSON.stringify({ "namaPasien": name, "dnaSequence": selectedFile, "namaPenyakit": disease });
       
        axios.post('http://localhost:8080/logs', json).then((res) => setResult(res.data)).catch((err) => {setError(err.response.status);});

    }

    function printMessage(e) {
        if (e === 406){
            return "The DNA sequence is not valid"          
        } else if (e === 424){
            return "The disease has not been inputed before"
        } 

    }

    return (
        <div>
        <div className='formrow'>
            <div className="output1"><p></p></div>
            <div className="output2-form">
                <form className='add-form' >
                    <div className='form-control'>
                        <p><label>Name</label></p>
                        <p><input type = 'text' placeholder='Name' className='input-form' value = {name} onChange={(e) => setName(e.target.value)}/></p>
                    </div>
                    <div className='form-control'>
                        <p><label>Disease Prediction</label></p>
                        <p><input type = 'text' placeholder='Disease Prediction' className='input-form' value = {disease} onChange={(e) => setDisease(e.target.value)}/></p>
                    </div>
                    <div className='form-control form-control check'>
                        <p><label>DNA Sequence</label></p>
                        <p><input type = 'file' onChange={(e) => processFile(e)} /></p>
                    </div>
                    <input type = 'submit' value='Submit' className='btn-submit' onClick = {submitForm}/>
                </form>
            </div>
            <div className="output1"><p></p></div>
        </div>
        <div>
            <div className='tes'>{result !== "{}" ? <Output text={result.tanggal + " - " + result.namaPasien + " - " + result.namaPenyakit + " - " + (result.kemiripan * 100).toFixed(2) + "% - " + JSON.stringify(result.hasil)} /> : 
            (error !== "" ? <Output text={printMessage(error)}/> : null )}</div>
        </div>
        </div>
    )
}

export default Form