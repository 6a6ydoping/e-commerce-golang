import React, { useState, useEffect } from 'react';
import axios from 'axios';

const SellingItems = () => {
  const [sellingItems, setSellingItems] = useState([]);

  useEffect(() => {
    axios.get('http://localhost:8000/menu')
      .then(res => setSellingItems(res.data))
      .catch(err => console.error(err));
  }, []);
  console.log(sellingItems);
  return (
    <ul>
      {sellingItems.map(item => (
        <li key={item.ID}>
          <h3>{item.Name}</h3>
          <p>{item.Price}</p>
        </li>
      ))}
    </ul>
  );
};

export default SellingItems;
