import React, { useState } from 'react';
import './App.css';
import Submit from './Submit'
import Request from './Request'

import { FormControl, FormControlLabel, FormLabel, RadioGroup, Radio, Button } from '@material-ui/core';


const App = () => {

  const api_url = process.env.REACT_APP_API_URL;

  const [radioState, setRadioState] = useState('submit');

  const [dataState, setDataState] = useState({
    data: '',
    hash: '',
  })

  const handleRadio = (e) => {
    setRadioState(e.target.value);
    setDataState({
      data: '',
      hash: '',
    })
  }

  const directButtonFunc = () => {
    switch (radioState) {
      case 'submit':
        return fetch(api_url + "/submit", {
          method: "POST",
          mode: "cors",
          headers: {
            "Content-Type": "application/json"
          },
          body: JSON.stringify({ Data: dataState.data })
        })
          .then(response => response.json())
          .then(data => {
            setDataState({
              ...dataState,
              hash: data.Hash
            });
          })
          .catch(err => {
            handleError(err)
          });
      case 'request':
        return fetch(api_url + "/request/" + dataState.hash, {
          method: "GET",
          mode: "cors",
          headers: {
            "Content-Type": "application/json"
          },
        })
          .then(response => response.json())
          .then(data => {
            setDataState({
              ...dataState,
              data: data.Data
            });
          })
          .catch(err => {
            handleError(err)
          });
      default:
        return null
    }
  }

  const handleError = (error) => {
    console.log(error.message);
  }

  return (
    <div className="App">
      <h1> Go-Bin Landing Page </h1>
      <FormControl component="fieldset">
        <FormLabel component="legend"> Select Function </FormLabel>
        <RadioGroup row aria-label="function" name="radioState" id="radioState" value={radioState} onChange={handleRadio}>
          <FormControlLabel name="radioState" value="submit" control={<Radio />} label="Submit" />
          <FormControlLabel name="radioState" value="request" control={<Radio />} label="Request" />
        </RadioGroup>
      </FormControl>
      <form>
        <Request dataState={dataState} setDataState={setDataState} radioState={radioState} />
        <Submit dataState={dataState} setDataState={setDataState} radioState={radioState} />
        <Button variant="contained" margin="normal" onClick={() => directButtonFunc()}> Go </Button>
      </form>
    </div>
  );
}

export default App;
