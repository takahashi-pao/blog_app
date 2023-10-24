import axios from 'axios';
import React, { useEffect, useState } from 'react';
import FileUploadComponent from './parts/register';
import { Link } from 'react-router-dom';

function Add() {

  useEffect(() => {

  }, []);

  return (
    <div>
      <Link to="/">Go To Home</Link>
      <FileUploadComponent key="UploadComponent"></FileUploadComponent>
    </div>
    
  );
}

export default Add;
