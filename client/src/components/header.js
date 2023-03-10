import React, { useState } from "react";
import SearchBar from "./searchBar";
import "./styles/header.css";
import "./styles/header.css";
function Header(props) {
    const [searchString, setSearchString] = useState("");

    const handleSearch = (query) => {
        console.log("HANDLE SEARCH" + query);
        setSearchString(query);
        props.onSearch(query);
    };

    return (
        <nav className="navbar">
            <div className="store_name__container">
                <svg
                    xmlns="http://www.w3.org/2000/svg"
                    viewBox="0 0 48 48"
                    height="100%"
                    width="100%"
                    className="NavBar_svg__wdLyI"
                    id="gamepad_logo"
                    style={{
                        fill: "#fff",
                    }}
                    {...props}
                >
                    <path d="M7.35 38q-1.9 0-2.975-1.275Q3.3 35.45 3.6 33.2L6 16.45q.4-2.65 2.475-4.55T13.2 10h21.65q2.65 0 4.725 1.9 2.075 1.9 2.475 4.55L44.4 33.2q.3 2.25-.775 3.525T40.65 38q-1.15 0-1.95-.375t-1.35-.925l-5.2-5.2h-16.3l-5.2 5.2q-.55.55-1.35.925T7.35 38Zm.9-3.2 6.3-6.3h18.9l6.3 6.3q.25.25.9.45.45 0 .675-.45.225-.45.125-.9l-2.4-16.95q-.25-1.75-1.475-2.85T34.85 13h-21.7q-1.5 0-2.725 1.1T8.95 16.95L6.55 33.9q-.1.45.125.9t.675.45q.35 0 .9-.45ZM35 26q.8 0 1.4-.6.6-.6.6-1.4 0-.8-.6-1.4-.6-.6-1.4-.6-.8 0-1.4.6-.6.6-.6 1.4 0 .8.6 1.4.6.6 1.4.6Zm-4.25-6.5q.8 0 1.4-.6.6-.6.6-1.4 0-.8-.6-1.4-.6-.6-1.4-.6-.8 0-1.4.6-.6.6-.6 1.4 0 .8.6 1.4.6.6 1.4.6ZM15 25.75h2.5V22h3.75v-2.5H17.5v-3.75H15v3.75h-3.75V22H15Zm9-1.65Z" />
                </svg>
                <h3 className="store_name">Game Store</h3>
            </div>
            <SearchBar onSearch={handleSearch} />
            <div className="end__container">
                <div className="github__container">
                    <svg
                        stroke="currentColor"
                        fill="currentColor"
                        strokeWidth={0}
                        viewBox="0 0 24 24"
                        height="1em"
                        width="1em"
                        xmlns="http://www.w3.org/2000/svg"
                        className="NavBar_gh__Z0SNl"
                        id="github_logo"
                        {...props}
                    >
                        <g stroke="none">
                            <path fill="none" d="M0 0h24v24H0z" />
                            <path d="M5.883 18.653c-.3-.2-.558-.455-.86-.816a50.32 50.32 0 0 1-.466-.579c-.463-.575-.755-.84-1.057-.949a1 1 0 0 1 .676-1.883c.752.27 1.261.735 1.947 1.588-.094-.117.34.427.433.539.19.227.33.365.44.438.204.137.587.196 1.15.14.023-.382.094-.753.202-1.095C5.38 15.31 3.7 13.396 3.7 9.64c0-1.24.37-2.356 1.058-3.292-.218-.894-.185-1.975.302-3.192a1 1 0 0 1 .63-.582c.081-.024.127-.035.208-.047.803-.123 1.937.17 3.415 1.096A11.731 11.731 0 0 1 12 3.315c.912 0 1.818.104 2.684.308 1.477-.933 2.613-1.226 3.422-1.096.085.013.157.03.218.05a1 1 0 0 1 .616.58c.487 1.216.52 2.297.302 3.19.691.936 1.058 2.045 1.058 3.293 0 3.757-1.674 5.665-4.642 6.392.125.415.19.879.19 1.38a300.492 300.492 0 0 1-.012 2.716 1 1 0 0 1-.019 1.958c-1.139.228-1.983-.532-1.983-1.525l.002-.446.005-.705c.005-.708.007-1.338.007-1.998 0-.697-.183-1.152-.425-1.36-.661-.57-.326-1.655.54-1.752 2.967-.333 4.337-1.482 4.337-4.66 0-.955-.312-1.744-.913-2.404a1 1 0 0 1-.19-1.045c.166-.414.237-.957.096-1.614l-.01.003c-.491.139-1.11.44-1.858.949a1 1 0 0 1-.833.135A9.626 9.626 0 0 0 12 5.315c-.89 0-1.772.119-2.592.35a1 1 0 0 1-.83-.134c-.752-.507-1.374-.807-1.868-.947-.144.653-.073 1.194.092 1.607a1 1 0 0 1-.189 1.045C6.016 7.89 5.7 8.694 5.7 9.64c0 3.172 1.371 4.328 4.322 4.66.865.097 1.201 1.177.544 1.748-.192.168-.429.732-.429 1.364v3.15c0 .986-.835 1.725-1.96 1.528a1 1 0 0 1-.04-1.962v-.99c-.91.061-1.662-.088-2.254-.485z" />
                        </g>
                    </svg>
                    <h3>
                        <a href="https://github.com/6a6ydoping" target="_blank">
                            <h3 className="github">6a6ydoping</h3>
                        </a>
                    </h3>
                </div>
                <div className="balance__container">
                    <svg
                        xmlns="http://www.w3.org/2000/svg"
                        viewBox="0 0 470 470"
                        xmlSpace="preserve"
                        className="NavBar_svg2__tmR5o"
                        id="balance_logo"
                        style={{
                            fill: "transparent",
                            stroke: "#fff",
                            strokeWidth: 34,
                        }}
                        {...props}
                    >
                        <path d="M323.85 131.05v-42.2C323.85 39.858 283.992 0 235 0s-88.85 39.858-88.85 88.85v42.2H72V470h326V131.05h-74.15zm-30 0h-117.7v-42.2C176.15 56.4 202.55 30 235 30s58.85 26.4 58.85 58.85v42.2z" />
                    </svg>
                    <h3 className="balance">Balance: 0</h3>
                </div>
            </div>
        </nav>
    );
}
export default Header;
