import React, { useState } from 'react';
import './App.css';


const Submit = () => {

  const API_URL = "http://localhost:8080"

  const [inputData, setInputData] = useState(null);

  const updateField = e => {
    setInputData(e.target.value);
  }

  const submitData = () => {
    return fetch(API_URL + "/submit", {
      method: "POST",
      mode: "cors",
      headers: {
        "Content-Type": "application/json"
      },
      body: JSON.stringify({ Data: inputData })
    })
    .then(response => response.json())
    .then(data => { console.log(data) })
    .catch(err => {
      handleError(err)
    });
  }

  const handleError = (error) => {
    console.log(error.message);
  }

  return (
    <div className="App">
      <form>
        <label>
          Data: <input type="text" name="inputData" onKeyDown={updateField} />
        </label>
        <input type="submit" value="Go" onClick={() => submitData()} />
      </form>
    </div>
  );
}

export default Submit;
