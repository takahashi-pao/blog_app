import React, { createContext, useEffect, useContext, useState } from 'react';
import { Link, useNavigate } from 'react-router-dom';
import { useIsDeleteMode } from './isDeleteMode';
import { useClickedIds } from '../home/parts/article_click';

type auth_response_props = {
  message: string,
  error: string,
  userId: string,
  isSignIn: boolean
}

const UserStatus = () => {
  const [isLoggedIn, setIsLoggedIn] = useState(false);
  const [userId, setUserId] = useState('');
  const {isDeleteMode, setIsDeleteMode} = useIsDeleteMode();
  const [deleteModeBtnLabel, setDeleteModeBtnLabel] = useState('');
  const { clickedIds, setClickedIds } = useClickedIds();

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
      setUserId(data.userId);
    } catch (error) {
      console.error('データを取得できませんでした:', error);
    }
  };

  const handleSignOut = async() => {
      // サインアウト処理を行う
      try {
        const response = await fetch('http://localhost:8080/SignOut', {
          method: 'GET',
        });       

        const data: auth_response_props = await response.json();
        if (data.error != "") {
            console.log(data.error)
            return
        }
        console.log(data.message)
        setIsLoggedIn(data.isSignIn);
        setUserId(data.userId);

      } catch (error) {
        console.error('サインアウトに失敗', error);        
      }
  };

  const handleDeleteBtnLabel = () => {
    if(isDeleteMode){
      // 削除モード→通常モード
      setIsDeleteMode(false)     
      return
    }

    // 通常モード→削除モード
    setIsDeleteMode(true)    
  }

  useEffect(() => {
    if(isDeleteMode){
      setDeleteModeBtnLabel("削除モード")
    }else{
      setDeleteModeBtnLabel("通常モード")
    }
    setClickedIds([])
  }, [isDeleteMode])

  useEffect(() => {
    // コンポーネントがマウントされたときにログイン状態を取得
    fetchUserStatus();
    setIsDeleteMode(false)
  }, []);

  return (
    <div>
      {isLoggedIn ? (
        <div>
          <div>
            <p></p>
          </div>
          
          <div>
              <p>サインイン中 - ユーザーID: {userId}</p>
              <div>            
                  <button onClick={handleSignOut}>サインアウト</button> {/* サインアウトボタン */}
              </div>
              <div>
                  <Link to="/register">Go To Add</Link>
                  <br></br>
                  <button onClick={handleDeleteBtnLabel}>{deleteModeBtnLabel}</button>
              </div>
          </div>
        </div>

      ) : (
        <div>
            <div>            
                <Link to="/signIn">サインイン</Link>
            </div>
        </div>
      )}
    </div>
  );
};

export default UserStatus;
