import React, { useState, useEffect } from "react";
import "./style.css";
import "./allItems.css";
import Header from "./header";
import { useNavigate } from "react-router-dom";
import SearchBar from "./searchBar";

function ItemList({ items }) {
    console.log("PROPS: " + JSON.stringify(items));

    return (
        <>
            <div>
                <ul>
                    {items.map((item) => (
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
