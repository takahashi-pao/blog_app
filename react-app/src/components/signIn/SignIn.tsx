import React, { ChangeEvent, useState } from 'react';
import { Link, useNavigate } from 'react-router-dom';

function SignIn() {
  const [userId, setUserId] = useState(''); // IDを管理
  const [password, setPassword] = useState(''); // パスワードを管理
  const [isPasswordPage, setPasswordPage] = useState(false); // パスワード入力画面の状態

  const navigate = useNavigate();

  const handleUserIdChange = (event: ChangeEvent<HTMLInputElement>) => {
    // ID入力欄の値が変更されたときに呼び出される関数
    setUserId(event.target.value);
  };

  const handleNextClick = () => {
    // ログインボタンがクリックされたときの処理
    if (userId) {
      // IDが入力されていれば、パスワード入力画面に遷移
      setPasswordPage(true);
    }
  };

  const handlePasswordChange = (event: ChangeEvent<HTMLInputElement>) => {
    // パスワード入力欄の値が変更されたときに呼び出される関数
    setPassword(event.target.value);
  };

  const handleSignIn = async() => {
    // ログインボタンがクリックされたときの処理
    if (password) {
      // パスワードが入力されていれば、確認リクエストを行う
      const formData = new FormData();
      formData.append('id', userId);
      formData.append('password', password);

      try {
        const response = await fetch('http://localhost:8080/SignIn', {
          method: 'POST',
          body: formData,
        });       

        const data = await response.json();
        console.log(data)
        navigate("/")
      } catch (error) {
        console.error('IDかパスワードが違います', error);        
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
          <button onClick={handleSignIn}>ログイン</button> {/* ログインボタン */}
        </div>
      ) : (
        <div>
          <input
            placeholder="IDを入力"
            value={userId}
            onChange={handleUserIdChange}
          />
          <Link to="/signUp">新規登録</Link>
          <button onClick={handleNextClick}>次へ</button> {/* 次へボタン */}
        </div>
      )}
    </div>
  );
}

export default SignIn;
