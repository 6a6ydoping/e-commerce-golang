import React, { useState, useEffect } from "react";
import "./style.css";
import "./allItems.css";
import "./searchBar.css";
import ItemList from "./itemList";

function SearchBar(props) {
    const [items, setItems] = useState([]);
    const [searchString, setSearchString] = useState("");

    const handleClick = async () => {
        if (searchString == "") {
            const response = await fetch(`http://localhost:8000/menu`);
            const json = await response.json();
            setItems(json);
        } else {
            const response = await fetch(
                `http://localhost:8000/menu?query=${searchString}`
            );
            const json = await response.json();
            setItems(json);
        }
    };

    const handleSubmit = async () => {};

    if (items.length === 0) {
        return (
            <>
                <link
                    rel="stylesheet"
                    href="https://fonts.googleapis.com/css2?family=Material+Symbols+Outlined:opsz,wght,FILL,GRAD@48,400,0,0"
                />
                <div className="searchBar__container">
                    <form className="searchBar__form" onSubmit={handleSubmit}>
                        <input
                            type="text"
                            placeholder="Enter name of the item..."
                            value={searchString}
                            onChange={(e) => setSearchString(e.target.value)}
                        />
                        <button type="button" onClick={handleClick}>
                            <span className="material-symbols-outlined">
                                search
                            </span>
                        </button>
                        <select name="cars" id="cars">
                            <option value="rating">Rating</option>
                            <option value="price">Price</option>
                        </select>
                    </form>
                    <p>No items found</p>
                </div>
            </>
        );
    }

    return (
        <>
            <link
                rel="stylesheet"
                href="https://fonts.googleapis.com/css2?family=Material+Symbols+Outlined:opsz,wght,FILL,GRAD@48,400,0,0"
            />
            <div className="searchBar__container">
                <form className="searchBar__form" onSubmit={handleSubmit}>
                    <input
                        type="text"
                        placeholder="Enter name of the item..."
                        value={searchString}
                        onChange={(e) => setSearchString(e.target.value)}
                    />
                    <button type="button" onClick={handleClick}>
                        <span className="material-symbols-outlined">
                            search
                        </span>
                    </button>
                    <select name="cars" id="cars">
                        <option value="rating">Rating</option>
                        <option value="price">Price</option>
                    </select>
                </form>
                <ItemList items={items} />
            </div>
        </>
    );
}

export default SearchBar;
