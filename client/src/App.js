import React, { useState } from 'react';
import './App.css';
import Submit from './Submit'
import Request from './Request'

const App = () => {

  return (
    <div className="App">
      <h1> Go-Bin Landing Page </h1>

      <Submit />
      <Request />

    </div>
  );
}

export default App;
