import React, { useState, useEffect } from 'react';
import { useClickedIds } from './article_click';
import { useIsDeleteMode } from '../../userStatus/isDeleteMode';
import { posix } from 'path';

type BackGroundImageComponentProps = {
  id: number;
  thumbnailFileName: string;
  children: React.ReactNode;
};

function BackGroundImageComponent({ id, thumbnailFileName, children }: BackGroundImageComponentProps) {
  const [imageData, setImageData] = useState<string | null>(null);
  const { clickedIds, setClickedIds } = useClickedIds();
  const {isDeleteMode, setIsDeleteMode} = useIsDeleteMode();

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

  const handleIsClicked = () => {
    if(clickedIds.includes(id) && !isDeleteMode){
      return 'clicked'
    }

    return 'unclicked '
  }

  return (
    <div className={`overflow-hidden`}>
      <div className={`${clickedIds.includes(id) && isDeleteMode ? 'content wd-100 hi-100 op-50 bg-black delete-mode-background-effect' : ''}`} key={id+'_deleteModeBackGroundEffect'}></div>
      <div className={`content-background ${handleIsClicked()}`} key={id}        
          style={{
          backgroundImage: `url(${imageData})`,
        }}>
        {children} {}
      </div>
    </div>

  );
}

export default BackGroundImageComponent;
