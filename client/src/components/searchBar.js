import React, { useState, useEffect } from "react";
import "./style.css";
import "./allItems.css";
import Header from "./header";
import { useNavigate } from "react-router-dom";
import ItemList from "./itemList";

function SearchBar() {
    const [searchQuery, setSearchQuery] = useState("");
    const [searchResult, setSearchResult] = useState("");
    const query = "";

    const handleSearch = async () => {
        const response = await fetch(`/api/search?q=${query}`);
        const results = await response.json();
        onSearch(results);
    };

    return (
        <>
            <link
                rel="stylesheet"
                href="https://fonts.googleapis.com/css2?family=Material+Symbols+Outlined:opsz,wght,FILL,GRAD@48,400,0,0"
            />
            <div className="searchBar__container">
                <form onSubmit={handleSubmit}>
                    <input
                        type="text"
                        placeholder="Enter name of the item..."
                        value={searchQuery}
                        onChange={(event) => setSearchQuery(event.target.value)}
                    />
                    <button type="submit">
                        {/* <span className="material-symbols-outlined">
                            search
                        </span> */}
                    </button>
                </form>
                <ItemList query={searchQuery} />
            </div>
        </>
    );
}

export default SearchBar;
