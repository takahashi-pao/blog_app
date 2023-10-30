import React, { ChangeEvent, useState } from 'react';
import { Link, useNavigate } from 'react-router-dom';
import './css/signUp.css';

function SignUp() {
  const [userId, setUserId] = useState(''); // IDを管理
  const [password, setPassword] = useState(''); // パスワードを管理
  const [passwordConfirm, setPasswordConfirm] = useState(''); // パスワードを管理
  const [passwordMatch, setPasswordMatch] = useState(false); // パスワードを管理
  const [isPasswordPage, setPasswordPage] = useState(false); // パスワード入力画面の状態

  const navigate = useNavigate();

  const handleUserIdChange = async(event: ChangeEvent<HTMLInputElement>) => {
    // ID入力欄の値が変更されたときに呼び出される関数
    setUserId(event.target.value);
    // 入力IDに重複がないかを確認
    if (event.target.value.length > 0){
        const response = await fetch(`http://localhost:8080/CheckExistId/${event.target.value}`);
        const result = await response.json();
        
        if (result.message){
            console.log(result.message)
        } else {
            console.log(result.error)
        }
    }
  };

  const handleLoginClick = () => {
    // ログインボタンがクリックされたときの処理
    if (userId) {
      // IDが入力されていれば、パスワード入力画面に遷移
      setPasswordPage(true);
    }
  };

  const handlePasswordChange = (event: ChangeEvent<HTMLInputElement>) => {
    // パスワード入力欄の値が変更されたときに呼び出される関数
    setPassword(event.target.value);
    checkPasswordMatch(event.target.value, passwordConfirm)
  };

  const handlePasswordConfirmChange = (event: ChangeEvent<HTMLInputElement>) => {
    // パスワード(確認)入力欄の値が変更されたときに呼び出される関数
    setPasswordConfirm(event.target.value);
    checkPasswordMatch(password, event.target.value)
  };

  const checkPasswordMatch = (pass1: string, pass2: string) => {
    setPasswordMatch(pass1 === pass2);
  };

  const handleSignUp = async () =>  {
    // 入力完了したときに呼び出される関数
    if (passwordMatch) {
        const formData = new FormData();
        formData.append('id', userId);
        formData.append('password', password);
  
        try {
          const response = await fetch('http://localhost:8080/SignUp', {
            method: 'POST',
            body: formData,
          });       
  
          const data = await response.json();
          navigate("/")
        } catch (error) {
          console.error('ユーザーの登録に失敗しました', error);        
        }
      }
  };

  return (    
    <div>
      <Link to="/">HOME</Link>
      <div>
        <p>ログイン</p>
        {isPasswordPage ? (<p>{userId}</p>) : (<p></p>)}        
      </div>
      {isPasswordPage ? ( // パスワード入力画面
        <div>
          <input
            type="password"
            placeholder="パスワードを入力"
            value={password}
            onChange={handlePasswordChange}
          />
          <input
            type="password"
            placeholder="確認"
            value={passwordConfirm}
            onChange={handlePasswordConfirmChange}
            className={passwordMatch ? 'valid' : ''}
          />
          <button onClick={handleSignUp}>登録</button> {/* ログインボタン */}
        </div>
      ) : (
        <div>
          <input
            placeholder="IDを入力"
            value={userId}
            onChange={handleUserIdChange}
          />
          <button onClick={handleLoginClick}>次へ</button> {/* 次へボタン */}
        </div>
      )}
    </div>
  );
}

export default SignUp;
