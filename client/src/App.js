import Header from './components/header';
import RegistrationForm from './components/registrationForm'
import './App.css';

// const [firstName, setFirstName] = useState(null);
// const [lastName, setLastName] = useState(null);
// const [email, setEmail] = useState(null);
// const [password,setPassword] = useState(null);
// const [confirmPassword,setConfirmPassword] = useState(null);



function App() {
  return (
    <div className="App">
      <Header/>
      <RegistrationForm/>
    </div>
  );
}


export default App;
