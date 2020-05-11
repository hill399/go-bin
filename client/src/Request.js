import React, { useState } from 'react';
import './App.css';


const Request = () => {

  const API_URL = "http://localhost:8080"

  const [requestKey, setRequestKey] = useState(null);

  const updateField = e => {
    setRequestKey(e.target.value);
  }

  const requestData = () => {
    return fetch(API_URL + "/request/" + requestKey, {
      method: "GET",
      mode: "cors",
      headers: {
        "Content-Type": "application/json"
      },
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
          Key: <input type="text" name="requestKey" onKeyDown={updateField} />
        </label>
        <input type="submit" value="Get" onClick={() => requestData()} />
      </form>
    </div>
  );
}

export default Request;
