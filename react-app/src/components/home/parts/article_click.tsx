import React, { createContext, useContext, useState } from 'react';

// コンテキストの型情報を定義
type ClickedIdsContextType = {
  clickedIds: number[];
  setClickedIds: React.Dispatch<React.SetStateAction<number[]>>;
};

const ClickedIdsContext = createContext<ClickedIdsContextType | null>(null);

export function useClickedIds() {
  const context = useContext(ClickedIdsContext);
  if (context === null) {
    throw new Error('useClickedIds must be used within a ClickedIdsProvider');
  }
  return context;
}

// ClickedIdsProvider の型情報を提供
type ClickedIdsProviderProps = {
  children: React.ReactNode;
};

export function ClickedIdsProvider({ children }: ClickedIdsProviderProps) {
  const [clickedIds, setClickedIds] = useState<number[]>([]);

  return (
    <ClickedIdsContext.Provider value={{ clickedIds, setClickedIds }}>
      {children}
    </ClickedIdsContext.Provider>
  );
}
