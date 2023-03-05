import { Container } from 'react-bootstrap';
import { BrowserRouter, Route, Routes } from 'react-router-dom';
import './App.css';
import Home from './components/home/Home';
import ChecklistPencairan from './components/kredit/checklistPencairan/ChecklistPencairan';
import DrawdownReport from './components/kredit/drawdownReport/DrawdownReport';
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
      <Container as="main">
        <Routes>
          <Route path="/" element={<ChecklistPencairan />} exact />
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
