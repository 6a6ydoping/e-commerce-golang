import React, { useState, useEffect } from "react";
import "./style.css";
import "./allItems.css";
import Header from "./header";
import { useNavigate } from "react-router-dom";
import SearchBar from "./searchBar";

function ItemList(props) {
    const [data, setData] = useState([]);

    useEffect(() => {
        fetch(`http://localhost:8000/menu?query=${props.query}`, {
            method: "GET",
            headers: {
                "Content-type": "application/json",
            },
        })
            .then((response) => {
                if (!response.ok) {
                    throw new Error(response.status);
                } else {
                    return response.json();
                }
            })
            .then((data) => {
                console.log(data);
                setData(data);
            })
            .catch((err) => {
                console.log(err.message);
            });
    }, []);

    return (
        <>
            <div>
                <ul>
                    {data.map((item) => (
                        <div className="item__form">
                            <h1>Name: {item.name} </h1>
                            <h3>Price: {item.price}</h3>
                        </div>
                    ))}
                </ul>
            </div>
        </>
    );
}

export default ItemList;
