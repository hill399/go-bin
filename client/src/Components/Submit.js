import React from 'react';
import { TextField } from '@material-ui/core';

const Submit = (props) => {

  const { dataState, setDataState, radioState } = props

  const isDisabled = () => {
    if (radioState === 'submit') {
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
          <TextField name="data"
            style={{ width: '80%' }} 
            margin="normal"
            multiline={true}
            rows={6}
            label="Data" 
            variant="outlined" 
            value={dataState.data} 
            disabled={isDisabled()} 
            onChange={updateField} />
      </div>
  );
}

export default Submit;
