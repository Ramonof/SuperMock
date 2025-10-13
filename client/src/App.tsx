import { Container, Stack } from '@chakra-ui/react'
import Navbar from './components/navbar'
import ProjectForm from './components/ProjectForm';
import ProjectList from './components/ProjectList';
import { Route, Routes } from 'react-router';
import Projects from './pages/Projects';
import Main from './pages/__root';
import SideBar from './components/SideBar';
import RequireAuth from './components/RequireAuth';
import Login from './components/Login';

export const BASE_URL = import.meta.env.MODE === "development" ? "http://localhost:8000/api/v1" : "/api";

export const ROLES = {
  'User': 2001,
  'Editor': 1984,
  'Admin': 5150
}

function App() {

  return (
    <Stack h="100vh">
      <Navbar />
      <SideBar />
      <Routes>
        <Route path='/' element={<Main />} />
        <Route path='/login' element={<Login />} />
        {/* <Route path='/projects' element={<Projects />} /> */}
        <Route element={<RequireAuth allowedRoles={[ROLES.User]} />}>
          <Route path="/projects" element={<Projects />} />
        </Route>
      </Routes>
      {/* <Container>
        <ProjectForm />
        <ProjectList />
      </Container> */}
    </Stack>
  )
}

export default App
