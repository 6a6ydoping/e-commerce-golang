import React, { useState, setState } from "react";
import { Link } from "react-router-dom";
import Header from "./header";
import SellingItems from "./sellingItems";
import "./style.css";

function Home() {
    const [searchString, setSearchString] = useState("");

    const handleSearch = (query) => {
        setSearchString(query);
    };

    return (
        <>
            <Header headerName="Home page" onSearch={handleSearch} />
            <SellingItems query={searchString} />
        </>
    );
}

export default Home;
