import React, { useState, setState } from "react";
import { Link } from "react-router-dom";
import Header from "./header";
import SellingItems from "./sellingItems";
import "./style.css";

function Home() {
    return (
        <>
            <Header headerName="Home page" />
            <SellingItems />
        </>
    );
}

export default Home;
