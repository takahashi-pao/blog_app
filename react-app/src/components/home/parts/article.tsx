import axios from 'axios';
import React, { useEffect, useState } from 'react';
import { useClickedIds } from './article_click';
import { useIsDeleteMode } from '../../userStatus/isDeleteMode';
import { GetByteSize } from '../../../common'
import { kMaxLength } from 'buffer';

type Member = {
  id: number;
  title: string;
  date: string;
  tag: string[];
}

function Article(member: Member) {
  const [members, setMembers] = useState<Member[]>([{ id: 0, title: '', date: '', tag: []}]);
  const [hoveredIds, setHoveredIds] = useState<number[]>([]);
  const { clickedIds, setClickedIds } = useClickedIds();
  const {isDeleteMode, setIsDeleteMode} = useIsDeleteMode();
  const [isDelete, setIsDelete] = useState<boolean>(false);

  const handleOnClick = (targetId:number) => {
    if(isDeleteMode && clickedIds.includes(targetId)){
      setClickedIds(
        clickedIds.filter((id, index) => (id !== targetId))
      )
    }else if(isDeleteMode){
      setClickedIds([...clickedIds,targetId]) 
    }else if(!clickedIds.includes(targetId)){
      setClickedIds([targetId])
    }
  }

  const handleIsClicked = () => {
    if(clickedIds.includes(member.id) && !isDeleteMode){
      return 'clicked'
    }

    return 'unclicked '
  }

  const moldTitleLength = () => {
    const byte: number = GetByteSize(member.title)
    console.log(member.id+":byte="+byte+" "+member.title.substring(0, 10)+"...")
    console.log(member.title+":"+truncateStringToBytes())
  }

  function truncateStringToBytes() {
    let truncated = '';
    let bytes = 0;
  
    for (let i = 0; i < member.title.length; i++) {
      const char = member.title.charAt(i);
      const charCode = char.charCodeAt(0);
  
      // UTF-8 のエンコード方式に基づいて、1 バイトから 4 バイトまでの範囲で文字をカウント
      bytes += charCode < 0x80 ? 1 : charCode < 0x800 ? 2 : charCode < 0x10000 ? 3 : 4;
  
      if (bytes > 22) {
        break;
      }
  
      truncated += char;
    }

    if (member.title.length > truncated.length){
      truncated += "..."
    }
  
    return truncated;
  }

  useEffect(() => {
    console.log(clickedIds)
    moldTitleLength()
  }, [clickedIds])

  // 削除モード管理
  useEffect(() => {   
    console.log(isDeleteMode)
  }, [isDeleteMode]);
  
  return (
    <div className='wd-100 hi-100'>
        <div className={`content-background-blur ${handleIsClicked()}`}></div>
        <div className={`content-background-window ${handleIsClicked()}`}></div>

        <div
        className={`close-button ${clickedIds.includes(member.id) ? isDeleteMode ? 'undisplayed' : 'displayed' : 'undisplayed'}`}
        onClick={() => setClickedIds([])}
        ></div>
        <div
        className={`content text-white ${hoveredIds.includes(member.id) ? 'hovered' : 'unhovered'}`}
        onMouseOver={() => setHoveredIds([member.id])}
        onMouseOut={() => setHoveredIds([])}
        onClick={() => handleOnClick(member.id)}
        >
            <div className={`${clickedIds.includes(member.id) ? isDeleteMode ? 'wd-100 hi-100 op-50 bg-black': '' : 'wd-100 hi-100 op-50 bg-black'}`}></div>
            <div className='discription'>                
                <p className='date text-white'>{member.date}</p>
                <p className='title text-white'>{member.title}</p>

                <div>
                {member.tag?.map((tagItem, index) => (
                <span className='tag text-white' key={index}>
                #{tagItem}
                </span>
                ))}
                </div>
            </div>
        </div>
    </div>
  );
}

export default Article;
