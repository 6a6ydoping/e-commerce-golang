import React, { useState, useEffect } from "react";
import axios from "axios";
import "./styles/sellingItems.css";

const SellingItems = (props) => {
    const [sellingItems, setSellingItems] = useState([]);
    const [loading, setLoading] = useState(true);

    useEffect(() => {
        const fetchSellingItems = async () => {
            try {
                console.log("SELLING ITEMS = " + props.query);
                const res = await axios.get(
                    `http://localhost:8000/menu?query=${props.query}`
                );
                setSellingItems(res.data);
                setLoading(false);
            } catch (err) {
                console.error(err);
            }
        };

        fetchSellingItems();
    }, [props.query]);

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
