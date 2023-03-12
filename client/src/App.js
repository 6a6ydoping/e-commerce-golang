import Header from "./components/header";
import RegistrationForm from "./components/registrationForm";
import Home from "./components/home";
import "./App.css";
import {
    BrowserRouter as Router,
    Routes,
    Route,
    useNavigate,
    Link,
    redirect,
    Navigate,
} from "react-router-dom";
import React, { useState, useEffect } from "react";
import SellingItems from "./components/sellingItems";
import AuthForm from "./components/authForm";
import ItemList from "./components/itemList";
import Profile from "./components/profile";
import SearchBar from "./components/searchBar";
import PageNotFound from "./components/pageNotFound";

function App() {
    const [items, setItems] = useState([]);
    return (
        <div className="App">
            <Router>
                <Routes>
                    <Route
                        exact
                        path="/register"
                        element={<RegistrationForm />}
                    />
                    <Route exact path="/auth" element={<AuthForm />} />
                    <Route exact path="/home" element={<Home />} />
                    <Route
                        exact
                        path="/menu"
                        element={
                            <SearchBar
                                onSubmit={(results) => setItems(results)}
                            />
                        }
                    />
                    <Route exact path="/profile" element={<Profile />} />
                    <Route exact path="/item" element={<Profile />} />
                    <Route
                        path="/"
                        element={<Navigate replace={true} to="/home" />}
                    />
                    <Route path="*" element={<PageNotFound />} />
                </Routes>
            </Router>
        </div>
    );
}

export default App;
