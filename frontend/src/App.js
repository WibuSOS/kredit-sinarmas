import { BrowserRouter, Route, Routes } from 'react-router-dom';
import './App.css';
import Login from './components/login/Login';
import Header from './components/shared/Header';
import Sidebar from './components/shared/Sidebar';
import UserContext, { useStore } from './Context';

function MultiRouter() {
  const { state } = useStore();

  if (!state.user?.token) {
    return <Login />
  }

  return (
    <BrowserRouter>
      <Header />
      <Sidebar />
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
