import React, {useState,setState} from 'react';
import { useNavigate } from 'react-router-dom';
import Header from './header';
import './style.css';
import axios from 'axios';
function RegistrationForm() {
    
    const [firstName, setFirstName] = useState(null);
    const [lastName, setLastName] = useState(null);
    const [email, setEmail] = useState(null);
    const [password,setPassword] = useState(null);
    const [confirmPassword,setConfirmPassword] = useState(null);
    let [userType, setUserType] = useState('Client');

    const navigate = useNavigate()

    const handleInputChange = (e) => {

        const {id , value} = e.target;
        if(id === "firstName"){
            setFirstName(value);
        }
        if(id === "lastName"){
            setLastName(value);
        }
        if(id === "email"){
            setEmail(value);
        }
        if(id === "password"){
            setPassword(value);
        }
        if(id === "confirmPassword"){
            setConfirmPassword(value);
        }
    }

    const handleOptionChange = (event) => {
        setUserType(event.target.value)
    }

    const handleSubmit = () => {
        fetch('http://localhost:8000/register', {
            method: 'POST',
            body: JSON.stringify({
              firstName: firstName,
              lastName: lastName,
              password: password,
              email: email,
              password: password,
              userType: userType,
            }),
            headers: {
              'Content-type': 'application/json',
            },
          })
             .then((response) => {
                if (!response.ok){
                    throw new Error(response.status)
                }else{
                    navigate('/auth')
                }
             })
             .then((data) => {
                console.log(userType);
                // Handle data
             })
             .catch((err) => {
                console.log(err.message);
             });
    }

    return(
        <>
        <Header headerName="Registration"/>
        <div className="form">
            <div className="form-body">
                <div className="username">
                    <label className="form__label" htmlFor="firstName">First Name </label>
                    <input className="form__input" type="text" value={firstName} onChange = {(e) => handleInputChange(e)} id="firstName" placeholder="First Name"/>
                </div>
                <div className="lastname">
                    <label className="form__label" htmlFor="lastName">Last Name </label>
                    <input  type="text" name="" id="lastName" value={lastName}  className="form__input" onChange = {(e) => handleInputChange(e)} placeholder="LastName"/>
                </div>
                <div className="email">
                    <label className="form__label" htmlFor="email">Email </label>
                    <input  type="email" id="email" className="form__input" value={email} onChange = {(e) => handleInputChange(e)} placeholder="Email"/>
                </div>
                <div className="password">
                    <label className="form__label" htmlFor="password">Password </label>
                    <input className="form__input" type="password"  id="password" value={password} onChange = {(e) => handleInputChange(e)} placeholder="Password"/>
                </div>
                <div className="confirm-password">
                    <label className="form__label" htmlFor="confirmPassword">Confirm Password </label>
                    <input className="form__input" type="password" id="confirmPassword" value={confirmPassword} onChange = {(e) => handleInputChange(e)} placeholder="Confirm Password"/>
                </div>
                <div className="user-type">
                    <input type="radio" id="choice1" name="userType" value="Client" onChange={handleOptionChange} checked={userType === "Client"}/>
                    <label for="choice1">Client</label>

                    <input type="radio" id="choice2" name="userType" value="Seller" onChange={handleOptionChange} checked={userType === "Seller"}/>
                    <label for="choice2">Seller</label>
                </div>
            </div>
            <div class="footer">
                <button onClick={()=>handleSubmit()} type="submit" className="btn">Register</button>
            </div>
        </div>
        </>
    )       
}

export default RegistrationForm