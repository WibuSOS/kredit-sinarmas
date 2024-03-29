import { Container } from 'react-bootstrap';
import { BrowserRouter, Route, Routes } from 'react-router-dom';
import './App.css';
import ChangePassword from './components/changePassword/ChangePassword';
import ChecklistPencairan from './components/kredit/checklistPencairan/ChecklistPencairan';
import DrawdownReport from './components/kredit/drawdownReport/DrawdownReport';
import Login from './components/login/Login';
import Header from './components/shared/Header';
import Sidebar from './components/shared/Sidebar';
import UserContext, { useStore } from './Context';

function MultiRouter() {
  const { state } = useStore();

  if (!state.loggedIn) {
    return <Login />
  }

  return (
    <BrowserRouter>
      <Header />
      <Sidebar />
      <Container as="main">
        <Routes>
          <Route path="/" element={<ChecklistPencairan />} exact />
          <Route path="/change_password" element={<ChangePassword />} exact />
          <Route path="/kredit/checklist_pencairan" element={<ChecklistPencairan />} exact />
          <Route path="/kredit/drawdown_report" element={<DrawdownReport />} exact />
        </Routes>
      </Container>
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
