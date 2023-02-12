import React, {useState, useEffect} from "react";
import './style.css';
import Header from './header';
import { useNavigate } from 'react-router-dom';

function AllItems(){
    const [data, setData] = useState([]);

    useEffect(() => {
        fetch('http://localhost:8000/menu', {
            method: 'GET',
            headers: {
                'Content-type': 'application/json',
            },
        })
            .then((response) => {
                if (!response.ok){
                    throw new Error(response.status)
                }else{
                    return response.json();
                }
            })
            .then((data) => {
                setData(data);
            })
            .catch((err) => {
                console.log(err.message);
            });
    }, []);

    return (
        <>
          <div>
            <p>All items</p>
            <ul>
              {data.map((item) => (
                <div>
                    <h1>Name: {item.Name} </h1>
                    <h3>Price: {item.Price}</h3>
                </div>

              ))}
            </ul>
          </div>
        </>
      );
}

export default AllItems;