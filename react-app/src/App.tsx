import axios from 'axios';
import React, { useEffect, useState } from 'react';

type Member = {
  id: number;
  title: string;
  date: string;
  tag: string[];
}

function App() {
  const [members, setMembers] = useState<Member[]>([{ id: 0, title: '', date: '', tag: [] }]);
  const [hoveredIds, setHoveredIds] = useState<number[]>([]);
  const [clickedIds, setClickedIds] = useState<number[]>([]);

  useEffect(() => {
    (async () => {
      const data = await axios.get('http://localhost:8080');
      setMembers(data.data);
    })();
  }, []);

  const handleMouseOver = (id: number) => {
    setHoveredIds([id]);
  };

  const handleMouseOut = (id: number) => {
    setHoveredIds([]);
  };

  const handleElementClick = (id: number) => {
    setClickedIds([id]);
  }
  const handleCloseElementClick = (id: number) => {
    setClickedIds([]);
  }

  return (
    <div className='content-wrap'>
      {members.map((member) => (
        <div className={`content-background ${clickedIds.includes(member.id) ? 'clicked' : 'unclicked'}`} key={member.id}>
          <div className={`content-background-blur ${clickedIds.includes(member.id) ? 'clicked' : 'unclicked'}`}></div>
          <div className={`content-background-window ${clickedIds.includes(member.id) ? 'clicked' : 'unclicked'}`}></div>

          <span
            className={`close-button ${clickedIds.includes(member.id) ? 'displayed' : 'undisplayed'}`}
            onClick={() => handleCloseElementClick(member.id)}
          ></span>
          <div
            className={`content text-white ${hoveredIds.includes(member.id) ? 'hovered' : 'unhovered'}`}
            onMouseOver={() => handleMouseOver(member.id)}
            onMouseOut={() => handleMouseOut(member.id)}
            onClick={() => handleElementClick(member.id)}
          >
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
      ))}
    </div>
  );
}

export default App;
