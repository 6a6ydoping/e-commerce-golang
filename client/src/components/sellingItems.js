import React, { useState, useEffect } from "react";
import axios from "axios";

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
        <ul>
            {sellingItems.map((item) => (
                <li key={item.ID}>
                    <h3>{item.name}</h3>
                    <p>{item.price}</p>
                </li>
            ))}
        </ul>
    );
};

export default SellingItems;
