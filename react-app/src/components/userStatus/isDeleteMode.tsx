import React, { createContext, useEffect, useContext, useState } from 'react';
import { Link, useNavigate } from 'react-router-dom';

// 削除モードコンテキスト型情報
type isDeleteModeContextType = {
    isDeleteMode: boolean;
    setIsDeleteMode: React.Dispatch<React.SetStateAction<boolean>>;
};

// 削除モードコンテキスト
const isDeleteModeContext = createContext<isDeleteModeContextType | null>(null);

// 削除モード制御用use関数
export function useIsDeleteMode() {
    const context = useContext(isDeleteModeContext);
    if (context === null) {
        throw new Error('useClickedIds must be used within a ClickedIdsProvider');
    }
    return context;
}

// 削除モードプロバイダプロパティ
type isDeleteModeProviderProps = {
    children: React.ReactNode;
};

// 削除モードDOMプロバイダ
export function IsDeleteModeProvider({ children }: isDeleteModeProviderProps) {
    const [isDeleteMode, setIsDeleteMode] = useState<boolean>(false);

    return (
        <isDeleteModeContext.Provider value={{ isDeleteMode, setIsDeleteMode }}>
        {children}     
        </isDeleteModeContext.Provider>
    );
}