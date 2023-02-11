import React from 'react';
import './style.css'
function Header(props) {
    return(
        <nav className="navbar-dark navbar reg-header">
            <div className="row col-12 d-flex justify-content-center text-white">
                <h3>{props.headerName}</h3>
            </div>
        </nav>
    )
}
export default Header;