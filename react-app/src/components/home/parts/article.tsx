import axios from 'axios';
import React, { useEffect, useState } from 'react';
import { useClickedIds, ClickedIdsProvider } from './article_click';

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

  useEffect(() => {
    
  }, []);
  
  return (
    <div className='wd-100 hi-100'>
        <div className={`content-background-blur ${clickedIds.includes(member.id) ? 'clicked' : 'unclicked'}`}></div>
        <div className={`content-background-window ${clickedIds.includes(member.id) ? 'clicked' : 'unclicked'}`}></div>

        <div
        className={`close-button ${clickedIds.includes(member.id) ? 'displayed' : 'undisplayed'}`}
        onClick={() => setClickedIds([])}
        ></div>
        <div
        className={`content text-white ${hoveredIds.includes(member.id) ? 'hovered' : 'unhovered'}`}
        onMouseOver={() => setHoveredIds([member.id])}
        onMouseOut={() => setHoveredIds([])}
        onClick={() => setClickedIds([member.id])}
        >
            <div className={`${clickedIds.includes(member.id) ? '' : 'wd-100 hi-100 op-50 bg-black'}`}></div>
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
