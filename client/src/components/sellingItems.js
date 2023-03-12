import React, { useState, useEffect } from "react";
import axios from "axios";
import "./styles/sellingItems.css";

const SellingItems = () => {
    const [sellingItems, setSellingItems] = useState([]);
    const [loading, setLoading] = useState(true);

    useEffect(() => {
        axios
            .get("http://localhost:8000/menu")
            .then((res) => {
                setSellingItems(res.data);
                setLoading(false);
            })
            .catch((err) => console.error(err));
    }, []);
    console.log(sellingItems);

    if (loading) {
        return <div>Loading...</div>;
    }

    return (
        <div class="selling__container">
            {sellingItems.map((item) => (
                // <li key={item.ID}>
                //     <h3>{item.name}</h3>
                //     <p>{item.price}</p>
                // </li>
                <div key={item.ID} class="itemBlock">
                    <h4 className="itemName">{item.name}</h4>
                    <p className="itemPrice">{item.price}</p>
                </div>
            ))}
        </div>
    );
};

export default SellingItems;
