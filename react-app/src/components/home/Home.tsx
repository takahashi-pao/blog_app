import axios from 'axios';
import React, { useEffect, useState } from 'react';
import BackGroundImageComponent from './parts/backGroundImage';
import Article from './parts/article';
import { ClickedIdsProvider } from './parts/article_click';
import { Link } from 'react-router-dom';
import UserStatus from '../userStatus/userStatus';

type Member = {
  id: number;
  title: string;
  date: string;
  tag: string[];
  thumbnail: string;
}

function Home() {
  const [members, setMembers] = useState<Member[]>([{ id: 0, title: '', date: '', tag: [], thumbnail: ''}]);

  useEffect(() => {
    (async () => {
      const data = await axios.get('http://localhost:8080/GetArticle');
      setMembers(data.data);
    })();
  }, []);

  return (    
    <ClickedIdsProvider>
      <div>
        <UserStatus></UserStatus>
      </div>
      
      <div className='content-wrap'>
        {members && members.map((member) => (
            <BackGroundImageComponent id={member.id} thumbnailFileName={member.thumbnail} key={member.id}>
              <Article id={member.id} title={member.title} date={member.date} tag={member.tag}></Article>
            </BackGroundImageComponent>
        ))}
      </div>
    </ClickedIdsProvider>
  );
}

export default Home;
