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

const DiseaseForm = () => {
    const [name, setName] = useState("");
    const [selectedFile, setSelectedFile] = useState("");
    var hi = "lol";

    const submitForm = () => {
    
     
        console.log('tessubmit');
        console.log(name);
        console.log(selectedFile);
        console.log(hi);
        //formData.append("namaPenyakit", name);
        //formData.append("dnaSequence", selectedFile);

        const json = JSON.stringify({ "namaPenyakit": name, "dnaSequence": selectedFile });
        //const res = await axios.post('https://httpbin.org/post', json);
        
        //const json = JSON.stringify({ "userId": 1,"id": 2,"title": selectedFile,"body": name });
        console.log(json);
      
        axios
            .post('http://localhost:8080/diseases', json)
            .then((res) => {
                console.log(res.data);
                console.log(res.status);
                console.log(res.statusText);
                console.log(res.headers);
                console.log(res.config);
            })
            .catch((err) => console.log(err));       

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
            </form>
        </div>
        <div className="output1"><p></p></div>
    </div>
    )
}

export default DiseaseForm