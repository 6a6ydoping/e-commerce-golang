import React, { useState, useEffect } from "react";
import "./style.css";
import "./allItems.css";
import "./searchBar.css";
import ItemList from "./itemList";

function SearchBar(props) {
    const options = ["Rating", "Price"];
    const [items, setItems] = useState([]);
    const [searchString, setSearchString] = useState("");
    const [selectedOption, setSelectedOption] = useState("rating");

    const handleFilterByChange = (e) => {
        const selected = e.target.value;
        setSelectedOption(selected);
    };

    const handleClick = async () => {
        if (searchString == "") {
            const response = await fetch(`http://localhost:8000/menu`);
            const json = await response.json();
            setItems(json);
        } else {
            const response = await fetch(
                `http://localhost:8000/menu?query=${searchString}`,
                {
                    method: "POST",
                    body: JSON.stringify({
                        searchString: searchString,
                        filterBy: selectedOption,
                    }),
                    headers: {
                        "Content-type": "application/json",
                    },
                }
            );
            const json = await response.json();
            setItems(json);
        }
    };

    const handleSubmit = async () => {};

    return (
        <>
            <div className="searchBar__container">
                <form className="searchBar__form" onSubmit={handleSubmit}>
                    <input
                        className="searchBar"
                        placeholder="Search games..."
                        value={searchString}
                        onChange={(e) => setSearchString(e.target.value)}
                    ></input>
                    <button type="button" onClick={handleClick}>
                        <svg
                            xmlns="http://www.w3.org/2000/svg"
                            viewBox="0 0 48 48"
                            className="NavBar_svg__wdLyI"
                            id="search_logo"
                            aria-label="Search"
                            style={{
                                fill: "#ccc",
                            }}
                            {...props}
                        >
                            <path d="M39.8 41.95 26.65 28.8q-1.5 1.3-3.5 2.025-2 .725-4.25.725-5.4 0-9.15-3.75T6 18.75q0-5.3 3.75-9.05 3.75-3.75 9.1-3.75 5.3 0 9.025 3.75 3.725 3.75 3.725 9.05 0 2.15-.7 4.15-.7 2-2.1 3.75L42 39.75Zm-20.95-13.4q4.05 0 6.9-2.875Q28.6 22.8 28.6 18.75t-2.85-6.925Q22.9 8.95 18.85 8.95q-4.1 0-6.975 2.875T9 18.75q0 4.05 2.875 6.925t6.975 2.875Z" />
                        </svg>
                    </button>
                </form>
            </div>
        </>
    );
}

export default SearchBar;
