import React, { useState } from 'react';
import axios from 'axios';
import Output from './Output';

const DiseaseForm = () => {
    const [name, setName] = useState("");
    const [selectedFile, setSelectedFile] = useState("");
    const [status, setStatus] = useState("");

    const submitForm = async event => {
        event.preventDefault();

        const json = JSON.stringify({ "namaPenyakit": name, "dnaSequence": selectedFile });
   
        axios.post('http://localhost:8080/diseases', json).then((res) => setStatus(res.status)).catch((err) => setStatus(err.response.status));

    }

    function processFile(e) {
        var file = e.target.files[0];
        var reader = new FileReader();
        reader.onload = function(e) {
            console.log(e.target.result)
            setSelectedFile(e.target.result)
        };
        reader.readAsText(file);
    }

    function printMessage(e) {
        if (e === 406){
            return "The DNA sequence is not valid"          
        } else if (e === 409){
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
                <div>{status !== "" ? <Output text={printMessage(status)} /> : null}</div>
            </form>
        </div>
        <div className="output1"><p></p></div>
    </div>
    )
}

export default DiseaseForm