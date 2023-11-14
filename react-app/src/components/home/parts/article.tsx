import axios from 'axios';
import React, { useEffect, useState } from 'react';
import { useClickedIds } from './article_click';
import { useIsDeleteMode } from '../../userStatus/isDeleteMode';

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

  useEffect(() => {
    console.log("delete")
  }, [isDelete])

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
                {member.tag.map((tagItem, index) => (
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
