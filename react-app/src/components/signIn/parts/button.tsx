import React, { createContext, useContext, useState } from 'react';

// コンテキストの型情報を定義
type SignInContextType = {
  isLogin: boolean;
  setLoginState: React.Dispatch<React.SetStateAction<boolean>>;
};

const SignInContext = createContext<SignInContextType>({
  isLogin: false,  // 初期値を設定
  setLoginState: () => {},  // 初期値を設定
});

export function useSignInState() {
  const context = useContext(SignInContext);
  if (context === null) {
    throw new Error('useIsLoginState must be used within a LoginProvider');
  }
  return context;
}

// SignInProvider の型情報を提供
type SignInProviderProps = {
  children: React.ReactNode;
};

export function LoginBtnProvider({ children }: SignInProviderProps) {
  const [isLogin, setLoginState] = useState<boolean>(false);  // 初期値を設定

  const handleUserRegistClick = () => {
    setLoginState(true);
  };

  const handleBackClick = () => {
    setLoginState(false);
  };

  const handleLoginClick = () => {
    
  };

  return (
    <SignInContext.Provider value={{ isLogin, setLoginState }}>
      {children}
      <fieldset>
        <div>
          <p>ログイン</p>
        </div>
        <div>
          <input></input>
        </div>
        <div>
          <button onClick={handleUserRegistClick}>Register</button>
        </div>
      </fieldset>
    </SignInContext.Provider>
  );
}
