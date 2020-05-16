import React from 'react';
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
      name="id" 
      label="ID" 
      variant="outlined" 
      value={dataState.id} 
      disabled={isDisabled()} 
      onChange={updateField} />
    </div>
  );
}

export default Request;
