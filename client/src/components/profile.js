import React, {useState, useEffect} from "react";
import './style.css';
import './allItems.css'
import Header from './header';
import { useNavigate } from 'react-router-dom';

function Profile(){
    const [data, setData] = useState([]);
    console.log(document.cookie);
    const token = document.cookie.split('token=')[1];
    console.log(token);
    useEffect(() => {
        fetch('http://localhost:8000/profile', {
            method: 'GET',
            headers: {
                'Cookie': 'token=' + document.cookie.split('token=')[1].split(';')[0],
            },
            credentials: 'include',
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
            <Header headerName="Profile"/>
            <ul>
              {data.map((profile) => (
                <div className="item__form">
                    <h1>Name: {profile.firstName} </h1>
                    <h3>Price: {profile.lastName}</h3>
                </div>

              ))}
            </ul>
          </div>
        </>
      );
}

export default Profile;