import Header from './components/header';
import RegistrationForm from './components/registrationForm'
import Home from './components/home'
import './App.css';
import {
  BrowserRouter as Router,
  Routes,
  Route,
  useNavigate,
  Link
} from "react-router-dom";
import SellingItems from './components/sellingItems';
import AuthForm from './components/authForm';
import AllItems from './components/allItems';


function App() {
  return (
    <div className="App">
      <Router>
        <Routes >
        <Route exact path="/register"  element={<RegistrationForm />}/>
        <Route exact path="/auth"  element={<AuthForm />}/>
        <Route exact path="/home" element={<Home/>}/>
        <Route exact path="/menu" element={<AllItems/>}/>
        </Routes>
    </Router>
    </div>
  );
}


export default App;
