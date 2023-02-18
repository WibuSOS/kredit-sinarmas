import { BrowserRouter, Route, Routes } from 'react-router-dom';
import './App.css';
// import Header from './components/Header';
import Login from './components/login/Login';
import UserContext, { useStore } from './Context';

function MultiRouter() {
  const { state } = useStore()

  if (!state.user?.token) {
    return (
      <Routes>
        <Route path="/" element={<Login />} exact />
      </Routes>
    )
  }

  return (
    <Routes>
      <Route path="/" element={console.log("ada token")} exact />
    </Routes>
  )
}

function App() {
  return (
    <UserContext>
      <BrowserRouter>
        {/* <Header /> */}
        <MultiRouter />
      </BrowserRouter>
    </UserContext>
  )
}

export default App;
