import { BrowserRouter, Route, Routes } from 'react-router-dom';
import './App.css';
// import Header from './components/Header';
import Login from './components/login/Login';
import UserContext, { useStore } from './Context';

function MultiRouter() {
  const { state } = useStore()

  if (!state.user?.token) {
    return <Login />
  }

  return (
    <BrowserRouter>
      {/* <Header /> */}
      <Routes>
        <Route path="/" element={console.log("ada token")} exact />
      </Routes>
    </BrowserRouter>
  )
}

function App() {
  return (
    <UserContext>
      <MultiRouter />
    </UserContext>
  )
}

export default App;
