import React, { useState, useEffect } from 'react';
import { useClickedIds, ClickedIdsProvider } from './article_click';

type BackGroundImageComponentProps = {
  id: number;
  thumbnailFileName: string;
  children: React.ReactNode;
};

function BackGroundImageComponent({ id, thumbnailFileName, children }: BackGroundImageComponentProps) {
  const [imageData, setImageData] = useState<string | null>(null);
  const { clickedIds, setClickedIds } = useClickedIds();

  useEffect(() => {
    // Ginサーバーから画像データを取得
    if (thumbnailFileName != '' && thumbnailFileName != null && thumbnailFileName != undefined) {
        fetch(`http://localhost:8080/GetThumbnail/${thumbnailFileName}`)
        .then(response => response.blob())
        .then(data => {
          setImageData(URL.createObjectURL(data));
        });
    }
  }, [thumbnailFileName]);

  return (
    <div className={`content-background ${clickedIds.includes(id) ? 'clicked' : 'unclicked'}`} key={id}        
        onClick={() => {if(!clickedIds.includes(id)) setClickedIds([id])}}
        style={{
        backgroundImage: `url(${imageData})`,
      }}>
      {children} {}
    </div>
  );
}

export default BackGroundImageComponent;
