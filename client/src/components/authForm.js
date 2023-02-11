import React, {useState,setState} from 'react';
import './style.css';
import Header from './header';
import { useNavigate } from 'react-router-dom';
function AuthForm() {
    const [email, setEmail] = useState(null);
    const [password,setPassword] = useState(null);
    const [userType, setUserType] = useState(null);

    const navigate = useNavigate();

    const handleInputChange = (e) => {

        const {id , value} = e.target;
        if(id === "email"){
            setEmail(value);
        }
        if(id === "password"){
            setPassword(value);
        }
        if(id === "choice1" || id === "choice2"){
            setUserType(value);
        }

    }
    const data = {
        email: email,
        password: password,
        userType: userType,}
    
    const handleSubmit = () => {
        fetch('http://localhost:8000/auth', {
            method: 'POST',
            body: JSON.stringify({
              email: email,
              password: password,
              userType: userType
            }),
            headers: {
              'Content-type': 'application/json',
            },
          })
             .then((response) => {
                if(!response.ok){
                    throw new Error(response.status)
                }else{
                    navigate('/home')
                }
             })
             .then((data) => {
                console.log(data);
                // Handle data
             })
             .catch((err) => {
                console.log(err.message);
             });
    }

    return(
        <>
        <Header headerName="Login"/>
       <div className="form">
         <div className="form-body">
            <div className="userEmail__form">
                <label className="userEmail__label form__label" htmlFor="userEmail">Email</label>
                <input className="userEmail__label__input form__input" type="text" id="userEmail" value={email} onChange = {(e) => handleInputChange(e)} placeholder="Email"></input>           
            </div>
            <div className="password__form"> 
                <label className="password__label form__label" htmlFor="password">Password</label>
                <input className="password__label__input form__input" type="password" id="password" value={password} onChange = {(e) => handleInputChange(e)} placeholder="Password"></input>           
            </div>

            <div className="user-type">
                    <input type="radio" id="choice1" name="userType" value="Client" checked/>
                    <label for="choice1">Client</label>

                    <input type="radio" id="choice2" name="userType" value="Seller"/>
                    <label for="choice2">Seller</label>
            </div>

            <div class="footer">
                <button onClick={()=>handleSubmit()} type="submit" className="btn">Register</button>
            </div>
            </div>
       </div>
       </>
    )       
}

export default AuthForm