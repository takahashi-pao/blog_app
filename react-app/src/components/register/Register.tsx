import axios from 'axios';
import React, { useEffect, useState } from 'react';
import FileUploadComponent from './parts/register';
import { Link, NavigateFunction, useNavigate } from 'react-router-dom';

type auth_response_props = {
  message: string,
  error: string,
  userId: string,
  isSignIn: boolean
}

function Register() {

  const [isLoggedIn, setIsLoggedIn] = useState(false);
  const navigate: NavigateFunction = useNavigate()

  // サーバーからログイン状態とユーザーIDを取得する非同期関数
  const fetchUserStatus = async () => {
    try {
      // ログイン状態とユーザーIDを取得するAPIエンドポイントへのリクエストを実行
      const response = await fetch('http://localhost:8080/IsLogin');
      const data: auth_response_props = await response.json();
      console.log(data)
      if(data.error != ""){
        console.log(data.error)
        return
      }
      // 取得したデータをstateに設定
      setIsLoggedIn(data.isSignIn);
    } catch (error) {
      console.error('データを取得できませんでした:', error);
    }
  };

  useEffect(() => {
    // コンポーネントがマウントされたときにログイン状態を取得
    fetchUserStatus();
    if (!isLoggedIn) {      
      navigate("/")
    }
  }, []);

  return (
    <div>
      <Link to="/">Go To Home</Link>
      <FileUploadComponent key="UploadComponent"></FileUploadComponent>
    </div>
    
  );
}

export default Register;
