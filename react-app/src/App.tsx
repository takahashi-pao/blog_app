import axios from 'axios';
import React, { useEffect, useState } from 'react';

type Member = {
  id: number;
  name: string;
  icon: string;
}

// var members = [
//   {
//     id: 1,
//     name: "test1",
//     icon: "👻"
//   },
//   {
//     id: 2,
//     name: "test2",
//     icon: "😈"
//   },
//   {
//     id: 3,
//     name: "test3",
//     icon: "🧠"
//   },
// ]

function App() {
  const [members, setMembers] =
    useState<Member[]>([{id: 0, name: "", icon: ""}])

  useEffect(() => {
    (
      async() => {
        const data = await axios.get("http://localhost:8080")
        setMembers(data.data)
      }
    )()
  }, [])

  return (
    <div>
      {members.map(member => (
        <p key={member.id}>
          <span>{member.name}</span>
          <span>{member.icon}</span>
        </p>
      ))}
    </div>
  );
}

export default App;
