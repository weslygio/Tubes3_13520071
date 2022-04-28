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