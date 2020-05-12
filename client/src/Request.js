import React from 'react';
import './App.css';
import { TextField } from '@material-ui/core';


const Request = (props) => {

  const { dataState, setDataState, radioState } = props

  const isDisabled = () => {
    if (radioState === 'request') {
      return false
    } else {
      return true
    }
  }

  const updateField = e => {
    setDataState({
      ...dataState,
      [e.target.name]: e.target.value
    });
  }

  return (
    <div>
      <TextField 
      style={{ width: '80%' }} 
      margin="normal" 
      name="hash" 
      label="Hash" 
      variant="outlined" 
      value={dataState.hash} 
      disabled={isDisabled()} 
      onChange={updateField} />
    </div>
  );
}

export default Request;
